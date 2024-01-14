package payment

import (
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/config"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/payment/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	// a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/payment")
	// routes.Use(a.AuthRequired)
	routes.POST("/make_collection", svc.MakeCollection)
	routes.POST("/send_payment", svc.SendPayment)
	routes.POST("/collection_callback", svc.ReceiveAndHandleCollectionCallback)
	routes.POST("/payment_callback", svc.ReceiveAndHandlePaymentCallback)
	routes.POST("/check_payment_subscription", svc.CheckPaymentSubscription)
	routes.POST("/check_property_hold_subscription", svc.CheckPropertyHoldSubscription)
	routes.POST("/check_phone_subscription", svc.CheckPhoneSubscription)
	routes.GET("/notifications/:id", svc.GetUserNotificationUpdates)
	routes.GET("/get_all_payment_transactions", svc.GetAllPaymentTransactions)
	routes.POST("/add_user_invite", svc.AppReferral)
	routes.POST("/verify_voucher_code", svc.CheckVoucherValidity)
	routes.GET("/get_user_referral/:id", svc.GetMyAppReferral)
	routes.PUT("/update_referral", svc.UpdateReferreral)
	routes.GET("/get_dash_metrics", svc.GetDashboardMetrics)
	routes.GET("/get_landlord_revenue/:id", svc.GetLandlordRevenue)
}

func (svc *ServiceClient) MakeCollection(ctx *gin.Context) {
	routes.MakeCollection(ctx, svc.Client)
}

func (svc *ServiceClient) SendPayment(ctx *gin.Context) {
	routes.SendPayment(ctx, svc.Client)
}

func (svc *ServiceClient) ReceiveAndHandleCollectionCallback(ctx *gin.Context) {
	routes.CollectionCallback(ctx, svc.Client)
}

func (svc *ServiceClient) ReceiveAndHandlePaymentCallback(ctx *gin.Context) {
	routes.PaymentCallback(ctx, svc.Client)
}

func (svc *ServiceClient) CheckPaymentSubscription(ctx *gin.Context) {
	routes.CheckSubscription(ctx, svc.Client)
}

func (svc *ServiceClient) CheckPhoneSubscription(ctx *gin.Context) {
	routes.CheckPhoneSubscription(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserNotificationUpdates(ctx *gin.Context) {
	routes.GetUserNotification(ctx, svc.Client)
}

func (svc *ServiceClient) GetAllPaymentTransactions(ctx *gin.Context) {
	routes.GetAllPaymentTransactions(ctx, svc.Client)
}

func (svc *ServiceClient) CheckPropertyHoldSubscription(ctx *gin.Context) {
	routes.CheckPropertyHoldSubscription(ctx, svc.Client)
}

func (svc *ServiceClient) CheckVoucherValidity(ctx *gin.Context) {
	routes.CheckVoucherValidity(ctx, svc.Client)
}

func (svc *ServiceClient) GetMyAppReferral(ctx *gin.Context) {
	routes.GetMyAppReferral(ctx, svc.Client)
}

func (svc *ServiceClient) AppReferral(ctx *gin.Context) {
	routes.AppReferral(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateReferreral(ctx *gin.Context) {
	routes.UpdateReferreral(ctx, svc.Client)
}

func (svc *ServiceClient) GetDashboardMetrics(ctx *gin.Context) {
	routes.GetDashboardMetrics(ctx, svc.Client)
}

func (svc *ServiceClient) GetLandlordRevenue(ctx *gin.Context) {
	routes.GetLandlordRevenue(ctx, svc.Client)
}


