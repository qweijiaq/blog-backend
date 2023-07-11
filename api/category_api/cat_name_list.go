package category_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type CatResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// CatNameListView 分类名称列表
// @Tags 分类管理
// @Summary 分类名称列表
// @Description 分类名称列表
// @Router /api/cat_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]CatResponse}
func (CatApi) CatNameListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	query := elastic.NewBoolQuery()
	agg := elastic.NewTermsAggregation().Field("category")
	result, err := global.ESClient.Search(models.ArticleModel{}.Index()).Query(query).Aggregation("tags", agg).Size(0).Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["tags"]
	var tagType T
	json.Unmarshal(byteData, &tagType)

	var catList = make([]CatResponse, 0)
	for _, bucket := range tagType.Buckets {
		catList = append(catList, CatResponse{
			Label: bucket.Key,
			Value: bucket.Key,
		})
	}
	res.OkWithData(catList, c)
}
