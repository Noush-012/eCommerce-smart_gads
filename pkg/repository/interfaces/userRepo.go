package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user domain.Users) error
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)

	SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error
}
