package routes

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, paymentHandler *handler.PaymentHandler,
	orderHandler *handler.OrderHandler) {

	// Signup
	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.UserSignup)
	}
	// Login
	login := api.Group("/login")
	{
		login.GET("/", userHandler.LoginPage)
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

		// products routes
		products := api.Group("/products")
		{
			products.GET("/", productHandler.ListProducts)                           // show products
			products.GET("/product-item/:product_id", productHandler.GetProductItem) // show product items of a product
		}
		// cart routes
		cart := api.Group("/cart")
		{
			cart.GET("/", userHandler.GetcartItems)
			cart.POST("/", userHandler.AddToCart)
			cart.PUT("/", userHandler.UpdateCart)
			cart.DELETE("/", userHandler.DeleteCartItem)
			// checkout
			cart.GET("/checkout", orderHandler.CheckOut)
			cart.GET("/payment-option", paymentHandler.GetAllPaymentOptions)
			cart.GET("/checkout/:id", orderHandler.PlaceCODOrder)
			cart.GET("/checkout/razorpay", orderHandler.RazorPayCheckout)
		}
		order := api.Group("/orders")
		{
			order.GET("/", orderHandler.GetAllOrderHistory)
			order.POST("/")
		}
		profile := api.Group("/profile")
		{
			profile.GET("/", userHandler.Profile)
			profile.GET("/address", userHandler.GetAllAddress)
			profile.POST("/address", userHandler.AddAddress)
			// edit address
			profile.PATCH("/address", userHandler.UpdateAddress)

		}

	}

}
