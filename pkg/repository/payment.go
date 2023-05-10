package repository

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"gorm.io/gorm"
)

type PaymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.PaymentRepository {
	return &PaymentDatabase{DB: db}
}

func (p *PaymentDatabase) GetAllPaymentOptions(ctx context.Context) (PaymentOptions []response.PaymentOptionResp, err error) {
	query := `SELECT * FROM payment_options`

	if err := p.DB.Raw(query).Scan(&PaymentOptions).Error; err != nil {
		return PaymentOptions, err
	}
	return PaymentOptions, nil

}
func (p *PaymentDatabase) GetPaymentOptionByID(ctx context.Context, id uint) (PayOption string, err error) {
	query := `SELECT name FROM payment_options WHERE id = $1`
	if err := p.DB.Raw(query, id).Scan(&PayOption).Error; err != nil {
		return PayOption, err
	}
	return PayOption, nil
}
