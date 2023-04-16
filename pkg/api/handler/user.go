package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/auth"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/req"
	request "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/req"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/verify"
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
		response := "Field should not be empty"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Copying
	var user domain.Users
	copier.Copy(&user, body)

	usr, err := u.toRepo.Login(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(usr.Phone)
	// Proceed for OTP if no error
	if _, err := verify.TwilioSendOTP("+91" + usr.Phone); err != nil {
		response := "failed to send otp"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := gin.H{"Successfuly send OTP to registered mobile number": usr.ID}
	c.JSON(http.StatusOK, response)

}

func (u *UserHandler) UserOTPVerify(c *gin.Context) {

	var body req.OTPVerify
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user = domain.Users{
		ID: body.UserID,
	}

	usr, err := u.toRepo.OTPLogin(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(body.OTP)
	// Verify otp
	err = verify.TwilioVerifyOTP("+91"+usr.Phone, body.OTP)
	if err != nil {
		response := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// setup JWT
	ok := auth.JwtCookieSetup(c, "user", usr.ID)
	if !ok {
		response := "failed to login"
		c.JSON(http.StatusInternalServerError, response)
		return

	}
	response := "Successfuly logged in"
	c.JSON(http.StatusOK, response)

}
