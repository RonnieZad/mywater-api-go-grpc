package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
)

type UpdateReferreralRequestBody struct {
	ReferralRecordId string `json:"referral_record_id"`
	UserId           string `json:"user_id"`
	Status           string `json:"status"`
	IsPaymentMade    bool   `json:"is_payment_made"`
}

func UpdateReferreral(ctx *gin.Context, c pb.PaymentServiceClient) {
	body := UpdateReferreralRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateReferreral(context.Background(), &pb.UpdateReferreralRequest{
		ReferralRecordId: body.ReferralRecordId,
		UserId:           body.UserId,
		Status:           body.Status,
		IsPaymentMade:    body.IsPaymentMade,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
