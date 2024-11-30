package routes

import (
	"context"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendTextMessageBody struct {
	UserId      string `json:"user_id"`
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
}

func SendTextMessage(ctx *gin.Context, c pb.AuthServiceClient) {
	body := SendTextMessageBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.SendTextMessage(context.Background(), &pb.SendTextMessageRequest{
		UserId:      body.UserId,
		Message:     body.Message,
		PhoneNumber: body.PhoneNumber,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
