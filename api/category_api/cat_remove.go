package category_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (CatApi) CatRemoveView(c *gin.Context) {
	var mr models.RemoveRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var cats []models.CategoryModel
	count := global.DB.Find(&cats, mr.IDList).RowsAffected
	if count == 0 {
		res.OkWithMessage("分类不存在", c)
		return
	}

	global.DB.Delete(cats)
	res.FailWithMessage(fmt.Sprintf("共删除%d个分类", count), c)
}
