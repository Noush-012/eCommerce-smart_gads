package request

type RazorpayReq struct {
	UserID          uint `json:"-"`
	PaymentOptionId uint `json:"payment_option_id" binding:"required"`
	PaymentMethodId uint `json:"payment_method_id" binding:"required"`
}
