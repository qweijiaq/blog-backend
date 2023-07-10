package user_api

import (
	"backend/global"
	"backend/models/res"
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
