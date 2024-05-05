package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model struct for storing user data
type User struct {
	Id                          uuid.UUID `json:"id" gorm:"primaryKey"`
	EmailAddress                string    `json:"email_address" gorm:"primaryKey"`
	PhoneNumber                 string    `json:"phone_number"`
	AccountNumber               int64     `json:"account_number"`
	PhoneNumberVerificationCode int64     `json:"phone_number_verification_code"`
	UserName                    string    `json:"name"`
	Password                    string    `json:"password"`
	Role                        string    `json:"role"`
	Permissions                 string    `json:"permissions"`
	IsKYBVerified               bool      `json:"kyb_verified" gorm:"default:false"`
	CompanyLogo                 string    `json:"company_logo"`
	CompanyName                 string    `json:"company_name"`
	CompanyWebsite              string    `json:"company_website"`
	CompanyAddress              string    `json:"company_address"`
	CompanyPhone                string    `json:"company_phone"`
	CompanyDescription          string    `json:"company_description"`
	CompanyEmail                string    `json:"company_email"`
	IsAccountActive             bool      `json:"is_account_active" gorm:"default:false"`
	IsActiveLocked              bool      `json:"is_account_locked" gorm:"default:false"`
	PasswordResetToken          int64     `json:"password_reset_token"`
	PasswordResetExpiryTime     string    `json:"password_reset_expiry_time"`
	gorm.Model                  `json:"-"`
}

type UserClient struct {
	Id                          uuid.UUID `json:"id" gorm:"primaryKey"`
	PhoneNumber                 string    `json:"phone_number" gorm:"primaryKey"`
	AccountNumber               int64     `json:"account_number"`
	PhoneNumberVerificationCode int64     `json:"phone_number_verification_code"`
	UserName                    string    `json:"name"`
	Password                    string    `json:"password"`
	Role                        string    `json:"role"`
	Permissions                 string    `json:"permissions"`
	IsAccountActive             bool      `json:"is_account_active" gorm:"default:false"`
	IsActiveLocked              bool      `json:"is_account_locked" gorm:"default:false"`
	PasswordResetToken          int64     `json:"password_reset_token"`
	PasswordResetExpiryTime     string    `json:"password_reset_expiry_time"`
	gorm.Model                  `json:"-"`
}

type KycVerificationMessage struct {
	Message            string `json:"message"`
	UserId             string `json:"user_id"`
	VerificationStatus bool   `json:"verification_status"`
}

type VerificationResultBody struct {
	DOB      string `json:"DOB"`
	FullName string `json:"FullName"`
	Gender   string `json:"Gender"`
	IDType   string `json:"IDType"`
	Actions  struct {
		LivenessCheck            string `json:"Liveness_Check"`
		RegisterSelfie           string `json:"Register_Selfie"`
		VerifyDocument           string `json:"Verify_Document"`
		HumanReviewCompare       string `json:"Human_Review_Compare"`
		ReturnPersonalInfo       string `json:"Return_Personal_Info"`
		SelfieToIDCardCompare    string `json:"Selfie_To_ID_Card_Compare"`
		HumanReviewLivenessCheck string `json:"Human_Review_Liveness_Check"`
	} `json:"Actions"`
	Country       string `json:"Country"`
	Document      string `json:"Document"`
	IDNumber      string `json:"IDNumber"`
	ResultCode    string `json:"ResultCode"`
	ResultText    string `json:"ResultText"`
	SmileJobID    string `json:"SmileJobID"`
	PartnerParams struct {
		JobID   string `json:"job_id"`
		UserID  string `json:"user_id"`
		JobType string `json:"job_type"`
	} `json:"PartnerParams"`
	ExpirationDate string `json:"ExpirationDate"`
	Timestamp      string `json:"timestamp"`
	Signature      string `json:"signature"`
	ImageLinks     struct {
		IdCardImage string `json:"id_card_image"`
		IdCardBack  string `json:"id_card_back"`
		SelfieImage string `json:"selfie_image"`
	} `json:"ImageLinks"`
}
