package response

type PaymentOptionResp struct {
	Id   uint   `json:"code"`
	Name string `json:"option"`
}

// razorpay
type RazorPayOrderResp struct {
	RazorpayKey     string      `json:"razorpay_key"`
	UserID          uint        `json:"user_id"`
	SGPay_id        uint        `json:"SGPay_id"`
	AmountToPay     uint        `json:"amount_to_pay"`
	RazorpayAmount  uint        `json:"razorpay_amount"`
	RazorpayOrderID interface{} `json:"razorpay_order_id"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
	OrderID         uint        `json:"order_id"`
	// CouponID uint `json:"coupon_id"`
}
