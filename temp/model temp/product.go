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
