package routes

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler) {

	// Signup
	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.UserSignup)
	}
	// Login
	login := api.Group("/login")
	{
		// Login with otp
		login.POST("/", userHandler.LoginSubmit)
		// OTP verfication
		login.POST("/otp-verify", userHandler.UserOTPVerify)
	}

	// Middleware routes
	api.Use(middleware.AuthenticateUser)
	{

		api.GET("/", userHandler.Home)
		api.GET("/logout", userHandler.LogoutUser)
	}

}
