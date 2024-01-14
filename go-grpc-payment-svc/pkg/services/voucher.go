package service

import (
	"context"
	"fmt"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"gorm.io/gorm"
	"net/http"
	"time"
	"github.com/google/uuid"
)


// Validate and redeem a voucher code
func redeemVoucher(s *Server, voucherCode string, userId string, amountToPay float64) {

	voucher := &models.Voucher{}
	err := s.H.DB.Where("code = ?", voucherCode).First(&voucher).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Print(err)
		}
		fmt.Print("Voucher not found")

	}

	// Check if the voucher has expired
	if time.Now().After(voucher.ExpiryTime) {
		fmt.Print("Voucher is expired")
	}

	// Check if the voucher has reached its usage limit
	if voucher.UsageControl > 0 {
		// Decrement the usage control by 1
		voucher.UsageControl--

		// Update the voucher in the database
		err := s.H.DB.Save(&voucher).Error
		if err != nil {
			fmt.Print(err)
		}

		// Create a new UserVoucherUsage object
		userVoucherUsage := &models.UserVoucherUsage{
			ID:        uuid.New(),
			UserID:    uuid.MustParse(userId),
			VoucherID: voucher.ID,
			IsUsed:    true,
			UsedAt:    time.Now(),
		}

		// Save the UserVoucherUsage object to the database
		err = s.H.DB.Save(&userVoucherUsage).Error
		if err != nil {
			fmt.Print(err)
		}

		// Check if the voucher has been fully redeemed
		if voucher.UsageControl == 0 {
			// Delete the voucher from the database
			err := s.H.DB.Delete(&voucher).Error
			if err != nil {
				fmt.Print("Voucher is invalid")
			}
		}
	}

	// Calculate the amount to pay after applying the voucher
	amountToPayAfterVoucher := amountToPay - voucher.RedeemAmount

	fmt.Print(amountToPayAfterVoucher)
}

// verify voucher
func (s *Server) CheckVoucherValidity(ctx context.Context, req *pb.CheckVoucherValidityRequest) (*pb.CheckVoucherValidityResponse, error) {

	// Retrieve the voucher from the database
	voucher := &models.Voucher{}
	err := s.H.DB.Where("code = ?", req.VoucherCode).First(&voucher).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.CheckVoucherValidityResponse{
				IsVoucherValid: false,
				Status:         http.StatusConflict,
				Error:          "Voucher not found",
			}, nil
		}
		return &pb.CheckVoucherValidityResponse{
			IsVoucherValid: false,
			Status:         http.StatusInternalServerError,
			Error:          "Failed to retrieve voucher",
		}, nil
	}

	// Parse the user ID from string to uuid.UUID
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.CheckVoucherValidityResponse{
			IsVoucherValid: false,
			Status:         http.StatusBadRequest,
			Error:          "Invalid user ID",
		}, nil
	}

	// Check if the voucher has expired
	if time.Now().After(voucher.ExpiryTime) {
		return &pb.CheckVoucherValidityResponse{
			IsVoucherValid: false,
			Status:         http.StatusConflict,
			Error:          "Voucher has expired",
		}, nil
	}

	// Check if the voucher has been fully redeemed
	if voucher.UsageControl > 0 {
		// Check if the current user has used the voucher
		var usage models.UserVoucherUsage
		err := s.H.DB.Where("user_id = ? AND voucher_id = ?", userId, voucher.ID).First(&usage).Error
		if err == nil && usage.IsUsed {
			return &pb.CheckVoucherValidityResponse{
				IsVoucherValid: false,
				Status:         http.StatusConflict,
				Error:          "You already used this voucher",
			}, nil
		}
	}

	// Calculate the amount to pay after applying the voucher
	amountToPayAfterVoucher := req.AmountToPay - voucher.RedeemAmount

	// If the voucher is allowed to be used, mark it as used for the current user
	if voucher.UsageControl == 0 {
		usage := models.UserVoucherUsage{
			ID:        uuid.New(),
			UserID:    userId,
			VoucherID: voucher.ID,
			IsUsed:    true,
			UsedAt:    time.Now(),
		}
		err := s.H.DB.Create(&usage).Error
		if err != nil {
			return &pb.CheckVoucherValidityResponse{
				IsVoucherValid: false,
				Status:         http.StatusInternalServerError,
				Error:          "Failed to mark voucher as used",
			}, nil
		}
	}

	return &pb.CheckVoucherValidityResponse{
		IsVoucherValid:          true,
		PercentageDiscount:      voucher.PercentageDiscount,
		AmountToPay:             req.AmountToPay,
		VoucherAmount:           voucher.RedeemAmount,
		AmountToPayAfterVoucher: amountToPayAfterVoucher,
		Status:                  http.StatusOK,
		Message:                 "Voucher is valid",
	}, nil
}