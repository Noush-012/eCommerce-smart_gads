package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AuthDatabase struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) interfaces.AuthRepository {
	return &AuthDatabase{DB: DB}
}

func (i *AuthDatabase) SaveUser(ctx context.Context, user domain.Users) error {
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password,created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now()
	err := i.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age,
		user.Email, user.Phone, user.Password, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save user %s %s", user.UserName, err)
	}
	return nil
}
