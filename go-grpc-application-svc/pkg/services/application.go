package services

import (
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/client"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-project/application-svc/pkg/pb"
)

type Server struct {
	H           db.Handler
	PropertySvc client.PropertyServiceClient
	AuthSvc     client.AuthServiceClient
	pb.ApplicationServiceServer
}

func (s *Server) mustEmbedUnimplementedApplicationServiceServer() {}
