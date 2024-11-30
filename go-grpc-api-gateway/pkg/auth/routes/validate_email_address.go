package routes

import (
	"context"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ValidateEmailBody struct {
	EmailAddress string `json:"email_address"`
	IpAddress    string `json:"ip_address"`
}

func ValidateEmail(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ValidateEmailBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ValidateEmail(context.Background(), &pb.ValidateEmailRequest{
		EmailAddress: b.EmailAddress,
		IpAddress:    b.IpAddress,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
