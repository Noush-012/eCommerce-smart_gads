package handler

import (
	"errors"
	"fmt"
	"net/http"

	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
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

// CreateNewCoupon godoc
// @summary api for admin to create a cooupon
// @id CreateNewCoupon
// @tags Admin CreateNewCoupon
// @Param input body   request.CreateCoupon{} true "inputs"
// @Router /admin/coupons [post]
// @Success 200 {object} response.Response{} "Coupon created successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Internal server error"
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
		return
	}
	response := response.SuccessResponse(200, "Coupon created successfully", nil)
	ctx.JSON(200, response)

}

// UpdateCoupon godoc
// @summary api for admin to update a cooupon
// @id UpdateCoupon
// @tags Admin UpdateCoupon
// @Param input body   request.UpdateCoupon{} true "inputs"
// @Router /admin/coupons [put]
// @Success 200 {object} response.Response{} "Coupon updated successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Internal server error"
func (c *CouponHandler) UpdateCoupon(ctx *gin.Context) {
	var body request.UpdateCoupon
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		ctx.JSON(400, response)
		return
	}

	if err := c.couponService.UpdateCoupon(ctx, body); err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), body)
		ctx.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Coupon updated successfully", nil)
	ctx.JSON(200, response)

}

// DeleteCoupon godoc
// @summary api for admin to delete a cooupon
// @id DeleteCoupon
// @tags Admin DeleteCoupon
// @Param id path int true "id"
// @Router /admin/coupons/{id} [delete]
// @Success 200 {object} response.Response{} "Coupon deleted successfully"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Internal server error"
func (c *CouponHandler) DeleteCoupon(ctx *gin.Context) {
	id, err := utils.StringToUint(ctx.Param("id"))
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), id)
		ctx.JSON(400, response)
		return
	}
	if err := c.couponService.DeleteCoupon(ctx, id); err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), id)
		ctx.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Coupon deleted successfully", nil)
	ctx.JSON(200, response)
}

// LisAllCoupon godoc
// @summary api for admin to list all coupons
// @id ListAllCoupons
// @tags Admin ListAllCoupons
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /admin/coupons [get]
// @Success 200 {object} response.Response{} "List of coupons"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Internal server error"
func (c *CouponHandler) ListAllCoupons(ctx *gin.Context) {

	count, err1 := utils.StringToUint(ctx.Query("count"))
	pageNumber, err2 := utils.StringToUint(ctx.Query("page_number"))
	fmt.Println(count, pageNumber)
	err1 = errors.Join(err1, err2)
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}
	coupons, err := c.couponService.GetAllCoupons(ctx, pagination)
	if err != nil {
		response := response.ErrorResponse(500, "Internal server error", err.Error(), nil)
		ctx.JSON(500, response)
		return
	}

	response := response.SuccessResponse(200, "List of coupons", coupons)
	ctx.JSON(200, response)
}
