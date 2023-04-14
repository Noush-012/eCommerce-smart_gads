package db

import (
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
func ConnToDB() {
	dsn := viper.GetString("DATABASE")

	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Faild to Connect Database")
		return
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
}
