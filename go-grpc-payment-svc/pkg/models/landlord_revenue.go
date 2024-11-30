package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LandlordRevenueDeposit struct {
	Id              uuid.UUID `json:"id" gorm:"primaryKey"`
	LandlordId      uuid.UUID `json:"landlordId"`
	PropertyId      uuid.UUID `json:"propertyId"`
	Currency        string    `json:"currency"`
	Amount          float64   `json:"amount"`
	TransactionId   int32     `json:"transanctionId"`
	TransactionType string    `json:"paymentType"`
	Reason          string    `json:"reason"`
	TransactionDate string    `json:"paymentDate"`
	Status          string    `json:"status"`
	gorm.Model      `json:"-"`
}

type LandlordRevenueWithdraw struct {
	Id              uuid.UUID `json:"id" gorm:"primaryKey"`
	LandlordId      uuid.UUID `json:"landlordId"`
	PropertyId      uuid.UUID `json:"propertyId"`
	Currency        string    `json:"currency"`
	Amount          float64   `json:"amount"`
	TransactionId   int32     `json:"transanctionId"`
	TransactionType string    `json:"paymentType"`
	Reason          string    `json:"reason"`
	TransactionDate string    `json:"paymentDate"`
	Status          string    `json:"status"`
	gorm.Model      `json:"-"`
}
