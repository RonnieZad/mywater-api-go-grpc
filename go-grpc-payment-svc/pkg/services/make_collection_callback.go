package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/utils"
	"github.com/ably/ably-go/ably"
	"github.com/google/uuid"
)

func (s *Server) CollectionCallback(ctx context.Context, req *pb.CollectionCallbackRequest) (*pb.CollectionCallbackResponse, error) {

	var callback models.CollectionCallbackBody

	collectionCallbackJson, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(collectionCallbackJson)

	// Unmarshal the CollectionCallbackJson
	if err := json.Unmarshal(collectionCallbackJson, &callback); err != nil {
		fmt.Println(err)
	}

	client, error := ably.NewRealtime(ably.WithKey("EIgGUg.TJrcFQ:ucexDmzPxVwZ5VQDrBDGgK3bLjHsWXQsjZTCcoV83Bg"))
	if error != nil {
		panic(error)
	}
	channel := client.Channels.Get("fundCollectionEvents")

	// Check if the callback event is "payment.status.changed"
	if callback.Data.Status == "successful" {

		fmt.Println("fundCollection Events successfully received")

		amount, err := strconv.ParseFloat(callback.Data.Amount, 64)
		if err != nil {
			fmt.Print("failed to convert amount")
			// handle error
		}

		//Check if user applied a voucher code
		if callback.Data.Metadata.VoucherCode != "" {
			redeemVoucher(s, callback.Data.Metadata.VoucherCode, callback.Data.Metadata.UserId, amount)
		}

		rewardReferrer(s, callback.Data.Metadata.UserId, amount*0.1)

		// Format the float64 without decimal points as a string
		formattedAmountStr := strconv.FormatFloat(amount, 'f', 0, 64)
		message := ""
		heading := ""

		user, error := s.AuthSvc.GetUser(callback.Data.Metadata.UserId)

		if error != nil {
			fmt.Print(error)

		} else if user.Status >= http.StatusNotFound {
			fmt.Print("no related user found")
		}

		if callback.Data.Metadata.PaymentReason == "property hold" {
			heading = "Your payment is successful"
			message = "We have reserved this property just for you. You have 12 hours to take action before it is released from your watch\n\nRegards,\nEnyumba Team"
			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		} else if callback.Data.Metadata.PaymentReason == "property tour access" {
			heading = "Your payment is successful"
			message = "Thank you for subsribing for property tour for this property, your payment was successful"
			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		} else if callback.Data.Metadata.PaymentReason == "phone number access" {
			heading = "Your payment is successful"
			message = "You can now access landlord phone numbers for all properties valid for a period of 30days"
			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		} else if callback.Data.Metadata.PaymentReason == "roommate listing" {
			heading = "Your payment is successful"
			message = "Thank you for your application. Our team will review it and call to confirm the photo capture exercise and assessment."
			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		} else if callback.Data.Metadata.PaymentReason == "landlord payment" {
			heading = "Your payment is successful"
			message = "Your landlord payment was received successfully"
			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		} else if callback.Data.Metadata.PaymentReason == "loan repayment" {
			heading = "Your payment is successful"
			message = "Your loan repayment was received successfully"
			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		} else if callback.Data.Metadata.PaymentReason == "security deposit" {
			heading = "Your payment is successful"
			message = "Your security deposit has been received successfully"
			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		} else if callback.Data.Metadata.PaymentReason == "landlord direct payment" {
			heading = "Your payment is successful"
			message = "Your landlord payment was received successfully"

			// func addLandlordRevenue(s *Server, landlordIdStr string, propertyIdStr string, transactionId int32, transactionDate string, amount float64, status string) error {

			//make record of landlord revenue entry
			addLandlordRevenue(s, callback.Data.Metadata.UserId, callback.Data.Metadata.PropertyId, callback.Data.ID, callback.Data.Created, amount, "Successful")

			mailErrr := utils.SendEmail(user.Data.EmailAddress, heading, "Congratulations", message)
			if mailErrr != nil {
				// Handle error
			}
			smsErrr := utils.SendSMS(callback.Data.PhoneNumber, message)
			if smsErrr != nil {
				// Handle error
			}
		}

		// Create a new PaymentMessage object
		msg := &models.PaymentMessage{
			Message:       message,
			Heading:       heading,
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

		error = channel.Publish(ctx, "fundCollectionEvents", string(msgJSON))
		if error != nil {
			panic(error)
		}

		//Get the property ID from the request.
		propertyId, err := uuid.Parse(callback.Data.Metadata.PropertyId)
		if err != nil {
			fmt.Print("failed to parse property id")
		}

		layout := "2006-01-02T15:04:05Z" // the format of the input string date
		t, err := time.Parse(layout, callback.Data.Created)
		if err != nil {
			panic(err)
		}

		expiryDate := ""

		if strings.Contains(callback.Data.Metadata.PaymentReason, "access") {
			t = t.AddDate(0, 0, 30) // add 30 days to the date
			expiryDate = t.Format(layout)
		}

		if callback.Data.Metadata.PaymentReason == "property hold" {
			t = t.Add(12 * time.Hour) // add 12 hours to the date
			expiryDate = t.Format(layout)
		}

		userId, err := uuid.Parse(callback.Data.Metadata.UserId)
		if err != nil {
			fmt.Print("error converting to UUID")
		}

		// Parse the expiry date string to a time.Time object
		expiryDateTime, err := time.Parse(layout, expiryDate)
		if err != nil {
			fmt.Println("Error parsing expiry date:", err)
			// Handle the error accordingly (e.g., return an error response)
		}

		paymentTransaction := &models.PaymentTransaction{
			Id:              uuid.New(),
			PropertyId:      propertyId,
			UserId:          userId,
			Amount:          amount,
			Currency:        callback.Data.Currency,
			TransactionType: "Mobile Money",
			PhoneNumber:     callback.Data.PhoneNumber,
			Reason:          callback.Data.Metadata.PaymentReason,
			TransactionDate: callback.Data.Created,
			TransactionId:   callback.Data.ID,
			ExpiryDate:      expiryDateTime,
			Status:          "Received",
		}

		s.H.DB.Create(&paymentTransaction)

		if callback.Data.Metadata.PaymentReason == "property hold" {

			//mark propety held
			propertyHoldStatus, err := s.PropertySvc.UpdateHoldPropertyStatus(callback.Data.Metadata.PropertyId, true)

			// UpdateHoldPropertyStatus
			if error != nil {
				fmt.Print(error)
			} else if propertyHoldStatus.Status >= http.StatusNotFound {
				fmt.Print("no related property found")
			}

			//hold property
			holdProperty, err := s.PropertySvc.HoldProperty(callback.Data.Metadata.UserId, callback.Data.Metadata.PropertyId)

			if err == nil {
				//handle error logic here
			}

			fmt.Print(holdProperty)
		}

		if callback.Data.Metadata.PaymentReason == "loan repayment" {

			//loan repayment here
			installmentCount, err := strconv.Atoi(callback.Data.Metadata.InstallmentCount)
			if err != nil {
				fmt.Println("Error:", err)
			}

			applicationDetail, err := s.ApplicationSvc.UpdateApplicationRepaymentDetails(callback.Data.Metadata.ApplicationId, int32(installmentCount), amount, true)

			if err != nil {
				fmt.Print(err)
			} else if applicationDetail.Status >= http.StatusNotFound {
				fmt.Print("no related application found")
			}
		}

		if callback.Data.Metadata.PaymentReason == "property tour access" {

			t = t.AddDate(0, 0, 30) // add 30 days to the date
			expiryDate = t.Format(layout)

			//mark propety held
			tourRef, err := s.ApplicationSvc.MakeTourRequestApplication(callback.Data.Metadata.PropertyId, callback.Data.Metadata.UserId, 1, expiryDate, true)

			// UpdateHoldPropertyStatus
			if err != nil {
				fmt.Print(err)
			} else if tourRef.Status >= http.StatusNotFound {
				fmt.Print("no related tour application found")
			}
		}

		if callback.Data.Metadata.PaymentReason == "roommate listing" {

			//update roommate payment status
			roomMateApplcationUpdate, err := s.ApplicationSvc.RoomMateApplicationUpdate(callback.Data.Metadata.UserId, true)

			// UpdateHoldPropertyStatus
			if err != nil {
				fmt.Print(err)
			} else if roomMateApplcationUpdate.Status >= http.StatusNotFound {
				fmt.Print("no related property found")
			}
		}

		if callback.Data.Metadata.PaymentReason == "security deposit" {

			//mark security deposit paid
			applicationDetail, err := s.ApplicationSvc.UpdateApplicationSecurityDepositPaidStatus(callback.Data.Metadata.ApplicationId, true)

			if err != nil {
				fmt.Print(err)
			} else if applicationDetail.Status >= http.StatusNotFound {
				fmt.Print("no related application found")
			}

			//set property availability to false
			propertyRef, error := s.PropertySvc.UpdatePropertyAvailabiltyStatus(callback.Data.Metadata.PropertyId, callback.Data.Metadata.UserId, false)

			if error != nil {
				fmt.Print(error)
			} else if propertyRef.Status >= http.StatusNotFound {
				fmt.Println("no related property found")
			}

			//automatically approve application
			approveRentApplicationDetail, err := s.ApplicationSvc.ApproveRentApplication(callback.Data.Metadata.UserId, callback.Data.Metadata.ApplicationId, 2)

			if err != nil {
				fmt.Print(err)
			} else if approveRentApplicationDetail.Status >= http.StatusNotFound {
				fmt.Print("no related application found")
			}
		}
	}

	// Check if the callback event is "payment.status.changed"
	if callback.Data.Status == "failed" {

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
			Heading:       "Your payment has failed",
			UserId:        callback.Data.Metadata.UserId,
			PropertyId:    callback.Data.Metadata.PropertyId,
			Amount:        formattedAmountStr,
			PaymentReason: callback.Data.Metadata.PaymentReason,
			PaymentStatus: false,
		}

		// Convert the message to a JSON string
		msgJSON, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		error = channel.Publish(ctx, "fundCollectionEvents", string(msgJSON))
		if error != nil {
			panic(error)
		}

		propertyId, err := uuid.Parse(callback.Data.Metadata.PropertyId)
		if err != nil {
			// handle error
		}

		fmt.Println(callback.Data.Metadata)

		userId, err := uuid.Parse(callback.Data.Metadata.UserId)
		if err != nil {
			fmt.Print("error converting to UUID")
		}

		paymentTransaction := &models.PaymentTransaction{
			Id:              uuid.New(),
			PropertyId:      propertyId,
			UserId:          userId,
			Amount:          amount,
			Currency:        callback.Data.Currency,
			TransactionType: "Mobile Money",
			PhoneNumber:     callback.Data.PhoneNumber,
			Reason:          callback.Data.Metadata.PaymentReason,
			TransactionDate: callback.Data.Created,
			TransactionId:   callback.Data.ID,
			Status:          "Failed",
		}
		s.H.DB.Create(&paymentTransaction)
	}

	return &pb.CollectionCallbackResponse{
		Status: http.StatusCreated,
		// Id:     order.Id,
	}, nil
}
