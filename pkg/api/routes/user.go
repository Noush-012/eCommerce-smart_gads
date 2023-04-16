package routes

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {

	// Signup
	r.POST("/signup", userHandler.UserSignup)

	// Login
	r.POST("/login", userHandler.LoginSubmit)

}
