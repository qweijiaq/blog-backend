package comment_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"backend/service/redis_server"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (CommentApi) CommentRemoveView(c *gin.Context) {
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
	// 统计评论下的子评论数 再把自己算上去
	subCommentList := FindSubCommentCount(commentModel)
	count := len(subCommentList) + 1
	redis_server.NewCommentCount().SetCount(commentModel.ArticleID, -count)

	// 判断是否是子评论
	if commentModel.ParentCommentID != nil {
		// 子评论
		// 找父评论，减掉对应的评论数
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	// 删除子评论以及当前评论
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	// 反转，然后一个一个删
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), c)
	return
}
