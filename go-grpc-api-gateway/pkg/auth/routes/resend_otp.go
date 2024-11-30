package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type ResendOTPBody struct {
	PhoneNumber string `json:"phone_number"`
}

func ResendOTP(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ResendOTPBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ResendOTP(context.Background(), &pb.ResendOTPRequest{
		PhoneNumber: b.PhoneNumber,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
