package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

func GetUserMetrics(ctx *gin.Context, c pb.AuthServiceClient) {

	res, err := c.GetUserMetrics(context.Background(), &pb.GetUserMetricsRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}