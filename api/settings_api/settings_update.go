package settings_api

import (
	"backend/config"
	"backend/core"
	"backend/global"
	"backend/models/res"

	"github.com/gin-gonic/gin"
)

// SettingsInfoUpdateView 修改某一项的配置信息
// @Tags 系统管理
// @Summary 修改某一项的配置信息
// @Description 修改某一项的配置信息 site email qq qiniu jwt
// @Param name path int true "name"
// @Router /api/settings/{name} [put]
// @Param token header string true "token"
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsURI
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = info
	case "upload":
		var info config.Upload
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Upload = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkContext(c)
}
