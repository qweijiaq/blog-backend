package tag_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" structs:"title"`
}

func (TagApi) TagCreateView(c *gin.Context) {
	var tr TagRequest
	err := c.ShouldBindJSON(&tr)
	if err != nil {
		res.FailWithError(err, &tr, c)
		return
	}

	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", tr.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}
	err = global.DB.Create(&models.TagModel{
		Title: tr.Title,
	}).Error
	if err != nil {
		res.FailWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)
}
