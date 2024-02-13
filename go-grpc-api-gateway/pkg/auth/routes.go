package auth

import (
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth/routes"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/account_registration", svc.Register)
	routes.POST("/account_login", svc.Login)
	routes.POST("/account_registration_client", svc.RegisterUserClient)
	routes.POST("/account_login_client", svc.LoginUserClient)
	routes.POST("/request_account_edit", svc.EditAccount)
	routes.POST("/phone_number_verification", svc.PhoneNumberVerification)
	routes.POST("/verify_phone_number_with_otp", svc.PhoneNumberVerificationWithOTP)
	routes.POST("/reset_password", svc.ResetPassword)
	routes.POST("/verify_reset_token", svc.VerifyResetToken)
	routes.POST("/resend_otp", svc.ResendOTP)
	routes.POST("/kyc_callback", svc.ReceiveAndHandleKYCCallback)
	routes.POST("/send_user_text_message", svc.SendTextMessage)
	routes.POST("/send_user_email", svc.SendEmailMessage)
	routes.POST("/send_user_push_message", svc.SendPushMessage)
	routes.POST("/validate_email", svc.ValidateEmailAddress)
	routes.PUT("/update_password", svc.UpdatePassword)
	routes.PUT("/update_user_docs", svc.UpdateUserDocs)
	routes.PUT("/update_phone_number", svc.UpdatePhoneNumber)
	routes.PUT("/deactivate_user", svc.DeactivateUser)
	routes.PUT("/lock_account", svc.LockAccount)
	routes.GET("/get_all_users", svc.GetAllUsers)
	routes.GET("/:id", svc.GetUser)
	routes.PUT("/update_user_device_uuid", svc.GetUserDeviceUUID)
	routes.PUT("/update_landlord_invite_status", svc.InvitedLandlord)
	routes.DELETE("/delete_user/:id", svc.DeleteUser)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) PhoneNumberVerification(ctx *gin.Context) {
	routes.PhoneNumberVerification(ctx, svc.Client)
}

func (svc *ServiceClient) PhoneNumberVerificationWithOTP(ctx *gin.Context) {
	routes.PhoneNumberVerificationWithOTP(ctx, svc.Client)
}

func (svc *ServiceClient) ResetPassword(ctx *gin.Context) {
	routes.ResetPassword(ctx, svc.Client)
}

func (svc *ServiceClient) VerifyResetToken(ctx *gin.Context) {
	routes.VerifyResetToken(ctx, svc.Client)
}

func (svc *ServiceClient) UpdatePassword(ctx *gin.Context) {
	routes.UpdatePassword(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateUserDocs(ctx *gin.Context) {
	routes.UpdateDocument(ctx, svc.Client)
}

func (svc *ServiceClient) ResendOTP(ctx *gin.Context) {
	routes.ResendOTP(ctx, svc.Client)
}

func (svc *ServiceClient) UpdatePhoneNumber(ctx *gin.Context) {
	routes.UpdatePhoneNumber(ctx, svc.Client)
}

func (svc *ServiceClient) DeactivateUser(ctx *gin.Context) {
	routes.DeactivateUser(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) {
	routes.DeleteUser(ctx, svc.Client)
}

func (svc *ServiceClient) LockAccount(ctx *gin.Context) {
	routes.LockAccount(ctx, svc.Client)
}

func (svc *ServiceClient) ReceiveAndHandleKYCCallback(ctx *gin.Context) {
	routes.KYCVerificationResultCallback(ctx, svc.Client)
}

func (svc *ServiceClient) GetUser(ctx *gin.Context) {
	routes.GetUser(ctx, svc.Client)
}

func (svc *ServiceClient) ValidateEmailAddress(ctx *gin.Context) {
	routes.ValidateEmail(ctx, svc.Client)
}

func (svc *ServiceClient) GetAllUsers(ctx *gin.Context) {
	routes.GetAllUsers(ctx, svc.Client)
}

func (svc *ServiceClient) SendTextMessage(ctx *gin.Context) {
	routes.SendTextMessage(ctx, svc.Client)
}

func (svc *ServiceClient) SendEmailMessage(ctx *gin.Context) {
	routes.SendEmailMessage(ctx, svc.Client)
}

func (svc *ServiceClient) SendPushMessage(ctx *gin.Context) {
	routes.SendPushMessage(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserDeviceUUID(ctx *gin.Context) {
	routes.GetUserDeviceUUID(ctx, svc.Client)
}

func (svc *ServiceClient) InvitedLandlord(ctx *gin.Context) {
	routes.InvitedLandlord(ctx, svc.Client)
}

func (svc *ServiceClient) EditAccount(ctx *gin.Context) {
	routes.EditAccount(ctx, svc.Client)
}

func (svc *ServiceClient) RegisterUserClient(ctx *gin.Context) {
	routes.RegisterUserClient(ctx, svc.Client)
}

func (svc *ServiceClient) LoginUserClient(ctx *gin.Context) {
	routes.LoginUserClient(ctx, svc.Client)
}

