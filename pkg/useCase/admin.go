package usecase

import (
	"context"
	"errors"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
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
