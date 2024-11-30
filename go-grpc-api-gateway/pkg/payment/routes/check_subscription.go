package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

type CheckSubscriptionStatusBody struct {
	UserId           string `json:"user_id"`
	PropertyId       string `json:"property_id"`
	SubscriptionType string `json:"subscription_type"`
}

func CheckSubscription(ctx *gin.Context, c pb.PaymentServiceClient) {
	body := CheckSubscriptionStatusBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CheckSubscription(context.Background(), &pb.CheckSubscriptionRequest{
		UserId:           body.UserId,
		PropertyId:       body.PropertyId,
		SubscriptionType: body.SubscriptionType,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
