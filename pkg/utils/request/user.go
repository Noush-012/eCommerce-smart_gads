package request

import "time"

type AddressPatchReq struct {
	ID           uint      `json:"address_id"`
	UserID       uint      `json:"-"`
	House        string    `json:"house"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip_code"`
	Country      string    `json:"country"`
	IsDefault    bool      `json:"is_default"`
	UpdatedAt    time.Time `json:"-"`
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

type AddToWishlist struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
	Quantity      uint `json:"quantity" binding:"required"`
}
