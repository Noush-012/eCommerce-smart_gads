package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderService interface {
	GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error)

	// Checkouts
	RazorPayCheckout(ctx context.Context, razorpay request.RazorpayReq) (razorpayOrder response.RazorPayOrderResp, err error)
	CheckoutOrder(ctx context.Context, userId uint, couponCode request.Coupon) (checkOut response.CartResp, err error)
	PlaceOrderByCOD(ctx context.Context, userId, PaymentMethodID uint, couponCode string) (shopOrderId uint, err error)

	UpdateOrderStatus(c context.Context, UpdateData request.UpdateStatus) error
	ReturnEligibilityCheck(c context.Context, data request.ReturnRequest) error
	GetAllReturnRequest(c context.Context, page request.ReqPagination) (ReturnRequests []response.ReturnRequests, err error)
}
