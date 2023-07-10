package menu_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var mr models.RemoveRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var MenuList []models.MenuModel
	count := global.DB.Find(&MenuList, mr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&MenuList).Association("Banners").Clear()
		if err != nil {
			return err
		}
		err = global.DB.Delete(&MenuList).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
