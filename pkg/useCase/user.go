package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
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
