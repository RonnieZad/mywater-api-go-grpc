package models

import (
	"github.com/google/uuid"
	"time"
)

// AppReferralStatus represents the status of the referral.
type AppReferralStatus string

const (
	StatusSubscribed AppReferralStatus = "subscribed"
	StatusPending    AppReferralStatus = "pending"
	StatusActive     AppReferralStatus = "active"
)

// AppReferral represents the referral information for a user.
type AppReferral struct {
	ID                 uuid.UUID         `json:"id" gorm:"type:uuid;primaryKey"`
	UserId             uuid.UUID         `json:"user_id"`
	InvitedUserId      uuid.UUID         `json:"invited_user_id"`
	DeviceSerialNumber string            `json:"device_serial_number"`
	Currency           string            `json:"currency" gorm:"default:UGX"`
	Earning            float64           `json:"earning" gorm:"default:0.0"`
	Status             AppReferralStatus `json:"status" gorm:"default:'pending'"`
	InvitationDate     time.Time         `json:"invitation_date" gorm:"default:null"`
}
