package main

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/routes"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/initializer"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/verify"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadViper()
	verify.SetClient()
}

func main() {
	uH := initializer.InitializeApi()

	g := gin.New()

	routes.UserRoutes(g.Group("/"), uH)

	g.Run() // listen and serve on 0.0.0.0:8080

}
