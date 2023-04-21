package handler

import (
	"net/http"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/auth"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/req"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/resp"
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
		responce := resp.ErrorResponse(http.StatusBadRequest, "Failed to create admin account", err.Error(), body)
		c.JSON(http.StatusBadRequest, responce)
		return
	}
	// Success response
	c.JSON(http.StatusOK, "Create admin account successful")
}

// Admin login submit
func (a *AdminHandler) AdminLoginSubmit(c *gin.Context) {

	//
	var body req.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := resp.ErrorResponse(400, "invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
	}

	// validate login data
	var admin domain.Admin
	copier.Copy(&admin, body)
	admin, err := a.adminService.Login(c, admin)
	if err != nil {
		response := resp.ErrorResponse(400, "Failed to login", err.Error(), admin)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Setup JWT
	if !auth.JwtCookieSetup(c, "admin-auth", admin.ID) {
		response := resp.ErrorResponse(500, "Generate JWT failure", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
	}

	// Success response
	response := resp.SuccessResponse(200, "successfully logged in", nil)
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

func (u *AdminHandler) LogoutAdmin(c *gin.Context) {
	c.SetCookie("admin-auth", "", -1, "", "", false, true)
	response := resp.SuccessResponse(http.StatusOK, "Log out successful", nil)
	c.JSON(http.StatusOK, response)
}
