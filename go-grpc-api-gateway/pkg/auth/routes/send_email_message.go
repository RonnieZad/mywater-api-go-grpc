package routes

import (
	"context"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendEmailMessageBody struct {
	UserId       string `json:"user_id"`
	Subject      string `json:"subject"`
	Message      string `json:"message"`
	EmailAddress string `json:"email_address"`
}

func SendEmailMessage(ctx *gin.Context, c pb.AuthServiceClient) {
	body := SendEmailMessageBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.SendEmailMessage(context.Background(), &pb.SendEmailMessageRequest{
		UserId:       body.UserId,
		Subject:      body.Subject,
		Message:      body.Message,
		EmailAddress: body.EmailAddress,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
