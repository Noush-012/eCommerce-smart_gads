package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSaveUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	// creating postgres instance with sqlMock
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create psql instance : %v", err)
	}
	defer db.Close()
	repo := NewAuthRepository(gormDB)

	// Test case 1 - Successful save
	user := domain.Users{
		UserName:  "Jose",
		FirstName: "Josekutty",
		LastName:  "Jose",
		Age:       22,
		Email:     "jose@abc.com",
		Phone:     "1234567890",
		Password:  "password",
	}
	// mock.ExpectQuery(`INSERT INTO users\(user_name, first_name, last_name, age, email, phone, password,created_at\)
	// VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8 \)`).
	mock.ExpectExec(`INSERT INTO users (user_name, first_name, last_name, age, email, phone, password, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`).
		WithArgs(user.UserName, user.FirstName, user.LastName, user.Age, user.Email, user.Phone, user.Password, time.Now())
	err = repo.SaveUser(context.Background(), user)
	fmt.Println("----------------", err)
	fmt.Println("----------------", t)
	assert.NoError(t, err)

	// Test case 2 - Duplicate username
	mock.ExpectExec(`INSERT INTO users (user_name, first_name, last_name, age, email, phone, password, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`).
		WithArgs(user.UserName, user.FirstName, user.LastName, user.Age, user.Email, user.Phone, user.Password, time.Now())
	err = repo.SaveUser(context.Background(), user)
	assert.Error(t, err)

	// user2 := domain.Users{
	// 	UserName:  "Jose",
	// 	FirstName: "Josekutty",
	// 	LastName:  "Jose",
	// 	Age:       22,
	// 	Email:     "",
	// 	Phone:     "1234567890",
	// 	Password:  "password",
	// }

	// mock.ExpectQuery(`INSERT INTO users\(user_name, first_name, last_name, age, email, phone, password,created_at\)
	// VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8 \)`).
	// 	WithArgs(user2.UserName, user2.FirstName, user2.LastName, user2.Age, user2.Email, user2.Phone, user2.Password, time.Now)
	// err = repo.SaveUser(context.TODO(), user2)
	// assert.Error(t, err)
}
