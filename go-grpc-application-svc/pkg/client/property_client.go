package client

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
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

func (c *PropertyServiceClient) GetLocksByGroupId(GroupId int32) (*pb.GetLocksByGroupIdResponse, error) {
	req := &pb.GetLocksByGroupIdRequest{
		GroupId: GroupId,
	}

	return c.Client.GetLocksByGroupId(context.Background(), req)
}

func (c *PropertyServiceClient) DeleteProperty(PropertyId string) (*pb.DeletePropertyResponse, error) {
	req := &pb.DeletePropertyRequest{
		PropertyId: PropertyId,
	}

	return c.Client.DeleteProperty(context.Background(), req)
}
