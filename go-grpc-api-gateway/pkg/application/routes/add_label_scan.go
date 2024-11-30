package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/application/pb"
	"github.com/gin-gonic/gin"
)

func AddLabelScan(ctx *gin.Context, c pb.ApplicationServiceClient) {
	var body pb.AddLabelScanRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddLabelScan(context.Background(), &body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}