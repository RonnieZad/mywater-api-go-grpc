package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/utils"
	"github.com/ably/ably-go/ably"
	"github.com/google/uuid"
)

// payment callback method here
func (s *Server) PaymentCallback(ctx context.Context, req *pb.PaymentCallbackRequest) (*pb.PaymentCallbackResponse, error) {

	var callback models.PaymentCallbackBody

	collectionCallbackJson, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the CollectionCallbackJson
	if err := json.Unmarshal(collectionCallbackJson, &callback); err != nil {
		fmt.Println(err)
	}

	client, error := ably.NewRealtime(ably.WithKey("EIgGUg.TJrcFQ:ucexDmzPxVwZ5VQDrBDGgK3bLjHsWXQsjZTCcoV83Bg"))
	if error != nil {
		panic(error)
	}

	channel := client.Channels.Get("fundPaymentEvents")

	// Check if the callback event is "payment.status.changed"
	if callback.Data.State == "processed" {

		if callback.Data.Metadata.PaymentReason == "landlord payment" {

			fmt.Println("application id: ", callback.Data.Metadata.ApplicationId)
			applyPropertyLock, lockError := s.ApplicationSvc.ManageUserLock(callback.Data.Metadata.ApplicationId)

			if lockError != nil {
				fmt.Print(lockError)
			} else if applyPropertyLock.Status >= http.StatusNotFound {
				fmt.Print("issue on creating lock card")
			}

			amount, err := strconv.ParseFloat(callback.Data.Amount, 64)
			if err != nil {
				fmt.Print("failed to convert amount")
				// handle error
			}

			// Format the float64 without decimal points as a string
			formattedAmountStr := strconv.FormatFloat(amount, 'f', 0, 64)

			// Create a new PaymentMessage object
			msg := &models.PaymentMessage{
				Message:       "Congratulations, you are almost there. Check out next steps for property location and access.\n\nHappy stay,\nEnyumba Team.",
				Heading:       "Your Rent payment is successful",
				UserId:        callback.Data.Metadata.UserId,
				PropertyId:    callback.Data.Metadata.PropertyId,
				Amount:        formattedAmountStr,
				PaymentReason: callback.Data.Metadata.PaymentReason,
				PaymentStatus: true,
			}

			message := fmt.Sprintf("Hello our esteemed Landlord, you have received a payment towards your house property, our Enyumba client will be moving in soon. Thank you for using Enyumba.")

			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}

			user, error := s.AuthSvc.GetUser(callback.Data.Metadata.UserId)

			if error != nil {
				fmt.Print(error)

			} else if user.Status >= http.StatusNotFound {
				fmt.Print("no related user found")
			}

			mailErrr := utils.SendEmail(user.Data.EmailAddress, "Landlord payment successful", "Congratulations", "We are pleased to inform you that your payment to your landlord has been successfully processed. To ensure a smooth experience, please open the Enyumba App, navigate to your property listing, and click 'Get Directions' to find the easiest route to your property. Upon arrival, use the app to confirm your arrival and any necessary details. If you have questions or need assistance, contact our customer support team at support@enyumba.com. Thank you for choosing Enyumba; we look forward to serving you again.")
			if mailErrr != nil {
				// Handle error
			}

			// Convert the message to a JSON string
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				panic(err)
			}

			error = channel.Publish(ctx, "fundPaymentEvents", string(msgJSON))
			if error != nil {
				panic(error)
			}

			// Get the property ID from the request.
			propertyId, err := uuid.Parse(callback.Data.Metadata.PropertyId)
			if err != nil {
				fmt.Print("failed to parse property id")
			}

			fmt.Println(callback.Data.Metadata)
			layout := "2006-01-02T15:04:05Z" // the format of the input string date
			t, err := time.Parse(layout, callback.Data.Created)
			if err != nil {
				panic(err)
			}
			t = t.AddDate(0, 0, 30) // add 30 days to the date
			expiryDate := t.Format(layout)

			userId, err := uuid.Parse(callback.Data.Metadata.UserId)
			if err != nil {
				fmt.Print("error converting to UUID")
			}

			applicationRef, error := s.ApplicationSvc.UpdateApplicationLandlordPaymentStatus(callback.Data.Metadata.ApplicationId, true)

			if error != nil {
				fmt.Print(error)
			} else if applicationRef.Status >= http.StatusNotFound {
				fmt.Println("no related application found")
			}

			// Parse the expiry date string to a time.Time object
			expiryDateTime, err := time.Parse(layout, expiryDate)
			if err != nil {
				fmt.Println("Error parsing expiry date:", err)
				// Handle the error accordingly (e.g., return an error response)
			}

			txnId, _ := strconv.Atoi(callback.Data.RemoteTransactionId)

			paymentTransaction := &models.PaymentTransaction{
				Id:              uuid.New(),
				PropertyId:      propertyId,
				UserId:          userId,
				Amount:          amount,
				Currency:        callback.Data.Currency,
				TransactionType: "Mobile Money",
				PhoneNumber:     "+256782333080",
				Reason:          callback.Data.Metadata.PaymentReason,
				TransactionDate: callback.Data.Created,
				TransactionId:   int32(txnId),
				ExpiryDate:      expiryDateTime,
				Status:          "Landlord payment successful",
			}

			s.H.DB.Create(&paymentTransaction)
		} else if callback.Data.Metadata.PaymentReason == "referral cashout" {
			markPrizeWithdrawn(s, callback.Data.Metadata.UserId, 0.0)
		}

	}

	// Check if the callback event is "payment.status.changed"
	if callback.Data.State == "processed_with_errors" {
		if callback.Data.Metadata.PaymentReason == "landlord payment" {

			amount, err := strconv.ParseFloat(callback.Data.Amount, 64)
			if err != nil {
				fmt.Print("failed to convert amount")
				// handle error
			}

			// Format the float64 without decimal points as a string
			formattedAmountStr := strconv.FormatFloat(amount, 'f', 0, 64)

			// Create a new PaymentMessage object
			msg := &models.PaymentMessage{
				Message:       "Your payment is unsuccessful. Check your balance and try again later. If it persists, please contact our support team.\n\nRegards,\nEnyumba Team.",
				Heading:       "Landord Payment has failed",
				UserId:        callback.Data.Metadata.UserId,
				PropertyId:    callback.Data.Metadata.PropertyId,
				Amount:        formattedAmountStr,
				PaymentReason: callback.Data.Metadata.PaymentReason,
				PaymentStatus: true,
			}

			// Convert the message to a JSON string
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				panic(err)
			}

			error = channel.Publish(ctx, "fundPaymentEvents", string(msgJSON))
			if error != nil {
				panic(error)
			}

			propertyId, err := uuid.Parse(callback.Data.Metadata.PropertyId)
			if err != nil {
				// handle error
			}

			userId, err := uuid.Parse(callback.Data.Metadata.UserId)
			if err != nil {
				fmt.Print("error converting to UUID")
			}

			txnId, _ := strconv.Atoi(callback.Data.RemoteTransactionId)

			paymentTransaction := &models.PaymentTransaction{
				Id:              uuid.New(),
				PropertyId:      propertyId,
				UserId:          userId,
				Amount:          amount,
				Currency:        callback.Data.Currency,
				TransactionType: "Mobile Money",
				PhoneNumber:     callback.Data.PhoneNumber,
				Reason:          callback.Data.Description,
				TransactionDate: callback.Data.Created,
				TransactionId:   int32(txnId),
				Status:          "landlord payment failed",
			}
			s.H.DB.Create(&paymentTransaction)
		}
	}

	return &pb.PaymentCallbackResponse{
		Status: http.StatusCreated,
	}, nil
}
