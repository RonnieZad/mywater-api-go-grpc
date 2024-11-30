package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
)

func FindPropertiesWithExpiredSubscriptions(s *Server) {

	var subscriptionDetails []models.PaymentTransaction

	if result := s.H.DB.Where("reason = ?", "property hold").Find(&subscriptionDetails); result.Error != nil {
		fmt.Println("Error querying database:", result.Error)
	}

	for _, subscriptionDetail := range subscriptionDetails {
		if subscriptionDetail.ExpiryDate.IsZero() {
			fmt.Println("Expiry date is not set for:", subscriptionDetail.PropertyId)
			continue
		}

		if time.Now().After(subscriptionDetail.ExpiryDate) {

			propertyRef, error := s.PropertySvc.ResetHoldPropertyStatus(subscriptionDetail.PropertyId.String())
			if error != nil {
				fmt.Println("Error resetting property status:", error)
			} else if propertyRef.Status >= http.StatusNotFound {
				fmt.Print("no related application found")
			}
		}
	}

	// return expiredPropertyIDs
}

func RunPropertyStatusUpdateTask(s *Server) {
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			FindPropertiesWithExpiredSubscriptions(s)
		}
	}
}
