package services

import (
	"context"

	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/google/uuid"
)

// validate authentication request
func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User
	var userClient models.UserClient

	userUUIDValue, err := uuid.Parse(claims.AccountId)
	if err != nil {

		// handle the erro
		// fmt,Print("error here, failed to aprse data")
	}

	if claims.UserRole == "client" {
		if result := s.H.DB.Where(&models.UserClient{Id: userUUIDValue}).First(&userClient); result.Error != nil {
			return &pb.ValidateResponse{
				Status: http.StatusNotFound,
				Error:  "User Not Found",
			}, nil
		}
	} else if result := s.H.DB.Where(&models.User{Id: userUUIDValue}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User Not Found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		// UserId: user.ID.String(),
	}, nil
}
