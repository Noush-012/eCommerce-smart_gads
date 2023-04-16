package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.Users) error
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
}
