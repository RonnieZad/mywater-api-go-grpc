package payment

import (
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/config"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.PaymentServiceClient
}

func InitServiceClient(c *config.Config) pb.PaymentServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.PaymentSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewPaymentServiceClient(cc)
}
