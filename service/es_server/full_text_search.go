package es_server

import (
	"backend/global"
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
	"strings"
)

type SearchData struct {
	KW    string `json:"kw"`
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))
			//if strings.TrimSpace(body) != "" {
			bodyList = append(bodyList, getBody(body))
			//}
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			KW:    id,
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
		})
	}
	b, _ := json.Marshal(searchDataList)
	fmt.Println(string(b))
	return searchDataList
}

func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}

func AsyncArticleByFullText(id, title, content string) {

	indexList := GetSearchIndexDataByContent(id, title, content)

	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Infof("%s 添加成功, 共 %d 条", title, len(result.Succeeded()))
}

// 删除文章全文数据
func DeleteFullTextByArticleId(id string) {
	boolSearch := elastic.NewTermQuery("kw", id)
	res, _ := global.ESClient.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	logrus.Infof("成功删除 %d 条记录", res.Deleted)
}

// 通过文章 ID 获取文章标题
//type Document struct {
//	Title string `json:"title"`
//}
//
//func GetArticleTitleByID(id string) (string, error) {
//
//	// 构造查询请求
//	res, err := global.ESClient.Get().
//		Index("article_index").
//		Id(id).
//		Do(context.Background())
//	if err != nil {
//		return "", fmt.Errorf("Error getting response: %s", err)
//	}
//
//	// 解析查询结果
//	var doc Document
//	err = json.Unmarshal(res.Source, &doc)
//	if err != nil {
//		return "", fmt.Errorf("Error parsing the response body: %s", err)
//	}
//
//	// 返回文章标题
//	return doc.Title, nil
//}
