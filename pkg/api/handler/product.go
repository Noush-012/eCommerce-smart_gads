package handler

import (
	"net/http"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/req"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/resp"
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

}

func (p *ProductHandler) AddProduct(c *gin.Context) {
	var body req.ReqProduct
	if err := c.ShouldBindJSON(&body); err != nil {
		responce := resp.ErrorResponse(400, "Invalid entry", err.Error(), body)
		c.JSON(http.StatusBadRequest, responce)
		return
	}

	var product domain.Product
	copier.Copy(&product, body)
	if err := p.ProductService.AddProduct(c, product); err != nil {
		response := resp.ErrorResponse(400, "failed to add product", err.Error(), body)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := resp.SuccessResponse(http.StatusOK, "Product added successful", product)
	c.JSON(http.StatusOK, response)

}

func (p *ProductHandler) AddBrand(c *gin.Context) {
	var ProductBrand domain.Brand

	// Get json and bind
	if err := c.ShouldBindJSON(&ProductBrand); err != nil {
		response := resp.ErrorResponse(http.StatusBadRequest, "Invalid entry", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call add brand usecase
	err := p.ProductService.AddBrand(c, ProductBrand)
	if err != nil {
		response := resp.ErrorResponse(400, "Failed to add brand", err.Error(), ProductBrand)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Success response
	response := resp.SuccessResponse(200, "Successfuly added a new brand in database", ProductBrand)
	c.JSON(200, response)

}

func (p *ProductHandler) GetAllBrands(c *gin.Context) {
	//
}
