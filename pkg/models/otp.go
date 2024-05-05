package models

import "time"

// PhoneNumberOTP defines a model for OTP verification codes associated with a phone number.
type PhoneNumberOTP struct {
	PhoneNumber                 string    `json:"phone_number" gorm:"primaryKey"`
	PhoneNumberVerificationCode int64     `json:"phone_number_verification_code"`
	OTPRequestCount             int       `json:"otp_request_count"`
	LastOTPRequestTime          time.Time `json:"last_otp_request_time"`
}
