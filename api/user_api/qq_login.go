package user_api

import (
	"backend/global"
	"backend/models"
	"backend/models/ctype"
	"backend/models/res"
	"backend/plugins/qq"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	ip, addr := utils.GetAddrByGin(c)

	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.OkWithMessage(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID
	// 根据 openID 判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", openID).Error
	if err != nil {
		// 不存在，则注册
		user = models.UserModel{
			NickName:   qqInfo.NickName,
			UserName:   openID,                           // 直接使用 openID
			Password:   utils.HashPwd(utils.RandStr(16)), // 随机生成16位
			Avatar:     qqInfo.Avatar,
			Addr:       addr, // 根据 IP 计算
			Token:      openID,
			IP:         ip,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("注册失败", c)
			return
		}
	}
	// 登录操作
	token, err := utils.GenToken(utils.JwtPayload{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserId:   user.ID,
		Avatar:   user.Avatar,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token 生成失败", c)
		return
	}

	global.DB.Create(&models.LoginDataModel{
		UserID:    user.ID,
		IP:        ip,
		Nickname:  user.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})
	res.OkWithData(token, c)
}
