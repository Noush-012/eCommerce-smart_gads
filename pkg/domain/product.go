package domain

import "time"

// Category struct
type Category struct {
	ID           uint        `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	ParentID     *uint       `json:"parent_id" gorm:"default:1;index"`
	Parent       *Category   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryName string      `json:"category_name" gorm:"unique;not null"`
	Products     []*Product  `json:"-" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Children     []*Category `json:"-" gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Product struct
type Product struct {
	ID            uint           `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Name          string         `json:"product_name" gorm:"not null;size:50"`
	Description   string         `json:"description" gorm:"not null;size:500"`
	CategoryID    uint           `json:"brand_id" gorm:"index;not null"`
	Category      *Category      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price         uint           `json:"price" gorm:"not null"`
	DiscountPrice uint           `json:"discount_price" gorm:"default:null"`
	Image         string         `json:"image" gorm:"not null"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"default:null"`
	Items         []*ProductItem `json:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// ProductItem struct
type ProductItem struct {
	ID             uint             `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	ProductID      uint             `json:"product_id" gorm:"index;not null"`
	QtyInStock     uint             `json:"qty_in_stock" gorm:"not null"`
	Price          uint             `json:"price" gorm:"not null"`
	SKU            string           `json:"sku" gorm:"unique;not null"`
	DiscountPrice  uint             `json:"discount_price" gorm:"default:null"`
	CreatedAt      time.Time        `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time        `json:"updated_at" gorm:"default:null"`
	Configurations []*ProductConfig `json:"-" gorm:"many2many:product_configurations;"`
	Images         []*ProductImage  `json:"-" gorm:"foreignKey:ProductItemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Variation struct
type Variation struct {
	ID            uint               `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	CategoryID    uint               `json:"category_id" gorm:"index;not null"`
	Category      *Category          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VariationName string             `json:"variation_name" gorm:"not null"`
	Options       []*VariationOption `json:"-" gorm:"foreignKey:VariationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// VariationOption struct
type VariationOption struct {
	ID             uint             `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	VariationID    uint             `json:"variation_id" gorm:"index;not null"`
	Variation      *Variation       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OptionValue    string           `json:"option_value" gorm:"not null"`
	Configurations []*ProductConfig `json:"-" gorm:"many2many:product_configurations;"`
}

// ProductConfig struct
type ProductConfig struct {
	ProductItemID     uint             `json:"product_id" gorm:"primaryKey;not null;autoIncrement"`
	VariationOptionID uint             `json:"variation_option_id" gorm:"primaryKey;not null"`
	ProductItem       *ProductItem     `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VariationOption   *VariationOption `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductImage struct {
	ID            uint        `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	ProductItemID uint        `json:"product_item_id" gorm:"not null"`
	ProductItem   ProductItem `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Image         string      `json:"image" gorm:"not null"`
}

// // Product category model
// type ProductCategory struct {
// 	ID           uint   `json:"id" gorm:"primaryKey;not null"`
// 	CategoryName string `json:"category_name" gorm:"not null" binding:"required,min=3,max=20"`
// }

// // Model of category / Brands
// type Brand struct {
// 	ID        uint   `json:"id" gorm:"primaryKey;not null"`
// 	BrandID   uint   `json:"brand_id"`
// 	Brand     *Brand `json:"-"`
// 	BrandName string `json:"brand_name" gorm:"unique;not null" binding:"required,min=1,max=30"`
// }

// // Product model
// type Product struct {
// 	ID            uint      `json:"id" gorm:"primaryKey;not null"`
// 	ProductName   string    `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
// 	Description   string    `json:"description" gorm:"not null" binding:"required,min=10,max=1000"`
// 	BrandID       uint      `json:"brand_id" binding:"omitempty,numeric"`
// 	Brand         Brand     `json:"-"`
// 	Price         uint      `json:"price" gorm:"not null" binding:"required,numeric"`
// 	DiscountPrice uint      `json:"discount_price"`
// 	Image         string    `json:"image" gorm:"not null"`
// 	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
// 	UpdatedAt     time.Time `json:"updated_at"`
// }

// type ProductItem struct {
// 	ID        uint `json:"id" gorm:"primaryKey;not null"`
// 	ProductID uint `json:"product_id" gorm:"not null" binding:"required,numeric"`
// 	Product   Product
// 	//images are stored in sperate table along with productItem Id
// 	QtyInStock    uint      `json:"qty_in_stock" gorm:"not null" binding:"required,numeric"` // no need of stockAvailble column , because from this qty we can get it
// 	Price         uint      `json:"price" gorm:"not null" binding:"required,numeric"`
// 	SKU           string    `json:"sku" gorm:"unique;not null"`
// 	DiscountPrice uint      `json:"discount_price"`
// 	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
// 	UpdatedAt     time.Time `json:"updated_at"`
// }

// // variation means color etc..
// type Variation struct {
// 	ID            uint   `json:"-" gorm:"primaryKey;not null"`
// 	BrandID       uint   `json:"category_id" gorm:"not null" binding:"required,numeric"`
// 	Brand         Brand  `json:"-"`
// 	VariationName string `json:"variation_name" gorm:"not null" binding:"required"`
// }

// type VariationOption struct {
// 	ID             uint      `json:"-" gorm:"primaryKey;not null"`
// 	VariationID    uint      `json:"variation_id" gorm:"not null" binding:"required,numeric"` // a specific field of variation like color/size
// 	Variation      Variation `json:"-"`
// 	VariationValue string    `json:"variation_value" gorm:"not null" binding:"required"` // the variations value like blue/XL
// }

// type ProductConfiguration struct {
// 	ProductItemID     uint            `json:"product_item_id" gorm:"not null"`
// 	ProductItem       ProductItem     `json:"-"`
// 	VariationOptionID uint            `json:"variation_option_id" gorm:"not null"`
// 	VariationOption   VariationOption `json:"-"`
// }

// type ProductImage struct {
// 	ID            uint        `json:"id" gorm:"primaryKey;not null"`
// 	ProductItemID uint        `jsong:"product_item_id" gorm:"not null"`
// 	ProductItem   ProductItem `json:"-"`
// 	Image         string      `json:"image" gorm:"not null"`
// }

// // Model of specify varient of product
// type ProductVarient struct {
// 	ID         uint `json:"id" gorm:"primaryKey;not null"`
// 	ProductID  uint `json:"product_id" gorm:"not null" binding:"required,numeric"`
// 	Product    Product
// 	QtyInStock uint      `json:"qty_in_stock" gorm:"not null" binding:"required,numeric"`
// 	Price      uint      `json:"price" gorm:"not null" binding:"required,numeric"`
// 	OfferPrice uint      `json:"offer_price"`
// 	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// }
