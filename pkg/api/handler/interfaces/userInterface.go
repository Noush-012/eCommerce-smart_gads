package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	LoginPage(c *gin.Context)
	UserSignup(c *gin.Context)
	LoginSubmit(c *gin.Context)
	UserOTPVerify(c *gin.Context)
	Home(c *gin.Context)
	GetcartItems(c *gin.Context)
	AddToCart(c *gin.Context)
	UpdateCart(c *gin.Context)
	DeleteCartItem(c *gin.Context)
	Profile(c *gin.Context)
	AddAddress(c *gin.Context)
	UpdateAddress(c *gin.Context)
	DeleteAddress(c *gin.Context)
	GetAllAddress(c *gin.Context)
	AddToWishlist(c *gin.Context)
	GetWishlist(c *gin.Context)
	DeleteFromWishlist(c *gin.Context)
	GetWalletHistory(c *gin.Context)
	LogoutUser(c *gin.Context)
}
