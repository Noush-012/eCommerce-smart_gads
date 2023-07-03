package repository

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("SQL mock failed %v ", err)
	}
	defer db.Close()
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to create psql instance ")
	}
	err = gormDb.Statement.Error
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}
	user := domain.Users{
		ID:       1,
		Email:    "test@xyz.com",
		Phone:    "1234567890",
		UserName: "test",
	}
	expectedError := errors.New("failed to get user")
	authDb := &userDatabase{DB: gormDb}
	mock.ExpectQuery(
		`SELECT * FROM users WHERE id = $1 OR email = $2 OR phone = $3 OR user_name = $4`).
		WithArgs(user.ID, user.Email, user.Phone, user.UserName).
		WillReturnRows(sqlmock.NewRows([]string{"ID", "UserName", "FirstName", "LastName", "Age", "Email", "Phone"}))
	result, resultErr := authDb.FindUser(context.Background(), user)
	assert.Equal(t, expectedError, resultErr)
	fmt.Println(result)

}
