package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type PaymentRepository interface {
	GetAllPaymentOptions(ctx context.Context) (PaymentOptions []response.PaymentOptionResp, err error)
	GetPaymentMethodByID(ctx context.Context, id uint) (payMethod string, err error)
	GetPaymentMethodByName(ctx context.Context, name string) (payMethod domain.PaymentMethod, err error)
	GetPaymentStatusByOrderId(ctx context.Context, orderId uint) (ok bool, err error)

	SavePaymentData(ctx context.Context, paymentData domain.PaymentDetails) error
}
