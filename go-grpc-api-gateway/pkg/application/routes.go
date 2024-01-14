package application

import (
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/application/routes"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/auth"
	"github.com/RonnieZad/nyumba-go-grpc-project/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/application")
	routes.Use(a.AuthRequired)
	routes.POST("/add_advert_label", svc.AddLabelAdvert)
	routes.POST("/record_label_scan", svc.RecordLabelScan)
	routes.GET("/get_user_label_scan/:id", svc.GetUserLabelScan)
	routes.GET("/get_user_label_advert/:id", svc.GetUserLabelAdvert)
	routes.GET("/get_metrics/:id", svc.GetCompanyDashboardAnalytic)
}

func (svc *ServiceClient) AddLabelAdvert(ctx *gin.Context) {
	routes.AddLabelAdvert(ctx, svc.Client)
}

func (svc *ServiceClient) RecordLabelScan(ctx *gin.Context) {
	routes.AddLabelScan(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserLabelScan(ctx *gin.Context) {
	routes.GetUserLabelScan(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserLabelAdvert(ctx *gin.Context) {
	routes.GetUserLabelAdvert(ctx, svc.Client)
}


func (svc *ServiceClient) GetCompanyDashboardAnalytic(ctx *gin.Context) {
	routes.GetCompanyDashboardAnalytic(ctx, svc.Client)
}
