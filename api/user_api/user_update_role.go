package user_api

import (
	"backend/global"
	"backend/models"
	"backend/models/ctype"
	"backend/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"` // 防止用户权限非法，管理员有能力修改
	UserID   uint       `json:"user_id" binding:"required" msg:"用户ID错误"`
}

func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var ur UserRole
	if err := c.ShouldBindJSON(&ur); err != nil {
		res.FailWithError(err, &ur, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, ur.UserID).Error
	if err != nil {
		res.FailWithMessage("用户id错误，该用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      ur.Role,
		"nick_name": ur.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("修改权限失败", c)
		return
	}
	res.OkWithMessage("修改权限成功", c)
}
