package repository

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AuthDatabase struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) interfaces.AuthRepository {
	return &AuthDatabase{DB: DB}
}
