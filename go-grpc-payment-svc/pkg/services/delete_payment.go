package service

import (
	"context"
	"net/http"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
)

// DeletePayment deletes all payments for a given user
func (s *Server) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {

	// Find all rent applications for the user
	var userPaymentTransactions []models.PaymentTransaction
	if result := s.H.DB.Where("user_id = ?", req.UserId).Find(&userPaymentTransactions); result.Error != nil {
		return &pb.DeletePaymentResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to fetch user House Reservation Tour",
		}, nil
	}

	// Delete all rent applications
	if err := s.H.DB.Delete(&userPaymentTransactions).Error; err != nil {
		return &pb.DeletePaymentResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to delete user payment",
		}, nil
	}

	return &pb.DeletePaymentResponse{
		Status:  http.StatusOK,
		Message: "User Payments Deleted Successfully",
	}, nil
}