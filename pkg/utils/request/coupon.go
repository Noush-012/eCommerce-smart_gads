package request

import "time"

type CreateCoupon struct {
	Code              string    `json:"code,omitempty"`
	MinOrderValue     float64   `json:"min_order_value,omitempty"`
	DiscountPercent   float64   `json:"discount_percent,omitempty"`
	DiscountMaxAmount float64   `json:"discount_max_amount,omitempty"`
	ValidTill         time.Time `json:"valid_till"`
}

type UpdateCoupon struct {
	ID                int       `json:"id" binding:"required"`
	Code              string    `json:"code,omitempty"`
	MinOrderValue     float64   `json:"min_order_value,omitempty"`
	DiscountPercent   float64   `json:"discount_percent,omitempty"`
	DiscountMaxAmount float64   `json:"discount_max_amount,omitempty"`
	ValidTill         time.Time `json:"valid_till"`
}
type Coupon struct {
	Coupon string `json:"coupon_code,omitempty"`
}
