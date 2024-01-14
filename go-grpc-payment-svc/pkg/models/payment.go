package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentTransaction struct {
	Id              uuid.UUID `json:"id" gorm:"primaryKey"`
	PropertyId      uuid.UUID `json:"propertyId"`
	UserId          uuid.UUID `json:"userId"`
	Currency        string    `json:"currency"`
	TransactionType string    `json:"paymentType"`
	Reason          string    `json:"reason"`
	PhoneNumber     string    `json:"phoneNumber"`
	TransactionDate string    `json:"paymentDate"`
	ExpiryDate      time.Time `json:"expiryDate"`
	Status          string    `json:"status"`
	IsVoucherUsed   bool      `json:"isVoucherUsed" gorm:"default:false"`
	VoucherCode     string    `json:"voucherCode"`
	VoucherCodeId   uuid.UUID `json:"voucherCodeId" gorm:"default:null"`
	VoucherAmount   float64   `json:"voucherAmount"`
	Amount          float64   `json:"amount"`
	AmountPaid      float64   `json:"amountPaid"`
	TransactionId   int32     `json:"transanctionId"`
	gorm.Model      `json:"-"`
}

type Property struct {
	Id                      uuid.UUID `json:"id" gorm:"primaryKey"`
	Name                    string    `json:"name"`
	PropertyType            int32     `json:"propertyType"`
	PropertyReference       string    `json:"propertyReference"`
	LocationText            string    `json:"locationText"`
	LocationCoordinatesLat  float64   `json:"locationCoordinateLat"`
	LocationCoordinatesLong float64   `json:"locationCoordinateLong"`
	RentPrice               float64   `json:"price"`
	ServiceChargeFee        float64   `json:"serviceChargeFee"`
	AgentFee                float64   `json:"agentFee"`
	LegalFee                float64   `json:"legalFee"`
	RentalDepositPeriod     int32     `json:"rentalDepositPeriod"`
	Description             string    `json:"description"`
	FeatureBedrooms         int32     `json:"featureBedrooms"`
	FeatureBathrooms        int32     `json:"featureBathrooms"`
	FeatureSize             int32     `json:"featureSize"`
	CoverPhoto              string    `json:"coverPhoto"`
	VideoPreview            string    `json:"videoUrl"`
	IsPropertyAvailable     bool      `json:"isPropertyAvailable" gorm:"default:true"`
	IsPropertyOnHold        bool      `json:"isPropertyOnHold" gorm:"default:false"`
	PropertyHeldBy          uuid.UUID `json:"propertyHeldBy"`
	LeaseTerms              string    `json:"leaseTerms"`
	// Images                  []PropertyImages        `json:"propertyImages"`
	// PropertyAbout           PropertyAbout           `json:"propertyAbout" gorm:"foreignKey:PropertyId"`
	// PropertyManagement      PropertyManagement      `json:"propertyManagement" gorm:"foreignKey:PropertyId"`
	// PropertyNearbySchools   []PropertyNearbySchool  `json:"propertyNearbySchool" gorm:"foreignKey:PropertyId"`
	// PropertyNeighbourhoods  []PropertyNeighbourhood `json:"propertyNeighbourhood" gorm:"foreignKey:PropertyId"`
	gorm.Model `json:"-"`
}

type HouseFinancingApplication struct {
	Id                        uuid.UUID `json:"id" gorm:"primaryKey"`
	ApplicationId             string    `json:"application_id"`
	UserId                    uuid.UUID `json:"user_id"`
	PropertyId                uuid.UUID `json:"property_id"`
	EarningWindowFrame        string    `json:"earning_window_frame"`
	NextExpectedPayDate       string    `json:"next_expected_pay_date"`
	IncomeRangeMinimum        float64   `json:"income_range_minimum"`
	IncomeRangeMaximum        float64   `json:"income_range_maximum"`
	RequestedAmount           float64   `json:"requested_amount"`
	RequestedAmountPercentage float64   `json:"requested_amount_percentage"`
	AmountQualified           float64   `json:"amount_qualified"`
	PaybackInterest           float64   `json:"payback_interest"`
	InterestAmount            float64   `json:"interest_amount"`
	OtherFeesAmount           float64   `json:"other_fees_amount"`
	TotalAmount               float64   `json:"total_amount"`
	RentFinancePeriod         int32     `json:"rent_finance_period"`
	ApprovalStatus            bool      `json:"approval_status"`
	ApprovedBy                uuid.UUID `json:"approved_by"`
	IsLandlordPaid            bool      `json:"is_landlord_paid" gorm:"default:false"`
	gorm.Model                `json:"-"`
}

type CollectionCallbackBody struct {
	Hook struct {
		ID      int32  `json:"id,omitempty"`
		Created string `json:"created,omitempty"`
		Updated string `json:"updated,omitempty"`
		Event   string `json:"event,omitempty"`
		Target  string `json:"target,omitempty"`
		User    int    `json:"user,omitempty"`
	} `json:"hook"`
	Data struct {
		ID           int32  `json:"id,omitempty"`
		Organization int    `json:"organization,omitempty"`
		Amount       string `json:"amount,omitempty"`
		Currency     string `json:"currency,omitempty"`
		PaymentType  string `json:"payment_type,omitempty"`
		Metadata     struct {
			UserId           string `json:"user_id"`
			PropertyId       string `json:"property_id"`
			PaymentReason    string `json:"payment_reason"`
			InstallmentCount string `json:"installment_count"`
			ApplicationId    string `json:"application_id"`
			TourDate         string `json:"tour_date"`
			TourType         string `json:"tour_type"`
			VoucherCode      string `json:"voucher_code"`
		} `json:"metadata,omitempty"`
		Description string      `json:"description,omitempty"`
		PhoneNumber string      `json:"phonenumber,omitempty"`
		Status      string      `json:"status,omitempty"`
		LastError   interface{} `json:"last_error,omitempty"`

		RejectedReason  interface{} `json:"rejected_reason,omitempty"`
		RejectedBy      interface{} `json:"rejected_by,omitempty"`
		RejectedTime    interface{} `json:"rejected_time,omitempty"`
		CancelledReason interface{} `json:"cancelled_reason,omitempty"`
		CancelledBy     interface{} `json:"cancelled_by,omitempty"`
		CancelledTime   interface{} `json:"cancelled_time,omitempty"`
		Created         string      `json:"created,omitempty"`
		Author          int         `json:"author,omitempty"`
		Modified        string      `json:"modified,omitempty"`
		UpdatedBy       interface{} `json:"updated_by,omitempty"`
		StartDate       string      `json:"start_date,omitempty"`
		MfsCode         string      `json:"mfs_code,omitempty"`
	} `json:"data"`
}

type PaymentCallbackBody struct {
	Hook struct {
		ID      int    `json:"id,omitempty"`
		Created string `json:"created,omitempty"`
		Updated string `json:"updated,omitempty"`
		Event   string `json:"event,omitempty"`
		Target  string `json:"target,omitempty"`
		User    int    `json:"user,omitempty"`
	} `json:"hook"`
	Data struct {
		ID           int    `json:"id,omitempty"`
		Organization int    `json:"organization,omitempty"`
		Amount       string `json:"amount,omitempty"`
		Currency     string `json:"currency,omitempty"`
		PaymentType  string `json:"payment_type,omitempty"`
		Metadata     struct {
			UserId        string `json:"user_id"`
			PropertyId    string `json:"property_id,omitempty"`
			ApplicationId string `json:"application_id,omitempty"`
			PaymentReason string `json:"payment_reason"`
		} `json:"metadata,omitempty"`
		Description         string      `json:"description,omitempty"`
		PhoneNumber         string      `json:"phonenumber,omitempty"`
		State               string      `json:"state,omitempty"`
		LastError           interface{} `json:"last_error,omitempty"`
		RejectedReason      interface{} `json:"rejected_reason,omitempty"`
		RejectedBy          interface{} `json:"rejected_by,omitempty"`
		RemoteTransactionId string      `json:"remote_transaction_id,omitempty"`
		RejectedTime        interface{} `json:"rejected_time,omitempty"`
		CancelledReason     interface{} `json:"cancelled_reason,omitempty"`
		CancelledBy         interface{} `json:"cancelled_by,omitempty"`
		CancelledTime       interface{} `json:"cancelled_time,omitempty"`
		Created             string      `json:"created,omitempty"`
		Author              int         `json:"author,omitempty"`
		Modified            string      `json:"modified,omitempty"`
		// UpdatedBy       interface{} `json:"updated_by,omitempty"`
		StartDate string `json:"start_date,omitempty"`
		MfsCode   string `json:"mfs_code,omitempty"`
	} `json:"data"`
}

type BeyonicCollection struct {
	MaxAttempts      int     `json:"max_attempts"`
	PhoneNumber      string  `json:"phonenumber"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	SendInstructions bool    `json:"send_instructions"`
	Metadata         struct {
		UserId           string  `json:"user_id"`
		PropertyId       string  `json:"property_id"`
		PaymentReason    string  `json:"payment_reason"`
		InstallmentCount *int32  `json:"installment_count,omitempty"` // Marking it as omitempty will handle the nil case
		ApplicationId    *string `json:"application_id,omitempty"`
		TourDate         *string `json:"tour_date,omitempty"`
		TourType         *int32  `json:"tour_type,omitempty"` // Marking it as omitempty will handle the nil case
		VoucherCode      *string `json:"voucher_code,omitempty"`
	} `json:"metadata"`
}

type BeyonicPayment struct {
	PhoneNumber string          `json:"phonenumber"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	Currency    string          `json:"currency"`
	Amount      float64         `json:"amount"`
	Description string          `json:"description"`
	PaymentType string          `json:"payment_type"`
	Metadata    PaymentMetadata `json:"metadata"`
}

type PaymentMetadata struct {
	UserId        string  `json:"user_id"`
	PropertyId    *string `json:"property_id,omitempty"`
	VoucherCode   *string `json:"voucher_code,omitempty"`
	PaymentReason string  `json:"payment_reason"`
	ApplicationId *string `json:"application_id,omitempty"`
}

type PaymentMessage struct {
	Message       string `json:"message"`
	Heading       string `json:"heading"`
	UserId        string `json:"user_id,omitempty"`
	PropertyId    string `json:"property_id,omitempty"`
	Amount        string `json:"amount_paid"`
	PaymentReason string `json:"payment_reason"`
	PaymentStatus bool   `json:"payment_status"`
}

// voucher model
type Voucher struct {
	ID                 uuid.UUID           `json:"id" gorm:"type:uuid;primaryKey"`
	Code               string              `json:"code" gorm:"unique"`
	TargetProduct      string              `json:"target_product"`
	ExpiryTime         time.Time           `json:"expiry_time" gorm:"default:null"`
	UsageControl       int                 `json:"usage_control" gorm:"default:0"`
	PercentageDiscount float64             `json:"percentage_discount" gorm:"default:0"`
	RedeemAmount       float64             `json:"redeem_amount" gorm:"default:0"`
	Shareable          bool                `json:"shareable" gorm:"default:false"`
	UsersWhoUsedIt     []*UserVoucherUsage `json:"-" gorm:"foreignKey:VoucherID"`
}

type UserVoucherUsage struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	VoucherID uuid.UUID `json:"voucher_id" gorm:"type:uuid"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
	UsedAt    time.Time `json:"used_at" gorm:"default:null"`
}

type DashboardMetricsCapacityStruct struct {
	PropertyCapacity      int32 `json:"property_qty"`
	RoommateHostCapacity  int32 `json:"roommate_host_qty"`
	RentFinancingCapacity int32 `json:"rent_financing_qty"`
	UserCapacity          int32 `json:"user_qty"`
	TransactionCapacity   int32 `json:"transaction_qty"`
}
