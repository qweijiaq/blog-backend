package images_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择图片ID"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// ImageUpdateView 更新图片
// @Tags 图片管理
// @Summary 更新图片
// @Description 更新图片
// @Param data body ImageUpdateRequest     true "图片ID"
// @Router /api/images [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var iur ImageUpdateRequest
	err := c.ShouldBindJSON(&iur)
	if err != nil {
		res.FailWithError(err, &iur, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, iur.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
	}
	err = global.DB.Model(&imageModel).Update("name", iur.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
}
