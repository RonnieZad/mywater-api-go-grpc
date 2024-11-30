package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type InvitedLandlordBody struct {
	UserId          string `json:"user_id"`
	InvitedLandlord bool   `json:"invited_landlord"`
}

func InvitedLandlord(ctx *gin.Context, c pb.AuthServiceClient) {
	b := InvitedLandlordBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.InvitedLandlord(context.Background(), &pb.InvitedLandlordRequest{
		UserId:          b.UserId,
		InvitedLandlord: b.InvitedLandlord,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
