package utils

import (
	"strconv"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/gin-gonic/gin"
)

func StringToUint(str string) (uint, error) {
	val, err := strconv.Atoi(str)
	return uint(val), err
}

func GetUserIdFromContext(ctx *gin.Context) uint {
	userIdStr := ctx.GetString("userId")
	userIdInt, _ := strconv.Atoi(userIdStr)
	return uint(userIdInt)
}
func GenerateSKU(prod request.ProductItemReq) (string, error) {
	// var sku string

	// if prod.Brand != "" {
	// 	sku = prod.Brand + "-"
	// 	if prod.Category != "" {
	// 		sku += prod.Category + "-"
	// 		if prod.SubCategory != "" {
	// 			sku += prod.SubCategory + "-"
	// 			if prod.Name != "" {
	// 				sku += prod.Name
	// 				return sku, nil
	// 			}
	return "", nil
}
