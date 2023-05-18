package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
)

type CouponRepository interface {
	GetCouponBycode(ctx context.Context, code uint) (coupon domain.Coupon, err error)
	GetCouponById(ctx context.Context, couponId uint) (coupon domain.Coupon, err error)
	CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error
	UpdateCoupon(ctx context.Context, couponData request.UpdateCoupon) error
}
