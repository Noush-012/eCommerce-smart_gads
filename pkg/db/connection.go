package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB variables
var (
	DB  *gorm.DB
	err error
)

// To connect database
func ConnToDB() (*gorm.DB, error) {
	dsn := viper.GetString("DATABASE")

	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Failed to Connect Database")
		return nil, errors.New("failed to connect database")
	}
	fmt.Println("Successfully Connected to database")

	// Migrate models
	err := DB.AutoMigrate(
		// Users
		domain.Users{},
		domain.Admin{},

		// Product tables
		domain.Product{},
		domain.ProductVarient{},
		domain.Brand{},
	)
	if err != nil {
		log.Fatal("Migration failed")
		return nil, nil
	}
	fmt.Println("DB migration success")
	return DB, nil

}
