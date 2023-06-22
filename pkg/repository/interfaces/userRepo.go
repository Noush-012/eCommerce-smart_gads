package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type UserRepository interface {
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)
	GetUserbyID(ctx context.Context, userId uint) (domain.Users, error)
	SaveAddress(ctx context.Context, userAddress request.Address) error
	UpdateAddress(ctx context.Context, userAddress request.AddressPatchReq) error
	DeleteAddress(ctx context.Context, userID, addressID uint) error
	GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error)
	GetEmailPhoneByUserId(ctx context.Context, userID uint) (contact response.UserContact, err error)
	GetDefaultAddress(ctx context.Context, userId uint) (address response.Address, err error)

	SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error
	GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error)

	GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error

	AddToWishlist(ctx context.Context, wishlistData request.AddToWishlist) error
	GetWishlist(ctx context.Context, userId uint) (wishlist []response.Wishlist, err error)
	DeleteFromWishlist(ctx context.Context, productId, userId uint) error

	GetWalletHistory(ctx context.Context, userId uint) (wallet []domain.Wallet, err error)
	CreditUserWallet(ctx context.Context, data domain.Wallet) error
}
