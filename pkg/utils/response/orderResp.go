package response

import (
	"time"
)

type ShopOrder struct {
	Id              uint      `json:"-" `
	OrderID         uint      `json:"order_id"`
	OrderDate       time.Time `json:"order_date"`
	OrderTotal      float64   `json:"order_total"`
	Shipping_id     uint      `json:"-"`
	ShippingAddress Address   `json:"shipping_address"`
	OrderStatusID   uint      `json:"-"`
	OrderStatus     string    `json:"order_status"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentStatus   string    `json:"payment_status"`
	TransactionID   string    `json:"transaction_id"`
}

type OrderHistory struct {
	ID             uint      `json:"order_id"`
	OrderDate      time.Time `json:"order_date"`
	OrderStatus    string    `json:"order_status"`
	DeliveryStatus string    `json:"delivery_status"`
	OrderTotal     float64   `json:"order_total"`
	PaymentMethod  string    `json:"payment_type"`
	PaymentStatus  string    `json:"payment_status"`
	// Rating         uint      `json:"rating"`
	// Actions        Actions   `json:"actions"`
}
type Actions struct {
	Id   uint   `json:"-" `
	Name string `json:"action_name"`
}

// razorpay
type RazorPayOrderResp struct {
	RazorpayKey     string      `json:"razorpay_key"`
	UserID          uint        `json:"user_id"`
	AmountToPay     uint        `json:"amount_to_pay"`
	RazorpayAmount  uint        `json:"razorpay_amount"`
	RazorpayOrderID interface{} `json:"razorpay_order_id"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
	OrderID         uint        `json:"order_id"`
	// CouponID uint `json:"coupon_id"`
}
