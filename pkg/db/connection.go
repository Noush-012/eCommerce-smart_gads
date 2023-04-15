package db

import (
	"errors"
	"fmt"

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
	errUSr := DB.AutoMigrate(domain.Users{})
	errAdm := DB.AutoMigrate(domain.Admin{})
	if errUSr == nil || errAdm == nil {
		fmt.Println("Migration success!")
	} else {
		fmt.Println("Migration Failed!", errUSr, errAdm)
	}
	return DB, nil
}
