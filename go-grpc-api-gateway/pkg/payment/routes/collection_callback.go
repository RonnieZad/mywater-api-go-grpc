package routes

import (
	"context"
	"fmt"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CollectionCallbackBody struct {
	Hook struct {
		ID      int    `json:"id,omitempty"`
		Created string `json:"created,omitempty"`
		Updated string `json:"updated,omitempty"`
		Event   string `json:"event,omitempty"`
		Target  string `json:"target,omitempty"`
		User    int    `json:"user,omitempty"`
	} `json:"hook"`
	Data struct {
		ID           int    `json:"id,omitempty"`
		Organization int    `json:"organization,omitempty"`
		Amount       string `json:"amount,omitempty"`
		Currency     string `json:"currency,omitempty"`
		PaymentType  string `json:"payment_type,omitempty"`
		Metadata     struct {
			UserId           string `json:"user_id"`
			PropertyId       string `json:"property_id"`
			PaymentReason    string `json:"payment_reason"`
			InstallmentCount string `json:"installment_count"`
			ApplicationId    string `json:"application_id"`
			TourDate         string `json:"tour_date"`
			TourType         string `json:"tour_type"`
			VoucherCode      string `json:"voucher_code"`
		} `json:"metadata,omitempty"`
		Description     string      `json:"description,omitempty"`
		PhoneNumber     string      `json:"phonenumber,omitempty"`
		Status          string      `json:"status,omitempty"`
		LastError       interface{} `json:"last_error,omitempty"`
		RejectedReason  interface{} `json:"rejected_reason,omitempty"`
		RejectedBy      interface{} `json:"rejected_by,omitempty"`
		RejectedTime    interface{} `json:"rejected_time,omitempty"`
		CancelledReason interface{} `json:"cancelled_reason,omitempty"`
		CancelledBy     interface{} `json:"cancelled_by,omitempty"`
		CancelledTime   interface{} `json:"cancelled_time,omitempty"`
		Created         string      `json:"created,omitempty"`
		Author          int         `json:"author,omitempty"`
		Modified        string      `json:"modified,omitempty"`
		UpdatedBy       interface{} `json:"updated_by,omitempty"`
		StartDate       string      `json:"start_date,omitempty"`
		MfsCode         string      `json:"mfs_code,omitempty"`
	} `json:"data"`
}

func CollectionCallback(ctx *gin.Context, c pb.PaymentServiceClient) {
	req := &pb.CollectionCallbackRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CollectionCallback(context.Background(), &pb.CollectionCallbackRequest{
		Hook: req.Hook,
		Data: req.Data,
	})

	fmt.Println(req)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
