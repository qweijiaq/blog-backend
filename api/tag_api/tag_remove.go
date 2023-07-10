package tag_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var mr models.RemoveRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var tags []models.TagModel
	count := global.DB.Find(&tags, mr.IDList).RowsAffected
	if count == 0 {
		res.OkWithMessage("标签不存在", c)
		return
	}

	global.DB.Delete(tags)
	res.FailWithMessage(fmt.Sprintf("共删除%d个标签", count), c)
}
