package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type PhoneNumberVerificationWithOTPRequestBody struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

func PhoneNumberVerificationWithOTP(ctx *gin.Context, c pb.AuthServiceClient) {
	body := PhoneNumberVerificationWithOTPRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.PhoneNumberVerificationWithOTP(context.Background(), &pb.PhoneNumberVerificationWithOTPRequest{
		PhoneNumber: body.PhoneNumber,
		Otp:         body.OTP,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
