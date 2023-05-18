package usecase

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
)

type couponUseCase struct {
	couponRepository interfaces.CouponRepository
}

func NewCouponUseCase(CouponRepo interfaces.CouponRepository) service.CouponService {
	return &couponUseCase{couponRepository: CouponRepo}
}

func (c *couponUseCase) CreateNewCoupon(ctx context.Context, CouponData request.CreateCoupon) error {
	if err := c.couponRepository.CreateNewCoupon(ctx, CouponData); err != nil {
		return err
	}
	return nil
}
