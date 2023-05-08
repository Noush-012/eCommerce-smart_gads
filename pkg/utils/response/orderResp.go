package response

import (
	"time"
)

type ShopOrder struct {
	Id              uint      `json:"-" `
	OrderID         uint      `json:"order_id"`
	OrderDate       time.Time `json:"order_date"`
	OrderTotal      float64   `json:"order_total"`
	ShippingAddress Address   `json:"shipping_address"`
	OrderStatusID   uint      `json:"-"`
	OrderStatus     string    `json:"order_status"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentStatus   string    `json:"payment_status"`
	TransactionID   string    `json:"transaction_id"`
}

type OrderHistory struct {
	OrderID        uint      `json:"order_id"`
	OrderDate      time.Time `json:"order_date"`
	OrderStatus    string    `json:"order_status"`
	DeliveryStatus string    `json:"delivery_status"`
	OrderTotal     float64   `json:"order_total"`
	Rating         uint      `json:"rating"`
	Actions        Actions   `json:"actions"`
}
type Actions struct {
	Id   uint   `json:"-" `
	Name string `json:"action_name"`
}
