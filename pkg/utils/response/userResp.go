package response

import "time"

type UserRespStrcut struct {
	ID          uint      `json:"id" copier:"must"`
	FirstName   string    `json:"first_name" copier:"must"`
	LastName    string    `json:"last_name" copier:"must"`
	Age         uint      `json:"age" copier:"must"`
	Email       string    `json:"email" copier:"must"`
	UserName    string    `json:"user_name" copire:"must"`
	Phone       string    `json:"phone" copier:"must"`
	BlockStatus bool      `json:"block_status" copier:"must"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// home page response
type ResUserHome struct {
	// Products []ResponseProduct `json:"products"`
	User UserRespStrcut `json:"user"`
}
