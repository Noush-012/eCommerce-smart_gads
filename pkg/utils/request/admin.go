package request

import "time"

type UpdateStatus struct {
	UserId   uint `json:"user_id" binding:"required,numeric"`
	StatusId uint `json:"status_id" binding:"required,numeric"`
	OrderId  uint `json:"order_id" binding:"required,numeric"`
}

type DateRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type UpdateDeliveryStatus struct {
	UserId   uint `json:"user_id" binding:"required,numeric"`
	StatusId uint `json:"status_id" binding:"required,numeric"`
	OrderId  uint `json:"order_id" binding:"required,numeric"`
}
