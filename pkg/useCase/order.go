package usecase

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/config"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderUseCase struct {
	OrderRepository interfaces.OrderRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository) service.OrderService {
	return &OrderUseCase{OrderRepository: repo}
}
func (o *OrderUseCase) CheckoutOrder(ctx context.Context, userId uint) (checkOut response.CartResp, err error) {
	checkOut, err = o.OrderRepository.CheckoutOrder(ctx, userId)
	if err != nil {
		return checkOut, err
	}
	return checkOut, nil
}
func (o *OrderUseCase) PlaceOrderByCOD(ctx context.Context, userId uint) (shopOrder response.ShopOrder, err error) {
	shopOrder, err = o.OrderRepository.PlaceCODOrder(ctx, userId)
	if err != nil {
		return shopOrder, err
	}
	// Clear user cart
	if err := o.OrderRepository.ClearUserCart(ctx, userId); err != nil {
		return shopOrder, err
	}
	return shopOrder, nil
}

func (o *OrderUseCase) RazorPayCheckout(ctx context.Context, razorpay request.RazorpayReq) (razorpayOrder response.RazorPayOrderResp, err error) {
	// find cart items
	CheckOut, err := o.OrderRepository.CheckoutOrder(ctx, razorpay.UserID)
	if err != nil {
		return razorpayOrder, err
	}
	// get user contact

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

	return razorpayOrder, nil

}

func (o *OrderUseCase) GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error) {
	orderHistory, err := o.OrderRepository.GetOrderHistory(ctx, page, userId)
	if err != nil {
		return orderHistory, err
	}
	return orderHistory, nil

}
