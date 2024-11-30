package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

type CheckVoucherValidityBody struct {
	UserId      string  `json:"user_id"`
	VoucherCode string  `json:"voucher_code"`
	AmountToPay float64 `json:"amount_to_pay"`
}

func CheckVoucherValidity(ctx *gin.Context, c pb.PaymentServiceClient) {
	body := CheckVoucherValidityBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CheckVoucherValidity(context.Background(), &pb.CheckVoucherValidityRequest{
		UserId:      body.UserId,
		VoucherCode: body.VoucherCode,
		AmountToPay: body.AmountToPay,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
