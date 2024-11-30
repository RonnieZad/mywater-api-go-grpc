package client

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
	"google.golang.org/grpc"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func InitAuthServiceClient(url string) AuthServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := AuthServiceClient{
		Client: pb.NewAuthServiceClient(cc),
	}

	return c
}

func (c *AuthServiceClient) GetUser(UserId string) (*pb.GetUserResponse, error) {
	req := &pb.GetUserRequest{
		UserId: UserId,
	}

	return c.Client.GetUser(context.Background(), req)
}

func (c *AuthServiceClient) GetUserClient(UserId string) (*pb.GetUserResponse, error) {
	req := &pb.GetUserRequest{
		UserId: UserId,
	}

	return c.Client.GetUser(context.Background(), req)
}

