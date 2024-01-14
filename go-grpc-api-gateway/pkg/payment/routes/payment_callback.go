package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"

	"github.com/gin-gonic/gin"
)

func PaymentCallback(ctx *gin.Context, c pb.PaymentServiceClient) {
	req := &pb.PaymentCallbackRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(req)

	fmt.Println("req.Data.Status")

	fmt.Println(req.Data)

	res, err := c.PaymentCallback(context.Background(), &pb.PaymentCallbackRequest{
		Hook: req.Hook,
		Data: req.Data,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
