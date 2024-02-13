package services

import (
	"context"

	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/utils"
	"github.com/google/uuid"
)

func (s *Server) AddLabelScan(ctx context.Context, req *pb.AddLabelScanRequest) (*pb.AddLabelScanResponse, error) {
	var labelScan models.LabelScan
	var advert models.Advert

	// Check if any records with the same AdvertId and UserId exist in the table.
	var existingRecords []models.LabelScan
	if result := s.H.DB.Where("user_uid = ? AND advert_id = ?", req.UserId, req.AdvertId).Find(&existingRecords); result.Error != nil {
		return &pb.AddLabelScanResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}

	if len(existingRecords) == 0 {

		if result := s.H.DB.Where("advert_id = ?", req.AdvertId).Find(&advert); result.Error != nil {
			return &pb.AddLabelScanResponse{
				Status: http.StatusNotFound,
				Error:  "No qr advert found with this id",
			}, nil
		}

		// No matching records found, proceed to add the new record.
		labelScan.Id = uuid.New()
		labelScan.UserUid = req.UserId
		labelScan.AdvertiserId = advert.AdvertiserId
		labelScan.AdvertId = req.AdvertId
		advert.ScanCount += 1

		s.H.DB.Save(&advert)

		//emit websocket to listening subscribed clients
		utils.SendWebsocket(ctx, "scanAdChannel")

		if result := s.H.DB.Create(&labelScan); result.Error != nil {
			return &pb.AddLabelScanResponse{
				Status: http.StatusConflict,
				Error:  result.Error.Error(),
			}, nil
		}

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
