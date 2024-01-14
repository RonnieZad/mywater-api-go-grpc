package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
)

func (s *Server) GetUserLabelScan(ctx context.Context, req *pb.GetUserLabelScanRequest) (*pb.GetUserLabelScanResponse, error) {
	// Find user by ID
	var labelScans []models.LabelScan
	var advert models.Advert
	

	if result := s.H.DB.Where("user_uid = ?", req.UserId).Order("created_at DESC").Find(&labelScans); result.Error != nil {
		return &pb.GetUserLabelScanResponse{
			Status: http.StatusNotFound,
			Error:  "No scans",
		}, nil
	}

	totalPoints := int64(0)
	// Convert payment transactions to protobuf message format
	var labelScansPB []*pb.LabelAdvert
	for _, labelScanData := range labelScans {

		if result := s.H.DB.Where("advert_id = ?", labelScanData.AdvertId).Find(&advert); result.Error != nil {
			return &pb.GetUserLabelScanResponse{
				Status: http.StatusNotFound,
				Error:  "No scans",
			}, nil
		}

		totalPoints += advert.RewardPoint

		user, error := s.AuthSvc.GetUser(labelScanData.AdvertiserId)

		if error != nil {
			fmt.Print(error)

		} else if user.Status >= http.StatusNotFound {
			fmt.Print("no related user found")
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
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d scans", len(labelScansPB)),
		Scans:   labelScansPB,
		TotalPoints:  totalPoints,
	}, nil
}
