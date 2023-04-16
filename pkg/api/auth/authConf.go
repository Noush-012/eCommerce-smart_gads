package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// ========================== JWT Token and cookie session  ========================== //

func JwtCookieSetup(c *gin.Context, name string, userId interface{}) bool {
	cookieTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId, // Store logged user info in token
		"exp":    cookieTime,
	})

	// Generate signed JWT token using env var of secret key
	if tokenString, err := token.SignedString([]byte(viper.GetString("DATABASE"))); err == nil {

		// Set cookie with signed string if no error
		c.SetCookie(name, tokenString, 10*60, "", "", false, true)

		fmt.Println("JWT sign & set Cookie successful")
		return true
	}
	fmt.Println("Failed JWT setup")
	return false

}
