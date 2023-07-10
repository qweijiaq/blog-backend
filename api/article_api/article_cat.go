package article_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CatResponse struct {
	Cat           string   `json:"cat"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"`
}

type CatType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

func (ArticleApi) ArticleCatView(c *gin.Context) {

	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("category", elastic.NewCardinalityAggregation().Field("category")).
		Size(0).
		Do(context.Background())
	cCat, _ := result.Aggregations.Cardinality("category")
	count := int64(*cCat.Value)

	agg := elastic.NewTermsAggregation().Field("category")

	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

	query := elastic.NewBoolQuery()

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("category", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	var catType CatType
	var catList = make([]*CatResponse, 0)
	_ = json.Unmarshal(result.Aggregations["category"], &catType)
	var catStringList []string
	for _, bucket := range catType.Buckets {

		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}

		catList = append(catList, &CatResponse{
			Cat:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})

		catStringList = append(catStringList, bucket.Key) // -2
	}

	var catModelList []models.CategoryModel
	global.DB.Find(&catModelList, "title in ?", catStringList)
	var catDate = map[string]string{}
	for _, model := range catModelList {
		catDate[model.Title] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	for _, response := range catList {
		response.CreatedAt = catDate[response.Cat]
	}

	res.OkWithList(catList, count, c)
}
