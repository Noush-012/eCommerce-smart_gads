package handler

import (
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponService service.CouponService
}

func NewCouponHandler(CouponUseCase service.CouponService) *CouponHandler {
	return &CouponHandler{
		couponService: CouponUseCase,
	}
}

func (c *CouponHandler) CreateNewCoupon(ctx *gin.Context) {
	var body request.CreateCoupon
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		ctx.JSON(400, response)
		return
	}
	if err := c.couponService.CreateNewCoupon(ctx, body); err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), body)
		ctx.JSON(500, response)
	}
	response := response.SuccessResponse(200, "Coupon created successfully", nil)
	ctx.JSON(200, response)

}
