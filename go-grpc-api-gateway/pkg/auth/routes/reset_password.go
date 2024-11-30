package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type ResetPassowrdRequestBody struct {
	Identifier string `json:"identifier"`
}

func ResetPassword(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ResetPassowrdRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ResetPassword(context.Background(), &pb.ResetPasswordRequest{
		Identifier: b.Identifier,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
