package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {

	// Create a new instance of the user handler
	authHandler := NewAuthHandler(nil)

	testCases := []struct {
		name           string
		requestBody    request.SignupUserData
		expectedStatus int
		expectedResp   response.Response
	}{
		{
			name: "Success signup",
			requestBody: request.SignupUserData{
				UserName:        "noush",
				FirstName:       "Noushad",
				LastName:        "Ibrahim",
				Age:             25,
				Phone:           "8606879012",
				Email:           "nous@abc.com",
				Password:        "password",
				ConfirmPassword: "password",
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
				FirstName:       "test",
				LastName:        "user",
				Age:             25,
				Phone:           "1234567890",
				Email:           "testuser@example.com",
				Password:        "password123",
				ConfirmPassword: "password123",
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
			// Create a new Gin router
			router := gin.Default()
			// Define a route to handle the test request
			router.POST("/signup", func(c *gin.Context) {
				// Bind the request body to the specified type
				var requestBody request.SignupUserData
				if err := c.ShouldBindJSON(&requestBody); err != nil {
					// Handle invalid input
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
					return
				}

				// Invoke the UserSignup handler with the mock gin.Context
				authHandler.UserSignup(c)
			})

			// Create a new HTTP request with the JSON body
			jsonData, _ := json.Marshal(tc.requestBody)
			request, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))

			// Create a new response recorder
			recorder := httptest.NewRecorder()

			// Serve the HTTP request
			router.ServeHTTP(recorder, request)

			// Check the response status code
			assert.Equal(t, tc.expectedStatus, recorder.Code)

			// Decode the response body into a response.Response object
			var resp response.Response
			err := json.Unmarshal(recorder.Body.Bytes(), &resp)
			assert.NoError(t, err)
			// fmt.Println("----------------", tc.expectedResp.Message, "++++++++", resp.Message)
			// Assert the expected response properties
			assert.Equal(t, tc.expectedResp.Message, resp.Message)
			assert.Equal(t, tc.expectedResp.Errors, resp.Errors)
		})
	}
}
