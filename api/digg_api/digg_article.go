package digg_api

import (
	"backend/models"
	"backend/models/res"
	"backend/service/redis_server"
	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 对长度校验
	// 查es
	redis_server.NewDigg().GetInfo()
	res.OkWithMessage("文章点赞成功", c)
}
