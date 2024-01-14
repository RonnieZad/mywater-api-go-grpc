package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context, c pb.AuthServiceClient) {

	res, err := c.GetAllUsers(context.Background(), &pb.GetAllUsersRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
	fmt.Print(&res)
}
