package services

import (
	"context"

	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/pb"
	// "github.com/RonnieZad/nyumba-go-grpc-auth-svc/pkg/utils"
	// "google.golang.org/grpc/metadata"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Find user by ID
	var user models.User

	if result := s.H.DB.Where("id", req.UserId).First(&user); result.Error != nil {
		return &pb.GetUserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	// Marshal user object to JSON
	data := &pb.UserDetail{
		Id:                 user.Id.String(),
		PhoneNumber:        user.PhoneNumber,
		Name:               user.UserName,
		EmailAddress:       user.EmailAddress,
		Role:               user.Role,
		CompanyName:        user.CompanyName,
		CompanyLogo:        user.CompanyLogo,
		CompanyWebsite:     user.CompanyWebsite,
		CompanyEmail:       user.EmailAddress,
		CompanyPhone:       user.CompanyPhone,
		CompanyAddress:     user.CompanyAddress,
		CompanyDescription: user.CompanyDescription,
	}

	token, _ := s.Jwt.GenerateToken(user.Id.String(), user.Role)

	return &pb.GetUserResponse{
		Status:  http.StatusOK,
		Message: "User Found",
		Token:   token,
		Data:    data,
	}, nil
}

func (s *Server) GetUserClient(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Find user by ID

	var userClient models.UserClient

	if result := s.H.DB.Where("id", req.UserId).First(&userClient); result.Error != nil {
		return &pb.GetUserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	// Marshal user object to JSON
	data := &pb.UserDetail{
		Id:          userClient.Id.String(),
		PhoneNumber: userClient.PhoneNumber,
		Name:        userClient.UserName,
		Role:        userClient.Role,
	}

	token, _ := s.Jwt.GenerateToken(userClient.Id.String(), userClient.Role)

	return &pb.GetUserResponse{
		Status:  http.StatusOK,
		Message: "User Found",
		Token:   token,
		Data:    data,
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {

	// Find all users
	var users []models.User
	if result := s.H.DB.Find(&users); result.Error != nil {
		return &pb.GetAllUsersResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to retrieve user",
		}, result.Error
	}

	var userList []*pb.UserDetail
	for _, user := range users {

		pbUser := &pb.UserDetail{
			Id:                 user.Id.String(),
			PhoneNumber:        user.PhoneNumber,
			EmailAddress:       user.EmailAddress,
			Role:               user.Role,
			CompanyName:        user.CompanyName,
			CompanyLogo:        user.CompanyLogo,
			CompanyWebsite:     user.CompanyWebsite,
			CompanyEmail:       user.EmailAddress,
			CompanyPhone:       user.CompanyPhone,
			CompanyAddress:     user.CompanyAddress,
			CompanyDescription: user.CompanyDescription,
		}

		userList = append(userList, pbUser)

	}

	return &pb.GetAllUsersResponse{
		Status:  http.StatusOK,
		Message: "Users Found",
		User:    userList,
	}, nil
}
