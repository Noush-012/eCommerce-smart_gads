package request

type AddToCartReq struct {
	UserID        uint `json:"-"`
	ProductItemID uint `json:"product_item_id" binding:"required"`
}
