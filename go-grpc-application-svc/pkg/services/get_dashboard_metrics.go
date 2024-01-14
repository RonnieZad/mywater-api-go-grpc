package services

import (
	"context"
	"fmt"
	"time"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/structpb"
)

// Get number of available property types for each category
func (s *Server) GetCompanyDashboardAnalytic(ctx context.Context, req *pb.GetCompanyDashboardAnalyticRequest) (*pb.GetCompanyDashboardAnalyticResponse, error) {

	advertisements, err := s.getAdvertisementsForAdvertiser(req.AdvertiserId)
    if err != nil {
        return nil, err
    }

    // Create a map to store statistics for each advertisement
	advertStats := make([]*structpb.Struct, 0)

	for _, advert := range advertisements {
        dailyCounts, err := s.getLabelScansCountForDay(req.AdvertiserId, advert.AdvertId)
        if err != nil {
            return nil, err
        }

        weeklyCounts, err := s.getLabelScansCountWeekly(req.AdvertiserId, advert.AdvertId)
        if err != nil {
            return nil, err
        }

        monthlyCounts, err := s.getLabelScansCountMonthly(req.AdvertiserId, advert.AdvertId)
        if err != nil {
            return nil, err
        }

        yearlyCounts, err := s.getLabelScansCountYearly(req.AdvertiserId, advert.AdvertId)
        if err != nil {
            return nil, err
        }

        // Convert map values to structpb.Values
        dailyValues := mapToStructpbValues(dailyCounts)
        weeklyValues := mapToStructpbValues(weeklyCounts)
        monthlyValues := mapToStructpbValues(monthlyCounts)
        yearlyValues := mapToStructpbValues(yearlyCounts)

        // Create a struct for each advertisement
        advertStruct := &structpb.Struct{
            Fields: map[string]*structpb.Value{
                "daily": {
                    Kind: &structpb.Value_StructValue{
                        StructValue: &structpb.Struct{
                            Fields: dailyValues,
                        },
                    },
                },
                "weekly": {
                    Kind: &structpb.Value_StructValue{
                        StructValue: &structpb.Struct{
                            Fields: weeklyValues,
                        },
                    },
                },
                "monthly": {
                    Kind: &structpb.Value_StructValue{
                        StructValue: &structpb.Struct{
                            Fields: monthlyValues,
                        },
                    },
                },
                "yearly": {
                    Kind: &structpb.Value_StructValue{
                        StructValue: &structpb.Struct{
                            Fields: yearlyValues,
                        },
                    },
                },
            },
			
			
        }
		advertStats = append(advertStats, advertStruct)
		
    }



	// Respond with the fetched statistics for daily, weekly, and monthly scans
	return &pb.GetCompanyDashboardAnalyticResponse{
		Message: "Metrics retrieved successfully",
		Analytics: advertStats,
	}, nil
}

// mapToStructpbValues converts a map[string]int to a map[string]*structpb.Value
func mapToStructpbValues(input map[string]int) map[string]*structpb.Value {
	result := make(map[string]*structpb.Value)
	for key, value := range input {
		result[key] = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(value)}}
	}
	return result
}

func (s *Server) getLabelScansCountForDay(advertiserId, advertId string) (map[string]int, error) {
	startTime1, endTime1 := s.getIntervalTime(0)
	startTime2 := startTime1.Add(4 * time.Hour)
	startTime3 := startTime1.Add(8 * time.Hour)
	startTime4 := startTime1.Add(12 * time.Hour)

	var results1, results2, results3, results4 []struct {
		Count int64
	}

	err := s.H.DB.Model(&models.LabelScan{}).
		Select("COUNT(*) as count").
		Where("advertiser_id = ? AND advert_id = ? AND created_at BETWEEN ? AND ?", advertiserId, advertId, startTime1, endTime1).
		Find(&results1).
		Error

	if err != nil {
		return nil, err
	}

	err = s.H.DB.Model(&models.LabelScan{}).
		Select("COUNT(*) as count").
		Where("advertiser_id = ? AND advert_id = ? AND created_at BETWEEN ? AND ?", advertiserId, advertId, startTime2, startTime3).
		Find(&results2).
		Error

	if err != nil {
		return nil, err
	}

	err = s.H.DB.Model(&models.LabelScan{}).
		Select("COUNT(*) as count").
		Where("advertiser_id = ? AND advert_id = ? AND created_at BETWEEN ? AND ?", advertiserId, advertId, startTime3, startTime4).
		Find(&results3).
		Error

	if err != nil {
		return nil, err
	}

	err = s.H.DB.Model(&models.LabelScan{}).
		Select("COUNT(*) as count").
		Where("advertiser_id = ? AND advert_id = ? AND created_at BETWEEN ? AND ?", advertiserId, advertId, startTime4, endTime1).
		Find(&results4).
		Error

	if err != nil {
		return nil, err
	}

	// Map the results to a map of intervals and their counts
	intervalCounts := map[string]int{
		"0-4 hours":   int(results1[0].Count),
		"4-8 hours":   int(results2[0].Count),
		"8-12 hours":  int(results3[0].Count),
		"12-24 hours": int(results4[0].Count),
	}

	return intervalCounts, nil
}

// getLabelScansCountWeekly returns the number of scans made on each weekly day
func (s *Server) getLabelScansCountWeekly(advertiserId, advertId string) (map[string]int, error) {
	startTime, endTime := s.getIntervalTime(1)
	var results []struct {
		DayOfWeek int
		Count     int64
	}

	// Assuming "created_at" is the timestamp field in your LabelScan model
	err := s.H.DB.Model(&models.LabelScan{}).
		Select("EXTRACT(DOW FROM created_at) as day_of_week, COUNT(*) as count").
		Where("advertiser_id = ? AND advert_id = ? AND created_at BETWEEN ? AND ?", advertiserId, advertId, startTime, endTime).
		Group("EXTRACT(DOW FROM created_at)").
		Find(&results).
		Error

	if err != nil {
		return nil, err
	}

	// Map the results to a map of weekdays and their counts
	weekdayCounts := make(map[string]int)
	for _, result := range results {
		weekday := time.Weekday(result.DayOfWeek).String()
		weekdayCounts[weekday] = int(result.Count)
	}

	return weekdayCounts, nil
}

// Fetch the number of scans made on each month
func (s *Server) getLabelScansCountMonthly(advertiserId, advertId string) (map[string]int, error) {
	// Assuming "created_at" is the timestamp field in your LabelScan model
	startTime, endTime := s.getIntervalTime(2)
	var results []struct {
		Month int
		Count int64
	}

	err := s.H.DB.Model(&models.LabelScan{}).
		Select("EXTRACT(MONTH FROM created_at) as month, COUNT(*) as count").
		Where("advertiser_id = ? AND advert_id = ? AND created_at BETWEEN ? AND ?", advertiserId, advertId, startTime, endTime).
		Group("EXTRACT(MONTH FROM created_at)").
		Find(&results).
		Error

	if err != nil {
		return nil, err
	}

	// Map the results to a map of months and their counts
	monthCounts := make(map[string]int)
	for _, result := range results {
		// Assuming you want the full month name
		monthName := time.Month(result.Month).String()
		monthCounts[monthName] = int(result.Count)
	}

	return monthCounts, nil
}

// Fetch the number of scans made on each year
func (s *Server) getLabelScansCountYearly(advertiserId, advertId string) (map[string]int, error) {
	// Assuming "created_at" is the timestamp field in your LabelScan model
	startTime, endTime := s.getIntervalTime(3)
	var results []struct {
		Year  int
		Count int64
	}

	err := s.H.DB.Model(&models.LabelScan{}).
		Select("EXTRACT(YEAR FROM created_at) as year, COUNT(*) as count").
		Where("advertiser_id = ? AND advert_id = ? AND created_at BETWEEN ? AND ?", advertiserId, advertId, startTime, endTime).
		Group("EXTRACT(YEAR FROM created_at)").
		Find(&results).
		Error

	if err != nil {
		return nil, err
	}

	// Map the results to a map of years and their counts
	yearCounts := make(map[string]int)
	for _, result := range results {
		yearCounts[fmt.Sprintf("%04d", result.Year)] = int(result.Count)
	}

	return yearCounts, nil
}

// Fetch the start and end time based on the specified time interval
func (s *Server) getIntervalTime(index int64) (time.Time, time.Time) {
	now := time.Now().UTC()
	switch index {
	case 0:
		startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		endTime := startTime.Add(24 * time.Hour)
		return startTime, endTime
	case 1:
		startTime := now.AddDate(0, 0, -int(now.Weekday()))
		startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
		return startTime, startTime.Add(7 * 24 * time.Hour)
	case 2:
		return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC), time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	case 3:
		return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC), time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	default:
		return time.Time{}, time.Time{}
	}
}


func (s *Server) getAdvertisementsForAdvertiser(advertiserID string) ([]models.Advert, error) {
    var adScans []models.Advert

    // Assuming you have a model named Advertisement
    err := s.H.DB.Model(&models.Advert{}).
        Where("advertiser_id = ?", advertiserID).
        Find(&adScans).
        Error

    if err != nil {
        return nil, err
    }

    return adScans, nil
}
