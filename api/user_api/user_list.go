package user_api

import (
	"backend/models"
	"backend/models/ctype"
	"backend/models/res"
	"backend/service/common"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.UserModel
	RoleId int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	// 判断是否是管理员
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

	var page UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		fmt.Println(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []UserResponse
	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(page.Role)}, common.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"nick_name"},
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 非管理员
			user.UserName = ""
		}
		// 数据脱敏
		user.Tel = utils.DesensitizationTel(user.Tel)
		user.Email = utils.DesensitizationEmail(user.Email)

		users = append(users, UserResponse{
			UserModel: user,
			RoleId:    int(user.Role),
		})
	}
	res.OkWithList(users, count, c)
}
