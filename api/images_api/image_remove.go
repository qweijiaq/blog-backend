package images_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ImageUpdateView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @Description 删除图片
// @Param data body models.RemoveRequest true "图片ID列表 -- 必须是数组类型"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var mr models.RemoveRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, mr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)

}
