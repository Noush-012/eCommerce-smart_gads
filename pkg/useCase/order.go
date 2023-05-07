package usecase

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderUseCase struct {
	OrderRepository interfaces.OrderRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository) service.OrderService {
	return &OrderUseCase{OrderRepository: repo}
}
func (o *OrderUseCase) CheckoutOrder(ctx context.Context, userId uint) (checkOut response.CartResp, err error) {
	checkOut, err = o.OrderRepository.CheckoutOrder(ctx, userId)
	if err != nil {
		return checkOut, err
	}
	return checkOut, nil
}
func (o *OrderUseCase) PlaceOrderByCOD(ctx context.Context, userId uint) (shopOrder response.ShopOrder, err error) {
	shopOrder, err = o.OrderRepository.PlaceCODOrder(ctx, userId)
	if err != nil {
		return shopOrder, err
	}
	// Clear user cart
	if err := o.OrderRepository.ClearUserCart(ctx, userId); err != nil {
		return shopOrder, err
	}
	return shopOrder, nil
}
