package models

import (
	"backend/global"
	"backend/models/ctype"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID        string `json:"id" structs:"id"`
	CreatedAt string `json:"created_at" structs:"created_at"`
	UpdatedAt string `json:"updated_at" structs:"updated_at"`

	Title    string `json:"title" structs:"title"`
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"`
	Abstract string `json:"abstract" structs:"abstract"`
	Content  string `json:"content,omit(list)" structs:"content"`

	LookCount    int `json:"look_count" structs:"look_count"`
	CommentCount int `json:"comment_count" structs:"comment_count"`
	DiggCount    int `json:"digg_count" structs:"digg_count"`
	CollectCount int `json:"collect_count" structs:"collect_count"`

	UserID       uint   `json:"user_id" structs:"user_id"`
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"`
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`

	Source string `json:"source" structs:"source"`
	Link   string `json:"link" structs:"link"`
	//Words          int             `json:"words"`

	BannerUrl string `json:"banner_url" structs:"banner_url"`
	BannerID  int    `json:"banner_id" structs:"banner_id"`

	Tags ctype.Array `json:"tags" structs:"tags"`

	Category ctype.Array `json:"category" structs:"category"`
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collect_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "text"
      },
      "source": { 
        "type": "text"
      },
      "link": { 
        "type": "text"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "keyword"
      },
	  "tags": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (a ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		a.RemoveIndex()
	}
	// 没有索引
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil
}

// RemoveIndex 删除索引
func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}

// Create 添加的方法
func (a *ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(a.Index()).
		BodyJson(a).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (a ArticleModel) ISExistData() bool {
	res, err := global.ESClient.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

// 文章是否存在
//func (a ArticleModel) IsExistData() bool {
//	res, err := global.ESClient.
//		Search(a.Index()).
//		Query(elastic.NewTermQuery("keyword", a.Title)).
//		Size(1).
//		Do(context.Background())
//	if err != nil {
//		logrus.Error(err.Error())
//		return false
//	}
//	if res.Hits.TotalHits.Value > 0 {
//		return true
//	}
//	return false
//}

func (a *ArticleModel) GetDataById(id string) error {
	res, err := global.ESClient.
		Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}

	json.Unmarshal(res.Source, a)

	return nil
}
