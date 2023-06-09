package routes

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler) {

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

	// Middleware routes
	api.Use(middleware.AuthenticateAdmin)
	{
		// Home
		api.GET("/", adminHandler.AdminHome)
		// Logout
		api.GET("/logout", adminHandler.LogoutAdmin)
		// Sales report
		api.GET("/sales-report", adminHandler.SalesReport)

		// Users dashboard
		user := api.Group("/users")
		{
			user.GET("/", adminHandler.ListUsers)
			user.PATCH("/block", adminHandler.BlockUser)
			user.GET("/orders", orderHandler.GetAllOrderHistory)
			user.PATCH("/orders", adminHandler.ChangeOrderStatus)
			user.GET("/return-orders", adminHandler.GetAllReturnOrder)
			user.PATCH("/return-orders/approval", adminHandler.ApproveReturnOrder)
			user.PATCH("/orders/delivery-update", adminHandler.UpdateDeliveryStatus)
		}

		// Brand
		brand := api.Group("/brands")
		{
			brand.GET("/", productHandler.GetAllBrands)
			brand.POST("/", productHandler.AddCategory)
		}

		// Product
		product := api.Group("/products")
		{
			// To list products
			product.GET("/", productHandler.ListProducts)
			// To list product items
			product.GET("/product-item/:product_id", productHandler.GetProductItem)
			// To add product
			product.POST("/", productHandler.AddProduct)
			// To update product
			product.PUT("/", productHandler.UpdateProduct)
			// To delete product
			product.DELETE("/", productHandler.DeleteProduct)
			// Add product item
			product.POST("/product-item", productHandler.AddProductItem)

		}

		coupons := api.Group("/coupons")
		{
			coupons.GET("/", couponHandler.ListAllCoupons)
			coupons.POST("/", couponHandler.CreateNewCoupon)
			coupons.PATCH("/", couponHandler.UpdateCoupon)
			coupons.DELETE("/:id", couponHandler.DeleteCoupon)
		}

	}
}
