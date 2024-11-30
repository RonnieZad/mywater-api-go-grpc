package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

func GetDashboardMetrics(ctx *gin.Context, c pb.PaymentServiceClient) {

	res, err := c.GetDashboardMetrics(context.Background(), &pb.GetDashboardMetricsRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
