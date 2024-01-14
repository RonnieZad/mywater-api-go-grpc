package main

import (
	"fmt"
	"log"

	"net"

	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/client"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/config"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	service "github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	propertySvc := client.InitPropertyServiceClient(c.PropertySvcUrl)
	applicationSvc := client.InitApplicationServiceClient(c.ApplicationSvcUrl)
	authSvc := client.InitAuthServiceClient(c.AuthSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Payment Service 💰 on Port", c.Port)

	s := service.Server{
		H:              h,
		PropertySvc:    propertySvc,
		ApplicationSvc: applicationSvc,
		AuthSvc:        authSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(grpcServer, &s)

	// Start the periodic task in the background
	go service.RunPropertyStatusUpdateTask(&s)

	// Start streaming new payment notifications to clients
	// go s.StartPaymentNotificationStreaming()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
