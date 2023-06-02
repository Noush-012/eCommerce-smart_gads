package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/config"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/gin-gonic/gin"
	razorpay "github.com/razorpay/razorpay-go"
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

func GenerateRazorPayOrder(amount float64, ReceiptId string) (razorpayOrderID interface{}, err error) {
	// Get razor pay api config
	razorPayKey := config.GetConfig().RazorPayKey
	razorPaySecret := config.GetConfig().RazorPaySecret

	// Create razorpay client instance
	rPayClient := razorpay.NewClient(razorPayKey, razorPaySecret)
	data := map[string]interface{}{
		"amount":   amount * 100,
		"currency": "INR",
		"receipt":  ReceiptId,
	}
	response, err := rPayClient.Order.Create(data, nil)
	if err != nil {
		return razorpayOrderID, err
	}
	razorpayOrderID = response["id"]

	return razorpayOrderID, nil
}

func VerifyRazorPayPayment(razorPayBody request.RazorpayVerifyReq) error {
	// Get razor pay api config
	razorPayKey := config.GetConfig().RazorPayKey
	razorPaySecret := config.GetConfig().RazorPaySecret

	// Verify signature
	data := razorPayBody.RazorpayOrderId + "|" + razorPayBody.PaymentID
	h := hmac.New(sha256.New, []byte(razorPaySecret))
	_, err := h.Write([]byte(data))
	if err != nil {
		return err
	}
	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(razorPayBody.Razorpay_signature)) != 1 {
		return err
	}
	// verify payment
	rPayClient := razorpay.NewClient(razorPayKey, razorPaySecret)

	// fetch payment and verify
	payment, err := rPayClient.Payment.Fetch(razorPayBody.PaymentID, nil, nil)
	if err != nil {
		return err
	}
	// check payment status
	if payment["status"] != "captured" {
		return fmt.Errorf("failed to verify payment \n razor pay payment with payment_id %v", razorPayBody.PaymentID)
	}
	return nil
}
