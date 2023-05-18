package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
)

type CouponService interface {
	CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error
}
