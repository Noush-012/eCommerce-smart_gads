package main

import (
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/db"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/initializer"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadViper()
}

func main() {
	db, err := db.ConnToDB()
	if err != nil {
		panic("Database connection failed")
	}

	var repo interfaces.UserRepository = repository.NewUserRepository(db)
	Uh := handler.NewUserHandler(repo)

	g := gin.New()

	g.POST("/signup", Uh.UserSignup)
	g.Run() // listen and serve on 0.0.0.0:8080

}
