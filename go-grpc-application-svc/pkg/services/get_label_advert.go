package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
)

func (s *Server) GetUserLabelAdvert(ctx context.Context, req *pb.GetUserLabelAdvertRequest) (*pb.GetUserLabelAdvertResponse, error) {
	// Find user by ID
	var labelAdverts []models.Advert

	if result := s.H.DB.Where("advertiser_id = ?", req.AdvertiserId).Order("created_at DESC").Find(&labelAdverts); result.Error != nil {
		return &pb.GetUserLabelAdvertResponse{
			Status: http.StatusNotFound,
			Error:  "You have no running promotions",
		}, nil
	}

	// Convert payment transactions to protobuf message format
	var labelScansPB []*pb.LabelAdvert
	for _, advertData := range labelAdverts {

		labelScansPB = append(labelScansPB, &pb.LabelAdvert{
			AdvertId:             advertData.AdvertId,
			AdvertPublicUrl:      advertData.AdvertPublicUrl,
			PromotionText:        advertData.Promotion,
			PromotionDescription: advertData.PromotionDescription,
			PictureUrl:           advertData.PictureUrl,
			ExpiryDate:           advertData.ExpiryDate,
			ScanCount:            advertData.ScanCount,
			CreationDate:         advertData.CreatedAt.Format("2006-01-02 15:04"),
		})
	}

	return &pb.GetUserLabelAdvertResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d adverts", len(labelScansPB)),
		Adverts: labelScansPB,
	}, nil
}
