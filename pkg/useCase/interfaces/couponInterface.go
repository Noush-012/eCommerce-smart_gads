package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
)

type CouponService interface {
	CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error
	GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error)

	UpdateCoupon(ctx context.Context, couponData request.UpdateCoupon) error
	DeleteCoupon(ctx context.Context, couponId uint) error
}
