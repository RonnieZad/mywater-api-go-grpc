package client

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ApplicationServiceClient struct {
	Client pb.ApplicationServiceClient
}

func InitApplicationServiceClient(url string) ApplicationServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ApplicationServiceClient{
		Client: pb.NewApplicationServiceClient(cc),
	}

	return c
}

func (c *ApplicationServiceClient) GetUserRentApplicationRequest(UserId string) (*pb.GetUserRentApplicationResponse, error) {
	req := &pb.GetUserRentApplicationRequest{
		UserId: UserId,
	}

	return c.Client.GetUserRentApplication(context.Background(), req)
}

func (c *ApplicationServiceClient) UpdateApplicationLandlordPaymentStatus(ApplicationId string, Status bool) (*pb.UpdateApplicationLandlordPaymentStatusResponse, error) {
	req := &pb.UpdateApplicationLandlordPaymentStatusRequest{
		ApplicationId: ApplicationId,
		Status:        Status,
	}

	return c.Client.UpdateApplicationLandlordPaymentStatus(context.Background(), req)
}

func (c *ApplicationServiceClient) UpdateApplicationSecurityDepositPaidStatus(ApplicationId string, Status bool) (*pb.UpdateApplicationSecurityDepositPaidStatusResponse, error) {
	req := &pb.UpdateApplicationSecurityDepositPaidStatusRequest{
		ApplicationId: ApplicationId,
		Status:        Status,
	}

	return c.Client.UpdateApplicationSecurityDepositPaidStatus(context.Background(), req)
}

func (c *ApplicationServiceClient) UpdateApplicationRepaymentDetails(ApplicationId string, InstallmentCount int32, InstallmentAmount float64, Status bool) (*pb.UpdateApplicationRepaymentDetailsResponse, error) {
	req := &pb.UpdateApplicationRepaymentDetailsRequest{
		ApplicationId:     ApplicationId,
		InstallmentCount:  InstallmentCount,
		InstallmentAmount: InstallmentAmount,
		Status:            Status,
	}

	return c.Client.UpdateApplicationRepaymentDetails(context.Background(), req)
}

func (c *ApplicationServiceClient) MakeTourRequestApplication(PropertyId string, UserId string, ReservationType int64, AppointmentDate string, PaymentStatus bool) (*pb.MakeTourRequestApplicationResponse, error) {
	req := &pb.MakeTourRequestApplicationRequest{
		PropertyId:      PropertyId,
		UserId:          UserId,
		ReservationType: ReservationType,
		AppointmentDate: AppointmentDate,
		PaymentStatus:   PaymentStatus,
	}

	return c.Client.MakeTourRequestApplication(context.Background(), req)
}

func (c *ApplicationServiceClient) RoomMateApplicationUpdate(UserId string, Status bool) (*pb.RoomMateApplicationUpdateResponse, error) {
	req := &pb.RoomMateApplicationUpdateRequest{
		UserId: UserId,
		Status: Status,
	}

	return c.Client.RoomMateApplicationUpdate(context.Background(), req)
}

func (c *ApplicationServiceClient) ApproveRentApplication(UserId string, ApplicationId string, ApplicationStatus int64) (*pb.ApproveRentApplicationResponse, error) {
	req := &pb.ApproveRentApplicationRequest{
		UserId:            UserId,
		ApplicationId:     ApplicationId,
		ApplicationStatus: ApplicationStatus,
	}

	return c.Client.ApproveRentApplication(context.Background(), req)
}

func (c *ApplicationServiceClient) ManageUserLock(ApplicationId string) (*pb.ManageUserLockResponse, error) {
	req := &pb.ManageUserLockRequest{
		ApplicationId: ApplicationId,
	}

	return c.Client.ManageUserLock(context.Background(), req)
}

func (c *ApplicationServiceClient) GetApplicationMetrics() (*pb.GetApplicationMetricsResponse, error) {
	req := &pb.GetApplicationMetricsRequest{}

	return c.Client.GetApplicationMetrics(context.Background(), req)
}
