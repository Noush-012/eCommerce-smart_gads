package handler

import (
	"errors"
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

// PlaceCODOrder godoc
// @summary api for user to place an order on cart with COD
// @security ApiKeyAuth
// @tags User
// @id PlaceCODOrder
// @Param id path uint true "Payment option ID"
// @Router /carts/place-order/cod [post]
// @Success 200 {object} response.Response{} "successfully order placed in COD"
// @Failure 400 {object} response.Response{}  "invalid input"
// @Failure 500 {object} response.Response{}  "Something went wrong! "
func (o *OrderHandler) PlaceCODOrder(c *gin.Context) {
	PaymentMethodID, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), PaymentMethodID)
		c.JSON(400, response)
		return
	}
	// get user from context
	userId := utils.GetUserIdFromContext(c)
	// Get coupon code
	// bind coupon
	var coupon request.Coupon
	err = c.ShouldBindJSON(&coupon)
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	// Place order and save
	shopOrder, err := o.OrderService.PlaceOrderByCOD(c, userId, PaymentMethodID, coupon.Coupon)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong! ", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	data := gin.H{
		"Success": "success",
		"OrderID": shopOrder,
	}
	response := response.SuccessResponse(200, "Order placed successfuly Order ID :", data)
	c.JSON(200, response)

}

// CheckoutCart godoc
// @summary api for user to checkout cart, apply voucher and proceed for payment
// @security ApiKeyAuth
// @tags User
// @id CheckoutCart
// @Router /carts/checkout [get]
// @Success 200 {object} response.Response{}  "Successfuly checked out"
// @Failure 500 {object} response.Response{}  "Something went wrong! "
func (o *OrderHandler) CheckOut(c *gin.Context) {
	// bind coupon
	var coupon request.Coupon
	err := c.ShouldBindJSON(&coupon)
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), nil)
		c.JSON(400, response)
		return
	}

	userId := utils.GetUserIdFromContext(c)

	CheckOut, err := o.OrderService.CheckoutOrder(c, userId, coupon)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly checked out", CheckOut)
	c.JSON(200, response)

}

// GetAllOrderHistory godoc
// @summary api for admin and user to get all order history made
// @security ApiKeyAuth
// @tags User
// @tags Admin
// @id orderHistory
// @Router /carts/orders [get]
// @Success 200 {object} response.Response{}  "Order history successful"
// @Failure 400 {object} response.Response{}  "Missing user id"
// @Failure 500 {object} response.Response{}  "Something went wrong! "
func (o *OrderHandler) GetAllOrderHistory(c *gin.Context) {
	var userId uint
	var err error
	// if url path is admin/users/orders
	if c.Request.URL.Path == "/admin/users/orders" {
		userId, err = utils.StringToUint(c.Query("userId"))

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

// RazorpayCheckout godoc
// @summary api for create an razorpay order
// @security ApiKeyAuth
// @tags User
// @id RazorpayPage
// @Param input body   request.RazorpayReq{} true "inputs"
// @Router /carts/checkout/razorpay [post]
// @Success 200 {object}  response.Response{} "Checkout successfull"
// @Failure 500 {object}  response.Response{} "Something went wrong!"
func (o *OrderHandler) RazorPayCheckout(c *gin.Context) {
	// get user from context
	userId := utils.GetUserIdFromContext(c)

	var body request.RazorpayReq
	id, err := utils.StringToUint(c.Param("payment_id"))
	if err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Missing or invalid input", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	// err = c.ShouldBindJSON(&body)
	couponCode := c.Query("coupon")

	// if err != nil {
	// 	response := response.ErrorResponse(http.StatusBadRequest, "Missing or invalid input", err.Error(), nil)
	// 	c.JSON(400, response)
	// 	return
	// }
	body.CouponCode = couponCode
	body.UserID = userId
	body.PaymentMethodId = id
	razorpayOrder, err := o.OrderService.RazorPayCheckout(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}

	// response := response.SuccessResponse(200, "Razorpay chekout successful", razorpayOrder)
	// fmt.Println("resp", razorpayOrder.RazorpayKey)
	c.HTML(200, "app.html", razorpayOrder)

}

// RazorpayVerify godoc
// @summary api user for verify razorpay payment
// @security ApiKeyAuth
// @tags User
// @id RazorpayVerify
// @Param payment_method_id formData uint true "Payment Method ID"
// @Router /carts/checkout/razorpay/success [post]
// @Failure 500 {object}  response.Response{} "Failed to verify razor pay order!"
// @Success 200 {object} response.Response{}  "successfully payment completed and order approved"
func (o *OrderHandler) RazorpayVerify(c *gin.Context) {
	// get user from context
	userId := utils.GetUserIdFromContext(c)
	razorPayPaymentId := c.Request.PostFormValue("razorpay_payment_id")
	razorPayOrderId := c.Request.PostFormValue("razorpay_order_id")
	razorpay_signature := c.Request.PostFormValue("razorpay_signature")
	Smart_gads_orderId, _ := utils.StringToUint(c.Request.PostFormValue("orderId"))
	payMethodId, _ := utils.StringToUint(c.Request.PostFormValue("payment_id"))

	body := request.RazorpayVerifyReq{
		UserID:             userId,
		PaymentMethodID:    payMethodId,
		PaymentID:          razorPayPaymentId,
		RazorpayOrderId:    razorPayOrderId,
		Razorpay_signature: razorpay_signature,
	}
	err := utils.VerifyRazorPayPayment(body)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to verify razor pay order!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	// Update order status and clear cart
	Updatebody := request.UpdateStatus{
		UserId:   userId,
		StatusId: 2, // ID 2 is for satus "placed"
		OrderId:  Smart_gads_orderId,
	}
	if err := o.OrderService.UpdateOrderStatus(c, Updatebody); err != nil {
		response := response.ErrorResponse(500, "Failed to update order status!", err.Error(), nil)
		c.JSON(500, response)
		return
	}

	// calling payment handler to save payment details
	c.Next()
}

// ReturnOrder godoc
// @summary api user for return order
// @security ApiKeyAuth
// @tags User
// @id ReturnOrder
// @Param payment_method_id formData uint true "Order ID"
func (o *OrderHandler) ReturnRequest(c *gin.Context) {
	// get user from context
	userId := utils.GetUserIdFromContext(c)
	var body request.ReturnRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to bind return request!", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	body.UserID = userId
	// Validate return request and save if valid
	err = o.OrderService.ReturnEligibilityCheck(c, body)
	if err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Alert", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Return request successful, Please wait for approval")
	c.JSON(200, response)

}

// CancellOrder godoc
// @summary api user for cancel order
// @security ApiKeyAuth
// @tags User
// @id CancellOrder
// @Param payment_method_id formData uint true "Order ID"
func (o *OrderHandler) CancellOrder(c *gin.Context) {
	// get user from context
	userId := utils.GetUserIdFromContext(c)
	var body request.CancelOrder
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to bind cancell order request!", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	body.UserID = userId
	err = o.OrderService.OrderCancellation(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Order cancellation successful")
	c.JSON(200, response)

}
