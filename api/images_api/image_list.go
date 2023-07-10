package images_api

import (
	"backend/models"
	"backend/models/res"
	"backend/service/common"

	"github.com/gin-gonic/gin"
)

// ImageUpdateView 查询图片
// @Tags 图片管理
// @Summary 查询图片
// @Description 查询图片
// @Param data query models.PageInfo true "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImagesApi) ImageListView(c *gin.Context) {
	var p models.PageInfo
	err := c.ShouldBindQuery(&p)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: p,
		Likes:    []string{"name"},
	})
	res.OkWithList(list, count, c)
	return
}
