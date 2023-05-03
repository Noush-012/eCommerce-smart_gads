package request

type AddressReq struct {
}

type AddToCartReq struct {
	UserID         uint    `json:"user_id"`
	ProductItemID  uint    `json:"product_item_id" binding:"required"`
	Quantity       uint    `json:"quantity" binding:"required"`
	Price          float64 `json:"-"`
	Discount_price uint    `json:"-"`
}

type UpdateCartReq struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
	Quantity      uint `json:"quantity" binding:"required"`
}

type DeleteCartItemReq struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
}
