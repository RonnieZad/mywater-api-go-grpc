package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	Role         string `json:"role"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		EmailAddress: b.EmailAddress,
		Password:     b.Password,
		PhoneNumber:  b.PhoneNumber,
		Role:         b.Role,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func LoginUserClient(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.LoginUserClient(context.Background(), &pb.LoginRequest{
		EmailAddress: b.EmailAddress,
		Password:     b.Password,
		PhoneNumber:  b.PhoneNumber,
		Role:         b.Role,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
