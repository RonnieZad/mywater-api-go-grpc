package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleAPIError(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return errors.New("bad request")
	case http.StatusUnauthorized:
		return errors.New("unauthorized")
	case http.StatusNotFound:
		return errors.New("not found")
	default:
		return fmt.Errorf("unexpected API error with status code: %d", resp.StatusCode)
	}
}

func initiateMobileMoneyPayment(paymentDetail models.BeyonicPayment) {

	jsonBody, err := json.Marshal(paymentDetail)
	if err != nil {
		fmt.Print(err)
	}

	url := "https://api.mfsafrica.com/api/payments"
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

	if res.StatusCode != 201 {
		handleAPIError(res, err)
	}

	body, _ := io.ReadAll(res.Body)
	fmt.Print(body)
	fmt.Println("payment initiated")
}

// send payment to a user like landlord
func (s *Server) SendPayment(ctx context.Context, req *pb.SendPaymentRequest) (*pb.SendPaymentResponse, error) {

	// if req.Metadata.PaymentReason == "referral cashout" {
	fmt.Print(req)
	// Check if a property ID is provided and validate the property status
	if req.Metadata.PropertyId != nil {
		property, err := s.PropertySvc.FindOne(*req.Metadata.PropertyId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to fetch property: %v", err)
		}
		if property.Status >= http.StatusNotFound {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid property status: %d", property.Status)
		}
	}

	payment := models.BeyonicPayment{
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Currency:    req.Currency,
		Amount:      req.AmountToPay,
		Description: req.Description,
		PaymentType: req.PaymentType,
		Metadata: models.PaymentMetadata{
			UserId:        req.Metadata.UserId,
			PropertyId:    req.Metadata.PropertyId,
			VoucherCode:   req.Metadata.VoucherCode,
			PaymentReason: req.Metadata.PaymentReason,
			ApplicationId: req.Metadata.ApplicationId,
		},
	}

	// Process payment if the property is found
	initiateMobileMoneyPayment(payment)

	return &pb.SendPaymentResponse{
		Status:  http.StatusCreated,
		Message: "Payment instructions initiated",
	}, nil
}
