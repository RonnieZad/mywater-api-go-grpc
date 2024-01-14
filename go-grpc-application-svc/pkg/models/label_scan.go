package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Advert struct {
	Id                   uuid.UUID `json:"id" gorm:"primaryKey"`
	AdvertId             string    `json:"advert_id"`
	AdvertiserId         string    `json:"advertiser_id"`
	RewardPoint          int64     `json:"reward_point"`
	AdvertPublicUrl      string    `json:"advert_public_url"`
	Promotion            string    `json:"promotion"`
	PromotionDescription string    `json:"promotion_description"`
	PictureUrl           string    `json:"picture_url"`
	ExpiryDate           string    `json:"expiry_date"`
	ScanCount            int64     `json:"scan_count" gorm:"default:0"`
	IsPending            bool      `json:"is_pending" gorm:"default:false"`
	gorm.Model           `json:"-"`
}

type LabelScan struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey"`
	UserUid      string    `json:"user_uid"`
	AdvertiserId string    `json:"advertiser_id"`
	AdvertId     string    `json:"advert_id"`
	gorm.Model   `json:"-"`
}
