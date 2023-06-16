package domain

import (
	"time"

	"gorm.io/gorm"
)

// User model
type Users struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserName    string    `json:"user_name" gorm:"not null" binding:"required,min=3,max=15"`
	FirstName   string    `json:"first_name" gorm:"not null" binding:"required,min=2,max=40"`
	LastName    string    `json:"last_name" gorm:"not null" binding:"required,min=1,max=40"`
	Age         uint      `json:"age" gorm:"not null" binding:"required,numeric"`
	Email       string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	Phone       string    `json:"phone" gorm:"unique;not null" binding:"required,min=10,max=10"`
	Password    string    `json:"password" gorm:"not null" binding:"required"`
	BlockStatus bool      `json:"block_status" gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Address struct {
	ID           uint      `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserID       uint      `json:"-"`
	Users        Users     `gorm:"foreignKey:UserID" json:"-"`
	House        string    `json:"house" gorm:"not null"`
	AddressLine1 string    `json:"address_line_1" gorm:"not null" binding:"required,min=2,max=40"`
	AddressLine2 string    `json:"address_line_2" gorm:"not null" binding:"required,min=2,max=40"`
	City         string    `json:"city" gorm:"not null" binding:"required,min=2,max=20"`
	State        string    `json:"state" gorm:"not null" binding:"required,min=2,max=20"`
	ZipCode      string    `json:"zip_code" gorm:"not null" binding:"required,min=2,max=10"`
	Country      string    `json:"country" gorm:"not null" binding:"required,min=2,max=20"`
	IsDefault    bool      `json:"-" gorm:"not null;default:false"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time
}

// cart model
type CartItem struct {
	ID            uint      `gorm:"primaryKey"`
	CartID        uint      `gorm:"not null"`
	ProductItemID uint      `gorm:"not null"`
	Quantity      uint      `gorm:"not null"`
	StockStatus   bool      `gorm:"not null;default:true"`
	Price         float64   `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Cart struct {
	ID     uint       `gorm:"primaryKey"`
	UserID uint       `gorm:"not null"`
	Items  []CartItem `gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Total  float64    `gorm:"default:0"`
}

type Wishlist struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	ProductItemID uint      `gorm:"not null"`
	Quantity      uint      `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

//	type Wallet struct {
//		ID           uint    `gorm:"primaryKey"`
//		UserID       uint    `gorm:"not null"`
//		Balance      float64 `gorm:"not null"`
//		Credit       uint
//		CreditRemark string
//		Debit        uint
//		DebitRemark  string
//		CreatedAt    time.Time
//		UpdatedAt    time.Time      `gorm:"not null"`
//		DeletedAt    gorm.DeletedAt `gorm:"index"`
//	}
type Wallet struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Balance   float64 `gorm:"not null"`
	Remark    string
	UpdatedAt time.Time
	CreatedAt time.Time
}
