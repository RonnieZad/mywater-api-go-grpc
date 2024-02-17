package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
)

func (s *Server) GetUserLabelScan(ctx context.Context, req *pb.GetUserLabelScanRequest) (*pb.GetUserLabelScanResponse, error) {
	// Find user by ID
	var labelScans []models.LabelScan

	if result := s.H.DB.Where("user_uid = ?", req.UserId).Order("created_at DESC").Find(&labelScans); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &pb.GetUserLabelScanResponse{
				Status: http.StatusNotFound,
				Error:  "No scans",
			}, nil
		}
		return nil, result.Error
	}

	var labelScansPB []*pb.LabelAdvert
	totalPoints := int64(0)

	for _, labelScanData := range labelScans {
		var advert models.Advert
		if result := s.H.DB.Where("advert_id = ?", labelScanData.AdvertId).First(&advert); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				continue // Skip this scan if the advert is not found
			}
			return nil, result.Error
		}

		totalPoints += advert.RewardPoint

		user, err := s.AuthSvc.GetUser(labelScanData.AdvertiserId)
		if err != nil {
			return nil, err
		}

		labelScansPB = append(labelScansPB, &pb.LabelAdvert{
			AdvertId:        labelScanData.AdvertId,
			AdvertPublicUrl: advert.AdvertPublicUrl,
			PromotionText:   advert.Promotion,
			PictureUrl:      advert.PictureUrl,
			ExpiryDate:      advert.ExpiryDate,
			CompanyName:     user.Data.CompanyName,
			CompanyLogo:     user.Data.CompanyLogo,
			CompanyWebsite:  user.Data.CompanyWebsite,
			CompanyEmail:    user.Data.CompanyEmail,
			CompanyPhone:    user.Data.CompanyPhone,
			CompanyAddress:  user.Data.CompanyAddress,
			ScanDate:        labelScanData.CreatedAt.Format("2006-01-02 15:04"),
		})
	}

	return &pb.GetUserLabelScanResponse{
		Status:      http.StatusOK,
		Message:     fmt.Sprintf("Found %d scans", len(labelScansPB)),
		Scans:       labelScansPB,
		TotalPoints: totalPoints,
	}, nil
}
