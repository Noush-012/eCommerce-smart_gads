package request

type RazorpayReq struct {
	UserID          uint   `json:"-"`
	PaymentMethodId uint   `json:"-"`
	CouponCode      string `json:"coupon_code"`
}

type RazorpayVerifyReq struct {
	UserID             uint   `json:"-"`
	PaymentMethodID    uint   `json:"payment_method_id"`
	PaymentID          string `json:"payment_Method_id"`
	RazorpayOrderId    string `json:"razorpay_order_id"`
	Razorpay_signature string `json:"Razorpay_signature"`
}
