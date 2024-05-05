package client

import (
	"context"
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"google.golang.org/grpc"
)

type PaymentServiceClient struct {
	Client pb.PaymentServiceClient
}

func InitPaymentServiceClient(url string) PaymentServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := PaymentServiceClient{
		Client: pb.NewPaymentServiceClient(cc),
	}

	return c
}

func (c *PaymentServiceClient) DeletePayment(UserId string) (*pb.DeletePaymentResponse, error) {
	req := &pb.DeletePaymentRequest{
		UserId: UserId,
	}

	return c.Client.DeletePayment(context.Background(), req)
}
