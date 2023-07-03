package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type AuthRepository interface {
	SaveUser(ctx context.Context, user domain.Users) error
}
