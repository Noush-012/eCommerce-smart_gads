package domain

import "time"

type PaymentDetails struct {
	ID              uint          `gorm:"primaryKey" json:"id,omitempty"`
	OrderID         uint          `json:"order_id,omitempty"`
	Order           ShopOrder     `gorm:"foreignKey:ShopOrderID" json:"-"`
	OrderTotal      float64       `json:"order_total"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	PaymentStatusID uint          `json:"payment_status_id,omitempty"`
	PaymentStatus   PaymentStatus `gorm:"foreignKey:PaymentStatusID" json:"-"`
	PaymentRef      string        `gorm:"unique"`
	UpdatedAt       time.Time
}

type PaymentOption struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	// PaymentMethod PaymentMethod `json:"payment_method" gorm:"foreignKey:PaymentMethodID"`
}
type PaymentMethod struct {
	Id              uint   `json:"id" gorm:"primaryKey"`
	PaymentOptionID uint   `json:"payment_option_id" gorm:"not null"`
	Name            string `json:"name" gorm:"not null"`
}
type PaymentStatus struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status" gorm:"not null"`
}
