package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"github.com/google/uuid"
)

func (s *Server) GetUserNotification(ctx context.Context, req *pb.GetUserNotificationRequest) (*pb.GetUserNotificationResponse, error) {
	// Find user by ID
	var paymentTransactions []models.PaymentTransaction

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		fmt.Println("failed to parse uuid from user")
		// handle error
	}

	if result := s.H.DB.Where("user_id = ?", userId).Order("created_at DESC").Find(&paymentTransactions); result.Error != nil {
		return &pb.GetUserNotificationResponse{
			Status: http.StatusNotFound,
			Error:  "No new notifications",
		}, nil
	}

	// Convert payment transactions to protobuf message format
	var paymentTransactionsPB []*pb.PaymentTransaction
	for _, pt := range paymentTransactions {
		paymentTransactionsPB = append(paymentTransactionsPB, &pb.PaymentTransaction{
			Id:              pt.Id.String(),
			PropertyId:      pt.PropertyId.String(),
			UserId:          pt.UserId.String(),
			Amount:          pt.Amount,
			Currency:        pt.Currency,
			TransactionType: pt.TransactionType,
			Reason:          pt.Reason,
			Status:          pt.Status,
			PhoneNumber:     pt.PhoneNumber,
			TransactionDate: pt.TransactionDate,
			ExpiryDate:      pt.ExpiryDate.Format("2006-01-02 15:04:05"),
			TransactionId:   pt.TransactionId,
		})
	}

	return &pb.GetUserNotificationResponse{
		Status:              http.StatusOK,
		Message:             fmt.Sprintf("Found %d payment transactions for user", len(paymentTransactions)),
		PaymentTransactions: paymentTransactionsPB,
	}, nil
}

func (s *Server) GetAllPaymentTransactions(ctx context.Context, req *pb.GetAllPaymentTransactionsRequest) (*pb.GetAllPaymentTransactionsResponse, error) {

	var paymentTransactions []models.PaymentTransaction

	if result := s.H.DB.Find(&paymentTransactions); result.Error != nil {
		return &pb.GetAllPaymentTransactionsResponse{
			Status: http.StatusNotFound,
			Error:  "Failed to retrieve payment history data",
		}, nil
	}

	// Convert payment transactions to protobuf message format
	var paymentTransactionsPB []*pb.PaymentTransaction
	for _, pt := range paymentTransactions {

		user, error := s.AuthSvc.GetUser(pt.UserId.String())

		if error != nil {
			fmt.Print(error)
			fmt.Print(pt.UserId.String())
		} else if user.Status >= http.StatusNotFound {
			fmt.Print("no related user found")
		}

		// fmt.Print(user.Data.Name)
		paymentTransactionsPB = append(paymentTransactionsPB, &pb.PaymentTransaction{
			Id:              pt.Id.String(),
			PropertyId:      pt.PropertyId.String(),
			UserId:          pt.UserId.String(),
			Amount:          pt.Amount,
			Currency:        pt.Currency,
			TransactionType: pt.TransactionType,
			Reason:          pt.Reason,
			Status:          pt.Status,
			PhoneNumber:     pt.PhoneNumber,
			TransactionDate: pt.TransactionDate,
			ExpiryDate:      pt.ExpiryDate.Format("2006-01-02 15:04:05"),
			PaidByName:      user.Data.Name,
			PaidByPhoto:     user.Data.ProfilePic,
			TransactionId:   pt.TransactionId,
		})
	}

	return &pb.GetAllPaymentTransactionsResponse{
		Status:              http.StatusOK,
		Message:             fmt.Sprintf("Found %d payment transactions", len(paymentTransactions)),
		PaymentTransactions: paymentTransactionsPB,
	}, nil
}
