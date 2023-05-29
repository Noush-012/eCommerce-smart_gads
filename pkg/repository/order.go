package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB              *gorm.DB
	PaymentDatabase interfaces.PaymentRepository
	couponDatabase  interfaces.CouponRepository
	userDatabase    interfaces.UserRepository
}

func NewOrderRepository(db *gorm.DB, paymentRepo interfaces.PaymentRepository, couponRepo interfaces.CouponRepository,
	userRepo interfaces.UserRepository) interfaces.OrderRepository {
	return &OrderDatabase{DB: db,
		PaymentDatabase: paymentRepo,
		couponDatabase:  couponRepo,
		userDatabase:    userRepo}
}

func (o *OrderDatabase) OrderStatus(ctx context.Context, id uint) (status string, err error) {
	query := `SELECT status AS order_status FROM order_statuses WHERE id = $1`
	if err := o.DB.Raw(query, id).Scan(&status).Error; err != nil {
		return status, err
	}
	return status, nil
}
func (o *OrderDatabase) GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error) {
	query := `SELECT id FROM carts WHERE user_id = $1`
	if err := o.DB.Raw(query, userId).Scan(&cartId).Error; err != nil {
		return cartId, err
	}
	return cartId, nil
}

func (o *OrderDatabase) GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	// get cartID by user id
	cartID, err := o.GetCartIdByUserId(ctx, userID)
	if err != nil {
		return CartItems, err
	}
	// get cartItems with cartID
	query := `SELECT ci.product_item_id, p.name,pi.price, pi.discount_price,ci.quantity,pi.qty_in_stock AS qty_left,
	pi.stock_status AS stock_status,
	pi.discount_price * ci.quantity AS sub_total
		FROM cart_items ci
		JOIN product_items pi ON ci.product_item_id = pi.id
		JOIN products p ON pi.product_id = p.id
		WHERE cart_id = $1
		ORDER BY ci.created_at DESC LIMIT $2 OFFSET $3`
	if err := o.DB.Raw(query, cartID, limit, offset).Scan(&CartItems).Error; err != nil {
		return CartItems, err
	}
	return CartItems, nil
}

// CheckoutOrder checks out an order for a given user with the provided coupon code.
func (o *OrderDatabase) CheckoutOrder(ctx context.Context, userId uint, couponCode string) (checkOut response.CartResp, err error) {
	var page request.ReqPagination
	page.PageNumber = 1
	page.Count = 50

	// Get cart items for the user
	cartItems, err := o.GetCartItemsbyUserId(ctx, page, userId)
	if err != nil {
		return checkOut, err
	}

	count := 0
	for _, v := range cartItems {
		if v.ProductItemID != 0 {
			count++

			checkOut.TotalPrice += float64(v.SubTotal)
			checkOut.TotalQty += v.Quantity
			checkOut.DiscountAmount += float64(v.Price - v.DiscountPrice)
		}
	}
	checkOut.TotalProductItems = uint(count)

	// Apply coupon
	AppliedCoupon, err := o.couponDatabase.ApplyCoupon(ctx, utils.ApplyCoupon{
		CouponCode: couponCode,
		UserId:     userId,
		TotalPrice: checkOut.TotalPrice,
	})
	if err != nil {
		return checkOut, err
	}
	checkOut.AppliedCouponCode = AppliedCoupon.CouponCode
	checkOut.AppliedCouponID = AppliedCoupon.CouponId
	checkOut.CouponDiscount = AppliedCoupon.CouponDiscount
	checkOut.FinalPrice = uint(AppliedCoupon.FinalPrice)

	// Get default address for the user
	query := `SELECT a.id, a.house, a.address_line1, a.address_line2, a.city, a.state, a.zip_code, a.country, a.is_default
	FROM addresses a
	WHERE a.is_default = true AND a.user_id = $1`
	var address response.Address
	if err := o.DB.Raw(query, userId).Scan(&address).Error; err != nil {
		return checkOut, err
	}
	checkOut.DefaultShipping = address

	if checkOut.TotalProductItems == 0 {
		return checkOut, errors.New("no items in cart to checkout")
	}

	return checkOut, nil
}

func (o *OrderDatabase) PlaceCODOrder(ctx context.Context, userId, PaymentMethodID uint, couponCode string) (OrderId uint, err error) {
	tnx := o.DB.Begin()
	checkOut, err := o.CheckoutOrder(ctx, userId, couponCode)
	if err != nil {
		return OrderId, err
	}
	if checkOut.TotalProductItems == 0 {
		return OrderId, errors.New("cart is empty, please add products")
	}

	order := domain.ShopOrder{
		UserID:           userId,
		OrderTotal:       checkOut.FinalPrice,
		ShippingID:       checkOut.DefaultShipping.ID,
		PaymentMethodID:  PaymentMethodID,
		OrderStatusID:    1,
		DeliveryStatusID: 1,
		CouponID:         0,
	}
	OrderId, err = o.SaveOrder(ctx, order)
	if err != nil {
		tnx.Rollback()
		return OrderId, err
	}
	// save order item
	for _, v := range checkOut.CartItems {
		if v.ProductItemID != 0 {
			orderLine := domain.OrderLine{
				ProductItemID: v.ProductItemID,
				ShopOrderID:   OrderId,
				Qty:           v.Quantity,
				Price:         v.DiscountPrice,
			}
			if err := o.SaveOrderLine(ctx, orderLine); err != nil {
				tnx.Rollback()
				return OrderId, err
			}

		}
	}
	// Save payment details
	if err := o.PaymentDatabase.SavePaymentData(ctx, domain.PaymentDetails{
		OrderID:         OrderId,
		OrderTotal:      checkOut.FinalPrice,
		PaymentMethodID: PaymentMethodID,
		PaymentStatusID: 1, // set payment status as pending ID 1 = "Pending"
		PaymentRef:      "",
		UpdatedAt:       time.Now(),
	}); err != nil {
		tnx.Rollback()
		return OrderId, err
	}

	if err = tnx.Commit().Error; err != nil {
		tnx.Rollback()
		return OrderId, err
	}
	return OrderId, nil

}

func (o *OrderDatabase) ClearUserCart(ctx context.Context, userId uint) error {
	// get cart id from user id
	cartId, err := o.GetCartIdByUserId(ctx, userId)
	if err != nil {
		return err
	}
	query := `DELETE from cart_items WHERE cart_id = $1`
	if err := o.DB.Exec(query, cartId).Error; err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	query := `SELECT so.id, so.order_date, os.status AS order_status, ds.status AS delivery_status,
	so.order_total, pm.payment_method, ps.status AS payment_status
	FROM shop_orders so
	JOIN order_statuses os ON os.id = so.order_status_id
	JOIN delivery_statuses ds ON ds.id = so.delivery_status_id
	JOIN payment_methods pm ON pm.id = so.payment_method_id
	LEFT JOIN payment_details pd ON pd.order_id = so.id
	LEFT JOIN payment_statuses ps ON ps.id = pd.payment_status_id
	WHERE so.user_id = $1 ORDER BY so.order_date DESC LIMIT $2 OFFSET $3;`
	if err := o.DB.Raw(query, userId, limit, offset).Scan(&orderHisory).Error; err != nil {
		return orderHisory, err
	}
	return orderHisory, nil

}

func (o *OrderDatabase) GetOrderByOrderId(ctx context.Context, OrderId uint) (orderData response.ShopOrder, err error) {
	query := `SELECT so.id, so.order_date,os.status AS order_status, so.order_total, pm.payment_method,
	so.shipping_id
	FROM shop_orders so
	JOIN order_statuses os ON os.id = so.order_status_id
	JOIN payment_methods pm ON pm.id = so.payment_method_id
	WHERE so.id = ?`
	if err := o.DB.Raw(query, OrderId).Scan(&orderData).Error; err != nil {
		return orderData, err

	}
	return orderData, nil
}

func (o *OrderDatabase) SaveOrder(ctx context.Context, order domain.ShopOrder) (OrderId uint, err error) {
	query := `INSERT INTO shop_orders(user_id,order_date,order_total,shipping_id,order_status_id,
	payment_method_id, coupon_id,delivery_status_id, delivery_updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING ID`
	order.OrderDate = time.Now()
	if order.PaymentMethodID == 1 {

		order.OrderStatusID = 1 // set default ID 1 is for placed for COD orders
	}
	if err := o.DB.Raw(query, order.UserID, order.OrderDate, order.OrderTotal, order.ShippingID, order.OrderStatusID,
		order.PaymentMethodID, order.CouponID, order.DeliveryStatusID, order.DeliveryUpdatedAt).Scan(&OrderId).Error; err != nil {
		return 0, err
	}
	return OrderId, nil
}

func (o *OrderDatabase) SaveOrderLine(ctx context.Context, orderLine domain.OrderLine) error {
	query := `INSERT INTO order_lines(product_item_id,shop_order_id, qty, price)
	VALUES ($1,$2,$3,$4)`
	if err := o.DB.Exec(query, orderLine.ProductItemID, orderLine.ShopOrderID, orderLine.Qty, orderLine.Price).Error; err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) ChangeOrderStatus(c context.Context, UpdateData request.UpdateStatus) error {

	query := `UPDATE shop_orders
	SET order_status_id = $1
	WHERE user_id = $2 AND id = $3`
	if err := o.DB.Exec(query, UpdateData.StatusId, UpdateData.UserId, UpdateData.OrderId).Error; err != nil {
		return err
	}
	//

	return nil
}

func (o *OrderDatabase) GetDeliveryDate(c context.Context, orderId uint) (time.Time, error) {
	var deliveryDate time.Time
	// id 2 for status "delivered"
	query := `SELECT delivery_updated_at FROM shop_orders WHERE delivery_status_id = 2
	AND id = $1`
	if err := o.DB.Raw(query, orderId).Scan(&deliveryDate).Error; err != nil {
		return deliveryDate, err
	}
	return deliveryDate, nil
}

func (o *OrderDatabase) SaveReturnRequest(c context.Context, data request.ReturnRequest) error {
	query := `INSERT INTO returns (shop_order_id, reason, requested_at)
	VALUES ($1, $2, $3)`
	requested_at := time.Now()
	if err := o.DB.Exec(query, data.OrderID, data.Reason, requested_at).Error; err != nil {
		return err
	}
	return nil

}

func (o *OrderDatabase) GetAllPendingReturnOrder(c context.Context, page request.ReqPagination) (ReturnRequests []response.ReturnRequests, err error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	query := `SELECT r.id AS return_id, so.user_id, r.requested_at, so.order_date,so.delivery_updated_at AS delivered_at,r.shop_order_id AS order_id,pm.payment_method,ps.status AS payment_status,
	r.reason,so.order_total, r.is_approved
	FROM returns r
	JOIN shop_orders so ON so.id = r.shop_order_id
	JOIN payment_methods pm ON pm.id = so.payment_method_id
	JOIN payment_details pd ON pd.order_id = so.id
	JOIN payment_statuses ps ON ps.id = pd.payment_status_id
	WHERE r.is_approved = false
	ORDER BY r.requested_at ASC LIMIT $1 OFFSET $2`
	err = o.DB.Raw(query, limit, offset).Scan(&ReturnRequests).Error
	if err != nil {
		return ReturnRequests, err
	}
	return ReturnRequests, nil
}

func (o *OrderDatabase) UpdateDeliveryStatus(c context.Context, UpdateData request.UpdateStatus) error {
	query := `UPDATE shop_orders
	SET delivery_status_id = $1, delivery_updated_at = $2
	WHERE id = $3`
	deliveryDate := time.Now()
	err := o.DB.Exec(query, UpdateData.StatusId, deliveryDate, UpdateData.OrderId).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) OrderCancellation(c context.Context, data request.CancelOrder) error {
	// check payment status
	payment, err := o.PaymentDatabase.GetPaymentDataByOrderId(c, data.OrderID)
	if err != nil {
		return err
	}
	// set order status cancel
	query := `UPDATE shop_orders
	SET order_status_id = (
	  SELECT ID FROM order_statuses WHERE status = 'cancelled'
	)
	WHERE id = $1;`
	err = o.DB.Exec(query, data.OrderID).Error
	if err != nil {
		return err
	}
	if payment.PaymentStatusID == 2 { // checking the status "Paid"

		// refund amount to the wallet
		o.userDatabase.CreditUserWallet(c, domain.Wallet{
			UserID:  data.UserID,
			Balance: float64(payment.OrderTotal),
			Remark:  "Cancelled order",
		})
	}

	return nil
}
