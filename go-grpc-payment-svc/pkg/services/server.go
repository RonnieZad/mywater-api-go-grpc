package service

import (
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/client"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
)

type Server struct {
	H              db.Handler
	PropertySvc    client.PropertyServiceClient
	ApplicationSvc client.ApplicationServiceClient
	AuthSvc        client.AuthServiceClient
	pb.PaymentServiceServer
}

func (s *Server) mustEmbedUnimplementedMakeCollectionServiceServer() {
	// Do nothing - this is just a helper method
}
