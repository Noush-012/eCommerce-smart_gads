package model

import "time"

type Product struct {
	ID            uint           `gorm:"primaryKey;not null"`
	Name          string         `gorm:"not null;size:50"`
	Description   string         `gorm:"not null;size:100"`
	CategoryID    uint           `gorm:"index;not null"`
	Category      *Category      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price         uint           `gorm:"not null"`
	DiscountPrice uint           `gorm:"default:null"`
	Image         string         `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"default:null"`
	Items         []*ProductItem `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// ProductItem struct
type ProductItem struct {
	ID             uint             `gorm:"primaryKey;not null"`
	ProductID      uint             `gorm:"index;not null"`
	QtyInStock     uint             `gorm:"not null"`
	Price          uint             `gorm:"not null"`
	SKU            string           `gorm:"unique;not null"`
	DiscountPrice  uint             `gorm:"default:null"`
	CreatedAt      time.Time        `gorm:"not null"`
	UpdatedAt      time.Time        `gorm:"default:null"`
	Configurations []*ProductConfig `gorm:"many2many:product_configurations;"`
	Images         []*ProductImage  `gorm:"foreignKey:ProductItemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Category struct
type Category struct {
	ID           uint        `gorm:"primaryKey;not null"`
	ParentID     *uint       `gorm:"index"`
	Parent       *Category   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryName string      `gorm:"unique;not null"`
	Products     []*Product  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Children     []*Category `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Variation struct
type Variation struct {
	ID            uint               `gorm:"primaryKey;not null"`
	CategoryID    uint               `gorm:"index;not null"`
	Category      *Category          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VariationName string             `gorm:"not null"`
	Options       []*VariationOption `gorm:"foreignKey:VariationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// VariationOption struct
type VariationOption struct {
	ID             uint             `gorm:"primaryKey;not null"`
	VariationID    uint             `gorm:"index;not null"`
	Variation      *Variation       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OptionValue    string           `gorm:"not null"`
	Configurations []*ProductConfig `gorm:"many2many:product_configurations;"`
}

// ProductConfig struct
type ProductConfig struct {
	ProductItemID     uint             `gorm:"primaryKey;not null"`
	VariationOptionID uint             `gorm:"primaryKey;not null"`
	ProductItem       *ProductItem     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VariationOption   *VariationOption `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductImage struct {
	ID            uint        `json:"id" gorm:"primaryKey;not null"`
	ProductItemID uint        `json:"product_item_id" gorm:"not null"`
	ProductItem   ProductItem `json:"-"`
	Image         string      `json:"image" gorm:"not null"`
}

// to store a review of a product
type Review struct {
	ID            uint        `json:"id" gorm:"primaryKey;not null"`
	ProductItem   ProductItem `json:"-" gorm:"foreignKey:ProductItemID;references:ID"`
	ProductItemID uint        `json:"product_item_id" gorm:"not null"`
	UserID        uint        `json:"user_id" gorm:"not null"`
	Rating        float32     `json:"rating" gorm:"not null"`
	Comment       string      `json:"comment"`
	CreatedAt     time.Time   `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// to store an order of a user with multiple order items
type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey;not null"`
	UserID      uint        `json:"user_id" gorm:"not null"`
	TotalAmount uint        `json:"total_amount" gorm:"not null"`
	Status      string      `json:"status" gorm:"not null"`
	OrderItems  []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time   `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// to store an item of an order
type OrderItem struct {
	ID             uint        `json:"id" gorm:"primaryKey;not null"`
	OrderID        uint        `json:"order_id" gorm:"not null"`
	Order          Order       `json:"-"`
	ProductItem    ProductItem `json:"product_item" gorm:"foreignKey:ProductItemID;references:ID"`
	ProductItemID  uint        `json:"product_item_id" gorm:"not null"`
	Qty            uint        `json:"qty" gorm:"not null"`
	Price          uint        `json:"price" gorm:"not null"`
	DiscountAmount uint        `json:"discount_amount"`
	CreatedAt      time.Time   `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

// User represents a user in the system
type User struct {
	ID           uint       `json:"id" gorm:"primaryKey;not null"`
	FirstName    string     `json:"first_name" gorm:"not null" binding:"required,min=1,max=50"`
	LastName     string     `json:"last_name" gorm:"not null" binding:"required,min=1,max=50"`
	Email        string     `json:"email" gorm:"unique;not null" binding:"required,email"`
	PasswordHash string     `json:"-" gorm:"not null"`
	CreatedAt    time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Roles        []UserRole `json:"roles" gorm:"many2many:user_roles;"`
}

// UserRole represents a role that can be assigned to a user
type UserRole struct {
	ID          uint   `json:"id" gorm:"primaryKey;not null"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	Users       []User `json:"users" gorm:"many2many:user_roles;"`
}

// to add product item of a existing product
// func (p *productDatabase) AddProductItem(ctx context.Context, productItem request.ProductItemReq) error {

// 	tnx := p.DB.Begin()

// 	var Prod_Item_ID uint

// 	// To check whether requesting product exist or not
// 	existingProduct, err := p.FindProductByID(ctx, productItem.ProductID)
// 	if err != nil {
// 		return err
// 	}
// 	if existingProduct.ID != productItem.ProductID {
// 		tnx.Rollback()
// 		return errors.New("product not exists belongs to requested product item")
// 	}

// 	query := `SELECT DISTINCT pi.id AS product_item_id FROM product_items pi INNER JOIN product_configs pc ON pi.id = pc.product_item_id
// 	WHERE pi.product_id = $1 AND pc.variation_option_id = $2`
// 	if err := tnx.Raw(query, productItem.ProductID, productItem.VariationOptionID).Scan(&Prod_Item_ID).Error; err != nil {
// 		tnx.Rollback()
// 		return fmt.Errorf("product item already exist in with given product configuration %v", err)
// 	}
// 	fmt.Println(Prod_Item_ID)
// 	query = `INSERT INTO product_items (product_id, qty_in_stock, price, sku, created_at) VALUES ($1, $2, $3, $4, $5)`
// 	createdAt := time.Now()
// 	productItem.Price = existingProduct.Price
// 	if err := p.DB.Exec(query, productItem.ProductID, productItem.QtyInStock, productItem.Price, productItem.SKU, createdAt).Error; err != nil {
// 		return fmt.Errorf("failed to add product item %v", err)
// 	}
// 	query = `INSERT INTO product_images (product_item_id, image) VALUES ($1 ,$2)`

// 	for _, img := range productItem.Images {
// 		err := tnx.Exec(query, Prod_Item_ID, img)
// 		if err != nil {
// 			tnx.Rollback()
// 			return fmt.Errorf("failed to add image for product item of product : %v", productItem.ProductID)
// 		}

// 	}
// 	query = `INSERT INTO product_configs (product_item_id, variation_option_id) VALUES ($1, $2)`
// 	err = tnx.Exec(query, Prod_Item_ID, productItem.VariationOptionID).Error
// 	if err != nil {
// 		tnx.Rollback()
// 		return fmt.Errorf("failed to save the product item for product with product_id %v", productItem.ProductID)
// 	}
// 	err = tnx.Commit().Error
// 	if err != nil {
// 		tnx.Rollback()
// 		return fmt.Errorf("failed to commit the transaction %v", err)
// 	}

// 	// query = `INSERT INTO product_configs (product_item_id, variation_option_id) VALUES ($1, $2)`
// 	// for _, varOpt := range productItem.VariationOptions {
// 	// 	err := tnx.Exec(query, Prod_Item_ID, varOpt)
// 	// 	if err != nil {
// 	// 		tnx.Rollback()
// 	// 		return fmt.Errorf("failed to add variation option for product item of product : %v", productItem.ProductID)
// 	// 	}

// 	return nil
// }

var a = `SELECT 
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
    JOIN variations v1 ON v1.category_id = c.parent_id AND v1.id = 1 
    JOIN product_configs pc1 ON pi.id = pc1.product_item_id 
    JOIN variation_options vo1 ON vo1.variation_id = v1.id AND vo1.id = pc1.variation_option_id 
    JOIN variations v2 ON v2.category_id = c.parent_id AND v2.id = 2 
    JOIN product_configs pc2 ON pi.id = pc2.product_item_id 
    JOIN variation_options vo2 ON vo2.variation_id = v2.id AND vo2.id = pc2.variation_option_id
WHERE 
    p.id = $1
    AND pc1.variation_option_id IN (SELECT id FROM variation_options WHERE variation_id = 1)
    AND pc2.variation_option_id IN (SELECT id FROM variation_options WHERE variation_id = 2);`
