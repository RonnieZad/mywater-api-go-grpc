package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model struct for storing user data
type User struct {
	Id                          uuid.UUID `json:"id" gorm:"primaryKey"`
	PhoneNumber                 string    `json:"phone_number" gorm:"primaryKey"`
	AccountNumber               int64     `json:"account_number"`
	PhoneNumberVerificationCode int64     `json:"phone_number_verification_code"`
	UserName                    string    `json:"name"`
	EmailAddress                string    `json:"email_address"`
	DateOfBirth                 string    `json:"date_of_birth"`
	Password                    string    `json:"password"`
	Role                        string    `json:"role"`
	Permissions                 string    `json:"permissions"`
	IsKYBVerified               bool      `json:"kyb_verified" gorm:"default:false"`
	CompanyLogo                 string    `json:"company_logo"`
	CompanyName                 string    `json:"company_name"`
	CompanyWebsite              string    `json:"company_website"`
	IsAccountActive             bool      `json:"is_account_active" gorm:"default:false"`
	IsActiveLocked              bool      `json:"is_account_locked" gorm:"default:false"`
	PasswordResetToken          int64     `json:"password_reset_token"`
	PasswordResetExpiryTime     string    `json:"password_reset_expiry_time"`
	gorm.Model                  `json:"-"`
}
