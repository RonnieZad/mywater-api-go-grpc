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

func redeemDataBundle(dataBundleDtails models.DataBundleRedemption) error {

	jsonBody, err := json.Marshal(dataBundleDtails)
	if err != nil {
		return err
	}

	url := "https://bundles.africastalking.com/mobile/data/request"
	method := "POST"
	apiKey := "e55d663040ccf2408e2f7f0d2e9277fe4d8c27c9"

	client := &http.Client{}
	request, error := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))

	if error != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("apiKey", apiKey)
	res, err := client.Do(request)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	// if res.StatusCode != 201 {
	// 	handleAPIError(res, err)
	// }

	fmt.Println("Data bundle redemption initiated")

	return nil

}

func (s *Server) ReedemDataBundle(ctx context.Context, req *pb.ReedemDataBundleRequest) (*pb.ReedemDataBundleResponse, error) {

	recipient := models.DataBundleRecipient{
		PhoneNumber: req.PhoneNumber,
		Unit:        "MB",
		Quantity:    50,
		Validity:    "Daily",
	}

	payment := models.DataBundleRedemption{
		Username:   "enyumba",
		Recipients: []models.DataBundleRecipient{recipient},
	}

	// message := "Hello, you have been requested to pay back our ka money. You will receive a payment request shortly to input your Mobile Money PIN."

	// smsErrr := utils.SendSMS(req.PhoneNumber, message)
	// if smsErrr != nil {
	// 	fmt.Println(smsErrr)

	// }

	//process payment if property is found
	redeemDataBundle(payment)

	return &pb.ReedemDataBundleResponse{
		Status:  http.StatusCreated,
		Message: "Data bundle redemption initiated",
	}, nil
}
