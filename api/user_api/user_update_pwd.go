package user_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"`
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`
}

func (UserApi) UserUpdatePwdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	var upr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&upr); err != nil {
		res.FailWithError(err, &upr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserId).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if !utils.CheckPwd(user.Password, upr.OldPwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := utils.HashPwd(upr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
	return
}
