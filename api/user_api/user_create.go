package user_api

import (
	"backend/models/ctype"
	"backend/models/res"
	"backend/service/user_server"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string `json:"nick_name" binding:"required" msg:"请输入昵称"`
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
	Role     int    `json:"role" binding:"required" msg:"请选择角色权限"`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var ucr UserCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		res.FailWithError(err, &ucr, c)
		return
	}
	err := user_server.UserService{}.CreateUser(ucr.UserName, ucr.NickName, ucr.Password, ctype.Role(ucr.Role), "", c.ClientIP())
	if err != nil {
		//global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("创建用户%s并入库成功", ucr.UserName), c)
	return
}
