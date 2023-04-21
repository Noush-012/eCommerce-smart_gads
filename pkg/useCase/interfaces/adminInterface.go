package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type AdminService interface {
	Signup(c context.Context, admin domain.Admin) error
	Login(c context.Context, admin domain.Admin) (domain.Admin, error)
}
