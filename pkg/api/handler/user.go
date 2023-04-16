package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	request "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/req"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	toRepo interfaces.UserService
}

func NewUserHandler(userUsecase interfaces.UserService) *UserHandler {
	return &UserHandler{toRepo: userUsecase}
}

// User signup handler
func (u *UserHandler) UserSignup(c *gin.Context) {
	var body request.SignupUserData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user domain.Users
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
	if err := u.toRepo.SignUp(c, user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "Account created successfuly"})

}

func (u *UserHandler) LoginSubmit(c *gin.Context) {
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
	// Copying
	var user domain.Users
	copier.Copy(&user, body)

	_, err := u.toRepo.Login(c, user)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// setup JWT

}
