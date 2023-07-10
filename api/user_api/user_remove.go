package user_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var mr models.RemoveRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, mr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO: 删除用户、消息表、评论表、用户收藏和发布的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)
}
