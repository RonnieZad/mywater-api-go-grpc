package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

type CheckPhoneSubscriptionStatusBody struct {
	UserId           string `json:"user_id"`
	SubscriptionType string `json:"subscription_type"`
}

func CheckPhoneSubscription(ctx *gin.Context, c pb.PaymentServiceClient) {
	body := CheckPhoneSubscriptionStatusBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CheckPhoneSubscription(context.Background(), &pb.CheckSubscriptionRequest{
		UserId:           body.UserId,

		SubscriptionType: body.SubscriptionType,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
