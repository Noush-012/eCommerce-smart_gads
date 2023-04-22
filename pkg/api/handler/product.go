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
		response := response.SuccessResponse(200, "there is no products to show", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	respones := response.SuccessResponse(200, "Product listed successfuly", products)
	c.JSON(http.StatusOK, respones)
}

func (p *ProductHandler) AddProduct(c *gin.Context) {
	var body request.ReqProduct
	if err := c.ShouldBindJSON(&body); err != nil {
		responce := response.ErrorResponse(400, "Invalid entry", err.Error(), body)
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
	response := response.SuccessResponse(http.StatusOK, "Product added successful", product)
	c.JSON(http.StatusOK, response)

}

func (p *ProductHandler) AddBrand(c *gin.Context) {
	var ProductBrand domain.Brand

	// Get json and bind
	if err := c.ShouldBindJSON(&ProductBrand); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, "Invalid entry", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call add brand usecase
	err := p.ProductService.AddBrand(c, ProductBrand)
	if err != nil {
		response := response.ErrorResponse(400, "Failed to add brand", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Success response
	response := response.SuccessResponse(200, "Successfuly added a new brand in database", ProductBrand)
	c.JSON(200, response)

}

func (p *ProductHandler) GetAllBrands(c *gin.Context) {
	//
}
