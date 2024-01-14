package routes

import (
	"context"
	"net/http"

	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	UserUid            string `json:"user_uid"`
	EmailAddress       string `json:"email_address"`
	Password           string `json:"password"`
	PhoneNumber        string `json:"phone_number"`
	Name               string `json:"user_name"`	
	Role               string `json:"role"`
	CompanyName        string `json:"company_name"`
	CompanyLogo        string `json:"company_logo"`
	CompanyWebsite     string `json:"company_website"`
	CompanyEmail       string `json:"company_email"`
	CompanyPhone       string `json:"company_phone"`
	CompanyAddress     string `json:"company_address"`
	CompanyDescription string `json:"company_description"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		UserUid:            body.UserUid,
		EmailAddress:       body.EmailAddress,
		Password:           body.Password,
		PhoneNumber:        body.PhoneNumber,
		Role:               body.Role,
		Name:               body.Name,
		CompanyName:        body.CompanyName,
		CompanyLogo:        body.CompanyLogo,
		CompanyWebsite:     body.CompanyWebsite,
		CompanyEmail:       body.CompanyEmail,
		CompanyPhone:       body.CompanyPhone,
		CompanyAddress:     body.CompanyAddress,
		CompanyDescription: body.CompanyDescription,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
