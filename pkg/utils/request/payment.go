package request

type RazorpayReq struct {
	UserID          uint `json:"-"`
	PaymentOptionId uint `json:"payment_option_id" binding:"required,numeric"`
	PaymentMethodId uint `json:"payment_method_id" binding:"required,numeric"`
}

type RazorpayVerifyReq struct {
	UserID            uint `json:"-"`
	RazorpayPaymentId uint `json:"payment_option_id" binding:"required,numeric"`
	RazorpayOrderId   uint `json:"razorpay_order_id" binding:"required,numeric"`
}
