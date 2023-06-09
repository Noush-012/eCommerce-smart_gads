// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/useCase/interfaces/userInterface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	request "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	response "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AddToWishlist mocks base method.
func (m *MockUserService) AddToWishlist(ctx context.Context, wishlistData request.AddToWishlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToWishlist", ctx, wishlistData)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToWishlist indicates an expected call of AddToWishlist.
func (mr *MockUserServiceMockRecorder) AddToWishlist(ctx, wishlistData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToWishlist", reflect.TypeOf((*MockUserService)(nil).AddToWishlist), ctx, wishlistData)
}

// Addaddress mocks base method.
func (m *MockUserService) Addaddress(ctx context.Context, address request.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addaddress", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Addaddress indicates an expected call of Addaddress.
func (mr *MockUserServiceMockRecorder) Addaddress(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addaddress", reflect.TypeOf((*MockUserService)(nil).Addaddress), ctx, address)
}

// DeleteAddress mocks base method.
func (m *MockUserService) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", ctx, userID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserServiceMockRecorder) DeleteAddress(ctx, userID, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserService)(nil).DeleteAddress), ctx, userID, addressID)
}

// DeleteFromWishlist mocks base method.
func (m *MockUserService) DeleteFromWishlist(ctx context.Context, productId, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFromWishlist", ctx, productId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromWishlist indicates an expected call of DeleteFromWishlist.
func (mr *MockUserServiceMockRecorder) DeleteFromWishlist(ctx, productId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromWishlist", reflect.TypeOf((*MockUserService)(nil).DeleteFromWishlist), ctx, productId, userId)
}

// GetAllAddress mocks base method.
func (m *MockUserService) GetAllAddress(ctx context.Context, userId uint) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAddress", ctx, userId)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAddress indicates an expected call of GetAllAddress.
func (mr *MockUserServiceMockRecorder) GetAllAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAddress", reflect.TypeOf((*MockUserService)(nil).GetAllAddress), ctx, userId)
}

// GetCartItemsbyCartId mocks base method.
func (m *MockUserService) GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) ([]response.CartItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItemsbyCartId", ctx, page, userID)
	ret0, _ := ret[0].([]response.CartItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItemsbyCartId indicates an expected call of GetCartItemsbyCartId.
func (mr *MockUserServiceMockRecorder) GetCartItemsbyCartId(ctx, page, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItemsbyCartId", reflect.TypeOf((*MockUserService)(nil).GetCartItemsbyCartId), ctx, page, userID)
}

// GetWalletHistory mocks base method.
func (m *MockUserService) GetWalletHistory(ctx context.Context, userId uint) ([]domain.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletHistory", ctx, userId)
	ret0, _ := ret[0].([]domain.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletHistory indicates an expected call of GetWalletHistory.
func (mr *MockUserServiceMockRecorder) GetWalletHistory(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletHistory", reflect.TypeOf((*MockUserService)(nil).GetWalletHistory), ctx, userId)
}

// GetWishlist mocks base method.
func (m *MockUserService) GetWishlist(ctx context.Context, userId uint) ([]response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishlist", ctx, userId)
	ret0, _ := ret[0].([]response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishlist indicates an expected call of GetWishlist.
func (mr *MockUserServiceMockRecorder) GetWishlist(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishlist", reflect.TypeOf((*MockUserService)(nil).GetWishlist), ctx, userId)
}

// Profile mocks base method.
func (m *MockUserService) Profile(ctx context.Context, userId uint) (response.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Profile", ctx, userId)
	ret0, _ := ret[0].(response.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Profile indicates an expected call of Profile.
func (mr *MockUserServiceMockRecorder) Profile(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Profile", reflect.TypeOf((*MockUserService)(nil).Profile), ctx, userId)
}

// RemoveCartItem mocks base method.
func (m *MockUserService) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCartItem", ctx, DelCartItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCartItem indicates an expected call of RemoveCartItem.
func (mr *MockUserServiceMockRecorder) RemoveCartItem(ctx, DelCartItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCartItem", reflect.TypeOf((*MockUserService)(nil).RemoveCartItem), ctx, DelCartItem)
}

// SaveCartItem mocks base method.
func (m *MockUserService) SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCartItem", ctx, addToCart)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCartItem indicates an expected call of SaveCartItem.
func (mr *MockUserServiceMockRecorder) SaveCartItem(ctx, addToCart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCartItem", reflect.TypeOf((*MockUserService)(nil).SaveCartItem), ctx, addToCart)
}

// UpdateAddress mocks base method.
func (m *MockUserService) UpdateAddress(ctx context.Context, address request.AddressPatchReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserServiceMockRecorder) UpdateAddress(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserService)(nil).UpdateAddress), ctx, address)
}

// UpdateCart mocks base method.
func (m *MockUserService) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, cartUpadates)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockUserServiceMockRecorder) UpdateCart(ctx, cartUpadates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockUserService)(nil).UpdateCart), ctx, cartUpadates)
}
