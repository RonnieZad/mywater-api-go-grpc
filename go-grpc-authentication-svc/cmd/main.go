package main

import (
	"fmt"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/client"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/config"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/services"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	paymentSvc := client.InitPaymentServiceClient(c.PaymentSvcUrl)
	propertySvc := client.InitPropertyServiceClient(c.PropertySvcUrl)
	applicationSvc := client.InitApplicationServiceClient(c.ApplicationSvcUrl)

	fmt.Println("Authentication Service 🔐 on Port", c.Port)

	s := services.Server{
		H:              h,
		Jwt:            jwt,
		PaymentSvc:     paymentSvc,
		PropertySvc:    propertySvc,
		ApplicationSvc: applicationSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
