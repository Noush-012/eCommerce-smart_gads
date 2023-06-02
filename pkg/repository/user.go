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

func (i *userDatabase) GetUserbyID(ctx context.Context, userId uint) (domain.Users, error) {
	var user domain.Users
	query := `SELECT * FROM users WHERE id = ?`
	if err := i.DB.Raw(query, userId).Scan(&user).Error; err != nil {
		return user, err
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

func (i *userDatabase) SaveAddress(ctx context.Context, userAddress request.Address) error {
	var defaultAddressID uint
	userAddress.CreatedAt = time.Now()
	query := `INSERT INTO addresses (user_id ,house,address_line1,address_line2,city,state,zip_code,country,created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	if err := i.DB.Raw(query, userAddress.UserID, userAddress.House, userAddress.AddressLine1,
		userAddress.AddressLine2, userAddress.City, userAddress.State, userAddress.ZipCode, userAddress.Country, userAddress.CreatedAt).Scan(&defaultAddressID).Error; err != nil {
		return err
	}

	// set as default if no existing default address
	query = `UPDATE addresses
	SET is_default = true
	WHERE user_id = $1
	AND is_default = false AND id = $2
	AND NOT EXISTS (
	  SELECT 1
	  FROM addresses
	  WHERE user_id = $1
	  AND is_default = true
	)`
	if err := i.DB.Exec(query, userAddress.UserID, defaultAddressID).Error; err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) UpdateAddress(ctx context.Context, userAddress request.AddressPatchReq) error {
	tnx := i.DB.Begin()
	// Set all addresses of the user to false except for the new address if new address is default
	if userAddress.IsDefault {
		resetQuery := `UPDATE addresses
				SET is_default = false
				WHERE user_id = $1`
		if err := tnx.Exec(resetQuery, userAddress.UserID).Error; err != nil {
			tnx.Rollback()
			return err
		}
	}
	query := `UPDATE addresses
	SET
		house = COALESCE(NULLIF($1, ''), house),
		address_line1 = COALESCE(NULLIF($2, ''), address_line1),
		address_line2 = COALESCE(NULLIF($3, ''), address_line2),
		city = COALESCE(NULLIF($4, ''), city),
		state = COALESCE(NULLIF($5, ''), state),
		zip_code = COALESCE(NULLIF($6, ''), zip_code),
		country = COALESCE(NULLIF($7, ''), country),
		is_default = COALESCE($8, is_default),
		updated_at = $11
	WHERE
		user_id = $9
		AND id = $10`
	userAddress.UpdatedAt = time.Now()
	if err := tnx.Exec(query, userAddress.House, userAddress.AddressLine1, userAddress.AddressLine2,
		userAddress.City, userAddress.State, userAddress.ZipCode,
		userAddress.Country, userAddress.IsDefault, userAddress.UserID, userAddress.ID, userAddress.UpdatedAt).Error; err != nil {
		tnx.Rollback()
		return err
	}
	err := tnx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
func (i *userDatabase) DeleteAddress(ctx context.Context, userID, addressID uint) error {
	query := `DELETE FROM addresses WHERE user_id = $1 AND id = $2`
	if err := i.DB.Exec(query, userID, addressID).Error; err != nil {
		return err
	}
	return nil
}

func (u *userDatabase) GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error) {
	query := `SELECT * FROM addresses WHERE user_id = ? ORDER BY is_default DESC, updated_at ASC`

	if err := u.DB.Raw(query, userId).Scan(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (i *userDatabase) SavetoCart(ctx context.Context, addToCart request.AddToCartReq) error {
	// get product item details
	query := `SELECT discount_price FROM product_items WHERE id = $1`
	if err := i.DB.Raw(query, addToCart.ProductItemID).Scan(&addToCart.Discount_price).Error; err != nil {
		return err
	}
	if addToCart.Discount_price == 0 {
		return errors.New("invalid product item id")
	}
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
	// Check if the product item already exist in cart
	query = `SELECT id FROM cart_items WHERE product_item_id = $1 AND cart_id = $2`
	var cartItemID int
	if err := i.DB.Raw(query, addToCart.ProductItemID, cartID).Scan(&cartItemID).Error; err != nil {
		return err
	}
	if cartItemID != 0 {
		query = `UPDATE cart_items SET quantity = quantity + $1, updated_at = $2 WHERE id = $3`
		UpdatedAt := time.Now()
		if err := i.DB.Exec(query, addToCart.Quantity, UpdatedAt, cartItemID).Error; err != nil {
			return fmt.Errorf("failed to save cart item %v", addToCart.ProductItemID)
		}
	} else {
		// insert product items to cart items
		query = `INSERT INTO cart_items (cart_id,product_item_id,quantity,price,created_at)
	VALUES ($1,$2, $3, $4, $5)`
		CreatedAt := time.Now()
		if err := i.DB.Exec(query, cartID, addToCart.ProductItemID, addToCart.Quantity, addToCart.Discount_price, CreatedAt).Error; err != nil {
			return fmt.Errorf("failed to save cart item %v", addToCart.ProductItemID)
		}
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

func (i *userDatabase) GetCartIdByUserId(ctx context.Context, userId uint) (cartId uint, err error) {
	query := `SELECT id FROM carts WHERE user_id = $1`
	if err := i.DB.Raw(query, userId).Scan(&cartId).Error; err != nil {
		return cartId, err
	}
	return cartId, nil
}

func (i *userDatabase) GetCartItemsbyUserId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	// get cartID by user id
	cartID, err := i.GetCartIdByUserId(ctx, userID)
	if err != nil {
		return CartItems, err
	}
	// get cartItems with cartID
	query := `SELECT ci.product_item_id, p.name,p.price,ci.price AS discount_price, 
	ci.quantity,pi.qty_in_stock AS qty_left,pi.stock_status, ci.price * ci.quantity AS sub_total
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

func (i *userDatabase) UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error {

	// get cartID by user id
	cartID, err := i.GetCartIdByUserId(ctx, cartUpadates.UserID)
	if err != nil {
		return err
	}
	// update cart
	query := `UPDATE cart_items 
	SET quantity = COALESCE($3, quantity),
	updated_at = $4
	WHERE id = $1 AND product_item_id = $2`
	// `UPDATE cart_items
	// SET product_item_id = COALESCE($1, product_item_id),
	// quantity = COALESCE($2, quantity)
	// WHERE id = $3`
	updatedAt := time.Now()

	if err := i.DB.Exec(query, cartID, cartUpadates.ProductItemID, cartUpadates.Quantity, updatedAt).Error; err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error {
	// get cartID by user id
	cartID, err := i.GetCartIdByUserId(ctx, DelCartItem.UserID)
	if err != nil {
		return err
	}
	// delete cartItems
	query := `DELETE FROM cart_items WHERE cart_id = $1 AND product_item_id = $2`
	if err := i.DB.Exec(query, cartID, DelCartItem.ProductItemID).Error; err != nil {
		return err
	}
	return nil

}
func (i *userDatabase) GetEmailPhoneByUserId(ctx context.Context, userID uint) (contact response.UserContact, err error) {
	// find data
	query := `SELECT email, phone FROM users WHERE id = ?`
	if err := i.DB.Raw(query, userID).Scan(&contact).Error; err != nil {
		return contact, err
	}
	return contact, nil
}

func (i *userDatabase) GetDefaultAddress(ctx context.Context, userId uint) (address response.Address, err error) {
	query := `SELECT a.id, a.house, a.address_line1, a.address_line2, a.city, a.state, a.zip_code, a.country, a.is_default
FROM addresses a
WHERE a.user_id = ? AND a.is_default = true`
	if err := i.DB.Raw(query, userId).Scan(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (i *userDatabase) AddToWishlist(ctx context.Context, wishlistData request.AddToWishlist) error {

	query := `INSERT INTO wishlists (user_id, product_item_id,quantity,created_at)
	VALUES ($1, $2, $3, $4)`
	CreatedAt := time.Now()
	if err := i.DB.Exec(query, wishlistData.UserID, wishlistData.ProductItemID, wishlistData.Quantity, CreatedAt).Error; err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) GetWishlist(ctx context.Context, userId uint) (wishlist []response.Wishlist, err error) {
	query := `SELECT w.product_item_id, p.name AS product_name,pi.discount_price AS price,w.quantity,p.image,
	vo1.option_value AS color,  vo2.option_value AS storage
	FROM wishlists w
	JOIN product_items pi ON pi.id = w.product_item_id
	JOIN products p ON p.id = pi.product_id
	JOIN product_configs pc1 ON pi.id = pc1.product_item_id AND pc1.variation_option_id IN (SELECT id FROM variation_options WHERE variation_id = 1)
		JOIN variation_options vo1 ON vo1.id = pc1.variation_option_id 
		JOIN product_configs pc2 ON pi.id = pc2.product_item_id AND pc2.variation_option_id IN (SELECT id FROM variation_options WHERE variation_id = 2)
		JOIN variation_options vo2 ON vo2.id = pc2.variation_option_id
	
	WHERE w.user_id = ?`
	if err := i.DB.Raw(query, userId).Scan(&wishlist).Error; err != nil {
		return wishlist, err
	}
	// fetch pictures
	// query = `SELECT image
	// FROM product_images
	// WHERE product_item_id = 14`

	// i.DB.Row(query)
	return wishlist, err

}

func (i *userDatabase) DeleteFromWishlist(ctx context.Context, productId, userId uint) error {
	query := `DELETE from wishlists WHERE product_item_id = $1 AND user_id = $2`
	err := i.DB.Exec(query, productId, userId).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) GetWalletHistory(ctx context.Context, userId uint) (wallet []domain.Wallet, err error) {
	query := `SELECT * FROM wallets WHERE user_id = ?`
	err = i.DB.Raw(query, userId).Scan(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, err
}

func (i *userDatabase) CreditUserWallet(ctx context.Context, data domain.Wallet) error {
	// check if the user already have wallet id
	wallet, err := i.GetWalletHistory(ctx, data.UserID)
	if err != nil {
		return err
	}
	if len(wallet) == 0 {

		// insert wallet
		query := `INSERT INTO wallets (user_id, balance, remark,updated_at,created_at) VALUES ($1, $2, $3, $4, $5)`
		data.CreatedAt = time.Now()
		err := i.DB.Exec(query, data.UserID, data.Balance, data.Remark, data.CreatedAt, data.CreatedAt).Error
		if err != nil {
			return err
		}

	} else {

		query := `UPDATE wallets
	SET balance = balance + $2, remark = $3, Updated_at = $4
	WHERE user_id = $1`
		data.UpdatedAt = time.Now()
		err := i.DB.Exec(query, data.UserID, data.Balance, data.Remark, data.UpdatedAt).Error
		if err != nil {
			return err
		}
	}
	return nil
}
