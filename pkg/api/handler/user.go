package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/helpers/request"
	helper "github.com/Noush-012/Project-eCommerce-smart_gads/useCase/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userCase helper.UserUseCase
}

// func NewUserHandler(userUsecase helper.UserUseCase) *UserHandler {
// 	return &UserHandler{userCase: userUsecase}
// }

// User signup handler
func (u *UserHandler) UserSignup(c *gin.Context) {
	var body request.SignupUserData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user := domain.Users{}
	// var user domain.Users
	if err := copier.Copy(&user, body); err != nil {
		fmt.Println("Copy failed")
	}

	// validate user struct
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	// Check the user already exist in DB and save user if not
	if err := u.userCase.SignUp(c, user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "Account created successfuly"})

}

func LoginSubmit(c *gin.Context) {
	var body request.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.Email == "" && body.Password == "" && body.UserName == "" {
		_ = errors.New("please enter user_name and password")
		response := "Invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
}
