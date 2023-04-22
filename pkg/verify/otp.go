package verify

import (
	"errors"
	"fmt"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	AUTHTOKEN  string
	SERVICESID string
	ACCOUNTSID string
	client     *twilio.RestClient
)

func SetClient() {
	// SERVICESID = viper.GetString("TWILIO_SERVICES_ID")
	// ACCOUNTSID = viper.GetString("TWILIO_ACCOUNT_SID")
	// AUTHTOKEN = viper.GetString("TWILIO_AUTHTOKEN")
	// SERVICESID = config.GetConfig().SERVICESID
	// ACCOUNTSID = config.GetConfig().ACCOUNTSID
	// AUTHTOKEN = config.GetConfig().AUTHTOKEN

	// client = twilio.NewRestClientWithParams(twilio.ClientParams{
	// 	Password: AUTHTOKEN,
	// 	Username: ACCOUNTSID,
	// })
	// if client != nil {
	// 	fmt.Println("Twilio connected")
	// }
}

func TwilioSendOTP(phoneNumber string) (string, error) {

	SERVICESID = config.GetConfig().SERVICESID
	ACCOUNTSID = config.GetConfig().ACCOUNTSID
	AUTHTOKEN = config.GetConfig().AUTHTOKEN

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: AUTHTOKEN,
		Username: ACCOUNTSID,
	})
	if client != nil {
		fmt.Println("Twilio connected")
	}

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(SERVICESID, params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func TwilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(SERVICESID, params)

	if err != nil {
		return err
	} else if *resp.Status != "approved" {
		return errors.New("OTP verification failed")
	}

	return nil
}
