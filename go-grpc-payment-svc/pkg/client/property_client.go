package client

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"google.golang.org/grpc"
)

type PropertyServiceClient struct {
	Client pb.PropertyServiceClient
}

func InitPropertyServiceClient(url string) PropertyServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := PropertyServiceClient{
		Client: pb.NewPropertyServiceClient(cc),
	}

	return c
}

func (c *PropertyServiceClient) FindOne(PropertyId string) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: PropertyId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *PropertyServiceClient) HoldProperty(UserId string, PropertyId string) (*pb.HoldPropertyResponse, error) {
	req := &pb.HoldPropertyRequest{
		UserId:     UserId,
		PropertyId: PropertyId,
	}

	return c.Client.HoldProperty(context.Background(), req)
}

func (c *PropertyServiceClient) UpdateHoldPropertyStatus(PropertyId string, Status bool) (*pb.UpdateHoldPropertyStatusResponse, error) {
	req := &pb.UpdateHoldPropertyStatusRequest{
		PropertyId: PropertyId,
		Status:     Status,
	}

	return c.Client.UpdateHoldPropertyStatus(context.Background(), req)
}

func (c *PropertyServiceClient) UpdatePropertyAvailabiltyStatus(PropertyId string, UserId string, Status bool) (*pb.UpdatePropertyAvailabiltyStatusResponse, error) {
	req := &pb.UpdatePropertyAvailabiltyStatusRequest{
		PropertyId: PropertyId,
		UserId:     UserId,
		Status:     Status,
	}

	return c.Client.UpdatePropertyAvailabiltyStatus(context.Background(), req)
}

func (c *PropertyServiceClient) ResetHoldPropertyStatus(PropertyId string) (*pb.ResetHoldPropertyStatusResponse, error) {
	req := &pb.ResetHoldPropertyStatusRequest{
		PropertyId: PropertyId,
	}

	return c.Client.ResetHoldPropertyStatus(context.Background(), req)
}

func (c *PropertyServiceClient) GetPropertyMetrics() (*pb.GetPropertyMetricsResponse, error) {
	req := &pb.GetPropertyMetricsRequest{}

	return c.Client.GetPropertyMetrics(context.Background(), req)
}
