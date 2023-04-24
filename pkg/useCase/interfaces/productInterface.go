package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type ProductService interface {
	AddProduct(ctx context.Context, product domain.Product) error
	AddBrand(ctx context.Context, brand domain.Brand) error
	GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	UpdateProduct(ctx context.Context, product domain.Product) error
}
