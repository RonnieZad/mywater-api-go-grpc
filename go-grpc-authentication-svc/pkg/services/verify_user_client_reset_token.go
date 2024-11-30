package services

import (
	"context"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"net/http"
	"time"
)

// Verifying reset token for secure authentication.
func (s *Server) VerifyResetToken(ctx context.Context, req *pb.VerifyResetTokenRequest) (*pb.VerifyResetTokenResponse, error) {
	var user models.User
	if result := s.H.DB.Where("password_reset_token = ?", req.PasswordResetToken).First(&user); result.Error != nil {
		return &pb.VerifyResetTokenResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid password reset token",
		}, nil
	}

	expiryTime, err := time.Parse(time.RFC3339, user.PasswordResetExpiryTime)
	if err != nil {
		// handle error
	}

	if expiryTime.Before(time.Now()) {
		return &pb.VerifyResetTokenResponse{
			Status:  http.StatusBadRequest,
			Message: "Password reset token has expired",
		}, nil
	}

	return &pb.VerifyResetTokenResponse{
		Status:  http.StatusOK,
		Message: "Password reset token is valid",
	}, nil
}
