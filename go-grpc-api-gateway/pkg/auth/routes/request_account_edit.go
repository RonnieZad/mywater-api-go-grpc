package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type EditAccountBody struct {
	UserId      string `json:"user_id"`
	Name        bool   `json:"name"`
	Nin         bool   `json:"nin"`
	DateOfBirth bool   `json:"date_of_birth"`
	PhoneNumber bool   `json:"phone_number"`
}

func EditAccount(ctx *gin.Context, c pb.AuthServiceClient) {
	body := EditAccountBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.EditAccount(context.Background(), &pb.EditAccountRequest{
		UserId:                       body.UserId,
		Name:                         body.Name,
		NationalIdentificationNumber: body.Nin,
		DateOfBirth:                  body.DateOfBirth,
		PhoneNumber:                  body.PhoneNumber,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
