package routes

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {

	// Signup
	r.POST("/signup", userHandler.UserSignup)

	// Login with otp
	r.POST("/login", userHandler.LoginSubmit)

	// OTP verfication
	r.POST("/login-otp", userHandler.UserOTPVerify)

}
