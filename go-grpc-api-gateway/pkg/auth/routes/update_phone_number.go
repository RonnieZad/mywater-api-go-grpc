package routes

import (
	"context"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdatePhoneNumberBody struct {
	UserId         string `json:"user_id"`
	NewPhoneNumber string `json:"new_phone_number"`
}

func UpdatePhoneNumber(ctx *gin.Context, c pb.AuthServiceClient) {
	b := UpdatePhoneNumberBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdatePhoneNumber(context.Background(), &pb.UpdatePhoneNumberRequest{
		UserId:         b.UserId,
		NewPhoneNumber: b.NewPhoneNumber,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
