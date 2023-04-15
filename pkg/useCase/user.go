package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo interfaces.UserRepository
}

func (u *UserUseCase) SignUp(ctx context.Context, user domain.Users) error {
	fmt.Println("Signup super")
	// Check the user already exist in DB
	ok := u.userRepo.FindByEmail(ctx, user.Email)
	if ok {
		fmt.Println("User already exists")
		return errors.New("user already exists")
	}
	// Hash user password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Println("Hashing failed")
		return errors.New("pass hashing failed")
	}
	user.Password = string(hashedPass)

	// Save user if not exist
	err = u.userRepo.SaveUser(ctx, user)
	if err != nil {
		return errors.New("failed to save user")
	}
	return nil
}
