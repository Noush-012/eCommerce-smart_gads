package usecase

import (
	"context"
	"errors"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type UserUseCase struct {
	userRepository  interfaces.UserRepository
	orderRepository interfaces.OrderRepository
}

func NewUserUseCase(repo interfaces.UserRepository, orderRepo interfaces.OrderRepository) service.UserService {
	return &UserUseCase{userRepository: repo,
		orderRepository: orderRepo}
}

func (u *UserUseCase) SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error {
	if err := u.userRepository.SavetoCart(ctx, addToCart); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {
	cartItems, err := u.userRepository.GetCartItemsbyUserId(ctx, page, userID)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (u *UserUseCase) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {
	if err := u.userRepository.UpdateCart(ctx, cartUpadates); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error {
	if err := u.userRepository.RemoveCartItem(ctx, DelCartItem); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) Profile(ctx context.Context, userId uint) (profile response.Profile, err error) {
	user, err := u.userRepository.GetUserbyID(ctx, userId)
	defaultAddress, err1 := u.userRepository.GetDefaultAddress(ctx, userId)
	page := request.ReqPagination{
		PageNumber: 1,
		Count:      5,
	}
	profile = response.Profile{
		ID:             user.ID,
		UserName:       user.UserName,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Age:            user.Age,
		Email:          user.Email,
		Phone:          user.Phone,
		DefaultAddress: defaultAddress,
	}
	orderHistory, err2 := u.orderRepository.GetOrderHistory(ctx, page, userId)
	profile.OrderHistory = append(profile.OrderHistory, orderHistory...)
	if err = errors.Join(err, err1, err2); err != nil {
		return profile, err
	}

	return profile, nil
}

func (u *UserUseCase) Addaddress(ctx context.Context, address request.Address) error {
	if err := u.userRepository.SaveAddress(ctx, address); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) UpdateAddress(ctx context.Context, address request.AddressPatchReq) error {
	if err := u.userRepository.UpdateAddress(ctx, address); err != nil {
		return err
	}
	return nil

}

func (u *UserUseCase) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	if err := u.userRepository.DeleteAddress(ctx, userID, addressID); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error) {
	address, err = u.userRepository.GetAllAddress(ctx, userId)
	if err != nil {
		return address, err
	}
	return address, nil
}

func (u *UserUseCase) AddToWishlist(ctx context.Context, wishlistData request.AddToWishlist) error {
	err := u.userRepository.AddToWishlist(ctx, wishlistData)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) GetWishlist(ctx context.Context, userId uint) (wishlist []response.Wishlist, err error) {
	wishlist, err = u.userRepository.GetWishlist(ctx, userId)
	if err != nil {
		return wishlist, err
	}
	return wishlist, nil
}

func (u *UserUseCase) DeleteFromWishlist(ctx context.Context, productId, userId uint) error {
	err := u.userRepository.DeleteFromWishlist(ctx, productId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) GetWalletHistory(ctx context.Context, userId uint) (wallet []domain.Wallet, err error) {
	wallet, err = u.userRepository.GetWalletHistory(ctx, userId)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}
