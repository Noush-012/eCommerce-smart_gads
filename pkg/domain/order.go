package domain

import (
	"time"
)

// Order model
type ShopOrder struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"-" gorm:"not null"`
	Users      Users     `gorm:"foreignKey:UserID" json:"-"`
	OrderDate  time.Time `json:"order_date" gorm:"not null"`
	OrderTotal uint      `json:"order_total" gorm:"not null"`
	ShippingID uint      `json:"shipping_id" gorm:"not null"`
	// Address           Address        `gorm:"foreignKey:AddressID" json:"-"`
	OrderStatusID   uint          `json:"order_status_id" gorm:"not null"`
	OrderStatus     OrderStatus   `gorm:"foreignKey:OrderStatusID" json:"-"`
	PaymentMethodID uint          `json:"payment_method_id" gorm:"not null"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"-"`
	// PaymentStatusID   uint           `json:"payment_status_id" gorm:"not null"`
	// PaymentStatus     PaymentStatus  `gorm:"foreignKey:PaymentStatusID" json:"-"`
	CouponID          uint           `json:"coupon_id"`
	DeliveryStatusID  int            `json:"delivery_status_id" gorm:"not null"`
	DeliveryStatus    DeliveryStatus `gorm:"foreignKey:DeliveryStatusID" json:"-"`
	DeliveryUpdatedAt time.Time      `json:"delivery_time"`
}
type OrderLine struct {
	ID            uint        `json:"id" gorm:"primaryKey;not null"`
	ProductItemID uint        `json:"product_item_id" gorm:"not null"`
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemID" json:"-"`
	ShopOrderID   uint        `json:"shop_order_id" gorm:"not null"`
	ShopOrder     ShopOrder   `gorm:"foreignKey:ShopOrderID" json:"-"`
	Qty           uint        `json:"qty" gorm:"not null"`
	Price         uint        `json:"price" gorm:"not null"`
}

type OrderStatus struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"name" gorm:"not null"`
}

type DeliveryStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}

type Return struct {
	ID          uint      `gorm:"primaryKey"`
	ShopOrderID int       `json:"order_id"`
	ShopOrder   ShopOrder `gorm:"foreignKey:ShopOrderID"`
	Reason      string    `json:"string"`
	IsApproved  bool      `json:"approved" gorm:"default:false"`
	RequestedAt time.Time `json:"requested_at"`
}
