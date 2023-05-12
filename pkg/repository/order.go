package repository

import (
	"context"
	"fmt"
	"time"

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
	// // get cartID by user id
	// cartId, err := i.GetCartIdByUserId(ctx, userId)
	if err != nil {
		return checkOut, err
	}
	var page request.ReqPagination
	page.PageNumber = 1
	page.Count = 5
	// get cartItems

	cartItems, err := o.GetCartItemsbyUserId(ctx, page, userId)
	fmt.Println(cartItems)
	if err != nil {
		return checkOut, err
	}
	count := 0
	for _, v := range cartItems {
		if v.ProductItemID != 0 {
			count++
		}
		checkOut.TotalPrice += v.SubTotal
		checkOut.TotalQty += v.Quantity
		checkOut.DiscountAmount += v.Price - v.DiscountPrice
	}
	checkOut.TotalProductItems = uint(count)
	// get default address
	query := `SELECT a.house,a.address_line1,a.address_line2,a.city,a.state,a.zip_code,a.country  
	FROM addresses a
	JOIN user_addresses ua on ua.address_id = a.id
	WHERE ua.is_default = true AND ua.user_id = $1;`
	var address response.Address
	if err := o.DB.Raw(query, userId).Scan(&address).Error; err != nil {
		return checkOut, err
	}
	checkOut.DefaultShipping = address
	return checkOut, nil

}

func (o *OrderDatabase) PlaceCODOrder(ctx context.Context, userId uint) (shopOrder response.ShopOrder, err error) {
	tnx := o.DB.Begin()
	checkOut, err := o.CheckoutOrder(ctx, userId)
	if err != nil {
		return shopOrder, err
	}
	shopOrder.OrderDate = time.Now()
	shopOrder.ShippingAddress.ID = checkOut.DefaultShipping.ID
	shopOrder.ShippingAddress = checkOut.DefaultShipping
	shopOrder.OrderTotal = float64(checkOut.TotalPrice)
	shopOrder.PaymentMethod = "Cash on delivery"
	shopOrder.PaymentStatus = "Payment pending"
	shopOrder.TransactionID = ""

	// save shop order data
	type OrderData struct {
		OrderID       uint
		OrderStatusID uint
	}
	var orderData OrderData
	query := `INSERT INTO shop_orders (user_id,order_date,order_total,shipping_id,
		order_status_id, payment_option_id,payment_method_id,payment_status_id) 
	VALUES ($1, $2, $3, $4, 
		(SELECT id AS order_status_id FROM order_statuses WHERE status = 'Placed'),
		(SELECT id FROM payment_options WHERE name = 'COD'),
		(SELECT id FROM payment_methods WHERE name = 'Cash on delivery'),
		(SELECT id FROM payment_statuses WHERE status = 'Payment pending')) RETURNING id, order_status_id`
	if err := tnx.Raw(query, userId, shopOrder.OrderDate, shopOrder.OrderTotal,
		shopOrder.ShippingAddress.ID).Scan(&orderData).Error; err != nil {
		tnx.Rollback()
		return shopOrder, err
	}
	shopOrder.OrderID = orderData.OrderID
	shopOrder.OrderStatusID = orderData.OrderStatusID

	shopOrder.OrderStatus, err = o.OrderStatus(ctx, shopOrder.OrderStatusID)
	if err != nil {
		return shopOrder, err
	}
	// // save payment details
	// query = `INSERT INTO payments (order_id,payment_method_id)
	// VALUES($1, (SELECT id FROM payment_methods WHERE name = 'cash on delivery'))`
	// if err := tnx.Exec(query, shopOrder.PaymentDetails.OrderID).Error; err != nil {
	// 	tnx.Rollback()
	// 	return shopOrder, err
	// }

	if err = tnx.Commit().Error; err != nil {
		tnx.Rollback()
		return shopOrder, err
	}
	return shopOrder, nil

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

	query := `SELECT so.id, so.order_date,os.status, so.order_total, po.name AS payment_type, pm.name AS payment_method, ps.status AS payment_status
	FROM shop_orders so
	JOIN order_statuses os ON os.id = so.order_status_id
	JOIN payment_options po ON so.payment_option_id = po.id
	JOIN payment_methods pm ON pm.id = so.payment_method_id 
	JOIN payment_statuses ps on ps.id = so.payment_status_id
	WHERE so.user_id = $1 ORDER BY so.order_date DESC LIMIT $2 OFFSET $3;`
	if err := o.DB.Raw(query, userId, limit, offset).Scan(&orderHisory).Error; err != nil {
		return orderHisory, err
	}
	return orderHisory, nil

}

func (o *OrderDatabase) GetOrderByOrderId(ctx context.Context, userId, OrderId uint) (orderData response.ShopOrder, err error) {
	query := `SELECT so.id, so.order_date,os.status, so.order_total, po.name AS payment_type, pm.name AS payment_method, ps.status AS payment_status
FROM shop_orders so
JOIN order_statuses os ON os.id = so.order_status_id
JOIN payment_options po ON so.payment_option_id = po.id
JOIN payment_methods pm ON pm.id = so.payment_method_id 
JOIN payment_statuses ps on ps.id = so.payment_status_id
WHERE so.user_id = $1 AND so.id = $2`
	if err := o.DB.Raw(query, userId, OrderId).Scan(&orderData).Error; err != nil {
		return orderData, err

	}
	return orderData, nil
}
