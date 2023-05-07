package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderRepository interface {
	GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error)
	CheckoutOrder(ctx context.Context, userId uint) (checkOut response.CartResp, err error)
	GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	PlaceCODOrder(ctx context.Context, userId uint) (shopOrder response.ShopOrder, err error)
	ClearUserCart(ctx context.Context, userId uint) error

	// Order
	OrderStatus(ctx context.Context, id uint) (status string, err error)
}
