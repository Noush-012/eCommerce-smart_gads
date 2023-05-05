package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(prodUseCase service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: prodUseCase,
	}
}

// ListProducts-User godoc
// @summary api for user to list all products
// @security ApiKeyAuth
// @tags User Products
// @id ListProducts-User
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /products [get]
// @Success 200 {object} response.Response{} ""Product listed successfuly""
// @Failure 500 {object} response.Response{}  "failed to get all products"
func (p *ProductHandler) ListProducts(c *gin.Context) {

	count, err1 := utils.StringToUint(c.Query("count"))
	pageNumber, err2 := utils.StringToUint(c.Query("page_number"))
	fmt.Println(count, pageNumber)
	err1 = errors.Join(err1, err2)
	if err1 != nil {
		response := response.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := request.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	products, err := p.ProductService.GetProducts(c, pagination)
	if err != nil {
		response := response.ErrorResponse(500, "failed to get all products", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if products == nil {
		response := response.SuccessResponse(200, "Oops ! no products to show", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	respones := response.SuccessResponse(200, "Product listed successfuly", products)
	c.JSON(http.StatusOK, respones)
}

// AddProduct godoc
// @summary api for admin to update a product
// @id AddProducts
// @tags Admin Product
// @Param input body  request.ProductReq{} true "inputs"
// @Router /admin/products [post]
// @Success 200 {object} response.Response{} "Product added successful"
// @Failure 400 {object} response.Response{} "Missing or invalid entry"
func (p *ProductHandler) AddProduct(c *gin.Context) {
	var body request.ProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		responce := response.ErrorResponse(400, "Missing or invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, responce)
		return
	}

	var product domain.Product
	copier.Copy(&product, body)
	if err := p.ProductService.AddProduct(c, product); err != nil {
		response := response.ErrorResponse(400, "failed to add product", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(http.StatusOK, "Product added successful", body)
	c.JSON(http.StatusOK, response)

}

// AddCategory Admin godoc
// @summary api for admin to add a parent category or child brand
// @id AddCategory
// @tags Admin Brand / Category
// @Param input body request.CategoryReq{} true "inputs"
// @Router /admin/products [post]
// @Success 200 {object} response.Response{} "Successfuly added a new brand/Category in database"
// @Failure 400 {object} response.Response{} "Missing or invalid entry"
func (p *ProductHandler) AddCategory(c *gin.Context) {
	var ProductBrand request.CategoryReq

	// Get json and bind
	if err := c.ShouldBindJSON(&ProductBrand); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Missing or invalid entry", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call add brand usecase
	err := p.ProductService.AddCategory(c, ProductBrand)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to add brand", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Success response
	response := response.SuccessResponse(200, "Successfuly added a new brand / category in database", ProductBrand)
	c.JSON(200, response)

}

// ListBrands-Admin godoc
// @summary api for admin to list all brands
// @security ApiKeyAuth
// @tags Product brands
// @id ListBrands-admin
// @Router /brands [get]
// @Success 200 {object} response.Response{} ""Successfuly listed all brands""
// @Failure 500 {object} response.Response{}  "Failed to get brands"
func (p *ProductHandler) GetAllBrands(c *gin.Context) {

	allBrands, err := p.ProductService.GetAllBrands(c)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to get brands", err.Error(), allBrands)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// Success response
	response := response.SuccessResponse(200, "Successfuly listed all brands", allBrands)
	c.JSON(200, response)
}

// UpdateProduct godoc
// @summary api for admin to update a product
// @id UpdateProduct
// @tags Admin Product
// @Param input body request.UpdateProductReq{} true "inputs"
// @Router /admin/products [put]
// @Success 200 {object} response.Response{} "Product updated successful"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
func (p *ProductHandler) UpdateProduct(c *gin.Context) {

	var body request.UpdateProductReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var product domain.Product
	copier.Copy(&product, &body)

	err := p.ProductService.UpdateProduct(c, product)
	if err != nil {
		response := response.ErrorResponse(400, "failed to update product", err.Error(), body)
		c.JSON(400, response)
		return
	}

	response := response.SuccessResponse(200, "Product updated successful", body)
	c.JSON(200, response)

	c.Abort()
}

// UpdateProduct godoc
// @summary api for admin to delete a product
// @id UpdateProduct
// @tags Admin Product
// @Param input body request.DeleteProductReq{} true "inputs"
// @Router /admin/products [put]
// @Success 200 {object} response.Response{} "Successfuly deleted product"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "Missing or invalid input"
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	var body request.DeleteProductReq
	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	productID := body.ID

	deletedProduct, err := p.ProductService.DeleteProduct(c, productID)
	if err != nil {
		response := response.ErrorResponse(500, "Failed to delete product", err.Error(), body)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := response.SuccessResponse(http.StatusOK, "Successfuly deleted product", deletedProduct)
	c.JSON(200, response)
}

// AddProductItem godoc
// @summary api for admin to add product item for particular product
// @id AddProductItem
// @tags Admin Product
// @Param input body request.ProductItemReq{} true "inputs"
// @Router /admin/products/product-item [post]
// @Success 200 {object} response.Response{} "Product item added successful"
// @Failure 400 {object} response.Response{} "Missing or invalid input"
// @Failure 500 {object} response.Response{} "failed to add product item"
func (p *ProductHandler) AddProductItem(c *gin.Context) {
	var body request.ProductItemReq

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Missing or invalid input", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := p.ProductService.AddProductItem(c, body)
	if err != nil {
		response := response.ErrorResponse(500, "failed to add product item", err.Error(), body)
		c.JSON(500, response)
		return
	}
	response := response.SuccessResponse(200, "Product item added successful", body)
	c.JSON(200, response)
	c.Abort()
}

/*
*
GetProductItem godoc
@Summary Handler to get the product item for a given product ID.
@Description This handler receives a GET request with a product ID as a parameter
and returns the corresponding product item. If the product ID is invalid or if
there are no product items for the given ID, it returns an error response.
@Tags Product
@Param product_id path uint true "Product ID"
@Success 200 {object} response.SuccessResponse "Successful response with product items"
@Failure 400 {object} response.ErrorResponse "Invalid input or error fetching product item"
@Router /products/{product_id} [get]
*/
func (p *ProductHandler) GetProductItem(c *gin.Context) {
	productID, err := utils.StringToUint(c.Param("product_id"))
	if err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "invalid param input", err.Error(), productID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	productItems, err := p.ProductService.GetProductItem(c, productID)
	if err != nil {
		response := response.ErrorResponse(400, "failed to get product item for given product id", err.Error(), productID)
		c.JSON(400, response)
		return
	}
	if productItems == nil {
		response := response.ErrorResponse(http.StatusBadRequest, "No product items for this product id", err.Error(), productID)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.SuccessResponse(200, "Fetching product item successful and listed below", productItems)
	c.JSON(200, response)

}
