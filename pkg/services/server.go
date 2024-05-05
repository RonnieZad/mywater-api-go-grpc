package services

import (
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/client"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/db"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
)

type Server struct {
	H              db.Handler
	PropertySvc    client.PropertyServiceClient
	ApplicationSvc client.ApplicationServiceClient
	PaymentSvc     client.PaymentServiceClient
	Jwt            utils.JwtWrapper
	pb.AuthServiceServer
}

// register function
func getRolePermissions(role string) string {
	switch role {
	case "admin":
		return "create user, view update"
	case "staff":
		return "view, update"
	case "agent":
		return "view, update"
	case "landlord":
		return "create user, view, update"
	case "client":
		return "create user, view, update"
	default:
		return ""
	}
}
