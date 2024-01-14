package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

func SendPayment(ctx *gin.Context, c pb.PaymentServiceClient) {
	var body pb.SendPaymentRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.SendPayment(context.Background(), &body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}