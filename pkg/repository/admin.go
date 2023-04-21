package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB: db}
}

func (a *adminDatabase) GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	query := `SELECT * FROM admins WHERE email = ? OR user_name = ?`
	if a.DB.Raw(query, admin.Email, admin.UserName).Scan(&admin).Error != nil {

		return admin, errors.New("failed to find admin")
	}
	return admin, nil
}

func (a *adminDatabase) SaveAdmin(ctx context.Context, admin domain.Admin) error {
	query := `INSERT INTO admins(user_name, email, password, created_at) VALUES ($1, $2, $3, $4)`
	createdAt := time.Now()
	if a.DB.Exec(query, admin.UserName, admin.Email, admin.Password, createdAt).Error != nil {
		return errors.New("failed to save admin")
	}
	return nil
}
