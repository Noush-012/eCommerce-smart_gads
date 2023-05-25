package utils

type ApplyCoupon struct {
	CouponId   uint   `json:"coupon_id"`
	CouponCode string `json:"coupon_code"`
	UserId     uint   `json:"user_id"`
	TotalPrice float64
}
type ApplyCouponResponse struct {
	CouponId       uint    `json:"coupon_id"`
	CouponCode     string  `json:"coupon_code"`
	CouponDiscount float64 `json:"coupon_discount"`
	FinalPrice     float64 `json:"final_price"`
}
