package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type ProductRepository interface {
	// Product CRUD section
	// func GetAllProducts(ctx context.Context, )([]resp.ResponseProduct, error)
	FindProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	SaveProduct(ctx context.Context, product domain.Product) error

	// Brand CRUD section
	FindBrand(ctx context.Context, brand domain.Brand) (domain.Brand, error)
	SaveBrand(ctx context.Context, brand domain.Brand) (err error)
}
