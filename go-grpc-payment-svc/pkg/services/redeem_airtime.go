package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/utils"
	"net/http"
)

func redeemAirtime(airtimeDetails models.AirtimeRedemption) error {
	jsonBody, err := json.Marshal(airtimeDetails)
	if err != nil {
		return err
	}

	url := "https://api.africastalking.com/version1/airtime/redeem"
	method := "POST"

	apiKey := "e55d663040ccf2408e2f7f0d2e9277fe4d8c27c9"

	client := &http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("apiKey", apiKey)
	res, err := client.Do(request)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Check response status code if needed
	// if res.StatusCode != http.StatusCreated {
	//     return errors.New("Airtime redemption failed")
	// }

	fmt.Println("Airtime redemption initiated")

	return nil
}

func (s *Server) RedeemAirtime(ctx context.Context, req *pb.RedeemAirtimeRequest) (*pb.RedeemAirtimeResponse, error) {

	recipient := models.Recipient{
		PhoneNumber:  req.PhoneNumber,
		CurrencyCode: "UGX",
		Amount:       req.AmountToPay,
	}

	payment := models.AirtimeRedemption{
		Username:   "enyumba",
		Recipients: []models.Recipient{recipient},
	}

	// message := "Hello, you have been requested to pay back our ka money. You will receive a payment request shortly to input your Mobile Money PIN."

	// smsErrr := utils.SendSMS(req.PhoneNumber, message)
	// if smsErrr != nil {
	// 	fmt.Println(smsErrr)

	// }

	//process payment if property is found
	redeemAirtime(payment)

	return &pb.RedeemAirtimeResponse{
		Status:  http.StatusCreated,
		Message: "Airtime redemption initiated",
	}, nil
}
