package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"gorm.io/gorm"
)

type couponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.CouponRepository {
	return &couponDatabase{DB: db}
}

// coupon management
// Fetch by coupon code
func (c *couponDatabase) GetCouponBycode(ctx context.Context, code string) (coupon domain.Coupon, err error) {
	query := `SELECT * FROM coupons WHERE code = ?`
	if err := c.DB.Raw(query, code).Scan(&coupon).Error; err != nil {
		return coupon, err
	}
	return coupon, nil
}
func (c *couponDatabase) GetCouponById(ctx context.Context, couponId uint) (coupon domain.Coupon, err error) {
	query := `SELECT * FROM coupons WHERE id = ?`
	if err := c.DB.Raw(query, couponId).Scan(&coupon).Error; err != nil {
		return coupon, err
	}
	return coupon, nil
}

func (c *couponDatabase) GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	query := `SELECT * FROM coupons ORDER BY id DESC LIMIT ? OFFSET ?`
	if err := c.DB.Raw(query, limit, offset).Scan(&coupon).Error; err != nil {
		return coupon, err
	}
	return coupon, nil
}

// Create a coupon
func (c *couponDatabase) CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error {

	query := `INSERT INTO coupons(code,min_order_value,discount_percent,discount_max_amount,valid_till)
	VALUES($1, $2, $3, $4, $5)`
	if err := c.DB.Exec(query, CouponData.Code, CouponData.MinOrderValue,
		CouponData.DiscountPercent, CouponData.DiscountMaxAmount, CouponData.ValidTill).Error; err != nil {
		return err
	}

	return nil
}

func (c *couponDatabase) UpdateCoupon(ctx context.Context, couponData request.UpdateCoupon) error {
	query := `UPDATE coupons
	SET code = COALESCE(NULLIF($1, ''), code),
	min_order_value = COALESCE(NULLIF($2, ''), min_order_value),
	discount_percent = COALESCE(NULLIF($3, ''), discount_percent),
	discount_max_amount = COALESCE(NULLIF($4, ''),discount_max_amount),
	valid_till = COALESCE(NULLIF($5, ''),valid_till)
	WHERE id = $6`

	if err := c.DB.Exec(query, couponData.Code, couponData.MinOrderValue, couponData.DiscountPercent, couponData.DiscountMaxAmount,
		couponData.DiscountMaxAmount, couponData.ValidTill, couponData.ID).Error; err != nil {
		return err
	}
	return nil
}

func (c *couponDatabase) DeleteCoupon(ctx context.Context, couponId uint) error {
	query := `DELETE FROM coupons WHERE id = ?`
	if err := c.DB.Exec(query, couponId).Error; err != nil {
		return err
	}
	return nil
}

func (c *couponDatabase) ApplyCoupon(ctx context.Context, data utils.ApplyCoupon) (AppliedCoupon utils.ApplyCouponResponse, err error) {

	// Get coupon and validate
	couponData, err := c.GetCouponBycode(ctx, data.CouponCode)
	if err != nil {
		return AppliedCoupon, err
	}
	// fmt.Println("--------------------->", couponData)
	if couponData.ValidTill.Before(time.Now()) {
		return AppliedCoupon, errors.New("coupon expired")
	}
	if data.TotalPrice < couponData.MinOrderValue {
		return AppliedCoupon, errors.New("unable to apply coupon. minimum order value not reached")
	}
	AppliedCoupon.CouponDiscount = data.TotalPrice * couponData.DiscountPercent / 100
	if AppliedCoupon.CouponDiscount > couponData.DiscountMaxAmount {
		AppliedCoupon.CouponDiscount = couponData.DiscountMaxAmount
	}
	AppliedCoupon.FinalPrice = data.TotalPrice - AppliedCoupon.CouponDiscount
	AppliedCoupon.CouponId = couponData.ID
	AppliedCoupon.CouponCode = couponData.Code

	// check coupon is valid or not
	return AppliedCoupon, nil

}
