package usecase

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
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

func (c *couponUseCase) UpdateCoupon(ctx context.Context, couponData request.UpdateCoupon) error {
	if err := c.couponRepository.UpdateCoupon(ctx, couponData); err != nil {
		return err
	}
	return nil
}

func (c *couponUseCase) DeleteCoupon(ctx context.Context, couponId uint) error {
	if err := c.couponRepository.DeleteCoupon(ctx, couponId); err != nil {
		return err
	}
	return nil
}

func (c *couponUseCase) GetAllCoupons(ctx context.Context, page request.ReqPagination) (coupon []domain.Coupon, err error) {
	if coupon, err = c.couponRepository.GetAllCoupons(ctx, page); err != nil {
		return nil, err
	}
	return coupon, nil
}
