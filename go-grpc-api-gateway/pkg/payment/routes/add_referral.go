package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

type AppReferralRequestBody struct {
	UserId        string  `json:"user_id"`
	InvitedUserId string  `json:"invited_user_id"`
	DeviceId      string  `json:"device_id"`
	Earning       float64 `json:"earning"`
	Currency      string  `json:"currency"`
	Status        string  `json:"status"`
}

func AppReferral(ctx *gin.Context, c pb.PaymentServiceClient) {
	body := AppReferralRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AppReferral(context.Background(), &pb.AppReferralRequest{
		UserId:        body.UserId,
		InvitedUserId: body.InvitedUserId,
		DeviceId:      body.DeviceId,
		Earning:       body.Earning,
		Currency:      body.Currency,
		Status:        body.Status,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
