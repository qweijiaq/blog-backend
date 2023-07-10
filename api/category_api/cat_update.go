package category_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (CatApi) CatUpdateView(c *gin.Context) {
	id := c.Param("id")
	var tr CatRequest
	err := c.ShouldBindJSON(&tr)
	if err != nil {
		res.FailWithError(err, &tr, c)
		return
	}

	var cat models.CategoryModel
	err = global.DB.Take(&cat, id).Error
	if err != nil {
		res.FailWithMessage("分类不存在", c)
		return
	}
	maps := structs.Map(&tr)
	err = global.DB.Model(&cat).Updates(maps).Error
	if err != nil {
		res.FailWithMessage("修改分类失败", c)
		return
	}
	res.OkWithMessage("修改分类成功", c)
}
