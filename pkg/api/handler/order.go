package handler

import (
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderUseCase service.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: orderUseCase,
	}
}

// PPlaceCODOrder godoc
// @summary api for user to place an order on cart with COD
// @security ApiKeyAuth
// @tags User Cart
// @id PlaceCODOrder
// @Param id path uint true "Payment option ID"
// @Router /carts/place-order/cod [post]
// @Success 200 {object} response.Response{} "successfully order placed in COD"
// @Failure 400 {object} response.Response{}  "invalid input"
// @Failure 500 {object} response.Response{}  "failed to save shop order"
func (o *OrderHandler) PlaceCODOrder(c *gin.Context) {
	PaymentOptionID, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), PaymentOptionID)
		c.JSON(400, response)
		return
	}
	// get user from context
	userId := utils.GetUserIdFromContext(c)
	// get final cart details
	shopOrder, err := o.OrderService.PlaceOrderByCOD(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong! ", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Order placed successfuly", shopOrder)
	c.JSON(200, response)

}

// CheckoutCart godoc
// @summary api for user to checkout cart and proceed for payment
// @security ApiKeyAuth
// @tags User Cart
// @id CheckoutCart
func (o *OrderHandler) CheckOut(c *gin.Context) {

	userId := utils.GetUserIdFromContext(c)
	CheckOut, err := o.OrderService.CheckoutOrder(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly checked out", CheckOut)
	c.JSON(200, response)

}
