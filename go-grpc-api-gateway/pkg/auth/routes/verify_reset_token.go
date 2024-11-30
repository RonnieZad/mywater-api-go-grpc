package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type VerifyResetTokenBody struct {
	PasswordResetToken string `json:"password_reset_token"`
}

func VerifyResetToken(ctx *gin.Context, c pb.AuthServiceClient) {
	b := VerifyResetTokenBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.VerifyResetToken(context.Background(), &pb.VerifyResetTokenRequest{
		PasswordResetToken: b.PasswordResetToken,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
