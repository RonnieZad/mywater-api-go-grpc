package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type UpdateDocumentRequestBody struct {
	UserId string `json:"user_id"`
	NationalIdUrl string `json:"national_id_url"`
	PaySlipIdUrl string `json:"pay_slip_id_url"`
	EmployeeIdUrl string `json:"employee_id_url"`

}

func UpdateDocument(ctx *gin.Context, c pb.AuthServiceClient) {
	b := UpdateDocumentRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateDocument(context.Background(), &pb.UpdateDocumentRequest{
		UserId: b.UserId,
		NationalIdUrl: b.NationalIdUrl,
		PaySlipIdUrl: b.PaySlipIdUrl,
		EmployeeIdUrl: b.EmployeeIdUrl,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
