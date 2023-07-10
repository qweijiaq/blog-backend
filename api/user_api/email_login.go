package user_api

import (
	"backend/global"
	"backend/models"
	"backend/models/ctype"
	"backend/models/res"
	log2 "backend/plugins/log"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLogin(c *gin.Context) {
	var er EmailLoginRequest
	err := c.ShouldBindJSON(&er)
	if err != nil {
		res.FailWithError(err, &er, c)
		return
	}

	log := log2.NewLogByGin(c)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ?", er.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		log.Warn(fmt.Sprintf("%s 用户名不存在", er.UserName))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isChecked := utils.CheckPwd(userModel.Password, er.Password)
	if !isChecked {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误 %s %s", er.UserName, er.Password))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := utils.GenToken(utils.JwtPayload{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserId:   userModel.ID,
		Avatar:   userModel.Avatar,
	})
	if err != nil {
		global.Log.Error(err)
		log.Error(fmt.Sprintf("token 生成失败 %s", err.Error()))
		res.FailWithMessage("token 生成失败", c)
		return
	}
	ip, addr := utils.GetAddrByGin(c)

	log = log2.New(c.ClientIP(), token)
	log.Info(fmt.Sprintf("%s 登录成功", er.UserName))

	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		Nickname:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})
	res.OkWithData(token, c)
}
