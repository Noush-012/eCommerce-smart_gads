package usecase

import (
	"context"
	"errors"
	"time"

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
	CouponRepo      interfaces.CouponRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, UserRepo interfaces.UserRepository,
	payMentRepo interfaces.PaymentRepository, couponRepo interfaces.CouponRepository) service.OrderService {
	return &OrderUseCase{
		OrderRepository: repo,
		UserRepository:  UserRepo,
		PayMentRepo:     payMentRepo,
		CouponRepo:      couponRepo}
}
func (o *OrderUseCase) CheckoutOrder(ctx context.Context, userId uint, couponCode request.Coupon) (checkOut response.CartResp, err error) {

	checkOut, err = o.OrderRepository.CheckoutOrder(ctx, userId, couponCode.Coupon)
	if err != nil {
		return checkOut, err
	}
	// check coupon eligibility

	return checkOut, nil
}
func (o *OrderUseCase) PlaceOrderByCOD(ctx context.Context, userId, PaymentMethodID uint, couponCode string) (shopOrderId uint, err error) {
	// verify id is for COD or not
	payMethod, err := o.PayMentRepo.GetPaymentMethodByID(ctx, PaymentMethodID)
	if err != nil {
		return shopOrderId, err
	}
	if payMethod == "Cash on delivery COD" {
		shopOrderId, err = o.OrderRepository.PlaceCODOrder(ctx, userId, PaymentMethodID, couponCode)
		if err != nil {
			return shopOrderId, err
		}
		// Clear user cart
		if err := o.OrderRepository.ClearUserCart(ctx, userId); err != nil {
			return shopOrderId, err
		}
	} else {
		return shopOrderId, errors.New("requestID is not for Cash on delivery")
	}

	return shopOrderId, nil
}

func (o *OrderUseCase) RazorPayCheckout(ctx context.Context, razorpay request.RazorpayReq) (razorpayOrder response.RazorPayOrderResp, err error) {
	// Verify requested payment id is for Razorpay
	payMethod, err := o.PayMentRepo.GetPaymentMethodByID(ctx, razorpay.PaymentMethodId)
	if err != nil {
		return razorpayOrder, err
	}
	if payMethod != "Razorpay" {
		return razorpayOrder, errors.New("requestID is not for Razorpay")
	}

	// Find cart items
	checkOut, err := o.OrderRepository.CheckoutOrder(ctx, razorpay.UserID, razorpay.CouponCode)
	if err != nil {
		return razorpayOrder, err
	}
	if checkOut.TotalProductItems == 0 {
		return razorpayOrder, errors.New("cart is empty")
	}

	// Get user contact
	contact, err := o.UserRepository.GetEmailPhoneByUserId(ctx, razorpay.UserID)
	if err != nil {
		return razorpayOrder, err
	}
	razorpayOrder = response.RazorPayOrderResp{
		Email:          contact.Email,
		Phone:          contact.Phone,
		RazorpayAmount: checkOut.FinalPrice,
		AmountToPay:    checkOut.FinalPrice,
		UserID:         razorpay.UserID,
		RazorpayKey:    config.GetConfig().RazorPayKey,
		SGPay_id:       2,
	}

	// Generate Razorpay order ID
	razorPayOrderId, err := utils.GenerateRazorPayOrder(checkOut.TotalPrice, "Test receipt")
	if err != nil {
		return razorpayOrder, err
	}
	razorpayOrder.RazorpayOrderID = razorPayOrderId

	// Save order as pending
	orderId, err := o.OrderRepository.SaveOrder(ctx, domain.ShopOrder{
		UserID:           razorpay.UserID,
		OrderTotal:       checkOut.FinalPrice,
		ShippingID:       checkOut.DefaultShipping.ID,
		PaymentMethodID:  razorpay.PaymentMethodId,
		DeliveryStatusID: 1, // id 1 is for status "Pending"
		OrderStatusID:    1, // id 1 is for status "Pending"
		CouponID:         0,
	})
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
func (o *OrderUseCase) UpdateOrderStatus(c context.Context, UpdateData request.UpdateStatus) error {
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

func (o *OrderUseCase) ReturnEligibilityCheck(c context.Context, data request.ReturnRequest) error {
	// check if the payment success or not for procced order
	ok, err := o.PayMentRepo.GetPaymentStatusByOrderId(c, data.OrderID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("payment failed ! not eligible to proceed")
	}
	// check if the delivery date under 10 day for eligitibiliy
	deliveryDate, err := o.OrderRepository.GetDeliveryDate(c, data.OrderID)
	if err != nil {
		return err
	}
	if deliveryDate.Add(time.Hour * 24 * 10).Before(time.Now()) {

		return errors.New("return option allowed only till 10 days of delivery time")
	}
	// save return data and send for approval
	if err := o.OrderRepository.SaveReturnRequest(c, data); err != nil {
		return err
	}
	return nil

}

func (o *OrderUseCase) GetAllReturnRequest(c context.Context, page request.ReqPagination) (ReturnRequests []response.ReturnRequests, err error) {
	ReturnRequests, err = o.OrderRepository.GetAllReturnOrder(c, page)
	if err != nil {
		return ReturnRequests, err
	}
	return ReturnRequests, nil
}
