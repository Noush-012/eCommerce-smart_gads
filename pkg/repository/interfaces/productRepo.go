package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type ProductRepository interface {
	// Product CRUD section
	GetAllProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error)
	FindProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	SaveProduct(ctx context.Context, product domain.Product) error
	FindProductByID(ctx context.Context, productID uint) (product domain.Product, err error)
	UpdateProduct(ctx context.Context, product domain.Product) error
	DeleteProduct(ctx context.Context, productID uint) (domain.Product, error)
	GetStockStatusByProductId(c context.Context, productId uint) (qtyLeft uint, err error)

	UpdateProductItem(ctx context.Context, UpdateData request.UpdateProductItemReq) error

	AddProductItem(ctx context.Context, productItem request.ProductItemReq) error
	GetProductItems(ctx context.Context, productId uint) ([]response.ProductItemResp, error)

	// Brand CRUD section
	FindBrand(ctx context.Context, brand request.CategoryReq) (request.CategoryReq, error)
	AddCategory(ctx context.Context, brand request.CategoryReq) (err error)
	GetAllBrand(ctx context.Context) (brand []response.Brand, err error)
}
