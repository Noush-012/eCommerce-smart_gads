package handler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/auth"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AdminHandler struct {
	adminService interfaces.AdminService
	orderService interfaces.OrderService
}

func NewAdminHandler(adminService interfaces.AdminService, orderUseCase interfaces.OrderService) *AdminHandler {
	return &AdminHandler{adminService: adminService,
		orderService: orderUseCase,
	}
}

// AdminSignUp godoc
// @summary api for admin to login
// @id AdminSignUp
// @tags Admin Login
// @Param input body domain.Admin{} true "inputs"
// @Router /admin/login [post]
// @Success 200 {object} response.Response{} "Create admin account successful"
// @Failure 400 {object} response.Response{} "Missing or invalid entry"
// @Failure 500 {object} response.Response{} "Failed to create admin account"
func (a *AdminHandler) AdminSignUp(c *gin.Context) {
	var body domain.Admin

	// Bind signup form data
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, "Missing or invalid entry")
		return
	}

	// Create admin profile{
	if err := a.adminService.Signup(c, body); err != nil {
		responce := response.ErrorResponse(http.StatusBadRequest, "Failed to create admin account", err.Error(), body)
		c.JSON(http.StatusBadRequest, responce)
		return
	}
	// Success response
	c.JSON(http.StatusOK, "Create admin account successful")
}

// AdminLogin godoc
// @summary api for admin to login
// @id AdminLogin
// @tags Admin Login
// @Param input body request.LoginData{} true "Credentials"
// @Router /admin/login [post]
// @Success 200 {object} response.Response{} "successfully logged in"
// @Failure 400 {object} response.Response{} "Missing or invalid entry"
// @Failure 400 {object} response.Response{} "Failed to login"
// @Failure 500 {object} response.Response{} "Generate JWT failure"
func (a *AdminHandler) AdminLoginSubmit(c *gin.Context) {

	//
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
	}

	// validate login data
	var admin domain.Admin
	copier.Copy(&admin, body)
	admin, err := a.adminService.Login(c, admin)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to login", err.Error(), admin)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Setup JWT
	if !auth.JwtCookieSetup(c, "admin-auth", admin.ID) {
		response := response.ErrorResponse(500, "Generate JWT failure", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
	}

	// Success response
	response := response.SuccessResponse(200, "successfully logged in", nil)
	c.JSON(http.StatusOK, response)
}

// AdminHome godoc
// @summary api admin home
// @id AdminHome
// @tags Admin Home
// @Router /admin [get]
// @Success 200 {object} response.Response{} "Welcome to Admin Home"
func (a *AdminHandler) AdminHome(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"message":    "Welcome to Admin Home",
	})
}

// Logout godoc
// @summary api for admin to logout
// @description admin can logout
// @security ApiKeyAuth
// @id AdminLogout
// @tags Admin Logout
// @Router /logout [post]
// @Success 200 "Log out successful"
func (a *AdminHandler) LogoutAdmin(c *gin.Context) {
	c.SetCookie("admin-auth", "", -1, "", "", false, true)
	response := response.SuccessResponse(http.StatusOK, "Log out successful", nil)
	c.JSON(http.StatusOK, response)
}

// ListUsers godoc
// @summary api for admin to list users
// @id ListUsers
// @tags Admin User
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /admin/users [get]
// @Success 200 {object} response.Response{} "List user successful"
// @Failure 500 {object} response.Response{} "failed to get all users"
func (a *AdminHandler) ListUsers(c *gin.Context) {

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

	users, err := a.adminService.GetAllUser(c, pagination)
	if err != nil {
		respone := response.ErrorResponse(500, "failed to get all users", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, respone)
		return
	}

	// check there is no usee
	if len(users) == 0 {
		response := response.SuccessResponse(200, "Oh...No user to show", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := response.SuccessResponse(200, "List user successful", users)
	c.JSON(http.StatusOK, response)

}

// BlockUser godoc
// @summary api for admin to block or unblock user
// @id BlockUser
// @tags Admin User
// @Param input body request.UserID{} true "inputs"
// @Router /admin/users/block [patch]
// @Success 200 {object} response.Response{} "Successfully changed user block_status"
// @Failure 400 {object} response.Response{} "invalid input"
func (a *AdminHandler) BlockUser(c *gin.Context) {

	var body request.UserID

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := a.adminService.BlockUser(c, body.UserID)
	if err != nil {
		response := response.ErrorResponse(400, "faild to change user block_status", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.SuccessResponse(200, "Successfully changed user block_status", body.UserID)
	// if successfully blocked or unblock user then response 200
	c.JSON(http.StatusOK, response)
}

func (a *AdminHandler) UserOrderHistory(c *gin.Context) {

	var body request.UserID
	c.ShouldBindJSON(&body)

}

// ChangeOrderStatus godoc
// @summary api for admin to change order status of user
// @id ChangeOrderStatus
// @tags Admin ChangeOrderStatus
// @Param input body request.UpdateOrderStatus{} true "inputs"
// @Router /admin/users/orders [patch]
// @Success 200 {object} response.Response{} "Order status updated successfully!"
// @Failure 400 {object} response.Response{} "Missing inputs"
// @Failure 500 {object} response.Response{}"Something went wrong!"
func (a *AdminHandler) ChangeOrderStatus(c *gin.Context) {
	var body request.UpdateOrderStatus
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response := response.ErrorResponse(400, "Missing inputs", err.Error(), body)
		c.JSON(400, response)
		return
	}
	err = a.orderService.UpdateOrderStatus(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Order status updated successfully!", nil)
	c.JSON(200, response)

}

// SalesReport godoc
// @summary api for admin to download sales report as csv format
// @id SalesReport
// @tags Admin SalesReport
// @Router /admin/sales-report [get]
// @Success 500 {object} response.Response{} "Something went wrong!"
// @Failure 500 {object} response.Response{} "Something went wrong! failed to generate sales report"
// @Failure 500 {object} response.Response{}"Something went wrong!"
func (a *AdminHandler) SalesReport(c *gin.Context) {
	salesReport, err := a.adminService.SalesReport(c)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	// set header for downloading browser
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename= smart_gads_sales_report.csv")
	wr := csv.NewWriter(c.Writer)

	headers := []string{"Order ID", "User ID", "Total", "Coupon Code", "Payment Method", "Order Status", "Delivery Status", "Order Date"}
	if err := wr.Write(headers); err != nil {
		response := response.ErrorResponse(500, "Something went wrong! failed to generate sales report", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	for _, sales := range salesReport {
		var row = []string{
			fmt.Sprintf("%v", sales.OrderID),
			fmt.Sprintf("%v", sales.UserID),
			fmt.Sprintf("%v", sales.TotalAmount),
			sales.CouponCode,
			sales.PaymentMethod,
			sales.OrderStatus,
			sales.DeliveryStatus,
			sales.OrderDate.Format("2006-01-02 15:04:05")}

		if err := wr.Write(row); err != nil {
			response := response.ErrorResponse(500, "Something went wrong! failed to generate sales report", err.Error(), nil)
			c.JSON(500, response)
			return
		}

	}
	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
}
