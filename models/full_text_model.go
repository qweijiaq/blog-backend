package models

import (
	"backend/global"
	"context"
	"github.com/sirupsen/logrus"
)

type FullTextModel struct {
	ID    string `json:"id" structs:"id"`
	KW    string `json:"kw" structs:"kw"` // 文章关联的 id
	Title string `json:"title" structs:"title"`
	Slug  string `json:"slug" structs:"slug"`
	Body  string `json:"body" structs:"body"`
}

func (FullTextModel) Index() string {
	return "full_text_index"
}

func (FullTextModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
	  "kw": {
      	"type": "keyword"
      },
      "title": { 
        "type": "text"
      },
      "slug": { 
        "type": "keyword"
      },
      "body": { 
        "type": "text"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (a FullTextModel) IndexExists() bool {
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
func (a FullTextModel) CreateIndex() error {
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
func (a FullTextModel) RemoveIndex() error {
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
