package handler

import (
	"errors"
	"net/http"

	handler "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	request "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userUsecase interfaces.UserService) handler.UserHandler {
	return &UserHandler{userService: userUsecase}
}

// Home godoc
// @summary api for user home page
// @description after user login user will seen this page with user informations
// @security ApiKeyAuth
// @id User Home
// @tags Home
// @Router / [get]
// @Success 200 "Welcome to home !"
func (u *UserHandler) Home(c *gin.Context) {

	response := response.SuccessResponse(200, "Welcome to home !", nil)
	c.JSON(http.StatusOK, response)
}

// Logout godoc
// @summary api for user to logout
// @description user can logout
// @security ApiKeyAuth
// @id UserLogout
// @tags User Logout
// @Router /logout [post]
// @Success 200 "Log out successful"
func (u *UserHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("user-auth", "", -1, "", "", false, true)
	response := response.SuccessResponse(http.StatusOK, "Log out successful", nil)
	c.JSON(http.StatusOK, response)
}

// GetCartItems godoc
// @summary api for user to get cart items
// @description user can get cart items
// @security ApiKeyAuth
// @id UserGetCartItems
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @tags User GetCartItems
// @Router /cart [get]
// @Success 200 "Successfuly get cart items"
func (u *UserHandler) GetcartItems(c *gin.Context) {
	var page request.ReqPagination
	count, err0 := utils.StringToUint(c.Query("count"))
	page_number, err1 := utils.StringToUint(c.Query("page_number"))
	err0 = errors.Join(err0, err1)
	if err0 != nil {
		response := response.ErrorResponse(400, "Missing or invalid inputs", err0.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	page.PageNumber = page_number
	page.Count = count

	userId := utils.GetUserIdFromContext(c)
	cartItems, err := u.userService.GetCartItemsbyCartId(c, page, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Get Cart Items successful", cartItems)
	c.JSON(http.StatusOK, response)
}

// AddToCart godoc
// @summary api for add product item to user cart
// @description user can add a stock in product to cart
// @security ApiKeyAuth
// @id AddToCart
// @tags User Cart
// @Param input body request.AddToCartReq{} true "Input Field"
// @Router /cart [post]
// @Success 200 "Successfuly added product item to cart "
// @Failure 400 "Failed to add product item in cart"
func (u *UserHandler) AddToCart(c *gin.Context) {
	var body request.AddToCartReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body.ProductItemID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		c.JSON(400, "No user id on context")
		return
	}
	if err := u.userService.SaveCartItem(c, body); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Failed to add product item in cart", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly added product item to cart ", body)
	c.JSON(200, response)

}

// UpdateCart godoc
// @summary api for update user cart
// @description user can update a stock in product to cart
// @security ApiKeyAuth
// @id UpdateCart
// @tags User Cart
// @Param input body request.UpdateCartReq{} true "Input Field"
// @Router /cart [post]
// @Success 200 "Successfuly updated product item in cart"
// @Failure 500 "Something went wrong!"
func (u *UserHandler) UpdateCart(c *gin.Context) {
	var body request.UpdateCartReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(400, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if body.UserID == 0 {
		response := response.ErrorResponse(400, "No user id on context", "", nil)
		c.JSON(400, response)
		return
	}
	if err := u.userService.UpdateCart(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly updated cart", body)
	c.JSON(200, response)

}

// DeleteCartItem godoc
// @summary api for delete product item from cart
// @description user can delete a stock in product to cart
// @security ApiKeyAuth
// @id DeleteCartItem
// @tags User Cart
// @Param input body request.DeleteCartItemReq{} true "Input Field"
// @Router /cart [delete]
// @Success 200 "Successfuly deleted product item from cart"
// @Failure 500 "Something went wrong!"
func (u *UserHandler) DeleteCartItem(c *gin.Context) {
	var body request.DeleteCartItemReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "invalid input", err.Error(), body)
		c.JSON(400, response)
		return
	}
	// get userId from context
	body.UserID = utils.GetUserIdFromContext(c)
	if err := u.userService.RemoveCartItem(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly removed item from cart", body)
	c.JSON(200, response)
}

// ! ***** for user profile ***** //
// Profile godoc
// @summary api for see user details
// @security ApiKeyAuth
// @id Account
// @tags User Account
// @Router /account [get]
// @Success 200 "Successfully user account details found"
// @Failure 500 {object} response.Response{} "faild to show user details"
func (u *UserHandler) Profile(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)

	user, err := u.userService.Profile(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Successfuly got profile", user)
	c.JSON(200, response)

}

// AddAddress godoc
// @summary api for adding a new address for user
// @description get a new address from user to store the the database
// @security ApiKeyAuth
// @id AddAddress
// @tags User Address
// @Param inputs body request.Address{} true "Input Field"
// @Router /account/address [post]
// @Success 200 {object} response.Response{} "Successfully address added"
// @Failure 400 {object} response.Response{} "inavlid input"
func (u *UserHandler) AddAddress(c *gin.Context) {
	var body request.Address
	userId := utils.GetUserIdFromContext(c)

	body.UserID = userId

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(400, response)
		return
	}
	if err := u.userService.Addaddress(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Save address successful", nil)
	c.JSON(200, response)

}

// UpdateAddress godoc
// @summary api for update user address
// @description user can update a address
// @security ApiKeyAuth
// @id UpdateAddress
// @tags User Address
// @Param input body request.AddressPatchReq{} true "Input Field"
// @Router /address [put]
// @Success 200 {object} response.Response{} "Address updated successfuly"
// @Failure 500 {object} response.Response{} "Something went wrong!"
func (u *UserHandler) UpdateAddress(c *gin.Context) {
	// Get user id from context
	userId := utils.GetUserIdFromContext(c)

	var body request.AddressPatchReq
	body.UserID = userId
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(400, response)
		return
	}
	if err := u.userService.UpdateAddress(c, body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address updated successfuly", nil)
	c.JSON(200, response)

}

// delete address
// @summary api for delete user address
// @description user can delete a address
// @security ApiKeyAuth
// @id DeleteAddress
// @tags User Address
// @Param id path string true "id"
// @Router /address [delete]
// @Success 200 {object} response.Response{} "Address deleted successfuly"
// @Failure 500 {object} response.Response{} "Something went wrong!"
func (u *UserHandler) DeleteAddress(c *gin.Context) {
	// Get user id from context
	userId := utils.GetUserIdFromContext(c)
	addressId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		response := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), nil)
		c.JSON(400, response)
		return
	}
	if err := u.userService.DeleteAddress(c, userId, addressId); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Address deleted successfuly", nil)
	c.JSON(200, response)
}

// GetAllAddress godoc
// @summary api for user to get all address
// @description user can get address
// @security ApiKeyAuth
// @id GetAllAddress
// @Param page_number query int false "Page Number"
// @tags User GetAllAddress
// @Router /address [get]
// @Success 200 {object} response.Response{} "Get all address successful"
// @Failure 500 {object} response.Response{} "Something went wrong!"
func (u *UserHandler) GetAllAddress(c *gin.Context) {
	// Get user id from context
	userId := utils.GetUserIdFromContext(c)
	if userId == 0 {
		response := response.ErrorResponse(500, "No user detected!", "", nil)
		c.IndentedJSON(400, response)
		return
	}
	address, err := u.userService.GetAllAddress(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Get all address successful", address)
	c.IndentedJSON(200, response)
}

// AddToWishlist godoc
// @summary api for user to add product to wishlist
// @description user can add product to wishlist
// @security ApiKeyAuth
// @id AddToWishlist
// @Param input body request.AddToWishlist{} true "Input Field"
// @tags User AddToWishlist
// @Router /wishlist/{product_id} [post]
// @Success 200 {object} response.Response{} "Add product to wishlist successful"
// @Failure 500 {object} response.Response{}  "Something went wrong!"
func (u *UserHandler) AddToWishlist(c *gin.Context) {
	var body request.AddToWishlist
	// Get user id from context
	userId := utils.GetUserIdFromContext(c)
	body.UserID = userId

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}

	err := u.userService.AddToWishlist(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Add product to wishlist successful", nil)
	c.IndentedJSON(200, response)

}

// GetWishlist godoc
// @summary api for user to get wishlist
// @description user can get wishlist
// @security ApiKeyAuth
// @id GetWishlist
// @tags User GetWishlist
// @Router /wishlist [get]
// @Success 200 {object} response.Response{} "Get wishlist successful"
// @Failure 500 {object} response.Response{}  "Something went wrong!"
func (u *UserHandler) GetWishlist(c *gin.Context) {
	// Get user id from context
	userId := utils.GetUserIdFromContext(c)
	wishlist, err := u.userService.GetWishlist(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}

	response := response.SuccessResponse(200, "Get wishlist successful", wishlist)
	c.IndentedJSON(200, response)
}

// DeleteFromWishlist godoc
// @summary api for user to delete product from wishlist
// @description user can delete product from wishlist
// @security ApiKeyAuth
// @id DeleteFromWishlist
// @tags User DeleteFromWishlist
// @Router /wishlist/{id} [delete]
// @Param id path string true "id"
// @Success 200 {object} response.Response{} "Delete product from wishlist successful"
// @Failure 500 {object} response.Response{}  "Something went wrong!"
func (u *UserHandler) DeleteFromWishlist(c *gin.Context) {
	// Get user id from context
	userId := utils.GetUserIdFromContext(c)
	// Get product id from path
	id, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		response := response.ErrorResponse(500, "Missing or invalid input", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	err = u.userService.DeleteFromWishlist(c, userId, id)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Delete product from wishlist successful", nil)
	c.IndentedJSON(200, response)
}

// GetWalletHistory godoc
// @summary api for user to get wallet history
// @description user can get wallet history
// @security ApiKeyAuth
// @id GetWalletHistory
// @tags User GetWalletHistory
// @Router /wallet/history [get]
// @Success 200 {object} response.Response{} "Get wallet history successful"
// @Failure 500 {object} response.Response{}  "Something went wrong!"
func (u *UserHandler) GetWalletHistory(c *gin.Context) {
	// Get user id from context
	userId := utils.GetUserIdFromContext(c)
	// Get product id from path
	// page, err := utils.StringToUint(c.Query("page"))
	// if err != nil {
	// 	page = 1
	// }
	// limit, err := utils.StringToUint(c.Query("limit"))
	// if err != nil {
	// 	limit = 10
	// }
	// Get wallet history
	history, err := u.userService.GetWalletHistory(c, userId)
	if err != nil {
		response := response.ErrorResponse(500, "Something went wrong!", err.Error(), nil)
		c.IndentedJSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Get wallet history successful", history)
	c.IndentedJSON(200, response)
}
