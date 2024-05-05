package client

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
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

func (c *PropertyServiceClient) DeletePropertyHeld(UserId string) (*pb.DeletePropertyHeldResponse, error) {
	req := &pb.DeletePropertyHeldRequest{
		UserId: UserId,
	}

	return c.Client.DeletePropertyHeld(context.Background(), req)
}

func (c *PropertyServiceClient) DeleteUserPropertyPreference(UserId string) (*pb.DeleteUserPropertyPreferenceResponse, error) {
	req := &pb.DeleteUserPropertyPreferenceRequest{
		UserId: UserId,
	}

	return c.Client.DeleteUserPropertyPreference(context.Background(), req)
}
