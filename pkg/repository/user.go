package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	repo "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repo.UserRepository {
	return &userDatabase{DB: DB}
}

func (i *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Check any of the user details matching with db user list
	query := `SELECT * FROM users WHERE id = ? OR email = ? OR phone = ? OR user_name = ?`
	if err := i.DB.Raw(query, user.ID, user.Email, user.Phone, user.UserName).Scan(&user).Error; err != nil {
		return user, errors.New("failed to get user")
	}
	return user, nil
}

func (i *userDatabase) SaveUser(ctx context.Context, user domain.Users) error {
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password,created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now()
	err := i.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age,
		user.Email, user.Phone, user.Password, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save user %s", user.UserName)
	}
	return nil
}

func (i *userDatabase) SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error {
	return nil
}
