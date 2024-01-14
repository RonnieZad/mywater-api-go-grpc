package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type GetUserDeviceUUIDBody struct {
	UserId     string `json:"user_id"`
	DeviceUuid string `json:"device_uuid"`
}

func GetUserDeviceUUID(ctx *gin.Context, c pb.AuthServiceClient) {
	body := GetUserDeviceUUIDBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.GetUserDeviceUUID(context.Background(), &pb.GetUserDeviceUUIDRequest{
		UserId:     body.UserId,
		DeviceUuid: body.DeviceUuid,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
