package services

import (
	"context"

	"net/http"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
)

// Handles the updating of a user's password.
func (s *Server) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	var user models.User
	if result := s.H.DB.Where("password_reset_token = ?", req.PasswordResetToken).First(&user); result.Error != nil {
		return &pb.UpdatePasswordResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid password reset token",
		}, nil
	}

	expiryTime, err := time.Parse(time.RFC3339, user.PasswordResetExpiryTime)
	if err != nil {
		// handle error
	}

	if expiryTime.Before(time.Now()) {
		return &pb.UpdatePasswordResponse{
			Status:  http.StatusBadRequest,
			Message: "Password reset token has expired",
		}, nil
	}

	// Hash the new password before saving it to the database
	hashedPassword := utils.HashPassword(req.NewPassword)

	if err != nil {
		return &pb.UpdatePasswordResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to hash new password",
		}, nil
	}

	user.Password = hashedPassword
	user.PasswordResetToken = 0
	user.PasswordResetExpiryTime = time.Time{}.Format(time.RFC3339)
	if result := s.H.DB.Save(&user); result.Error != nil {
		return &pb.UpdatePasswordResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update user's password",
		}, nil
	}

	return &pb.UpdatePasswordResponse{
		Status:  http.StatusOK,
		Message: "Password updated successfully",
	}, nil
}
