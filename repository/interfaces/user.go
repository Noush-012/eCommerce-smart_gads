package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	SaveUser(ctx context.Context, user domain.Users) error
	Delete(ctx context.Context, user domain.Users) error
	FindByEmail(ctx context.Context, id string) bool
}
