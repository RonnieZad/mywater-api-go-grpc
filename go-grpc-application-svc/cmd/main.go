package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/client"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/config"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
	service "github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/services"
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
	authSvc := client.InitAuthServiceClient(c.AuthSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Application Service 📃 on Port", c.Port)

	s := service.Server{
		H:           h,
		PropertySvc: propertySvc,
		AuthSvc:     authSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterApplicationServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
