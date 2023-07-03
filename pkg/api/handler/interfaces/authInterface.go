package interfaces

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	LoginPage(c *gin.Context)
	UserSignup(c *gin.Context)
	LoginSubmit(c *gin.Context)
	UserOTPVerify(c *gin.Context)
}
