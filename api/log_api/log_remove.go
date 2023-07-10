package log_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"backend/plugins/log"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (LogApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var list []log.LogModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&list)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)

}
