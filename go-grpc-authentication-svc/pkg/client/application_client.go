package client

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
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

func (c *ApplicationServiceClient) DeleteRentApplication(UserId string) (*pb.DeleteRentApplicationResponse, error) {
	req := &pb.DeleteRentApplicationRequest{
		UserId: UserId,
	}

	return c.Client.DeleteRentApplication(context.Background(), req)
}

func (c *ApplicationServiceClient) DeleteTourApplication(UserId string) (*pb.DeleteTourApplicationResponse, error) {
	req := &pb.DeleteTourApplicationRequest{
		UserId: UserId,
	}

	return c.Client.DeleteTourApplication(context.Background(), req)
}
