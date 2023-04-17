package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// ========================== JWT Token and cookie session  ========================== //

func JwtCookieSetup(c *gin.Context, name string, userId interface{}) bool {
	//time = 10 mins
	cookieTime := time.Now().Add(10 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId, // Store logged user info in token
		"exp":    cookieTime,
	})

	// Generate signed JWT token using env var of secret key
	if tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY"))); err == nil {

		// Set cookie with signed string if no error time = 10 mins
		c.SetCookie(name, tokenString, 10*60, "", "", false, true)

		fmt.Println("JWT sign & set Cookie successful")
		return true
	}
	fmt.Println("Failed JWT setup")
	return false

}

func ValidateToken(tokenString string) (jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(viper.GetString("SECRET_KEY")), nil
		},
	)
	if err != nil || !token.Valid {
		fmt.Println("not valid token")
		return jwt.StandardClaims{}, errors.New("not valid token")
	}

	// then parse the token to claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		fmt.Println("can't parse the claims")
		return jwt.StandardClaims{}, errors.New("can't parse the claims")
	}

	return *claims, nil
}
