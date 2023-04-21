package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type ProductService interface {
	AddProduct(ctx context.Context, product domain.Product) error
	AddBrand(ctx context.Context, brand domain.Brand) error
}
