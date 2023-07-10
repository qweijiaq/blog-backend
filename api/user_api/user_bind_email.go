package user_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"backend/plugins/email"
	"backend/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	var ber BindEmailRequest
	if err := c.ShouldBindJSON(&ber); err != nil {
		res.FailWithError(err, &ber, c)
		return
	}
	session := sessions.Default(c)
	if ber.Code == nil {
		// 第一次，后台发验证码
		// 生成 4 位验证码，将生成的验证码传入 session
		code := utils.Code(4)
		session.Set("valid_code", code)
		err := session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = email.NewCode().Send(ber.Email, "你的验证码是: "+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}
	// 第二次，用户输入邮箱，验证码，密码
	code := session.Get("valid_code")
	// 校验验证码
	if code != *ber.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	// 修改用户的邮箱
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserId).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if len(ber.Password) < 4 {
		res.FailWithMessage("密码长度小于4位", c)
		return
	}
	hashPwd := utils.HashPwd(ber.Password)
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    ber.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 完成绑定
	res.OkWithMessage("邮箱绑定成功", c)
}
