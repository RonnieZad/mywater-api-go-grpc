package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordBody struct {
	PasswordResetToken string `json:"password_reset_token"`
	NewPassword        string `json:"new_password"`
}

func UpdatePassword(ctx *gin.Context, c pb.AuthServiceClient) {
	b := UpdatePasswordBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdatePassword(context.Background(), &pb.UpdatePasswordRequest{
		PasswordResetToken: b.PasswordResetToken,
		NewPassword:        b.NewPassword,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
