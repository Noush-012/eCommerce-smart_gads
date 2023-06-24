package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSaveUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create psql instance : %v", err)
	}
	err = gormDB.Statement.Error
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}

	// Create a new instance of the AuthDatabase with the mock DB
	authDB := &AuthDatabase{DB: gormDB}

	// Set up the expected query and its arguments
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password, created_at) 
			  VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)`
	createdAt := time.Now()
	user := domain.Users{
		UserName:  "testuser",
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "test@example.com",
		Phone:     "1234567890",
		Password:  "password",
	}

	// Set up the expected query execution and its result
	mock.ExpectQuery(query).
		WithArgs(user.UserName, user.FirstName, user.LastName, user.Age, user.Email, user.Phone, user.Password, createdAt)

	// Call the SaveUser function
	err = authDB.SaveUser(context.Background(), user)

	// Verify the expectations
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
