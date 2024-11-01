package controllers

import (
	"yky-gin/services"
	"yky-gin/utils/resp"

	"github.com/gin-gonic/gin"
)

type AppConfigController struct {
	appConfigService *services.AppConfigService
}

func NewAppConfigController(appConfigService *services.AppConfigService) *AppConfigController {
	return &AppConfigController{
		appConfigService: appConfigService,
	}
}
func (p *AppConfigController) GetAppConfig(c *gin.Context) {
	appConfig, _ := p.appConfigService.GetAppConfig()
	resp.RespHelper.OK(c, appConfig)
}
