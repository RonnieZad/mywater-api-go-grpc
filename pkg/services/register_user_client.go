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

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Check if user already exists
	var userDetails models.User

	if result := s.H.DB.Where(&models.User{EmailAddress: req.EmailAddress}).First(&userDetails); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "User Already Exists",
		}, nil
	}

	picture := ""

	if req.Role == "client" {
		picture = req.ProfilePic
	} else {
		picture = req.CompanyLogo
	}

	accountNumber := utils.GenerateAccountNumber()
	user := &models.User{
		Id:                 uuid.New(),
		AccountNumber:      accountNumber,
		UserName:           req.Name,
		EmailAddress:       req.EmailAddress,
		PhoneNumber:        req.PhoneNumber,
		CompanyName:        req.CompanyName,
		CompanyWebsite:     req.CompanyWebsite,
		CompanyAddress:     req.CompanyAddress,
		CompanyPhone:       req.CompanyPhone,
		CompanyDescription: req.CompanyDescription,
		CompanyEmail:       req.CompanyEmail,
		CompanyLogo:        picture,
		Role:               req.Role,
		Password:           utils.HashPassword(req.Password),
	}

	// Check user role and set permissions accordingly
	switch req.Role {
	case "admin":
		user.Role = "admin"
		user.Permissions = getRolePermissions(user.Role)

	case "staff":
		user.Role = "staff"
		user.Permissions = getRolePermissions(user.Role)

	case "agent":
		user.Role = "agent"
		user.Permissions = getRolePermissions(user.Role)

	case "landlord":
		user.Role = "landlord"
		user.Permissions = getRolePermissions(user.Role)

	case "client":
		user.Role = "client"
		user.Permissions = getRolePermissions(user.Role)
	}

	// Create user in the database
	s.H.DB.Create(&user)

	// Generate JWT token for user
	token, _ := s.Jwt.GenerateToken(user.Id.String(), user.Role)

	// Marshal user object to JSON
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

	return &pb.RegisterResponse{
		Status:  http.StatusCreated,
		Message: "Account Has Been Created",
		Token:   token,
		Data:    &dataStruct,
	}, nil
}

func (s *Server) RegisterUserClient(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Check if user already exists
	var userDetails models.UserClient

	if result := s.H.DB.Where(&models.UserClient{PhoneNumber: req.PhoneNumber}).First(&userDetails); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "User Already Exists",
		}, nil
	}

	accountNumber := utils.GenerateAccountNumber()
	user := &models.UserClient{
		Id:            uuid.New(),
		PhoneNumber:   req.PhoneNumber,
		AccountNumber: accountNumber,
		Role:          req.Role,
		Password:      utils.HashPassword(req.Password),
	}

	// Create user in the database
	s.H.DB.Create(&user)

	// Generate JWT token for user
	token, _ := s.Jwt.GenerateToken(user.Id.String(), user.Role)

	// Marshal user object to JSON
	userJSON, _ := json.Marshal(&models.UserClient{
		Id:              user.Id,
		PhoneNumber:     user.PhoneNumber,
		AccountNumber:   user.AccountNumber,
		Role:            req.Role,
		IsAccountActive: user.IsAccountActive,
	})

	// Create a new Struct message
	var dataStruct structpb.Struct
	// Unmarshal the JSON object into the Struct message
	jsonpb.UnmarshalString(string(userJSON), &dataStruct)

	return &pb.RegisterResponse{
		Status:  http.StatusCreated,
		Message: "Account Has Been Created",
		Token:   token,
		Data:    &dataStruct,
	}, nil
}

func (s *Server) LoginUserClient(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.UserClient

	if result := s.H.DB.Where(&models.UserClient{PhoneNumber: req.PhoneNumber}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User Not Found",
		}, nil
	}

	fmt.Print("req.PhoneNumber")
	fmt.Print(req.PhoneNumber)

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

	userJSON, _ := json.Marshal(&models.UserClient{
		Id:              user.Id,
		PhoneNumber:     user.PhoneNumber,
		AccountNumber:   user.AccountNumber,
		UserName:        user.UserName,
		Role:            req.Role,
		IsAccountActive: user.IsAccountActive,
	})

	var dataStruct structpb.Struct
	jsonpb.UnmarshalString(string(userJSON), &dataStruct)
	token, _ := s.Jwt.GenerateToken(user.Id.String(), user.Role)

	return &pb.LoginResponse{
		Status:  http.StatusOK,
		Message: "Login Successful",
		Token:   token,
		Data:    &dataStruct,
	}, nil
}
