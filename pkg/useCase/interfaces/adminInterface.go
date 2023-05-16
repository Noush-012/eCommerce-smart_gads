package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type AdminService interface {
	Signup(c context.Context, admin domain.Admin) error
	Login(c context.Context, admin domain.Admin) (domain.Admin, error)
	GetAllUser(c context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error)
	BlockUser(c context.Context, userID uint) error
	GetUserOrderHistory(c context.Context, userId uint) (orderHistory []domain.ShopOrder, err error)
}
