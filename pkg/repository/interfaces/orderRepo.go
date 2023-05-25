package interfaces

import (
	"context"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderRepository interface {
	GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error)
	CheckoutOrder(ctx context.Context, userId uint, couponCode string) (checkOut response.CartResp, err error)
	GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	PlaceCODOrder(ctx context.Context, userId, PaymentMethodID uint, couponCode string) (OrderId uint, err error)
	ClearUserCart(ctx context.Context, userId uint) error

	// Order
	SaveOrder(ctx context.Context, order domain.ShopOrder) (OrderId uint, err error)
	SaveOrderLine(ctx context.Context, orderLine domain.OrderLine) error
	OrderStatus(ctx context.Context, id uint) (status string, err error)
	GetOrderByOrderId(ctx context.Context, OrderId uint) (orderData response.ShopOrder, err error)
	GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error)

	ChangeOrderStatus(c context.Context, UpdateData request.UpdateStatus) error
	UpdateDeliveryStatus(c context.Context, UpdateData request.UpdateStatus) error
	GetDeliveryDate(c context.Context, orderId uint) (time.Time, error)
	SaveReturnRequest(c context.Context, data request.ReturnRequest) error
	GetAllReturnOrder(c context.Context, page request.ReqPagination) (ReturnRequests []response.ReturnRequests, err error)
}
