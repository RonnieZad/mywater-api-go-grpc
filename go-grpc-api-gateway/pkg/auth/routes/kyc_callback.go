package routes

import (
	"context"
	"fmt"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerificationResultBody struct {
	DOB      string `json:"DOB"`
	FullName string `json:"FullName"`
	Gender   string `json:"Gender"`
	IDType   string `json:"IDType"`
	Actions  struct {
		LivenessCheck            string `json:"Liveness_Check"`
		RegisterSelfie           string `json:"Register_Selfie"`
		VerifyDocument           string `json:"Verify_Document"`
		HumanReviewCompare       string `json:"Human_Review_Compare"`
		ReturnPersonalInfo       string `json:"Return_Personal_Info"`
		SelfieToIDCardCompare    string `json:"Selfie_To_ID_Card_Compare"`
		HumanReviewLivenessCheck string `json:"Human_Review_Liveness_Check"`
	} `json:"Actions"`
	Country       string `json:"Country"`
	Document      string `json:"Document"`
	IDNumber      string `json:"IDNumber"`
	ResultCode    string `json:"ResultCode"`
	ResultText    string `json:"ResultText"`
	SmileJobID    string `json:"SmileJobID"`
	PartnerParams struct {
		JobID  string `json:"job_id"`
		UserID string `json:"user_id"`
		// JobType int32  `json:"job_type"`
	} `json:"PartnerParams"`
	ExpirationDate string `json:"ExpirationDate"`
	Timestamp      string `json:"timestamp"`
	Signature      string `json:"signature"`
	ImageLinks     struct {
		IdCardImage string `json:"id_card_image"`
		IdCardBack  string `json:"id_card_back"`
		SelfieImage string `json:"selfie_image"`
	} `json:"ImageLinks"`
}

func KYCVerificationResultCallback(ctx *gin.Context, c pb.AuthServiceClient) {
	req := &pb.KYCVerificationResultRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.KYCVerificationResultCallback(context.Background(), &pb.KYCVerificationResultRequest{
		Dob:            req.Dob,
		Fullname:       req.Fullname,
		Gender:         req.Gender,
		Idtype:         req.Idtype,
		Country:        req.Country,
		Document:       req.Document,
		Idnumber:       req.Idnumber,
		Resultcode:     req.Resultcode,
		Resulttext:     req.Resulttext,
		Smilejobid:     req.Smilejobid,
		Expirationdate: req.Expirationdate,
		Actions:        req.Actions,
		Partnerparams:  req.Partnerparams,
		Imagelinks:     req.Imagelinks,
	})

	fmt.Println(req)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
