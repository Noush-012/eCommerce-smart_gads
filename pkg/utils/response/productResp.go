package response

import "time"

type ResponseProduct struct {
	ID            uint      `json:"id"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description" `
	BrandID       uint      `json:"brand_id"`
	BrandName     string    `json:"brand_name"`
	Price         uint      `json:"price"`
	DiscountPrice uint      `json:"discount_price"`
	Image         string    `json:"image"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Brand struct {
	ID                 uint   `json:"Brand_id"`
	BrandName          string `json:"Brand_name"`
	BrandID            uint   `json:"master_category_id"`
	MasterCategoryName string `json:"master_category_name"`
}
