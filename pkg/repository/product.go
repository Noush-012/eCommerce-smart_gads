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

func (p *productDatabase) FindBrand(ctx context.Context, brand request.CategoryReq) (request.CategoryReq, error) {
	query := `SELECT * FROM categories WHERE id = ? OR category_name=?`

	if p.DB.Raw(query).Scan(&brand).Error != nil {
		return brand, errors.New("failed to get brand")

	}
	return brand, nil
}

// To add brand
func (p *productDatabase) AddCategory(ctx context.Context, brand request.CategoryReq) (err error) {
	if brand.ParentID == 0 {
		query := `INSERT INTO categories (category_name) VALUES ($1)`
		err = p.DB.Exec(query).Error
	} else {
		query := `INSERT INTO categories (parent_id, category_name)VALUES($1,$2)`
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
	if err := p.DB.Exec(query, productID).Error; err != nil {
		return domain.Product{}, fmt.Errorf("failed to delete error : %v", err)
	}
	return existingProduct, nil
}

// to add product item of a existing product
func (p *productDatabase) AddProductItem(ctx context.Context, productItem request.ProductItemReq) error {
	tnx := p.DB.Begin()

	var Prod_Item_ID uint
	// to check whether product and its variant already exists.
	existingProduct, err := p.FindProductByID(ctx, productItem.ProductID)
	if err != nil {
		return err
	}
	if existingProduct.ID != productItem.ProductID {
		tnx.Rollback()
		return errors.New("product not exists belongs to requested product item")
	}

	query := `SELECT DISTINCT pi.id AS product_item_id FROM product_items pi INNER JOIN product_configs pc on pi.id = pc.product_item_id
	WHERE pi.id = $1 AND pc.variation_option_id = $2`

	for _, vOptionID := range productItem.Configurations {
		if vOptionID.VariationOptionID == 0 {
			tnx.Rollback()
			return errors.New("invalid or missing product configuration")
		}
		if err := tnx.Raw(query, productItem.ProductID, vOptionID.VariationOptionID).Scan(&Prod_Item_ID).Error; err != nil {
			tnx.Rollback()
			return fmt.Errorf("%v product item already exist in with VariationOptionID: %v", err, vOptionID.VariationOptionID)
		}

	}
	// to save product item to database
	query = `INSERT INTO product_items (product_id,qty_in_stock,price,discount_price,sku,created_at) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	createdAt := time.Now()
	if err := tnx.Raw(query, productItem.ProductID, productItem.QtyInStock, productItem.Price,
		productItem.DiscountPrice, productItem.SKU, createdAt).Scan(&productItem.ProductItemId).Error; err != nil {
		tnx.Rollback()
		return fmt.Errorf("failed to add product item %v", err)
	}

	// to save product configuration
	// query = `INSERT INTO product_configs (product_item_id,variation_option_id) VALUES ($1, $2)`
	query = `INSERT INTO product_configs (product_item_id, variation_option_id)
	VALUES ($1, $2)`
	for _, vOption := range productItem.Configurations {
		if err := tnx.Exec(query, productItem.ProductItemId, vOption.VariationOptionID).Error; err != nil {
			tnx.Rollback()
			return fmt.Errorf("failed to add product configuration %v, VariationOptionID : %v", err, vOption.VariationOptionID)
		}
	}

	// to save images seperatly for each color variant

	query = `INSERT INTO product_images  (product_item_id,image) 
	VALUES ($1, $2)`
	for _, image := range productItem.Images {
		if err := tnx.Exec(query, productItem.ProductItemId, image).Error; err != nil {
			tnx.Rollback()
			return fmt.Errorf("failed to add product images %v", err)
		}
	}
	if err := tnx.Commit().Error; err != nil {
		tnx.Rollback()
		return fmt.Errorf("failed to commit the transaction %v", err)
	}
	return nil
}

// to list product item
func (p *productDatabase) GetProductItems(ctx context.Context, productId uint) ([]response.ProductItemResp, error) {
	// Check product id exist or not
	var ProductItems []response.ProductItemResp
	var Prod_Item_ID uint

	dbProd, err := p.FindProductByID(ctx, productId)
	if err != nil {
		return ProductItems, err
	}
	if dbProd.ID == 0 {
		return ProductItems, errors.New("invalid product id")
	}
	// to fetch product item id
	query := `SELECT id FROM product_items WHERE product_id = $1`
	if err := p.DB.Raw(query, productId).Scan(&Prod_Item_ID).Error; err != nil {
		return ProductItems, fmt.Errorf("failed to get product item id %v", err)
	}

	// get product items from database
	query = `SELECT 
    p.id AS product_id,
    pi.id AS product_item_id,
    pi.qty_in_stock AS stock_available,
    p.name AS product_name, 
    c.category_name AS brand,
    p.description,
    vo1.option_value AS color,
    vo2.option_value AS storage,
    p.price,
    pi.discount_price AS offer_price 
FROM 
    products p 
    JOIN categories c ON c.id = p.category_id 
    JOIN product_items pi ON pi.product_id = p.id 
    JOIN product_configs pc1 ON pi.id = pc1.product_item_id AND pc1.variation_option_id IN (SELECT id FROM variation_options WHERE variation_id = 1)
    JOIN variation_options vo1 ON vo1.id = pc1.variation_option_id 
    JOIN product_configs pc2 ON pi.id = pc2.product_item_id AND pc2.variation_option_id IN (SELECT id FROM variation_options WHERE variation_id = 2)
    JOIN variation_options vo2 ON vo2.id = pc2.variation_option_id
WHERE 
    p.id = $1;`
	if err := p.DB.Raw(query, productId).Scan(&ProductItems).Error; err != nil {
		return ProductItems, fmt.Errorf("failed to get product items %v", err)
	}

	// fetch product item images
	query = `SELECT
	pimg.image
	FROM
	product_images pimg
	WHERE product_item_id = $1`
	for i := range ProductItems {
		ProductItems[i].Images = []string{}
		p.DB.Raw(query, Prod_Item_ID).Scan(&ProductItems[i].Images)
	}
	// // append images to product items
	// for i, _ := range ProductItems {
	// 	ProductItems[i].Images = append(ProductItems[i].Images, dbProd.Image)
	// }
	return ProductItems, nil
}
