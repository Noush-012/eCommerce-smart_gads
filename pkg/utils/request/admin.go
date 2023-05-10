package request

type UpdateOrderStatus struct {
	UserId   uint `json:"user_id" binding:"required,numeric"`
	StatusId uint `json:"status_id" binding:"required,numeric"`
	OrderId  uint `json:"order_id" binding:"required,numeric"`
}
