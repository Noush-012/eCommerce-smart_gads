package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB: db}
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
	query := `SELECT ci.product_item_id, p.name,p.price,ci.price AS discount_price, 
	ci.quantity,pi.qty_in_stock AS qty_left, pi.stock_status AS stock_status, ci.price * ci.quantity AS sub_total
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

func (o *OrderDatabase) CheckoutOrder(ctx context.Context, userId uint) (checkOut response.CartResp, err error) {
	var page request.ReqPagination
	page.PageNumber = 1
	page.Count = 5

	// get cartItems
	cartItems, err := o.GetCartItemsbyUserId(ctx, page, userId)
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>", cartItems)
	if err != nil {
		return checkOut, err
	}

	count := 0
	for _, v := range cartItems {
		if v.ProductItemID != 0 {
			count++

			checkOut.TotalPrice += v.SubTotal
			checkOut.TotalQty += v.Quantity
			checkOut.DiscountAmount += v.Price - v.DiscountPrice
		}
	}
	checkOut.TotalProductItems = uint(count)
	// get default address
	query := `SELECT a.id,a.house,a.address_line1,a.address_line2,a.city,a.state,a.zip_code,a.country ,a.is_default
	FROM addresses a
	WHERE a.is_default = true AND a.user_id = $1`
	var address response.Address
	if err := o.DB.Raw(query, userId).Scan(&address).Error; err != nil {
		return checkOut, err
	}
	checkOut.DefaultShipping = address
	return checkOut, nil

}

func (o *OrderDatabase) PlaceCODOrder(ctx context.Context, userId uint, PaymentMethodID uint) (OrderId uint, err error) {
	tnx := o.DB.Begin()
	checkOut, err := o.CheckoutOrder(ctx, userId)
	if err != nil {
		return OrderId, err
	}
	if checkOut.TotalProductItems == 0 {
		return OrderId, errors.New("cart is empty, please add products")
	}

	order := domain.ShopOrder{
		UserID:          userId,
		OrderTotal:      float64(checkOut.TotalPrice),
		ShippingID:      checkOut.DefaultShipping.ID,
		PaymentMethodID: PaymentMethodID,
		CouponID:        0,
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
	LEFT JOIN payment_details pd ON pd.id = so.id
	LEFT JOIN payment_statuses ps ON ps.id = pd.payment_status_id
	WHERE so.user_id = $1 ORDER BY so.order_date DESC LIMIT $2 OFFSET $3;`
	if err := o.DB.Raw(query, userId, limit, offset).Scan(&orderHisory).Error; err != nil {
		return orderHisory, err
	}
	return orderHisory, nil

}

func (o *OrderDatabase) GetOrderByOrderId(ctx context.Context, OrderId uint) (orderData response.ShopOrder, err error) {
	query := `SELECT so.id, so.order_date,os.status, so.order_total, po.name AS payment_type, pm.name AS payment_method, ps.status AS payment_status
FROM shop_orders so
JOIN order_statuses os ON os.id = so.order_status_id
JOIN payment_options po ON so.payment_option_id = po.id
JOIN payment_methods pm ON pm.id = so.payment_method_id 
JOIN payment_statuses ps on ps.id = so.payment_status_id
WHERE so.id = $1`
	if err := o.DB.Raw(query, OrderId).Scan(&orderData).Error; err != nil {
		return orderData, err

	}
	return orderData, nil
}

func (o *OrderDatabase) SaveOrder(ctx context.Context, order domain.ShopOrder) (OrderId uint, err error) {
	query := `INSERT INTO shop_orders(user_id,order_date,order_total,shipping_id,order_status_id,
	payment_method_id,coupon_id,delivery_status_id,delivery_updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING ID`
	order.OrderDate = time.Now()
	if order.PaymentMethodID == 1 {

		order.OrderStatusID = 1 // set default ID 1 is for placed for COD orders
	}
	order.OrderStatusID = 1    // set default ID 1 is for pending
	order.DeliveryStatusID = 1 // set default ID 1 is for pending
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

func (o *OrderDatabase) ChangeOrderStatus(c context.Context, UpdateData request.UpdateOrderStatus) error {

	query := `UPDATE shop_orders
	SET order_status_id = $1
	WHERE user_id = $2 AND id = $3`
	if err := o.DB.Exec(query, UpdateData.StatusId, UpdateData.UserId, UpdateData.OrderId).Error; err != nil {
		return err
	}
	//

	return nil
}
