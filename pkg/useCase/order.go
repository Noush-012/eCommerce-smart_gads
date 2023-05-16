package usecase

import (
	"context"
	"errors"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/config"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderUseCase struct {
	OrderRepository interfaces.OrderRepository
	UserRepository  interfaces.UserRepository
	PayMentRepo     interfaces.PaymentRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, UserRepo interfaces.UserRepository, payMentRepo interfaces.PaymentRepository) service.OrderService {
	return &OrderUseCase{
		OrderRepository: repo,
		UserRepository:  UserRepo,
		PayMentRepo:     payMentRepo}
}
func (o *OrderUseCase) CheckoutOrder(ctx context.Context, userId uint) (checkOut response.CartResp, err error) {
	checkOut, err = o.OrderRepository.CheckoutOrder(ctx, userId)
	if err != nil {
		return checkOut, err
	}
	return checkOut, nil
}
func (o *OrderUseCase) PlaceOrderByCOD(ctx context.Context, userId uint, PaymentMethodID uint) (shopOrder uint, err error) {
	// verify id is for COD or not
	payMethod, err := o.PayMentRepo.GetPaymentMethodByID(ctx, PaymentMethodID)
	if err != nil {
		return shopOrder, err
	}
	if payMethod == "Cash on delivery COD" {
		shopOrder, err = o.OrderRepository.PlaceCODOrder(ctx, userId, PaymentMethodID)
		if err != nil {
			return shopOrder, err
		}
		// Clear user cart
		if err := o.OrderRepository.ClearUserCart(ctx, userId); err != nil {
			return shopOrder, err
		}
	} else {
		return shopOrder, errors.New("requestID is not for Cash on delivery")
	}

	return shopOrder, nil
}

func (o *OrderUseCase) RazorPayCheckout(ctx context.Context, razorpay request.RazorpayReq) (razorpayOrder response.RazorPayOrderResp, err error) {
	// verify requested payment id is for razorpay
	payMethod, err := o.PayMentRepo.GetPaymentMethodByID(ctx, razorpay.PaymentMethodId)
	if err != nil {
		return razorpayOrder, err
	}
	if payMethod != "Razorpay" {
		return razorpayOrder, errors.New("requestID is not for Razorpay")
	}

	// find cart items
	CheckOut, err := o.OrderRepository.CheckoutOrder(ctx, razorpay.UserID)
	if err != nil {
		return razorpayOrder, err
	}
	if CheckOut.TotalProductItems == 0 {
		return razorpayOrder, errors.New("cart is empty")
	}

	// get user contact
	contact, err := o.UserRepository.GetEmailPhoneByUserId(ctx, razorpay.UserID)
	if err != nil {
		return razorpayOrder, err
	}
	razorpayOrder.Email = contact.Email
	razorpayOrder.Phone = contact.Phone

	// generate razorpay order id
	razorPayOrderId, err := utils.GenerateRazorPayOrder(CheckOut.TotalPrice, "Test reciept")
	if err != nil {
		return razorpayOrder, err
	}
	razorpayOrder.RazorpayOrderID = razorPayOrderId
	razorpayOrder.AmountToPay = CheckOut.TotalPrice
	razorpayOrder.RazorpayAmount = CheckOut.TotalPrice
	razorpayOrder.UserID = razorpay.UserID
	razorpayOrder.RazorpayKey = config.GetConfig().RazorPayKey

	// save order as pending
	Order := domain.ShopOrder{
		UserID:          razorpay.UserID,
		OrderTotal:      float64(CheckOut.TotalPrice),
		ShippingID:      CheckOut.DefaultShipping.ID,
		PaymentMethodID: razorpay.PaymentMethodId,
		CouponID:        0,
	}
	orderId, err := o.OrderRepository.SaveOrder(ctx, Order)
	if err != nil {
		return razorpayOrder, err
	}
	razorpayOrder.OrderID = orderId

	return razorpayOrder, nil

}

func (o *OrderUseCase) GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error) {
	orderHistory, err := o.OrderRepository.GetOrderHistory(ctx, page, userId)
	if err != nil {
		return orderHistory, err
	}
	return orderHistory, nil

}
func (o *OrderUseCase) UpdateOrderStatus(c context.Context, UpdateData request.UpdateOrderStatus) error {
	// // check if the payment success or not for procced order
	// ok, err := o.PayMentRepo.GetPaymentStatusByOrderId(c, UpdateData.OrderId)
	// if err != nil {
	// 	return err
	// }
	// if !ok {
	// 	return errors.New("payment failed ! not eligible to proceed")
	// }
	err := o.OrderRepository.ChangeOrderStatus(c, UpdateData)
	if err != nil {
		return err
	}
	// clear user cart
	if UpdateData.StatusId == 2 { // checking the order status is "placed" to clear cart items (ID 2 is for placed)
		if err := o.OrderRepository.ClearUserCart(c, UpdateData.UserId); err != nil {
			return err
		}
	}
	return nil
}
