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
	repo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) service.UserService {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) SignUp(ctx context.Context, user domain.Users) error {
	// Check if user already exist
	usr, err := u.repo.FindUser(ctx, user)
	if err != nil {
		return err
	}

	if usr.ID == 0 {
		// Hash user password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Hashing failed")
			return err
		}
		user.Password = string(hashedPass)

		// Save user if not exist
		err = u.repo.SaveUser(ctx, user)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v user already exists", usr.UserName)
	}

	return nil
}

func (u *UserUseCase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	usr, err := u.repo.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if usr.ID == 0 {
		return user, errors.New("user not exist")
	}
	// check password with hashed pass
	if bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Password)) != nil {
		return user, errors.New("password incorrect")
	}

	return usr, nil

}
func (u *UserUseCase) OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	usr, err := u.repo.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if usr.ID == 0 {
		return user, errors.New("user not exist")
	}
	return usr, nil
}
