package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB: db}
}

// func GetAllProducts(ctx context.Context, page req.ReqPagination) ([]resp.ResponseProduct, error) {
// 	limit := page.Count
// 	offset := page.PageNumber

// 	// Alias : p = product, c = category
// 	query := `SELECT p.id,`
// }

func (p *productDatabase) GetProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products where id = ? product_name = ?`
	if p.DB.Raw(query, product.ID, product.ProductName).Scan(&product).Error != nil {
		return product, errors.New("failure to get product")
	}
	return product, nil
}

func (p *productDatabase) SaveProduct(ctx context.Context, product domain.Product) error {
	query := `INSERT INTO products (product_name, description, brand_id, price, image, created_at)
VALUES($1, $2, $3, $4, $5, $6)`
	createdAt := time.Now()
	if p.DB.Exec(query, product.ProductName, product.Description,
		product.BrandID, product.Price, product.Image, createdAt).Error != nil {
		return errors.New("failed to save product on database")
	}
	return nil
}

// find product by name
func (p *productDatabase) FindProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products WHERE id = ? OR product_name=?`
	if p.DB.Raw(query, product.ID, product.ProductName).Scan(&product).Error != nil {
		return product, errors.New("faild to get product")
	}
	return product, nil
}

// ================ Brand CRUD ================ //

func (p *productDatabase) FindBrand(ctx context.Context, brand domain.Brand) (domain.Brand, error) {
	query := `SELECT * FROM brands WHERE id = ? OR brand_name=?`

	if p.DB.Raw(query, brand.ID, brand.BrandName).Scan(&brand).Error != nil {
		return brand, errors.New("failed to get brand")

	}
	return brand, nil
}

// To add brand
func (p *productDatabase) SaveBrand(ctx context.Context, brand domain.Brand) (err error) {
	if brand.BrandID == 0 {
		query := `INSERT INTO brands (brand_name) VALUES ($1)`
		err = p.DB.Exec(query, brand.BrandName).Error
	} else {
		query := `INSERT INTO brands (brand_id, brand_name)VALUES($1,$2)`
		err = p.DB.Exec(query, brand.BrandID, brand.BrandName).Error
	}
	if err != nil {
		return errors.New("failed to save brand")
	}
	return nil
}
func (p *productDatabase) GetAllBrand(ctx context.Context) ([]response.Brand, error) {
	// var brands []response.Brand
	// query := `SELECT `

	return nil, nil
}

// get all products from database
func (c *productDatabase) GetAllProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	// aliase :: p := product; c := category
	querry := `SELECT p.id, p.product_name, p.description, p.price, p.offer_price, p.image, p.brand_id, 
	p.image, c.brand_name, p.created_at, p.updated_at  
	FROM products p LEFT JOIN brands c ON p.brand_id=c.id 
	ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	if c.DB.Raw(querry, limit, offset).Scan(&products).Error != nil {
		return products, errors.New("failed to get products from database")
	}

	return products, nil
}
