package message_api

import (
	"backend/models"
	"backend/models/res"
	"backend/service/common"

	"github.com/liu-cn/json-filter/filter"

	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	cr.Sort = "created_at desc"
	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.MessageModel, 0)
		res.OkWithList(list, count, c)
		return
	}
	res.OkWithList(data, count, c)
}
