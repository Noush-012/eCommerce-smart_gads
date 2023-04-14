package main

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/db"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadViper()
	db.ConnToDB()
}

func main() {
	g := gin.Default()
	userHandler := &handler.UserHandler{}

	g.POST("/signup", userHandler.UserSignup)
	g.Run() // listen and serve on 0.0.0.0:8080

}
