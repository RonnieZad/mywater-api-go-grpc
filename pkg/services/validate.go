package services

import (
	"context"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"

	"github.com/zerobounce/zerobouncego"
)

// phone number verification: sends otp to user
func (s *Server) ValidateEmail(ctx context.Context, req *pb.ValidateEmailRequest) (*pb.ValidateEmailResponse, error) {

	zerobouncego.SetApiKey("c27e593d32c74175b7d042e54539b13c")

	// For Querying a single E-Mail and IP
	// IP can also be an empty string
	response, error_ := zerobouncego.Validate(req.EmailAddress, req.IpAddress)

	if error_ != nil {
		return &pb.ValidateEmailResponse{
			IsValid: false,
			Message: "something went wrong",
			Error:   error_.Error(),
		}, nil
	}

	// Now you can check status
	if response.Status == zerobouncego.S_VALID {

		return &pb.ValidateEmailResponse{
			IsValid: true,
			Message: "Valid Email Address",
		}, nil
	}

	// Now you can check status
	if response.Status == zerobouncego.S_INVALID {
		return &pb.ValidateEmailResponse{
			IsValid: false,
			Message: "Invalid Email Address",
		}, nil
	}

	// Invalid email address
	return &pb.ValidateEmailResponse{
		IsValid: false,
		Message: "The email address is invalid",
	}, nil
}
