package services

import (
	"context"

	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/utils"

	"github.com/google/uuid"
	"net/http"
)

func (s *Server) AddLabelAdvert(ctx context.Context, req *pb.AddLabelAdvertRequest) (*pb.AddLabelAdvertResponse, error) {
	var advert models.Advert

	applicationNumber := utils.GenerateRandomNumber()
	advertId := fmt.Sprint("MW-", applicationNumber)
	advert.Id = uuid.New()
	advert.AdvertId = advertId
	advert.AdvertiserId = req.AdvertiserId
	advert.RewardPoint = 100
	advert.AdvertPublicUrl = req.AdvertPublicUrl
	advert.Promotion = req.PromotionText
	advert.PromotionDescription = req.PromotionDescription
	advert.PictureUrl = req.PictureUrl
	advert.ExpiryDate = req.ExpiryDate

	if result := s.H.DB.Create(&advert); result.Error != nil {
		return &pb.AddLabelAdvertResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.AddLabelAdvertResponse{
		Status:  http.StatusCreated,
		Message: "Advert set up successfully",
	}, nil
}
