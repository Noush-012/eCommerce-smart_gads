package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ========================== JWT Token and cookie session  ========================== //

func JwtCookieSetup(c *gin.Context, name string, userId uint) bool {
	//time = 10 mins
	cookieTime := time.Now().Add(10 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprint(userId),
		ExpiresAt: cookieTime,
	})

	// Generate signed JWT token using env var of secret key
	if tokenString, err := token.SignedString([]byte(config.GetJWTConfig())); err == nil {

		// Set cookie with signed string if no error time = 10 hours
		c.SetCookie(name, tokenString, 10*3600, "", "", false, true)

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

			return []byte(config.GetJWTConfig()), nil
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
