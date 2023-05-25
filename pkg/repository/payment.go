package repository

import (
	"context"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
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
func (p *PaymentDatabase) GetPaymentMethodByID(ctx context.Context, id uint) (payMethod string, err error) {
	query := `SELECT payment_method FROM payment_methods WHERE id = $1`
	if err := p.DB.Raw(query, id).Scan(&payMethod).Error; err != nil {
		return payMethod, err
	}
	return payMethod, nil
}
func (p *PaymentDatabase) GetPaymentMethodByName(ctx context.Context, name string) (payMethod domain.PaymentMethod, err error) {
	query := `SELECT payment_method,id FROM payment_methods WHERE payment_method LIKE ?`

	if err := p.DB.Raw(query, name).Scan(&payMethod).Error; err != nil {
		return payMethod, err
	}
	return payMethod, nil
}

func (p *PaymentDatabase) SavePaymentData(ctx context.Context, paymentData domain.PaymentDetails) error {

	query := `INSERT INTO payment_details (order_id, order_total,payment_method_id,payment_status_id,
		payment_ref, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	updatedAt := time.Now()

	if err := p.DB.Exec(query, paymentData.OrderID, paymentData.OrderTotal, paymentData.PaymentMethodID,
		paymentData.PaymentStatusID, paymentData.PaymentRef, updatedAt).Error; err != nil {
		return err
	}
	return nil
}

func (p *PaymentDatabase) GetPaymentStatusByOrderId(ctx context.Context, orderId uint) (ok bool, err error) {
	var status string
	query := `SELECT ps.status
	FROM payment_details pd 
	JOIN payment_statuses ps ON ps.id = pd.payment_status_id
	WHERE order_id = ?`
	if err := p.DB.Raw(query, orderId).Scan(&status).Error; err != nil {
		return false, err
	}
	if status == "Paid" {

		return true, nil
	}
	return false, nil

}
func (p *PaymentDatabase) UpdatePaymentStatus(ctx context.Context, statusId, orderId uint) error {
	query := `UPDATE payment_details SET payment_status_id = $1 WHERE order_id = $2`
	err := p.DB.Exec(query, statusId, orderId).Error
	if err != nil {
		return err
	}
	return nil
}
