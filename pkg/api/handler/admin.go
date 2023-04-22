package handler

import (
	"errors"
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
}

func NewAdminHandler(adminService interfaces.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// Admin signup
func (a *AdminHandler) AdminSignUp(c *gin.Context) {
	var body domain.Admin

	// Bind signup form data
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid entry")
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

// Admin login submit
func (a *AdminHandler) AdminLoginSubmit(c *gin.Context) {

	//
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid entry", err.Error(), body)
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

// Homepage
func (a *AdminHandler) AdminHome(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"message":    "Welcome to Admin Home",
	})
}

// logout

func (a *AdminHandler) LogoutAdmin(c *gin.Context) {
	c.SetCookie("admin-auth", "", -1, "", "", false, true)
	response := response.SuccessResponse(http.StatusOK, "Log out successful", nil)
	c.JSON(http.StatusOK, response)
}

// list users
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

	response := response.SuccessResponse(200, "successfully got all users", users)
	c.JSON(http.StatusOK, response)

}

func (a *AdminHandler) BlockUser(ctx *gin.Context) {

	var body request.Block

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err := a.adminService.BlockUser(ctx, body.UserID)
	if err != nil {
		response := response.ErrorResponse(400, "faild to change user block_status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.SuccessResponse(200, "Successfully changed user block_status", body.UserID)
	// if successfully blocked or unblock user then response 200
	ctx.JSON(http.StatusOK, response)
}
