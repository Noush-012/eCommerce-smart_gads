package repository

import (
	"context"
	"testing"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
)

// MockAuthDatabase is a mock implementation of AuthDatabase.
type MockAuthDatabase struct {
	SaveUserFunc func(ctx context.Context, user domain.Users) error
}

func (m *MockAuthDatabase) SaveUser(ctx context.Context, user domain.Users) error {
	if m.SaveUserFunc != nil {
		return m.SaveUserFunc(ctx, user)
	}
	return nil
}

func TestSaveUser(t *testing.T) {
	// Mock data
	mockUser := domain.Users{
		UserName:  "Test",
		FirstName: "Test",
		LastName:  "User",
		Age:       25,
		Email:     "user@example.com",
		Phone:     "1234567890",
		Password:  "password123",
	}

	// Create a mock AuthDatabase
	mockDB := &MockAuthDatabase{}

	// Set the mock implementation for SaveUserFunc
	mockDB.SaveUserFunc = func(ctx context.Context, user domain.Users) error {
		// Perform the necessary assertions
		if user.UserName != mockUser.UserName {
			t.Errorf("Expected user name: %s, got: %s", mockUser.UserName, user.UserName)
		}
		// Add more assertions for other user fields

		// Return nil to indicate success
		return nil
	}

	// Call the SaveUser method
	err := mockDB.SaveUser(context.Background(), mockUser)
	if err != nil {
		t.Errorf("SaveUser returned an error: %v", err)
	}
}
