package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/req"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/resp"
)

type productUseCase interface {
	// products
	GetProducts(ctx context.Context, pagination req.ReqPagination) (products []resp.ResponseProduct, err error)
}
