package usecase

import (
	"context"
	"errors"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminService struct {
	adminRepository   interfaces.AdminRepository
	orderRepositoty   interfaces.OrderRepository
	PaymentRepository interfaces.PaymentRepository
}

func NewAdminService(repo interfaces.AdminRepository, orderRepo interfaces.OrderRepository, PaymentRepo interfaces.PaymentRepository) service.AdminService {
	return &adminService{adminRepository: repo,
		orderRepositoty:   orderRepo,
		PaymentRepository: PaymentRepo}
}
func (a *adminService) Signup(c context.Context, admin domain.Admin) error {
	if dbAdmin, err := a.adminRepository.GetAdmin(c, admin); err != nil {
		return err
	} else if dbAdmin.ID != 0 {
		return errors.New("user already exists")
	}

	// Hash password
	PassHash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return errors.New("hashing failed")
	}

	admin.Password = string(PassHash)
	return a.adminRepository.SaveAdmin(c, admin)

}

func (a *adminService) Login(c context.Context, admin domain.Admin) (domain.Admin, error) {
	// Check admin exist in db
	dbAdmin, err := a.adminRepository.GetAdmin(c, admin)
	if err != nil {
		return admin, err
	}
	// compare password with hash password
	if bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password)) != nil {
		return admin, errors.New("wrong password")
	}
	return dbAdmin, nil

}

func (a *adminService) GetAllUser(c context.Context, page request.ReqPagination) (users []response.UserRespStrcut, err error) {
	users, err = a.adminRepository.GetAllUser(c, page)

	if err != nil {
		return nil, err
	}

	// if no error then copy users details to an array response struct
	var response []response.UserRespStrcut
	copier.Copy(&response, &users)

	return response, nil
}

// to block or unblock a user
func (a *adminService) BlockUser(c context.Context, userID uint) error {

	return a.adminRepository.BlockUser(c, userID)
}

// Get user order history
func (a *adminService) GetUserOrderHistory(c context.Context, userId uint) (orderHistory []domain.ShopOrder, err error) {
	orderHistory, err = a.adminRepository.GetUserOrderHistory(c, userId)
	if err != nil {
		return orderHistory, err
	}
	return orderHistory, err
}

// Generate sales report
func (a *adminService) SalesReport(c context.Context, daterange request.DateRange) (salesReport []domain.SalesReport, err error) {

	salesReport, err = a.adminRepository.GenerateSalesReport(c, daterange)
	if err != nil {
		return salesReport, err
	}
	return salesReport, nil
}

func (a *adminService) GetAllReturnOrders(c context.Context) {

}

func (o *adminService) UpdateDeliveryStatus(c context.Context, UpdateData request.UpdateStatus) error {

	err := o.orderRepositoty.UpdateDeliveryStatus(c, UpdateData)
	if err != nil {
		return err
	}
	// get payment method for order id
	order, err := o.orderRepositoty.GetOrderByOrderId(c, UpdateData.OrderId)
	if err != nil {
		return err
	}
	// update payment data as paid for COD orders if status delivered
	if UpdateData.StatusId == 2 && order.PaymentMethod == "Cash on delivery COD" {
		// ID 2 is for status "PAID"
		err = o.PaymentRepository.UpdatePaymentStatus(c, 2, UpdateData.OrderId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *adminService) ApproveReturnOrder(c context.Context, data request.ApproveReturnRequest) error {
	// get payment data
	// ID 2 is for status "Paid"
	payment, err := o.PaymentRepository.GetPaymentDataByOrderId(c, data.OrderID)

	if err != nil {
		return err
	}

	data.OrderTotal = payment.OrderTotal
	err = o.adminRepository.ApproveReturnOrder(c, data)
	if err != nil {
		return err
	}
	return nil
}
