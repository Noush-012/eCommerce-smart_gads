package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	mock "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/mock/useCaseMock"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {

	// ctx := context.Background()

	testCase := map[string]struct {
		signupData    request.SignupUserData
		buildStub     func(useCaseMock *mock.MockAuthService, signupData request.SignupUserData)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"Valid Signup": {
			signupData: request.SignupUserData{
				UserName:        "noush",
				FirstName:       "Noushad",
				LastName:        "Ibrahim",
				Age:             24,
				Phone:           "8606879012",
				Email:           "noush@abc.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			buildStub: func(useCaseMock *mock.MockAuthService, signupData request.SignupUserData) {
				// copying signupData to domain.user for pass to Mock usecase
				var user domain.Users
				if err := copier.Copy(&user, signupData); err != nil {
					fmt.Println("Copy failed")
				}
				useCaseMock.EXPECT().SignUp(gomock.Any(), user).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"Without email": {
			signupData: request.SignupUserData{
				UserName:        "noush",
				FirstName:       "Noushad",
				LastName:        "Ibrahim",
				Age:             24,
				Phone:           "8606879012",
				Email:           "",
				Password:        "password",
				ConfirmPassword: "password",
			},
			buildStub: func(useCaseMock *mock.MockAuthService, signupData request.SignupUserData) {
				// not expecting calls to usecase
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				responseStruct, err := getResponseStructFromResponseBody(responseRecorder.Body)
				assert.Nil(t, err)
				expectedMessage := "Invalid input"
				assert.Equal(t, expectedMessage, responseStruct.Message)
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mock.NewMockAuthService(ctrl)
			test.buildStub(mockUseCase, test.signupData)

			AuthHandler := NewAuthHandler(mockUseCase)

			server := gin.Default()
			server.POST("/signup", AuthHandler.UserSignup)

			jsonData, err := json.Marshal(test.signupData)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)
			test.checkResponse(t, responseRecorder)

		})

	}
}

// convert / un marshal response body to response.Response struct
func getResponseStructFromResponseBody(responseBody *bytes.Buffer) (responseStruct response.Response, err error) {
	data, err := io.ReadAll(responseBody)
	json.Unmarshal(data, &responseStruct)
	return
}
