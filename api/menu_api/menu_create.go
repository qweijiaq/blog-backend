package menu_api

import (
	"backend/global"
	"backend/models"
	"backend/models/ctype"
	"backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                // 切换的时间，单位：秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    // 切换的时间，单位：秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          // 具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var mr MenuRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithError(err, &mr, c)
		return
	}
	// 重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? and path = ?", mr.Title, mr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("重复的菜单", c)
		return
	}
	// 创建 banner 数据入库
	MenuModel := models.MenuModel{
		Title:        mr.Title,
		Path:         mr.Path,
		Slogan:       mr.Slogan,
		Abstract:     mr.Abstract,
		AbstractTime: mr.AbstractTime,
		BannerTime:   mr.BannerTime,
		Sort:         mr.Sort,
	}
	err = global.DB.Create(&MenuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(mr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel
	for _, sort := range mr.ImageSortList {
		// 这里需要判断 image_id 是否真正有这张图片
		err = global.DB.Take(&models.BannerModel{}, sort.ImageID).Error
		if err != nil {
			res.FailWithMessage(fmt.Sprintf("ID为%d的图片不存在", sort.ImageID), c)
			return
		}
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   MenuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// 给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
