package routes

import (
	"context"
	"fmt"
	// "fmt"

	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	// "github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context, c pb.AuthServiceClient) {
	userId := ctx.Param("id")

	fmt.Print("userId: ", userId)
	// b := DeleteUserBody{}

	// if err := ctx.BindJSON(&b); err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	res, err := c.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		UserId: userId,
		// UserName:    b.UserName,
		// PhoneNumber: b.PhoneNumber,
	})

	// fmt.Print(b)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
