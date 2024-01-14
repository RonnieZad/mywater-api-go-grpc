package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"net/http"
)

func initiateMobileMoneyCollection(paymentDetail models.BeyonicCollection) {

	jsonBody, err := json.Marshal(paymentDetail)
	if err != nil {
		fmt.Print(err)
	}

	url := "https://api.mfsafrica.com/api/collectionrequests"
	method := "POST"

	client := &http.Client{}
	request, error := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))

	if error != nil {
		panic(error)
	}

	request.Header.Add("Authorization", "Token e55d663040ccf2408e2f7f0d2e9277fe4d8c27c9")
	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// if res.StatusCode != 201 {
	// 	handleAPIError(res, err)
	// }

	// body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Mobile money collection initiated")

}

func (s *Server) MakeCollection(ctx context.Context, req *pb.MakeCollectionRequest) (*pb.MakeCollectionResponse, error) {

	property, err := s.PropertySvc.FindOne(req.Metadata.PropertyId)

	if req.Metadata.PaymentReason != "roommate listing" {
		if err != nil {
			return &pb.MakeCollectionResponse{Status: http.StatusBadRequest, Error: err.Error(), Message: "You cannot pay for this service now"}, nil
		} else if property.Status >= http.StatusNotFound {
			return &pb.MakeCollectionResponse{Status: property.Status, Error: property.Error, Message: "You cannot pay for this service now"}, nil
		}
	}

	payment := models.BeyonicCollection{
		MaxAttempts:      0,
		PhoneNumber:      req.PhoneNumber,
		Amount:           req.AmountToPay,
		Currency:         req.Currency,
		SendInstructions: true,
		Metadata: struct {
			UserId           string  `json:"user_id"`
			PropertyId       string  `json:"property_id"`
			PaymentReason    string  `json:"payment_reason"`
			InstallmentCount *int32  `json:"installment_count,omitempty"` // Marking it as omitempty will handle the nil case
			ApplicationId    *string `json:"application_id,omitempty"`
			TourDate         *string `json:"tour_date,omitempty"`
			TourType         *int32  `json:"tour_type,omitempty"` // Marking it as omitempty will handle the nil case
			VoucherCode      *string `json:"voucher_code,omitempty"`
		}{
			UserId:           req.Metadata.UserId,
			PropertyId:       req.Metadata.PropertyId,
			PaymentReason:    req.Metadata.PaymentReason,
			InstallmentCount: req.Metadata.InstallmentCount,
			ApplicationId:    req.Metadata.ApplicationId,
			TourDate:         req.Metadata.TourDate,
			TourType:         req.Metadata.TourType,
			VoucherCode:      req.Metadata.VoucherCode,
		},
	}

	if req.Metadata.PaymentReason == "security deposit" {
		if property.Data.IsPropertyAvailable == true {
			initiateMobileMoneyCollection(payment)
			return &pb.MakeCollectionResponse{
				Status:  http.StatusCreated,
				Message: "Security deposit collection initiated",
			}, nil
		} else {
			return &pb.MakeCollectionResponse{Status: http.StatusNotFound, Error: "Property is not available for booking", Message: "Property is not available for booking"}, nil
		}
	}

	//process payment if property is found
	initiateMobileMoneyCollection(payment)

	return &pb.MakeCollectionResponse{
		Status:  http.StatusCreated,
		Message: "Payment instructions initiated",
	}, nil
}
