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
	c.HTML(http.StatusOK, "checkout.html", response)

}

func (o *OrderHandler) GetAllOrderHistory(c *gin.Context) {
	var userId uint
	var err error
	// if url path is admin/users/orders
	if c.Request.URL.Path == "/admin/users/orders" {
		userId, err = utils.StringToUint(c.Query("userId"))
		fmt.Println(userId)
		if err != nil {
			response := response.ErrorResponse(400, "Missing user id", err.Error(), nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		// get user from context
		userId = utils.GetUserIdFromContext(c)
	}

	count, err1 := utils.StringToUint(c.Query("count"))
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))

	err1 = errors.Join(err1, err2)
	if err1 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	orderHistory, err := o.OrderService.GetOrderHistory(c, pagination, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), orderHistory)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Order history successful", orderHistory)
	c.JSON(200, response)

}

func (o *OrderHandler) RazorPayCheckout(c *gin.Context) {
	// get user from context
	userId := utils.GetUserIdFromContext(c)

	// Verify payment request id is razorpay
	var body request.RazorpayReq
	if err := c.BindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	body.UserID = userId
	razorpayOrder, err := o.OrderService.RazorPayCheckout(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Razorpay chekout successful", razorpayOrder)
	c.JSON(200, response)

}

func (o *OrderHandler) RazorpayVerify(c *gin.Context) {
	// get user from context
	// userId := utils.GetUserIdFromContext(c)
	// var body request.RazorpayVerifyReq
	// if err := c.BindJSON(&body); err != nil {
	// 	response := "invalid input"
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// utils.VerifyRazorPayPayment(body.RazorpayOrderId, body.RazorpayPaymentId)

}
