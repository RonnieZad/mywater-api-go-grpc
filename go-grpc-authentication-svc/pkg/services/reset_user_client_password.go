package services

import (
	"context"
	"strconv"

	"fmt"
	"net/http"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
)

// reset user password
func (s *Server) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	// Check if the user provided a valid phone number or email address
	var user models.User
	var identifierType string

	if utils.IsValidPhoneNumber(req.Identifier) {
		identifierType = "phone_number"
	} else if utils.IsValidEmail(req.Identifier) {
		identifierType = "email"
	} else {
		return &pb.ResetPasswordResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid phone number or email address",
		}, nil
	}

	//was initially fmt.Sprintf("%s = ?", identifierType), req.Identifier
	// Check if the user with the given identifier exists in the database
	if result := s.H.DB.Where(&models.User{PhoneNumber: req.Identifier}).First(&user); result.Error != nil {
		fmt.Print(result)
		return &pb.ResetPasswordResponse{
			// error: http.StatusBadRequest,
			Status:  http.StatusBadRequest,
			Message: "User with the given identifier does not exist",
		}, nil
	}

	// Generate a new password reset token and update the user's record in the database
	resetToken := utils.GenerateRandomNumber()
	// if err != nil {
	// 	return &pb.ResetPasswordResponse{
	// 		Status:  http.StatusInternalServerError,
	// 		Message: "Failed to generate password reset token",
	// 	}, nil
	// }
	expiryTime := time.Now().Add(time.Hour * 24) // Password reset token is valid for 24 hours

	user.PasswordResetToken = resetToken
	user.PasswordResetExpiryTime = expiryTime.Format(time.RFC3339)
	if result := s.H.DB.Save(&user); result.Error != nil {
		return &pb.ResetPasswordResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update user record in the database",
		}, nil
	}

	// Send the password reset instructions to the user's phone number or email address
	if identifierType == "phone_number" {
		// Send SMS message to the user's phone number
		message := fmt.Sprintf("Use this token to reset your password: %s for Enyumba App", strconv.Itoa(int(resetToken)))

		fmt.Print(message)
		if err := utils.SendSMS(user.PhoneNumber, message); err != nil {
			return &pb.ResetPasswordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to send password reset instructions to the user's phone number",
			}, nil
		}
	} else {
		// Send email to the user's email address
		// resetLink := fmt.Sprintf("https://example.com/reset-password?token=%s", resetToken)
		// message := fmt.Sprintf("Click on the following link to reset your password: %s", resetLink)
		// if err := utils.SendEmail(user.Email, "Password Reset Instructions", message); err != nil {
		// 	return &pb.ResetPasswordResponse{
		// 		Status:  http.StatusInternalServerError,
		// 		Message: "Failed to send password reset instructions to the user's email address",
		// 	}, nil
		// }
	}

	return &pb.ResetPasswordResponse{
		Status:  http.StatusOK,
		Message: "Password reset instructions sent successfully",
	}, nil
}
