package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
)

type GeocodingResponse struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
	} `json:"results"`
	Status string `json:"status"`
}

func ReverseGeocode(latitude, longitude float64) (string, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%f,%f&key=AIzaSyAnVPTpHLJa_3numhZsFeNPTmv_YH9CyZ4", latitude, longitude)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data GeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	fmt.Println(data)

	if len(data.Results) > 0 {
		for _, component := range data.Results[0].AddressComponents {
			for _, t := range component.Types {
				if t == "locality" {
					return component.LongName, nil
				}
			}
		}
	}

	return "", fmt.Errorf("city not found")
}

func (s *Server) GetUserLabelAdvert(ctx context.Context, req *pb.GetUserLabelAdvertRequest) (*pb.GetUserLabelAdvertResponse, error) {
	// Find user by ID
	var labelAdverts []models.Advert
	var labelScans []models.LabelScan

	if result := s.H.DB.Where("advertiser_id = ?", req.AdvertiserId).Order("created_at DESC").Find(&labelAdverts); result.Error != nil {
		return &pb.GetUserLabelAdvertResponse{
			Status: http.StatusNotFound,
			Error:  "You have no running promotions",
		}, nil
	}

	cityScanCounts := make(map[string]int)

	// Convert payment transactions to protobuf message format
	var labelScansPB []*pb.LabelAdvert
	for _, advertData := range labelAdverts {
		if result := s.H.DB.Where("advert_id = ?", advertData.AdvertId).Find(&labelScans); result.Error != nil {
			return &pb.GetUserLabelAdvertResponse{
				Status: http.StatusNotFound,
				Error:  "You have no running promotions",
			}, nil
		}

		for _, labelScan := range labelScans {
			if labelScan.UserLatitude != "" {
				userLatitude, err := strconv.ParseFloat(labelScan.UserLatitude, 64)
				if err != nil {
					fmt.Println("Error converting latitude to float64:", err)
					continue
				}

				userLongitude, err := strconv.ParseFloat(labelScan.UserLongitude, 64)
				if err != nil {
					fmt.Println("Error converting longitude to float64:", err)
					continue
				}

				location, err := ReverseGeocode(userLatitude, userLongitude)
				if err != nil {
					location = "Unknown"
				}

				cityScanCounts[location] += int(advertData.ScanCount)
			}
		}

		var cityScans []*pb.CityScan
		for city, scanCount := range cityScanCounts {
			cityScans = append(cityScans, &pb.CityScan{
				City:      city,
				ScanCount: int32(scanCount),
			})
		}

		labelScansPB = append(labelScansPB, &pb.LabelAdvert{
			AdvertId:             advertData.AdvertId,
			AdvertPublicUrl:      advertData.AdvertPublicUrl,
			PromotionText:        advertData.Promotion,
			PromotionDescription: advertData.PromotionDescription,
			PictureUrl:           advertData.PictureUrl,
			ExpiryDate:           advertData.ExpiryDate,
			ScanCount:            advertData.ScanCount,
			CreationDate:         advertData.CreatedAt.Format("2006-01-02 15:04"),
			CityScans:            cityScans,
		})
	}

	return &pb.GetUserLabelAdvertResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d adverts", len(labelScansPB)),
		Adverts: labelScansPB,
	}, nil
}
