package application

import (
	"fmt"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/config"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/application/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ApplicationServiceClient
}

func InitServiceClient(c *config.Config) pb.ApplicationServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ApplicationSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewApplicationServiceClient(cc)
}
