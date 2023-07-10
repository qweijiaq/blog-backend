package article_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"backend/service/es_server"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
	Cat string `json:"cat" form:"cat"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_server.CommList(es_server.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "abstract", "content"},
		Tag:      cr.Tag,
		Cat:      cr.Cat,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
