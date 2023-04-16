package main

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/routes"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/db"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/initializer"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository"
	usecase "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/verify"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadViper()
	verify.SetClient()
}

func main() {
	db, err := db.ConnToDB()
	if err != nil {
		panic("Database connection failed")
	}

	newRepo := repository.NewUserRepository(db)
	newHandler := usecase.NewUserUseCase(newRepo)
	uH := handler.NewUserHandler(newHandler)

	g := gin.New()

	routes.UserRoutes(g, uH)
	g.Run() // listen and serve on 0.0.0.0:8080

}
