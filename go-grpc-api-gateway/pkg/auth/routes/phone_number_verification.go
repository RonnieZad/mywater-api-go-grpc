package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type PhoneNumberVerificationRequestBody struct {
	PhoneNumber string `json:"phone_number"`
}

func PhoneNumberVerification(ctx *gin.Context, c pb.AuthServiceClient) {
	body := PhoneNumberVerificationRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.PhoneNumberVerification(context.Background(), &pb.PhoneNumberVerificationRequest{
		PhoneNumber: body.PhoneNumber,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
