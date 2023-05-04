package domain

import (
	"time"

	"gorm.io/gorm"
)

// Order model
type ShopOrder struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"-"`
	OrderDate     time.Time `json:"order_date" gorm:"not null"`
	OrderTotal    float64   `json:"order_total" gorm:"not null"`
	PaymentTypeID uint      `json:"payment_type_id" gorm:"not null"`
	ShippingID    uint      `json:"shipping_id" gorm:"not null"`
	OrderStatusID uint      `json:"order_status_id" gorm:"not null"`
}

type Payment struct {
	gorm.Model
	OrderID       uint
	Amount        float64
	PaymentType   string
	TransactionID string
	PaymentDate   time.Time
}
type PaymentOption struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
type PaymentMethod struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
type OrderStatus struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"name" gorm:"not null"`
}
