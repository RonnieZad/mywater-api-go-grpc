package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Server) AddLabelScan(ctx context.Context, req *pb.AddLabelScanRequest) (*pb.AddLabelScanResponse, error) {
	var labelScan models.LabelScan
	var advert models.Advert

	// Check if any records with the same AdvertId and UserId exist in the table.
	var existingRecord models.LabelScan
	err := s.H.DB.Where("user_uid = ? AND advert_id = ?", req.UserId, req.AdvertId).First(&existingRecord).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.AddLabelScanResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, err
		}
	}

	if  existingRecord.UserUid == "" && existingRecord.AdvertId == "" {
		// No matching records found, proceed to add the new record.
		err := s.H.DB.Where("advert_id = ?", req.AdvertId).First(&advert).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &pb.AddLabelScanResponse{
					Status: http.StatusNotFound,
					Error:  "No qr advert found with this id",
				}, nil
			}
			return &pb.AddLabelScanResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, err
		}

		labelScan.Id = uuid.New()
		labelScan.UserUid = req.UserId
		labelScan.AdvertiserId = advert.AdvertiserId
		labelScan.AdvertId = req.AdvertId
		labelScan.UserLatitude = req.UserLatitude
		labelScan.UserLongitude = req.UserLongitude
		advert.ScanCount++

		tx := s.H.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		if err := tx.Save(&advert).Error; err != nil {
			tx.Rollback()
			return &pb.AddLabelScanResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, err
		}

		if err := tx.Create(&labelScan).Error; err != nil {
			tx.Rollback()
			return &pb.AddLabelScanResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, err
		}

		if err := tx.Commit().Error; err != nil {
			return &pb.AddLabelScanResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, err
		}

		// Emit websocket to listening subscribed clients
		utils.SendWebsocket(ctx, "scanAdChannel")

		return &pb.AddLabelScanResponse{
			Status:  http.StatusCreated,
			Message: "Qr scanned successfully",
		}, nil
	}

	// Matching records already exist, return an error.
	return &pb.AddLabelScanResponse{
		Status: http.StatusConflict,
		Error:  "You have scanned this code already",
	}, nil
}
