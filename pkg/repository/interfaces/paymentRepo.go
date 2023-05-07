package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type PaymentRepository interface {
	GetAllPaymentOptions(ctx context.Context) (PaymentOptions []response.PaymentOptionResp, err error)
}
