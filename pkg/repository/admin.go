package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB: db}
}

func (a *adminDatabase) GetAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	query := `SELECT * FROM admins WHERE email = ? OR user_name = ?`
	if a.DB.Raw(query, admin.Email, admin.UserName).Scan(&admin).Error != nil {

		return admin, errors.New("failed to find admin")
	}
	return admin, nil
}

func (a *adminDatabase) SaveAdmin(ctx context.Context, admin domain.Admin) error {
	query := `INSERT INTO admins(user_name, email, password, created_at) VALUES ($1, $2, $3, $4)`
	createdAt := time.Now()
	if a.DB.Exec(query, admin.UserName, admin.Email, admin.Password, createdAt).Error != nil {
		return errors.New("failed to save admin")
	}
	return nil
}

func (a *adminDatabase) BlockUser(ctx context.Context, userID uint) error {
	// Check user if valid or not
	var user domain.Users
	query := `SELECT * FROM users WHERE id=?`
	a.DB.Raw(query, userID).Scan(&user)
	if user.Email == "" { // check user email with user ID
		return errors.New("invalid user id user doesn't exist")
	}

	query = `UPDATE users SET block_status = $1 WHERE id = $2`
	if a.DB.Exec(query, !user.BlockStatus, userID).Error != nil {
		return fmt.Errorf("failed to update user block_status to %v", !user.BlockStatus)
	}
	return nil
}

func (a *adminDatabase) GetAllUser(ctx context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	query := `SELECT * FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	err = a.DB.Raw(query, limit, offset).Scan(&users).Error

	return users, err
}

func (a *adminDatabase) GetUserOrderHistory(c context.Context, userId uint) (orderHistory []domain.ShopOrder, err error) {
	query := `SELECT *
	FROM shop_orders so 
	WHERE user_id = $1`
	if err := a.DB.Raw(query, userId).Scan(&orderHistory).Error; err != nil {
		return orderHistory, err
	}
	return orderHistory, nil
}

func (a *adminDatabase) ChangeOrderStatus(c context.Context, UpdateData request.UpdateOrderStatus) (order response.ShopOrder, err error) {

	query := `UPDATE shop_orders
	SET order_status_id = $1
	WHERE user_id = $2 AND id = $3`
	if err := a.DB.Exec(query, UpdateData.StatusId, UpdateData.UserId, UpdateData.OrderId).Error; err != nil {
		return order, err
	}
	query = `SELECT so.id, so.order_date,os.status, so.order_total, po.name AS payment_type, pm.name AS payment_method, ps.status AS payment_status
	FROM shop_orders so
	JOIN order_statuses os ON os.id = so.order_status_id
	JOIN payment_options po ON so.payment_option_id = po.id
	JOIN payment_methods pm ON pm.id = so.payment_method_id 
	JOIN payment_statuses ps on ps.id = so.payment_status_id
	WHERE so.user_id = $1 AND so.id = $2`
	if err := a.DB.Raw(query, UpdateData.UserId, UpdateData.OrderId).Scan(&order).Error; err != nil {
		return order, err
	}

	return order, err
}
