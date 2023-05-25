package request

type ReturnRequest struct {
	UserID  uint   `json:"-"`
	OrderID uint   `json:"order_id"`
	Reason  string `json:"reason"  binding:"omitempty,min=4,max=15"`
}
