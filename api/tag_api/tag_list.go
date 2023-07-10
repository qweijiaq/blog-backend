package tag_api

import (
	"backend/models"
	"backend/models/res"
	"backend/service/common"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	cr.Sort = "created_at desc"
	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.TagModel, 0)
		res.OkWithList(list, count, c)
		return
	}
	res.OkWithList(data, count, c)
}
