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
	ShippingAddress Address   `json:"shipping_address" gorm:"-"`
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

type ReturnRequests struct {
	UserID        uint      `json:"user_id"`
	OrderId       uint      `json:"order_id"`
	RequestedAt   time.Time `json:"requested_at"`
	OrderDate     time.Time `json:"order_date"`
	DeliveredAt   time.Time `json:"delivered_at"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	Reason        string    `json:"reason"`
	OrderTotal    uint      `json:"order_total"`
	IsApproved    bool      `json:"is_approved"`
}
