//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/config"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/db"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository"
	usecase "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase"
	"github.com/google/wire"
)

func InitiateAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnToDB,

		// Repositories
		repository.NewAdminRepository,
		repository.NewProductRepository,
		repository.NewUserRepository,
		repository.NewPaymentRepository,
		repository.NewOrderRepository,

		// Usecases
		usecase.NewAdminService,
		usecase.NewProductUseCase,
		usecase.NewUserUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewOrderUseCase,

		// Handlers
		handler.NewAdminHandler,
		handler.NewProductHandler,
		handler.NewUserHandler,
		handler.NewPaymentHandler,
		handler.NewOrderHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
