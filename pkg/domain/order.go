package domain

import (
	"time"
)

// Order model
type ShopOrder struct {
	Id              uint          `json:"id" gorm:"primaryKey"`
	UserID          uint          `json:"-" gorm:"not null"`
	OrderDate       time.Time     `json:"order_date" gorm:"not null"`
	OrderTotal      float64       `json:"order_total" gorm:"not null"`
	ShippingID      uint          `json:"shipping_id" gorm:"not null"`
	OrderStatusID   uint          `json:"order_status_id" gorm:"not null"`
	PaymentOptionID uint          `json:"payment_option_id" gorm:"not null"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentStatusID uint          `json:"payment_status_id"`
	TransactionID   string        `json:"transaction_id" gorm:"default:null"`
	PaymentDate     time.Time     `json:"payment_date" gorm:"default:null"`
	OrderStatus     OrderStatus   `json:"order_status" gorm:"foreignKey:OrderStatusID"`
	PaymentStatus   PaymentStatus `json:"payment_status" gorm:"foreignKey:PaymentStatusID"`
	PaymentOption   PaymentOption `json:"payment_option" gorm:"foreignKey:PaymentOptionID"`
	PaymentMethod   PaymentMethod `json:"payment_method" gorm:"foreignKey:PaymentMethodID"`
}
type OrderLine struct {
	ID            uint      `json:"id" gorm:"primaryKey;not null"`
	ProductItemID uint      `json:"proudct_item_id" gorm:"not null"`
	ShopOrderID   uint      `json:"shop_order_id" gorm:"not null"`
	ShopOrder     ShopOrder `json:"-"`
	Qty           uint      `json:"qty" gorm:"not null"`
	Price         uint      `json:"price" gorm:"not null"`
}

type OrderStatus struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"name" gorm:"not null"`
}
