package repository

import (
	"context"
	"errors"
	"fmt"
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

func (p *productDatabase) GetProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products where id = ? product_name = ?`
	if p.DB.Raw(query, product.ID, product.Name).Scan(&product).Error != nil {
		return product, errors.New("failure to get product")
	}
	return product, nil
}

func (p *productDatabase) SaveProduct(ctx context.Context, product domain.Product) error {
	query := `INSERT INTO products (name, description, category_id, price, image, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	createdAt := time.Now()
	if p.DB.Exec(query, product.Name, product.Description, product.CategoryID, product.Price, product.Image, createdAt).Error != nil {
		return errors.New("failed to save product on database")
	}
	return nil
}

// find product by name
func (p *productDatabase) FindProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `SELECT * FROM products WHERE id = ? OR name=?`
	if p.DB.Raw(query, product.ID, product.Name).Scan(&product).Error != nil {
		return product, errors.New("faild to get product")
	}
	return product, nil
}

// find product by id
func (p *productDatabase) FindProductByID(ctx context.Context, productID uint) (product domain.Product, err error) {
	query := `SELECT * FROM products WHERE id = $1`
	err = p.DB.Raw(query, productID).Scan(&product).Error
	if err != nil {
		return product, fmt.Errorf("failed find product with prduct_id %v", productID)
	}
	return product, nil
}

// ================ Brand CRUD ================ //

func (p *productDatabase) FindBrand(ctx context.Context, brand domain.Category) (domain.Category, error) {
	query := `SELECT * FROM brands WHERE id = ? OR brand_name=?`

	if p.DB.Raw(query).Scan(&brand).Error != nil {
		return brand, errors.New("failed to get brand")

	}
	return brand, nil
}

// To add brand
func (p *productDatabase) SaveBrand(ctx context.Context, brand domain.Category) (err error) {
	if brand.ID == 0 {
		query := `INSERT INTO brands (brand_name) VALUES ($1)`
		err = p.DB.Exec(query).Error
	} else {
		query := `INSERT INTO brands (brand_id, brand_name)VALUES($1,$2)`
		err = p.DB.Exec(query).Error
	}
	if err != nil {
		return errors.New("failed to save brand")
	}
	return nil
}

func (p *productDatabase) GetAllBrand(ctx context.Context) (brand []response.Brand, err error) {
	// get all brands from database
	query := `SELECT c.id, c.category_name FROM categories c OFFSET 1`
	if p.DB.Raw(query).Scan(&brand).Error != nil {
		return brand, fmt.Errorf("failed to get brands data from db")
	}
	return brand, nil
}

// get all products from database
func (p *productDatabase) GetAllProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	// aliase :: p := product; c := category
	query := `SELECT p.id, p.name, p.description, c.category_name, p.price, p.discount_price,
	 p.image,  p.created_at, p.updated_at  
	FROM products p LEFT JOIN categories c ON p.category_id = c.id 
	ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	if p.DB.Raw(query, limit, offset).Scan(&products).Error != nil {
		return products, errors.New("failed to get products from database")
	}

	return products, nil
}

// update product
func (p *productDatabase) UpdateProduct(ctx context.Context, product domain.Product) error {
	existingProduct, err := p.FindProductByID(ctx, product.ID)
	if err != nil {
		return err
	}
	if product.Name == "" {
		product.Name = existingProduct.Name
	}
	if product.Description == "" {
		product.Description = existingProduct.Description
	}
	if product.Price == 0 {
		product.Price = existingProduct.Price
	}
	if product.Image == "" {
		product.Image = existingProduct.Image
	}
	if product.CategoryID == 0 {
		product.CategoryID = existingProduct.CategoryID
	}
	query := `UPDATE products SET name = $1, description = $2, category_id = $3,
	price = $4, image = $5, updated_at = $6 WHERE id = $7`

	updatedAt := time.Now()

	if p.DB.Exec(query, product.Name, product.Description, product.CategoryID,
		product.Price, product.Image, updatedAt, product.ID).Error != nil {
		return errors.New("failed to update product")
	}

	return nil
}

func (p *productDatabase) DeleteProduct(ctx context.Context, productID uint) (domain.Product, error) {
	// Check requested product is exist or not
	var existingProduct domain.Product
	existingProduct, err := p.FindProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	} else if existingProduct.Name == "" {
		return domain.Product{}, errors.New("invalid product_id")
	}

	//delete query
	query := `DELETE FROM products WHERE id = $1`
	if p.DB.Exec(query, productID).Error != nil {
		return domain.Product{}, fmt.Errorf("failed to delete error : %v", err)
	}
	return existingProduct, nil
}
