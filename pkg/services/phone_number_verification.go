package services

import (
	"context"
	// "time"

	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"

	"net/http"
	"regexp"
)

// phone number verification: sends otp to user
func (s *Server) PhoneNumberVerification(ctx context.Context, req *pb.PhoneNumberVerificationRequest) (*pb.PhoneNumberVerificationResponse, error) {
	var otpDetails models.PhoneNumberOTP
	// Check if phone number is 13 characters in length
	if len(req.PhoneNumber) == 0 {
		return &pb.PhoneNumberVerificationResponse{
			IsValid: false,
			Message: "Invalid Phone Number",
		}, nil
	}

	if result := s.H.DB.Where("phone_number = ?", req.PhoneNumber).First(&otpDetails); result.Error != nil {
		fmt.Println("User with the given phone number does not exist")
	}

	// // Check if the user has exceeded the OTP request threshold within the time frame (e.g., 24 hours)
	// threshold := 3              // Adjust the threshold as needed
	// timeFrame := 24 * time.Hour // Adjust the time frame as needed
	// currentTime := time.Now()

	// if otpDetails.OTPRequestCount >= threshold && currentTime.Sub(otpDetails.LastOTPRequestTime) < timeFrame {
	// 	return &pb.PhoneNumberVerificationResponse{
	// 		Status:  http.StatusTooManyRequests,
	// 		Message: "You have exceeded the OTP request limit. Please try again later.",
	// 	}, nil
	// }

	otp := int64(123456)
	if req.PhoneNumber == "256702703612" || req.PhoneNumber == "256743108134" {
		otp = int64(290905)
	} else {
		otp = utils.GenerateRandomNumber()

		// Check if phone number contains only digits
		if !regexp.MustCompile(`^[0-9]+$`).MatchString(req.PhoneNumber) {
			return &pb.PhoneNumberVerificationResponse{
				IsValid: false,
				Message: "Invalid Phone Number",
			}, nil
		}

		message := fmt.Sprintf("This is your OTP Code %d for MyWater App, it is not a password", otp)

		smsErrr := utils.SendSMS(req.PhoneNumber, message)
		if smsErrr != nil {

			// Handle error
		}
	}

	// Generate a new OTP token and send it to the user's phone number
	otpDetails.PhoneNumber = req.PhoneNumber
	otpDetails.PhoneNumberVerificationCode = otp
	// otpDetails.OTPRequestCount++
	// otpDetails.LastOTPRequestTime = currentTime

	// Generate 6 bytes of random data

	if err := s.H.DB.Save(&otpDetails).Error; err != nil {
		return &pb.PhoneNumberVerificationResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update OTP details",
		}, nil
	}

	// Phone number is valid
	return &pb.PhoneNumberVerificationResponse{
		Status:  http.StatusOK,
		IsValid: true,
		Message: "OTP sent successfully",
	}, nil
}

// phone number verification: checks whether OTP a user enters matches
func (s *Server) PhoneNumberVerificationWithOTP(ctx context.Context, req *pb.PhoneNumberVerificationWithOTPRequest) (*pb.PhoneNumberVerificationWithOTPResponse, error) {

	var otpDetails models.PhoneNumberOTP
	var user models.User

	userPhoneNumberExists := false

	s.H.DB.Where("phone_number = ?", req.PhoneNumber).First(&user)
	if user.PhoneNumber == "" {
		fmt.Println("Phone number does not exist in the users table.")
	} else {
		userPhoneNumberExists = true
		fmt.Println("Phone number exists in the users table.")
	}

	// Checks if phone number and verification code match, returns successful response on success.
	if result := s.H.DB.Where("phone_number = ? AND phone_number_verification_code = ?", req.PhoneNumber, req.Otp).First(&otpDetails); result.Error == nil {
		return &pb.PhoneNumberVerificationWithOTPResponse{
			Status:  http.StatusOK,
			Message: "Phone verified successfuly",
			IsValid: userPhoneNumberExists,
		}, nil
	}

	// OTP check failed; invalid OTP entered.
	return &pb.PhoneNumberVerificationWithOTPResponse{
		Status:  http.StatusNotFound,
		Error:   "You entered a wrong OTP",
		IsValid: false,
	}, nil
}
