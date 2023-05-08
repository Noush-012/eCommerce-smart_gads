package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user domain.Users) error
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)
	GetUserbyID(ctx context.Context, userId uint) (domain.Users, error)
	SaveAddress(ctx context.Context, userAddress domain.Address) error
	GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error)

	SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error
	GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error)

	GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error
}
