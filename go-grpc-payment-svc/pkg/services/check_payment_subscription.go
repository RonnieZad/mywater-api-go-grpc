package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"github.com/google/uuid"
)

func (s *Server) CheckSubscription(ctx context.Context, req *pb.CheckSubscriptionRequest) (*pb.CheckSubscriptionResponse, error) {

	var subscriptionDetail models.PaymentTransaction

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		fmt.Print("error converting to UUID")
	}

	propertyId, err := uuid.Parse(req.PropertyId)
	if err != nil {
		fmt.Print("error converting to UUID")
	}

	result := s.H.DB.Where("user_id = ? AND property_id = ? AND reason = ?", userId, propertyId, req.SubscriptionType).First(&subscriptionDetail)

	if result.RowsAffected == 0 {
		// No subscription found for the given UserId and PropertyId
		return &pb.CheckSubscriptionResponse{
			Status: http.StatusConflict,
			Error:  "Subscription not found",
		}, nil
	} else if result.Error != nil {
		// Database error
		return &pb.CheckSubscriptionResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	} else {

		// Calculate the number of days left until the subscription expiry date
		if subscriptionDetail.Reason == "property tour access" {
			accessDaysLeft := int(subscriptionDetail.ExpiryDate.Sub(time.Now()).Hours() / 24)
			accessTimeLeftString := strconv.Itoa(accessDaysLeft)
			if accessDaysLeft > 0 {
				// Subscription has not expired yet
				return &pb.CheckSubscriptionResponse{
					SubscriptionType:     subscriptionDetail.Reason,
					Status:               http.StatusOK,
					IsSubscriptionActive: true,
					DaysLeft:             accessTimeLeftString,
				}, nil
			} else {
				// Subscription has expired
				return &pb.CheckSubscriptionResponse{
					Status: http.StatusConflict,
					Error:  "Subscription has expired",
				}, nil
			}
		} else {
			daysLeft := subscriptionDetail.ExpiryDate.Sub(time.Now()).Round(time.Second)
			timeLeftString := daysLeft.String()
			if daysLeft > 0 {
				// Subscription has not expired yet
				return &pb.CheckSubscriptionResponse{
					SubscriptionType:     subscriptionDetail.Reason,
					Status:               http.StatusOK,
					IsSubscriptionActive: true,
					DaysLeft:             timeLeftString,
				}, nil
			} else {
				// Subscription has expired
				return &pb.CheckSubscriptionResponse{
					Status: http.StatusConflict,
					Error:  "Subscription has expired",
				}, nil
			}
		}
	}
}

func (s *Server) CheckPropertyHoldSubscription(ctx context.Context, req *pb.CheckPropertyHoldSubscriptionRequest) (*pb.CheckPropertyHoldSubscriptionResponse, error) {

	if req.UserId != "" {

		userId, err := uuid.Parse(req.UserId)
		if err != nil {
			return &pb.CheckPropertyHoldSubscriptionResponse{
				Status: http.StatusBadRequest,
				Error:  "Invalid user ID",
			}, nil
		}

		propertyId, err := uuid.Parse(req.PropertyId)
		if err != nil {
			return &pb.CheckPropertyHoldSubscriptionResponse{
				Status: http.StatusBadRequest,
				Error:  "Invalid property ID",
			}, nil
		}

		activeSubscription := ""
		activeTimeLeft := -1 * time.Hour

		subscriptionDetails := make([]models.PaymentTransaction, 0)
		result := s.H.DB.Where("property_id = ? AND reason = ?", propertyId, req.SubscriptionType).Find(&subscriptionDetails)
		if result.Error != nil {
			return &pb.CheckPropertyHoldSubscriptionResponse{
				Status: http.StatusInternalServerError,
				Error:  result.Error.Error(),
			}, nil
		}

		if len(subscriptionDetails) == 0 {
			return &pb.CheckPropertyHoldSubscriptionResponse{
				Status: http.StatusNotFound,
				Error:  "Subscription not found",
			}, nil
		}

		for _, subscriptionDetail := range subscriptionDetails {
			// Parse the expiry date string into a time.Time object

			if time.Now().Before(subscriptionDetail.ExpiryDate) {
				timeLeft := subscriptionDetail.ExpiryDate.Sub(time.Now())
				if subscriptionDetail.UserId == userId {
					timeLeft := subscriptionDetail.ExpiryDate.Sub(time.Now())

					if activeSubscription == "" || timeLeft > activeTimeLeft {
						activeSubscription = subscriptionDetail.Reason
						activeTimeLeft = timeLeft
					}
				} else {
					if activeSubscription == "" || timeLeft > activeTimeLeft {
						activeSubscription = subscriptionDetail.Reason
						activeTimeLeft = timeLeft
					}
					hoursLeft := int(activeTimeLeft.Hours())
					minutesLeft := int(activeTimeLeft.Minutes()) % 60
					return &pb.CheckPropertyHoldSubscriptionResponse{
						Status:   http.StatusConflict,
						Error:    "Property is held by another user",
						DaysLeft: fmt.Sprintf("%02dh%02dm", hoursLeft, minutesLeft),
					}, nil
				}
			}
		}

		if activeSubscription != "" {
			hoursLeft := int(activeTimeLeft.Hours())
			minutesLeft := int(activeTimeLeft.Minutes()) % 60

			return &pb.CheckPropertyHoldSubscriptionResponse{
				SubscriptionType:     activeSubscription,
				Status:               http.StatusOK,
				IsSubscriptionActive: true,
				DaysLeft:             fmt.Sprintf("%02dh%02dm", hoursLeft, minutesLeft),
			}, nil
		}

	} else {
		fmt.Print("we are inside the else statement")
		propertyId, err := uuid.Parse(req.PropertyId)
		if err != nil {
			return &pb.CheckPropertyHoldSubscriptionResponse{
				Status: http.StatusBadRequest,
				Error:  "Invalid property ID",
			}, nil
		}

		activeSubscription := ""
		activeTimeLeft := -1 * time.Hour

		subscriptionDetails := make([]models.PaymentTransaction, 0)
		result := s.H.DB.Where("property_id = ? AND reason = ?", propertyId, req.SubscriptionType).Find(&subscriptionDetails)
		if result.Error != nil {
			return &pb.CheckPropertyHoldSubscriptionResponse{
				Status: http.StatusInternalServerError,
				Error:  result.Error.Error(),
			}, nil
		}

		if len(subscriptionDetails) == 0 {
			return &pb.CheckPropertyHoldSubscriptionResponse{
				Status: http.StatusNotFound,
				Error:  "Subscription not found",
			}, nil
		}

		if activeSubscription != "" {
			hoursLeft := int(activeTimeLeft.Hours())
			minutesLeft := int(activeTimeLeft.Minutes()) % 60

			return &pb.CheckPropertyHoldSubscriptionResponse{
				SubscriptionType:     activeSubscription,
				Status:               http.StatusOK,
				IsSubscriptionActive: true,
				DaysLeft:             fmt.Sprintf("%02dh%02dm", hoursLeft, minutesLeft),
			}, nil
		}

	}

	return &pb.CheckPropertyHoldSubscriptionResponse{
		Status: http.StatusConflict,
		Error:  "No active subscription found",
	}, nil
}

func (s *Server) CheckPhoneSubscription(ctx context.Context, req *pb.CheckSubscriptionRequest) (*pb.CheckSubscriptionResponse, error) {
	var subscriptionDetails []models.PaymentTransaction

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		fmt.Print("error converting to UUID")
	}

	// Retrieve all subscription details for the given UserId and SubscriptionType
	result := s.H.DB.Where("user_id = ? AND reason = ?", userId, req.SubscriptionType).Find(&subscriptionDetails)

	if result.RowsAffected == 0 {
		// No subscription found for the given UserId and SubscriptionType
		return &pb.CheckSubscriptionResponse{
			Status: http.StatusConflict,
			Error:  "Subscription not found",
		}, nil
	} else if result.Error != nil {
		// Database error
		return &pb.CheckSubscriptionResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	} else {
		// Iterate through the subscription details to find the first non-expired subscription
		var firstNonExpiredSubscription *models.PaymentTransaction
		for _, subscriptionDetail := range subscriptionDetails {
			// Check if the subscription has expired
			if subscriptionDetail.ExpiryDate.After(time.Now()) {
				firstNonExpiredSubscription = &subscriptionDetail
				break
			}
		}

		if firstNonExpiredSubscription != nil {
			// Calculate the number of days left until the subscription expiry date
			daysLeft := int(firstNonExpiredSubscription.ExpiryDate.Sub(time.Now()).Hours() / 24)
			timeLeftString := strconv.Itoa(daysLeft)

			// Subscription has not expired yet
			return &pb.CheckSubscriptionResponse{
				SubscriptionType:     firstNonExpiredSubscription.Reason,
				Status:               http.StatusOK,
				IsSubscriptionActive: true,
				DaysLeft:             timeLeftString,
			}, nil
		} else {
			// All subscriptions have expired
			return &pb.CheckSubscriptionResponse{
				Status: http.StatusConflict,
				Error:  "All subscriptions have expired",
			}, nil
		}
	}
}
