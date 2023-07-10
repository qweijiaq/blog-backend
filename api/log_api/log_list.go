package log_api

import (
	"backend/models"
	"backend/models/res"
	"backend/plugins/log"
	"backend/service/common"

	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInfo
	Level log.Level `form:"level"`
}

// LogListView 日志列表
// @Tags 日志管理
// @Summary 日志列表
// @param data query LogRequest false "查询参数"
// @Param level query int false "日志等级"
// @Router /api/logs [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[log.LogModel]}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	list, count, _ := common.ComList(log.LogModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
	return
}
