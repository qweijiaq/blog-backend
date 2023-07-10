package settings_api

import (
	"backend/global"
	"backend/models/res"

	"github.com/gin-gonic/gin"
)

// SettingsSiteInfoView 显示网站信息
// @Tags 系统管理
// @Summary 显示网站信息
// @Description 显示网站信息
// @Router /api/settings/site [get]
// @Produce json
// @Success 200 {object} res.Response{data=config.SiteInfo}
func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
