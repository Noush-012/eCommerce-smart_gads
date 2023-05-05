package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.Users) error
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
	OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error)
	Profile(ctx context.Context, userId uint) (domain.Users, error)
	Addaddress(ctx context.Context, address domain.Address) error

	SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error

	GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error
}
