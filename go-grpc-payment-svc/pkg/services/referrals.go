package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// rewardReferrer rewards the referrer with a certain amount.
func rewardReferrer(s *Server, userID string, amountToPay float64) error {
	// Parse the provided userID to a UUID
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid userID format: %v", err)
	}

	// Check if the referral exists for the given userID
	appReferral, err := getReferralByUserID(s, parsedUserID)
	if err != nil {
		return fmt.Errorf("failed to retrieve referral: %v", err)
	}

	// Check the referral status
	if appReferral.Status != models.StatusActive {
		return errors.New("referral is not in active status")
	}

	// Validate the amountToPay
	if amountToPay <= 0 {
		return errors.New("amountToPay must be greater than 0")
	}

	// Update referral details
	appReferral.Earning += amountToPay
	appReferral.Status = models.StatusSubscribed

	// Save the updated referral
	if err := s.H.DB.Save(appReferral).Error; err != nil {
		return fmt.Errorf("failed to save referral: %v", err)
	}

	return nil
}

func markPrizeWithdrawn(s *Server, userID string, amountToPay float64) error {
	// Parse the provided userID to a UUID
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid userID format: %v", err)
	}

	// Check if the referral exists for the given userID
	appReferral, err := getReferralByUserID(s, parsedUserID)
	if err != nil {
		return fmt.Errorf("failed to retrieve referral: %v", err)
	}

	// Check the referral status
	if appReferral.Status != models.StatusActive {
		return errors.New("referral is not in active status")
	}

	// Update referral details
	appReferral.Earning = 0.0
	appReferral.Status = models.StatusSubscribed

	// Save the updated referral
	if err := s.H.DB.Save(appReferral).Error; err != nil {
		return fmt.Errorf("failed to save referral: %v", err)
	}

	return nil
}

// getReferralByUserID retrieves the referral for a given userID.
func getReferralByUserID(s *Server, userID uuid.UUID) (*models.AppReferral, error) {
	appReferral := &models.AppReferral{}
	err := s.H.DB.Where("invited_user_id = ?", userID).First(appReferral).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("referral not found")
		}
		return nil, fmt.Errorf("error retrieving referral: %v", err)
	}
	return appReferral, nil
}

func (s *Server) AppReferral(ctx context.Context, req *pb.AppReferralRequest) (*pb.AppReferralResponse, error) {
	appReferralId := uuid.New()
	referralStatus := models.StatusActive

	invitedUserId, err := uuid.Parse(req.InvitedUserId)

	if err != nil {
		fmt.Print("error converting to UUID")
	}

	userId, uuidErr := uuid.Parse(req.UserId)
	if uuidErr != nil {
		fmt.Print("error converting to UUID")
		referralStatus = models.StatusPending
	}

	// Check if an AppReferral record already exists for the given InvitedUserId and DeviceUUID
	var existingAppReferral models.AppReferral
	if result := s.H.DB.Where("device_serial_number = ?", req.DeviceId).First(&existingAppReferral); result.Error != nil {
		// If not found, create a new AppReferral object with the DeviceUUID as a temporary UserId.

		appReferral := &models.AppReferral{
			ID:                 appReferralId,
			UserId:             userId,
			InvitedUserId:      invitedUserId,
			Earning:            0,
			DeviceSerialNumber: req.DeviceId,
			Status:             referralStatus,
			InvitationDate:     time.Now(),
		}

		// Save the AppReferral in the database
		if err := s.H.DB.Create(&appReferral).Error; err != nil {
			return &pb.AppReferralResponse{
				Status: http.StatusInternalServerError,
				Error:  "Failed to create app referral",
			}, nil
		}

		return &pb.AppReferralResponse{
			Status:           http.StatusCreated,
			Message:          "App referral created successfully",
			ReferralRecordId: appReferralId.String(),
		}, nil
	}

	fmt.Print("app referral id")
	fmt.Print(appReferralId.String())

	// If an AppReferral record already exists, you can choose to update its fields or ignore the request
	// Here, we'll just return the existing record's ID without making any changes.
	return &pb.AppReferralResponse{
		Status:           http.StatusOK,
		Message:          "App referral already exists",
		ReferralRecordId: appReferralId.String(),
	}, nil
}

// GetAppReferral returns the app referral for a given user
func (s *Server) GetMyAppReferral(ctx context.Context, req *pb.GetMyAppReferralRequest) (*pb.GetMyAppReferralResponse, error) {
	// Find app referral by ID
	var appReferral []models.AppReferral
	if result := s.H.DB.Where("user_id = ?", req.UserId).Find(&appReferral); result.Error != nil {
		return &pb.GetMyAppReferralResponse{
			Status: http.StatusNotFound,
			Error:  "App referral not found",
		}, nil
	}

	var appReferrals []*pb.MyAppReferrals
	totalEarning := 0.0
	for _, appReferral := range appReferral {

		userReferral := ""

		if appReferral.Status == "pending" {
			userReferral = "Pending"
		} else if appReferral.Status == "active" {
			userReferral = "Pending Payment"
		} else if appReferral.Status == "subscribed" {
			userReferral = "Paid"
		}

		user, error := s.AuthSvc.GetUser(appReferral.InvitedUserId.String())

		userName := "-"
		if error != nil {
			fmt.Print(error)
			userName = "-"

		} else if user.Status >= http.StatusNotFound {
			fmt.Print("no related user found")
			userName = "User"
		} else {
			userName = user.Data.Name
		}

		appReferrals = append(appReferrals, &pb.MyAppReferrals{
			InvitedUserId: userName,
			Status:        userReferral,
			Earning:       appReferral.Earning,
			ReferralDate:  appReferral.InvitationDate.Format("2006-01-02"),
			Currency:      "UGX",
		})

		// Calculate the total earning by summing up the Earning field
		totalEarning += appReferral.Earning

	}

	return &pb.GetMyAppReferralResponse{
		TotalEarning:   totalEarning,
		Message:        "App referral found",
		Status:         http.StatusOK,
		MyAppReferrals: appReferrals,
	}, nil
}

// UpdateReferreral allows users to change their phone number.
func (s *Server) UpdateReferreral(ctx context.Context, req *pb.UpdateReferreralRequest) (*pb.UpdateReferreralResponse, error) {
	// Validate ReferralRecordId and UserId
	recordID, err := uuid.Parse(req.ReferralRecordId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ReferralRecordId: %v", err)
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UserId: %v", err)
	}

	// Check if the referral record exists
	var referralRecord models.AppReferral
	result := s.H.DB.Where(&models.AppReferral{ID: recordID}).First(&referralRecord)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Referral Not Found")
		}
		return nil, status.Errorf(codes.Internal, "failed to fetch referral: %v", result.Error)
	}

	// Update referral details
	referralRecord.Status = models.StatusActive
	referralRecord.UserId = userID

	// Set earning if payment is made
	if req.IsPaymentMade {
		const earningAmount = 10000 // You can set this to the appropriate value
		referralRecord.Earning = earningAmount
	}

	// Save the updated referral
	if result := s.H.DB.Save(&referralRecord); result.Error != nil {
		return nil, status.Errorf(codes.Internal, "failed to update referral: %v", result.Error)
	}

	return &pb.UpdateReferreralResponse{
		Status:  int64(codes.OK),
		Message: "Special Invite has been set up",
	}, nil
}
