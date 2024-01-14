package routes

import (
	"context"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendPushMessageBody struct {
	UserId     string `json:"user_id"`
	Subject    string `json:"subject"`
	Message    string `json:"message"`
	DeviceUuid string `json:"device_uuid"`
}

func SendPushMessage(ctx *gin.Context, c pb.AuthServiceClient) {
	body := SendPushMessageBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.SendPushMessage(context.Background(), &pb.SendPushMessageRequest{
		UserId:     body.UserId,
		Subject:    body.Subject,
		Message:    body.Message,
		DeviceUuid: body.DeviceUuid,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
