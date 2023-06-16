package routes

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler interfaces.UserHandler, productHandler *handler.ProductHandler,
	paymentHandler *handler.PaymentHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler) {

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
			products.GET("/brands", productHandler.GetAllBrands)
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
			cart.GET("/payment-option", paymentHandler.GetAllPaymentOptions)
			cart.POST("/checkout", orderHandler.CheckOut)
			cart.POST("/checkout/:id", orderHandler.PlaceCODOrder)

			// Razorpay
			cart.GET("/checkout/razorpay/:payment_id", orderHandler.RazorPayCheckout)
			cart.POST("/checkout/razorpay/success", orderHandler.RazorpayVerify, paymentHandler.SavePaymentDetails)

		}

		wishList := api.Group("/wishlist")
		{
			wishList.POST("/", userHandler.AddToWishlist)
			wishList.GET("/", userHandler.GetWishlist)
			wishList.DELETE("/:id", userHandler.DeleteFromWishlist)
		}
		order := api.Group("/orders")
		{
			order.GET("/", orderHandler.GetAllOrderHistory)
			order.POST("/return", orderHandler.ReturnRequest)
			order.PATCH("/cancel", orderHandler.CancellOrder)
		}
		profile := api.Group("/profile")
		{
			profile.GET("/", userHandler.Profile)
			profile.GET("/address", userHandler.GetAllAddress)
			profile.POST("/address", userHandler.AddAddress)
			// edit address
			profile.PUT("/address", userHandler.UpdateAddress)
			profile.DELETE("/address:id", userHandler.DeleteAddress)

		}
		coupon := api.Group("/coupons")
		{
			coupon.GET("/", couponHandler.ListAllCoupons)
		}

		wallet := api.Group("/wallet")
		{
			wallet.GET("/history", userHandler.GetWalletHistory)
		}

	}

}
