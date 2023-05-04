package usecase

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type PaymentUseCase struct {
	PaymentRepository interfaces.PaymentRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository) service.PaymentService {
	return &PaymentUseCase{PaymentRepository: repo}
}

func (p *PaymentUseCase) GetAllPaymentOptions(ctx context.Context) (PaymentOptions []response.PaymentOptionResp, err error) {
	if PaymentOptions, err = p.PaymentRepository.GetAllPaymentOptions(ctx); err != nil {
		return PaymentOptions, err
	}
	return PaymentOptions, nil
}
