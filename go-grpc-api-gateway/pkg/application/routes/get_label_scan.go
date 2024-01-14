package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/application/pb"
	"github.com/gin-gonic/gin"
)


func GetUserLabelScan(ctx *gin.Context, c pb.ApplicationServiceClient) {
	id := ctx.Param("id")

	res, err := c.GetUserLabelScan(context.Background(), &pb.GetUserLabelScanRequest{
		UserId: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}