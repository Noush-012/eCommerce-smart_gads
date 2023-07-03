package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	mock "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/mock/repoMock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctrl)
	mockAuthRepository := mock.NewMockAuthRepository(ctrl)
	authUsecase := NewAuthUseCase(mockAuthRepository, mockUserRepository)

	ctx := context.Background()

	// Test data
	user := domain.Users{
		UserName:  "Jose",
		FirstName: "Jose",
		LastName:  "kutty",
		Age:       30,
		Email:     "jose@example.com",
		Phone:     "1234567890",
		Password:  "password",
	}

	// Test case : 1 "Success"
	t.Run("Success signup", func(t *testing.T) {
		mockUserRepository.EXPECT().FindUser(ctx, user).Return(domain.Users{}, nil)
		mockAuthRepository.EXPECT().SaveUser(ctx, gomock.Any()).Return(nil)

		err := authUsecase.SignUp(ctx, user)

		assert.NoError(t, err)
	})

	// Test case : 2 "User already exists"
	existingUser := user
	existingUser.ID = 1
	t.Run("User already exists, should return an error", func(t *testing.T) {
		mockUserRepository.EXPECT().FindUser(ctx, user).Return(existingUser, nil)

		err := authUsecase.SignUp(ctx, user)

		assert.EqualError(t, err, "Jose user already exists")
	})

	// Test case : 1 "Failed to save"
	// Passing null value to produce database error
	user.Email = ""
	t.Run("Error saving user, should return the error", func(t *testing.T) {
		expectedErr := errors.New("failed to save user")
		mockUserRepository.EXPECT().FindUser(ctx, user).Return(domain.Users{}, nil)
		fmt.Println("------------------------", user)
		mockAuthRepository.EXPECT().SaveUser(ctx, user).Return(expectedErr)

		err := authUsecase.SignUp(ctx, user)
		fmt.Println("----------------", err)
		assert.EqualError(t, err, expectedErr.Error())
	})

}

// func TestOTPLogin(t *testing.T) {
// 	ctrl := gomock.NewController(t)

// 	userRepoMock := mock.NewMockUserRepository(ctrl)
// 	authRepoMock := mock.NewMockAuthRepository(ctrl)
// 	AuthUseCase := NewAuthUseCase(authRepoMock, userRepoMock)

// }
