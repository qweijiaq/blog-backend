package menu_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var mr MenuRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithError(err, &mr, c)
		return
	}
	id := c.Param("id")
	// 先把之前的 banner 清空
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	err = global.DB.Model(&menuModel).Association("Banners").Clear()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片清空失败", c)
		return
	}
	// 如果选择了 banner，那就添加
	if len(mr.ImageSortList) > 0 {
		var bannerList []models.MenuBannerModel
		for _, sort := range mr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}
	// 普通更新
	maps := structs.Map(&mr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}

	res.OkWithMessage("修改菜单成功", c)
}
