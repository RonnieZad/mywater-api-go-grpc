package service

import (
	"errors"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"

	"github.com/google/uuid"
)

func addLandlordRevenue(s *Server, landlordIdStr string, propertyIdStr string, transactionId int32, transactionDate string, amount float64, status string) error {

	var landlordRevenue models.LandlordRevenueDeposit

	userId, err := uuid.Parse(landlordIdStr)
	if err != nil {
		return errors.New("invalid user ID")

	}

	propertyId, err := uuid.Parse(propertyIdStr)
	if err != nil {
		return errors.New("invalid propety ID")

	}

	landlordRevenue.Id = uuid.New()
	landlordRevenue.LandlordId = userId
	landlordRevenue.PropertyId = propertyId
	landlordRevenue.Currency = "UGX"
	landlordRevenue.Amount = amount
	landlordRevenue.Reason = "rent payment"
	landlordRevenue.TransactionId = transactionId
	landlordRevenue.TransactionType = "mobile money"
	landlordRevenue.TransactionDate = transactionDate
	landlordRevenue.Status = status

	if result := s.H.DB.Create(&landlordRevenue); result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return errors.New(
		"landlord revenue added successfully")
}
