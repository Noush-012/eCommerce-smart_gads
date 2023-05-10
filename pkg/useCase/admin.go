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
	adminRepository interfaces.AdminRepository
}

func NewAdminService(repo interfaces.AdminRepository) service.AdminService {
	return &adminService{adminRepository: repo}
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

func (a *adminService) UpdateOrderStatus(c context.Context, UpdateData request.UpdateOrderStatus) (UpdatedOrder response.ShopOrder, err error) {
	UpdatedOrder, err = a.adminRepository.ChangeOrderStatus(c, UpdateData)
	if err != nil {
		return UpdatedOrder, err
	}
	return UpdatedOrder, nil
}
