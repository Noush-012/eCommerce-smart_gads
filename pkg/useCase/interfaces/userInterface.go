package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type UserUseCase interface {
	SignUp(ctx context.Context, user domain.Users) error
}
