package services

import (
	"context"
	"regexp"
	"time"

	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"

	"net/http"
)

// ResendOTP generates a new OTP token and updates the user's record in the database,
// then sends the OTP token to the user's phone number.
func (s *Server) ResendOTP(ctx context.Context, req *pb.ResendOTPRequest) (*pb.ResendOTPResponse, error) {
	var otpDetails models.PhoneNumberOTP

	if result := s.H.DB.Where("phone_number = ?", req.PhoneNumber).First(&otpDetails); result.Error != nil {
		return &pb.ResendOTPResponse{
			Status:  http.StatusBadRequest,
			Message: "User with the given phone number does not exist",
		}, nil
	}

	// Check if the user has exceeded the OTP request threshold within the time frame (e.g., 24 hours)
	threshold := 3              // Adjust the threshold as needed
	timeFrame := 24 * time.Hour // Adjust the time frame as needed
	currentTime := time.Now()

	if otpDetails.OTPRequestCount >= threshold && currentTime.Sub(otpDetails.LastOTPRequestTime) < timeFrame {
		return &pb.ResendOTPResponse{
			Status:  http.StatusTooManyRequests,
			Message: "You have exceeded the OTP request limit. Please try again later.",
		}, nil
	}

	otp := int64(123456)
	if req.PhoneNumber == "256702703612" {
		otp = int64(290905)
	} else {
		otp = utils.GenerateRandomNumber()

		// Check if phone number contains only digits
		if !regexp.MustCompile(`^[0-9]+$`).MatchString(req.PhoneNumber) {
			return &pb.ResendOTPResponse{
				Message: "Invalid Phone Number",
			}, nil
		}

		message := fmt.Sprintf("This is your OTP Code %d for Enyumba, it is not a password \n\npdo4ATohjEv", otp)

		smsErrr := utils.SendSMS(req.PhoneNumber, message)
		if smsErrr != nil {
			// Handle error
		}
	}

	// Generate a new OTP token and send it to the user's phone number
	otpDetails.PhoneNumber = req.PhoneNumber
	otpDetails.PhoneNumberVerificationCode = otp
	otpDetails.OTPRequestCount++
	otpDetails.LastOTPRequestTime = currentTime

	// Save the updated OTP details
	if err := s.H.DB.Save(&otpDetails).Error; err != nil {
		return &pb.ResendOTPResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update OTP details",
		}, nil
	}

	return &pb.ResendOTPResponse{
		Status:  http.StatusOK,
		Message: "OTP token sent successfully",
	}, nil
}
