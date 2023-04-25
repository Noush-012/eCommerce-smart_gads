package routes

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler) {

	// Login
	login := api.Group("/login")
	{
		login.POST("/", adminHandler.AdminLoginSubmit)
	}

	// Signup
	signup := api.Group("/signup", adminHandler.AdminSignUp)
	{
		signup.POST("/")
	}

	// Middleware
	api.Use(middleware.AuthenticateAdmin)
	{
		api.GET("/", adminHandler.AdminHome)
		api.GET("/logout", adminHandler.LogoutAdmin)

		// Sales report

		// Users dashboard
		user := api.Group("/users")
		{
			user.GET("/", adminHandler.ListUsers)
			user.PATCH("/block", adminHandler.BlockUser)
		}

		// Brand
		brand := api.Group("/brands")
		{
			brand.GET("/")
			brand.POST("/", productHandler.AddBrand)
		}

		// Product
		product := api.Group("/products")
		{
			// To list products
			product.GET("/", productHandler.ListProducts)
			// To add product
			product.POST("/", productHandler.AddProduct)
			// To update product
			product.PUT("/", productHandler.UpdateProduct)
			// To delete product
			product.DELETE("/", productHandler.DeleteProduct)
			// Get product item
			// Add product item

			// Order

		}
	}
}
