package services

import (
	"context"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
)

// Get number of available property types for each category
func (s *Server) GetUserMetrics(ctx context.Context, req *pb.GetUserMetricsRequest) (*pb.GetUserMetricsResponse, error) {

	var users []models.User
	result := s.H.DB.Find(&users)
	if result.Error != nil {
		// return nil, status.Errorf(codes.Internal, pt.errorMessage)
	}

	return &pb.GetUserMetricsResponse{
		Message:    "Metrics retrieved successfully",
		TotalUsers: int32(len(users)),
	}, nil
}
