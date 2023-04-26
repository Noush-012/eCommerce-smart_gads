package request

type ProductReq struct {
	Name        string `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Description string `json:"description" gorm:"not null" binding:"required,min=10,max=1000"`
	CategoryID  uint   `json:"brand_id" binding:"required"`
	Price       uint   `json:"price" gorm:"not null" binding:"required,numeric"`
	Image       string `json:"image" gorm:"not null" binding:"required"`
}

type UpdateProductReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"product_name,omitempty"`
	Description string `json:"description,omitempty"`
	CategoryID  uint   `json:"brand_id,omitempty"`
	Price       uint   `json:"price,omitempty"`
	Image       string `json:"image,omitempty"`
}

type DeleteProductReq struct {
	ID uint `json:"Prod_id"`
}
