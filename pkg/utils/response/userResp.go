package response

import (
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

type Profile struct {
	ID             uint           `json:"id"`
	UserName       string         `json:"user_name" copire:"must"`
	FirstName      string         `json:"first_name" copier:"must"`
	LastName       string         `json:"last_name" copier:"must"`
	Age            uint           `json:"age" copier:"must"`
	Email          string         `json:"email" copier:"must"`
	Phone          string         `json:"phone" copier:"must"`
	DefaultAddress Address        `json:"default_address"`
	OrderHistory   []OrderHistory `json:"order_history"`
}

type UserRespStrcut struct {
	ID          uint      `json:"id" copier:"must"`
	FirstName   string    `json:"first_name" copier:"must"`
	LastName    string    `json:"last_name" copier:"must"`
	Age         uint      `json:"age" copier:"must"`
	Email       string    `json:"email" copier:"must"`
	UserName    string    `json:"user_name" copire:"must"`
	Phone       string    `json:"phone" copier:"must"`
	BlockStatus bool      `json:"block_status" copier:"must"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// home page response
type ResUserHome struct {
	// Products []ResponseProduct `json:"products"`
	User UserRespStrcut `json:"user"`
}

// cart item reponse
type CartItemResp struct {
	ProductItemID uint   `json:"id"`
	Name          string `json:"product_name"`
	Price         uint   `json:"price"`
	DiscountPrice uint   `json:"discount_price"`
	Quantity      uint   `json:"quantity"`
	QtyLeft       uint   `json:"quantity_left"`
	StockStatus   bool   `json:"stock_status"`
	SubTotal      uint   `json:"sub_total"`
}

type CartResp struct {
	CartItems         []CartItemResp `json:"-"`
	TotalProductItems uint           `json:"total_product_items"`
	TotalQty          uint           `json:"total_qty"`
	TotalPrice        float64        `json:"total_price"`
	DiscountAmount    float64        `json:"discount"`
	AppliedCouponID   uint           `json:"applied_coupon_id"`
	AppliedCouponCode string         `json:"applied_coupon_code"`
	CouponDiscount    float64        `json:"coupon_discount"`
	FinalPrice        uint           `json:"final_price"`
	DefaultShipping   Address        `json:"default_shipping"`
}
type Address struct {
	ID           uint   `json:"address_id"`
	UserID       uint   `json:"-"`
	House        string `json:"house"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	Country      string `json:"country"`
	IsDefault    bool   `json:"is_default"`
}

type CheckoutOrder struct {
	UserID         uint           `json:"-"`
	CartItemResp   []CartItemResp `json:"cart_items"`
	TotalQty       uint           `json:"total_qty"`
	TotalPrice     uint           `json:"total_price"`
	Discount       uint           `json:"discount"`
	DefaultAddress domain.Address `json:"address"`
}

type UserContact struct {
	Email string
	Phone string
}
type Wishlist struct {
	ProductItemId uint   `json:"product_item_id"`
	ProductName   string `json:"product_name"`
	Color         string `json:"color"`
	Storage       string `json:"storage"`
	Price         uint   `json:"price"`
	Quantity      uint   `json:"quantity"`
	Image         string `json:"image"`
}
