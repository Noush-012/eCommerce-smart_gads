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
	uH, aH, pH := initializer.InitializeApi()

	g := gin.New()

	routes.UserRoutes(g.Group("/"), uH)
	routes.AdminRoutes(g.Group("/admin"), aH, pH)

	// port := ":3000"
	g.Run()

}
