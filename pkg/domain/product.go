package domain

import "time"

// Product category model
type ProductCategory struct {
	ID           uint   `json:"id" gorm:"primaryKey;not null"`
	CategoryName string `json:"category_name" gorm:"not null" binding:"required,min=3,max=20"`
}

// Model of category / Brands
type Brand struct {
	ID        uint   `json:"id" gorm:"primaryKey;not null"`
	BrandID   uint   `json:"brand_id"`
	Brand     *Brand `json:"-"`
	BrandName string `json:"brand_name" gorm:"unique;not null" binding:"required,min=1,max=30"`
}

// Product model
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;not null"`
	ProductName string    `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Description string    `json:"description" gorm:"not null" binding:"required,min=10,max=100"`
	BrandID     uint      `json:"brand_id" binding:"omitempty,numeric"`
	Brand       Brand     `json:"-"`
	Price       uint      `json:"price" gorm:"not null" binding:"required,numeric"`
	OfferPrice  uint      `json:"discount_price"`
	Image       string    `json:"image" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Model of specify varient of product
type ProductVarient struct {
	ID         uint `json:"id" gorm:"primaryKey;not null"`
	ProductID  uint `json:"product_id" gorm:"not null" binding:"required,numeric"`
	Product    Product
	QtyInStock uint      `json:"qty_in_stock" gorm:"not null" binding:"required,numeric"`
	Price      uint      `json:"price" gorm:"not null" binding:"required,numeric"`
	OfferPrice uint      `json:"offer_price"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//
