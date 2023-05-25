package usecase

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type PaymentUseCase struct {
	PaymentRepository interfaces.PaymentRepository
	OrderRepository   interfaces.OrderRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository, orderPepo interfaces.OrderRepository) service.PaymentService {
	return &PaymentUseCase{PaymentRepository: repo,
		OrderRepository: orderPepo}
}

func (p *PaymentUseCase) GetAllPaymentOptions(ctx context.Context) (PaymentOptions []response.PaymentOptionResp, err error) {
	if PaymentOptions, err = p.PaymentRepository.GetAllPaymentOptions(ctx); err != nil {
		return PaymentOptions, err
	}
	return PaymentOptions, nil
}

func (p *PaymentUseCase) GetPaymentMethodByID(ctx context.Context, id uint) (payMethod string, err error) {
	payMethod, err = p.PaymentRepository.GetPaymentMethodByID(ctx, id)
	if err != nil {
		return payMethod, err
	}
	return payMethod, nil
}
func (p *PaymentUseCase) UpdatePaymentStatus(ctx context.Context, statusId, orderId uint) error {
	return nil
}

func (p *PaymentUseCase) SavePaymentDetails(ctx context.Context, paymentData domain.PaymentDetails) error {
	// get order data with orderId
	order, err := p.OrderRepository.GetOrderByOrderId(ctx, paymentData.OrderID)
	if err != nil {
		return err
	}
	paymentData.OrderTotal = uint(order.OrderTotal)

	if err := p.PaymentRepository.SavePaymentData(ctx, paymentData); err != nil {
		return err
	}
	return nil
}
