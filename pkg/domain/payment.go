package domain

import "time"

type PaymentDetails struct {
	ID      uint `gorm:"primaryKey" json:"id,omitempty"`
	OrderID uint `json:"order_id,omitempty"`
	// ShopOrder       ShopOrder     `gorm:"foreignKey:ShopOrderID" json:"-"`
	OrderTotal      uint          `json:"order_total"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	PaymentStatusID uint          `json:"payment_status_id,omitempty"`
	PaymentStatus   PaymentStatus `gorm:"foreignKey:PaymentStatusID" json:"-"`
	PaymentRef      string        `gorm:"default:null;uniqueIndex" json:"payment_ref"`
	UpdatedAt       time.Time
}

type PaymentStatus struct {
	ID     uint   `gorm:"primaryKey" json:"id,omitempty"`
	Status string `json:"payment_status,omitempty"`
}

type PaymentMethod struct {
	ID            uint `gorm:"primaryKey"`
	PaymentMethod string
}
