package handler

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentService service.PaymentService
}

func NewPaymentHandler(payUseCase service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		PaymentService: payUseCase,
	}
}

// GetAllPaymentOptions godoc
// @summary api for user to get all options for payment
// @security ApiKeyAuth
// @tags User
// @id AllPayment
// @Success 200 {object} response.Response{} "Payment option successfull"
// @Failure 500 {object} response.Response{}  "Something went wrong!"
func (p *PaymentHandler) GetAllPaymentOptions(c *gin.Context) {
	payOptions, err := p.PaymentService.GetAllPaymentOptions(c)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong! ", err.Error(), nil)
		c.JSON(500, response)
	}
	response := response.SuccessResponse(200, "Payment option successfull", payOptions)
	c.JSON(200, response)
}

func (p *PaymentHandler) SavePaymentDetails(c *gin.Context) {
	Smart_gads_orderId, _ := utils.StringToUint(c.Request.PostFormValue("orderId"))
	payMethodId, _ := utils.StringToUint(c.Request.PostFormValue("payment_id"))
	razorPayOrderId := c.Request.PostFormValue("razorpay_order_id")
	data := domain.PaymentDetails{
		OrderID:         Smart_gads_orderId,
		PaymentMethodID: payMethodId,
		PaymentStatusID: 2, // id 2 is for status "Paid"
		PaymentRef:      razorPayOrderId,
	}
	if err := p.PaymentService.SavePaymentDetails(c, data); err != nil {
		response := response.ErrorResponse(500, "Something went wrong! ", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Payment data updated successfull", nil)
	c.JSON(200, response)
}
