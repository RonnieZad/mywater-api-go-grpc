package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

func GetLandlordRevenue(ctx *gin.Context, c pb.PaymentServiceClient) {
	id := ctx.Param("id")

	res, err := c.GetLandlordRevenue(context.Background(), &pb.GetLandlordRevenueRequest{
		LandlordId: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
