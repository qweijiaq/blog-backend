package category_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"github.com/gin-gonic/gin"
)

type CatRequest struct {
	Title string `json:"title" binding:"required" structs:"title"`
}

func (CatApi) CatCreateView(c *gin.Context) {
	var tr CatRequest
	err := c.ShouldBindJSON(&tr)
	if err != nil {
		res.FailWithError(err, &tr, c)
		return
	}

	var cat models.CategoryModel
	err = global.DB.Take(&cat, "title = ?", tr.Title).Error
	if err == nil {
		res.FailWithMessage("该分类已存在", c)
		return
	}
	err = global.DB.Create(&models.CategoryModel{
		Title: tr.Title,
	}).Error
	if err != nil {
		res.FailWithMessage("添加分类失败", c)
		return
	}
	res.OkWithMessage("添加分类成功", c)
}
