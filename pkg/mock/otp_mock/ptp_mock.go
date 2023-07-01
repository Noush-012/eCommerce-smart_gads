package mock

import "github.com/golang/mock/gomock"

// MockTwilioClient is a mock implementation of the TwilioClient struct
type MockTwilioClient struct {
	mockCtrl  *gomock.Controller
	mock      *MockTwilioClient
	SendOTPFn func(phone string) (interface{}, error)
}

// NewMockTwilioClient creates a new instance of the mock TwilioClient struct
func NewMockTwilioClient(ctrl *gomock.Controller) *MockTwilioClient {
	return &MockTwilioClient{
		mockCtrl: ctrl,
		mock:     NewMockTwilioClient(ctrl),
	}
}
