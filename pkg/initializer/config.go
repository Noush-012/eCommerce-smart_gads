package initializer

import (
	"fmt"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/db"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository"
	usecase "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase"
	"github.com/spf13/viper"
)

func LoadViper() {
	viper.SetConfigType("env")  // set the file type
	viper.SetConfigFile(".env") // set the file name and path
	err := viper.ReadInConfig() // read the config file
	if err != nil {             // handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

func InitializeApi() *handler.UserHandler {
	db, err := db.ConnToDB()
	if err != nil {
		panic("Database connection failed")
	}

	newRepo := repository.NewUserRepository(db)
	newHandler := usecase.NewUserUseCase(newRepo)
	uH := handler.NewUserHandler(newHandler)
	return uH
}
