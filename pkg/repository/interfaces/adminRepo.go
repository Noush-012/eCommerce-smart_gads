package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type AdminRepository interface {
	GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	SaveAdmin(ctx context.Context, admin domain.Admin) error
	GetAllUser(ctx context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error)
	BlockUser(ctx context.Context, userID uint) error
	GetUserOrderHistory(c context.Context, userId uint) (orderHistory []domain.ShopOrder, err error)
	ChangeOrderStatus(c context.Context, UpdateData request.UpdateOrderStatus) (order response.ShopOrder, err error)
}
