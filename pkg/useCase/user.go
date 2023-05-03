package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/verify"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) service.UserService {
	return &UserUseCase{userRepository: repo}
}

func (u *UserUseCase) SignUp(ctx context.Context, user domain.Users) error {
	// Check if user already exist
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return err
	}

	if DBUser.ID == 0 {
		// Hash user password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Hashing failed")
			return err
		}
		user.Password = string(hashedPass)

		// Save user if not exist
		err = u.userRepository.SaveUser(ctx, user)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v user already exists", DBUser.UserName)
	}

	return nil
}

func (u *UserUseCase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	// Check if the user blocked by admin
	if DBUser.BlockStatus {
		return user, errors.New("user blocked by admin")
	}

	if _, err := verify.TwilioSendOTP("+91" + DBUser.Phone); err != nil {
		// response := response.ErrorResponse(500, "failed to send otp", err.Error(), nil)
		// c.JSON(http.StatusInternalServerError, response)
		return user, fmt.Errorf("failed to send otp %v",
			err)
	}
	// check password with hashed pass
	if bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password)) != nil {
		return user, errors.New("password incorrect")
	}

	return DBUser, nil

}
func (u *UserUseCase) OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	return DBUser, nil
}

func (u *UserUseCase) SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error {
	if err := u.userRepository.SavetoCart(ctx, addToCart); err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {
	cartItems, err := u.userRepository.GetCartItemsbyCartId(ctx, page, userID)
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

func (u *UserUseCase) Profile(ctx context.Context, userId uint) (domain.Users, error) {
	user, err := u.userRepository.GetUserbyID(ctx, userId)
	if err != nil {
		return user, err
	}
	return user, nil
}
