package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mock "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/mock/useCaseMock"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func serSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockAuthService(ctrl)
	// Create a new instance of the user handler
	authHandler := NewAuthHandler(c)

	testCases := []struct {
		name           string
		requestBody    request.SignupUserData
		buildStub      func(authUsecase mock.MockAuthService)
		expectedStatus int
		expectedResp   response.Response
	}{
		{
			name: "Success signup",
			requestBody: request.SignupUserData{
				UserName:        "noush",
				FirstName:       "Noushad",
				LastName:        "Ibrahim",
				Age:             24,
				Phone:           "8606879012",
				Email:           "noush@abc.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			buildStub: func(authUsecase mock.MockAuthService) {
				authUsecase.EXPECT().SignUp(gomock.Any(), request.SignupUserData{
					UserName:        "noush",
					FirstName:       "Noushad",
					LastName:        "Ibrahim",
					Age:             24,
					Phone:           "8606879012",
					Email:           "noush@abc.com",
					Password:        "password",
					ConfirmPassword: "password",
				}).Times(1).Return(response.Response{
					Message: "Account created successfuly",
					Errors:  nil,
				})
			},
			expectedStatus: http.StatusOK,
			expectedResp: response.Response{
				Message: "Account created successfuly",
				Errors:  nil,
			},
		},
		{
			name: "Empty username",
			requestBody: request.SignupUserData{
				UserName:        "",
				FirstName:       "Noushad",
				LastName:        "Ibrahim",
				Age:             24,
				Phone:           "8606879012",
				Email:           "noush@abc.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			buildStub: func(authUsecase mock.MockAuthService) {
				authUsecase.EXPECT().SignUp(gomock.Any(), request.SignupUserData{
					UserName:        "",
					FirstName:       "Noushad",
					LastName:        "Ibrahim",
					Age:             24,
					Phone:           "8606879012",
					Email:           "noush@abc.com",
					Password:        "password",
					ConfirmPassword: "password",
				}).Times(1).Return(response.Response{
					Message: "Invalid input",
					Errors:  "",
				})
			},
			expectedStatus: http.StatusBadRequest,
			expectedResp: response.Response{
				Message: "Invalid input",
				Errors:  "",
			},
		},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// passing mock use case to buildStub function which is
			tc.buildStub(*c)
			// Create a new Gin router
			engine := gin.Default()
			// Define a route to handle the test request
			// engine.POST("/signup", func(c *gin.Context) {
			// 	// Bind the request body to the specified type
			// 	var requestBody request.SignupUserData
			// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
			// 		// Handle invalid input
			// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			// 		return
			// 	}

			// 	// Invoke the UserSignup handler with the mock gin.Context
			// 	authHandler.UserSignup(c)
			// })
			engine.POST("/signup", authHandler.UserSignup)

			// Create a new HTTP request with the JSON body
			jsonData, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))

			// Create a new response recorder
			recorder := httptest.NewRecorder()

			engine.ServeHTTP(recorder, req)

			// Check the response status code
			assert.Equal(t, tc.expectedStatus, recorder.Code)

			// Decode the response body into a response.Response
			var resp response.Response
			err := json.Unmarshal(recorder.Body.Bytes(), &resp)
			assert.NoError(t, err)
			fmt.Printf("type of actual data %t\n", resp.Data)
			data, ok := resp.Data.(map[string]interface{})
			if ok {
				userData := request.SignupUserData{
					UserName:  data["user_name"].(string),
					FirstName: data["f_name"].(string),
					LastName:  data["l_name"].(string),
					Email:     data["email"].(string),
					Phone:     data["phone"].(string),
				}
				if !reflect.DeepEqual(tc.expectedResp, userData) {
					t.Errorf("got %q,but want %q", userData, tc.expectedResp)
				}
			} else {
				t.Errorf("actual.Data is not of type map[string]interface{}")
			}

			// Assert the expected response properties
			assert.Equal(t, tc.expectedResp.Message, resp.Message)
			assert.Equal(t, tc.expectedResp.Errors, resp.Errors)
		})
	}
}
