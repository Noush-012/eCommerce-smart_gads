package request

type ReqProduct struct {
	ProductName string `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Description string `json:"description" gorm:"not null" binding:"required,min=10,max=1000"`
	BrandID     uint   `json:"brand_id" binding:"required"`
	Price       uint   `json:"price" gorm:"not null" binding:"required,numeric"`
	Image       string `json:"image" gorm:"not null" binding:"required"`
}

type ReqProductUpdate struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name,omitempty"`
	Description string `json:"description,omitempty"`
	BrandID     uint   `json:"brand_id,omitempty"`
	Price       uint   `json:"price,omitempty"`
	Image       string `json:"image,omitempty"`
}
