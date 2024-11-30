package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	// "github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type DeactivateUserBody struct {
	UserId string `json:"user_id"`
}

func DeactivateUser(ctx *gin.Context, c pb.AuthServiceClient) {
	b := DeactivateUserBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.DeactivateUser(context.Background(), &pb.DeactivateUserRequest{
		UserId: b.UserId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
