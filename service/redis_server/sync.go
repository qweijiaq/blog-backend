package redis_server

import (
	"backend/global"
	"backend/models"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func DataSync() {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	diggInfo := NewDigg().GetInfo()
	lookInfo := NewArticleLook().GetInfo()
	commentInfo := NewCommentCount().GetInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)

		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment
		if article.DiggCount == newDigg && article.LookCount == newLook && article.CommentCount == newComment {
			logrus.Info(article.Title, "点赞数、浏览数和评论数无变化")
			continue
		}
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    newDigg,
				"look_count":    newLook,
				"comment_count": newComment,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Infof("%s: 数据同步成功， 点赞数为 %d, 浏览数为 %d, 评论数为 %d", article.Title, newDigg, newLook, newComment)
	}
	NewDigg().Clear()
	NewArticleLook().Clear()
	NewCommentCount().Clear()

}
