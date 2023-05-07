package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type ProductService interface {
	AddProduct(ctx context.Context, product domain.Product) error
	AddCategory(ctx context.Context, Category request.CategoryReq) error
	GetAllBrands(ctx context.Context) (brand []response.Brand, err error)
	GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	UpdateProduct(ctx context.Context, product domain.Product) error
	DeleteProduct(ctx context.Context, productID uint) (domain.Product, error)

	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItem(ctx context.Context, productId uint) (ProductItems []response.ProductItemResp, err error)

	SKUhelper(ctx context.Context, productId uint) (interface{}, error)
}
