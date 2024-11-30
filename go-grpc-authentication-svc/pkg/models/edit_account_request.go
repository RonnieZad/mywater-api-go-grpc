package models

import (
	"github.com/google/uuid"
)

// PhoneNumberOTP defines a model for OTP verification codes associated with a phone number.
type EditAccountRequest struct {
	ID                           uuid.UUID `json:"id" gorm:"primaryKey"`
	UserId                       uuid.UUID `UserId:"id" `
	Name                         bool      `json:"name"`
	NationalIdentificationNumber bool      `json:"national_identification_number"`
	DateOfBirth                  bool      `json:"date_of_birth"`
	PhoneNumber                  bool      `json:"phone_number"`
	IsApproved                   bool      `json:"is_approved"`
}
