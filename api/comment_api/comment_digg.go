package comment_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"backend/service/redis_server"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (CommentApi) CommentDigg(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}

	redis_server.NewCommentDigg().Set(fmt.Sprintf("%d", cr.ID))

	res.OkWithMessage("评论点赞成功", c)
	return

}
