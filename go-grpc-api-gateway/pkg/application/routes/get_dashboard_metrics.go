package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/application/pb"
	"github.com/gin-gonic/gin"
)

func GetCompanyDashboardAnalytic(ctx *gin.Context, c pb.ApplicationServiceClient) {
	id := ctx.Param("id")

	res, err := c.GetCompanyDashboardAnalytic(context.Background(), &pb.GetCompanyDashboardAnalyticRequest{
		AdvertiserId: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
