package settings_api

import (
	"backend/global"
	"backend/models/res"

	"github.com/gin-gonic/gin"
)

type SettingsURI struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
// @Tags 系统管理
// @Summary 显示某一项的配置信息
// @Description 显示某一项的配置信息 site email qq qiniu jwt
// @Param name path string true "name"
// @Param token header string true "token"
// @Router /api/settings/{name} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsURI
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	case "upload":
		res.OkWithData(global.Config.Upload, c)

	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}

}
