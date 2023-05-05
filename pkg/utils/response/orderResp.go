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
	TransactionID   string    `json:"transaction_id"`
}
