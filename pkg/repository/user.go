package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	repo "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repo.UserRepository {
	return &userDatabase{DB: DB}
}

func (i *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Check any of the user details matching with db user list
	query := `SELECT * FROM users WHERE id = ? OR email = ? OR phone = ? OR user_name = ?`
	if err := i.DB.Raw(query, user.ID, user.Email, user.Phone, user.UserName).Scan(&user).Error; err != nil {
		return user, errors.New("failed to get user")
	}
	return user, nil
}

func (i *userDatabase) SaveUser(ctx context.Context, user domain.Users) error {
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password,created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now()
	err := i.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age,
		user.Email, user.Phone, user.Password, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save user %s", user.UserName)
	}
	return nil
}

func (i *userDatabase) SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error {
	// get product item details
	query := `SELECT discount_price FROM product_items WHERE id = $1`
	if err := i.DB.Raw(query, addToCart.ProductItemID).Scan(&addToCart.Discount_price).Error; err != nil {
		return err
	}
	fmt.Println(addToCart.Discount_price)
	// get cart id with user id
	query = `SELECT id FROM carts WHERE user_id = $1`
	var cartID int
	if err := i.DB.Raw(query, addToCart.UserID).Scan(&cartID).Error; err != nil {
		return err
	}
	if cartID == 0 {
		// create a cart for user with userID if not exist
		query = `INSERT INTO carts (user_id) VALUES ($1) RETURNING id`
		if err := i.DB.Raw(query, addToCart.UserID).Scan(&cartID).Error; err != nil {
			return err
		}
	}
	// insert product items to cart items
	query = `INSERT INTO cart_items (cart_id,product_item_id,quantity,price,created_at)
	VALUES ($1,$2, $3, $4, $5)`
	CreatedAt := time.Now()
	if err := i.DB.Exec(query, cartID, addToCart.ProductItemID, addToCart.Quantity, addToCart.Discount_price, CreatedAt).Error; err != nil {
		return fmt.Errorf("failed to save cart item %v", addToCart.ProductItemID)
	}
	var cartItems []domain.CartItem
	if err := i.DB.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return err
	}

	// Calculate the new total based on the updated cart items
	var total float64
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Price
	}
	if err := i.DB.Exec("UPDATE carts SET total = $1 WHERE user_id = $2", total, addToCart.UserID).Error; err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {
	var cartID uint
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	// get cart id with user id
	query := `SELECT id FROM carts WHERE user_id = $1`
	if err := i.DB.Raw(query, userID).Scan(&cartID).Error; err != nil {
		return CartItems, err
	}
	// get cartItems with cartID
	query = `SELECT ci.product_item_id, p.name,p.price,ci.price AS discount_price, 
	ci.quantity,pi.qty_in_stock AS qty_left, ci.price * ci.quantity AS sub_total
	FROM cart_items ci
	JOIN product_items pi ON ci.product_item_id = pi.id
	JOIN products p ON pi.product_id = p.id
	WHERE cart_id = $1
	ORDER BY ci.created_at DESC LIMIT $2 OFFSET $3`
	if err := i.DB.Raw(query, cartID, limit, offset).Scan(&CartItems).Error; err != nil {
		return CartItems, err
	}
	return CartItems, nil
}
