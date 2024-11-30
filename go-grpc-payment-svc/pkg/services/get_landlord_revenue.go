package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"github.com/google/uuid"
)

func (s *Server) GetLandlordRevenue(ctx context.Context, req *pb.GetLandlordRevenueRequest) (*pb.GetLandlordRevenueResponse, error) {
	// Find user by ID
	var landlordRevenueDeposit []models.LandlordRevenueDeposit
	var totalAmount float64

	userId, err := uuid.Parse(req.LandlordId)
	if err != nil {
		fmt.Println("failed to parse uuid from user")
		// handle error
	}

	if result := s.H.DB.Where("landlord_id = ?", userId).Order("created_at DESC").Find(&landlordRevenueDeposit); result.Error != nil {
		return &pb.GetLandlordRevenueResponse{
			Status: http.StatusNotFound,
			Error:  "No revenue",
		}, nil
	}

	s.H.DB.Table("landlord_revenue_deposits").Select("sum(amount) as total").Scan(&totalAmount)

	return &pb.GetLandlordRevenueResponse{
		Status:       http.StatusOK,
		Message:      "Retrieved successfully",
		TotalRevenue: totalAmount,
	}, nil
}
