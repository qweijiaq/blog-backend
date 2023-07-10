package tag_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var tr TagRequest
	err := c.ShouldBindJSON(&tr)
	if err != nil {
		res.FailWithError(err, &tr, c)
		return
	}

	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}
	maps := structs.Map(&tr)
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		res.FailWithMessage("修改标签失败", c)
		return
	}
	res.OkWithMessage("修改标签成功", c)
}
