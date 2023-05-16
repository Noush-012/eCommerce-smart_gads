package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderRepository interface {
	GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error)
	CheckoutOrder(ctx context.Context, userId uint) (checkOut response.CartResp, err error)
	GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	PlaceCODOrder(ctx context.Context, userId uint, PaymentMethodID uint) (OrderId uint, err error)
	ClearUserCart(ctx context.Context, userId uint) error

	// Order
	SaveOrder(ctx context.Context, order domain.ShopOrder) (OrderId uint, err error)
	SaveOrderLine(ctx context.Context, orderLine domain.OrderLine) error
	OrderStatus(ctx context.Context, id uint) (status string, err error)
	GetOrderByOrderId(ctx context.Context, OrderId uint) (orderData response.ShopOrder, err error)
	GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error)

	ChangeOrderStatus(c context.Context, UpdateData request.UpdateOrderStatus) error
}
