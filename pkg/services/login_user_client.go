package services

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/types/known/structpb"
)

// login user
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	fmt.Print("it is messsing up here")

	//check if given user exists by scanning user phone number
	if result := s.H.DB.Where(&models.User{EmailAddress: req.EmailAddress}).First(&user); result.Error != nil {

		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User Acount Not Found",
		}, nil
	}

	// check user role vs system given

	if req.Role != user.Role {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "You do not have permission to use this feature or app",
		}, nil

	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Wrong Password, Try again",
		}, nil
	}

	///plan to rename id to user_id
	userJSON, _ := json.Marshal(&models.User{
		Id:                 user.Id,
		PhoneNumber:        user.PhoneNumber,
		AccountNumber:      user.AccountNumber,
		UserName:           user.UserName,
		EmailAddress:       user.EmailAddress,
		Role:               user.Role,
		CompanyLogo:        user.CompanyLogo,
		IsAccountActive:    user.IsAccountActive,
		CompanyName:        user.CompanyName,
		CompanyWebsite:     user.CompanyWebsite,
		CompanyEmail:       user.EmailAddress,
		CompanyPhone:       user.CompanyPhone,
		CompanyAddress:     user.CompanyAddress,
		CompanyDescription: user.CompanyDescription,
	})

	// Create a new Struct message
	var dataStruct structpb.Struct
	// Unmarshal the JSON object into the Struct message
	jsonpb.UnmarshalString(string(userJSON), &dataStruct)
	token, _ := s.Jwt.GenerateToken(user.Id.String(), user.Role)

	return &pb.LoginResponse{
		Status:  http.StatusOK,
		Message: "Logged In Successfuly",
		Token:   token,
		Data:    &dataStruct,
	}, nil
}
