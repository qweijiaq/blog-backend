package article_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"backend/service/es_server"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

// ArticleCollCreateView 用户收藏文章，或取消收藏
func (ArticleApi) ArticleCollectView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

	model, err := es_server.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	var coll models.UserCollectsModel
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserId, cr.ID).Error
	var num = -1
	if err != nil {
		// 没有找到 收藏文章
		global.DB.Create(&models.UserCollectsModel{
			UserID:    claims.UserId,
			ArticleID: cr.ID,
		})
		// 给文章的收藏数 +1
		num = 1
	} else {
		// 取消收藏
		// 文章数 -1
		global.DB.Where("user_id =?", claims.UserId).Delete(&coll)
	}

	// 更新文章收藏数
	err = es_server.ArticleUpdate(cr.ID, map[string]any{
		"collect_count": model.CollectCount + num,
	})
	if num == 1 {
		res.OkWithMessage("收藏文章成功", c)
	} else {
		res.OkWithMessage("取消收藏成功", c)
	}
}
