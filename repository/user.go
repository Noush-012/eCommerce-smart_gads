package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func (i *userDatabase) SaveUser(ctx context.Context, user domain.Users) error {
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password,created_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	createdAt := time.Now()
	err := i.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age, user.Email, user.Phone, user.Password, createdAt)
	if err != nil {
		return fmt.Errorf("failed to save user %s", user.UserName)
	}
	return nil
}

func (i *userDatabase) FindByEmail(ctx context.Context, id string) bool {
	query := `SELECT * FROM users WHERE email = ?`
	var result string
	i.DB.Raw(query, id).Scan(&result)
	return result != ""
}
