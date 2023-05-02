package request

type AddToCartReq struct {
	UserID         uint    `json:"user_id"`
	ProductItemID  uint    `json:"product_item_id" binding:"required"`
	Quantity       uint    `json:"quantity" binding:"required"`
	Price          float64 `json:"-"`
	Discount_price uint    `json:"-"`
}

type UpdateCartReq struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"id" binding:"required"`
	Quantity      uint `json:"quantity" binding:"required"`
}
