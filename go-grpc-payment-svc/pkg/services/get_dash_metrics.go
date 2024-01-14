package service

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"google.golang.org/protobuf/types/known/structpb"
)

// Get number of available property types for each category
func (s *Server) GetDashboardMetrics(ctx context.Context, req *pb.GetDashboardMetricsRequest) (*pb.GetDashboardMetricsResponse, error) {

	//get transaction volume
	var totalAmount float64

	s.H.DB.Table("payment_transactions").Select("sum(amount) as total").Scan(&totalAmount)

	//get property capacity
	property, err := s.PropertySvc.GetPropertyMetrics()

	if err != nil {

		fmt.Println("Error getting property metrics")
	}

	//get appliocation capacity
	application, err := s.ApplicationSvc.GetApplicationMetrics()

	if err != nil {

		fmt.Println("Error getting property metrics")
	}

	user, err := s.AuthSvc.GetUserMetrics()

	if err != nil {

		fmt.Println("Error getting property metrics")
	}

	fmt.Println(user.TotalUsers)

	// TotalRentFinancingApplications: int32(len(houseFinanceApplications)),
	// TotalRoommateApplications:      int32(len(roomMateApplications)),

	return &pb.GetDashboardMetricsResponse{
		Message: "Metrics retrieved successfully",
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"property_qty": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(property.TotalProperties),
					},
				},
				"roommate_host_qty": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(application.TotalRoommateApplications),
					},
				},
				"rent_financing_qty": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(application.TotalRentFinancingApplications),
					},
				},
				"user_qty": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: float64(user.TotalUsers),
					},
				},
				"transaction_qty": {
					Kind: &structpb.Value_NumberValue{
						NumberValue: totalAmount,
					},
				},
			},
		},
	}, nil
}
