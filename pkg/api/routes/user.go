package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// userHandler := handler.NewUserHandler()
	// Signup
	// signup := api.Group("/signup")
	// {
	// 	signup.POST("/")
	// }

	r.POST("/signup")

}
