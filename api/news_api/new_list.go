package news_api

import (
	"backend/models/res"
	"backend/service/redis_server"
	"backend/utils"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewResponse struct {
	Code int                    `json:"code"`
	Data []redis_server.NewData `json:"data"`
	Msg  string                 `json:"msg"`
}

const newAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

func (NewApi) NewListView(c *gin.Context) {
	var cr params
	var headers header
	err := c.ShouldBindJSON(&cr)
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if cr.Size == 0 {
		cr.Size = 1
	}

	key := fmt.Sprintf("%s-%d", cr.ID, cr.Size)
	newsData, _ := redis_server.GetNews(key)
	if len(newsData) != 0 {
		res.OkWithData(newsData, c)
		return
	}

	httpResponse, err := utils.Post(newAPI, cr, structs.Map(headers), timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	var response NewResponse
	byteData, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}
	res.OkWithData(response.Data, c)
	redis_server.SetNews(key, response.Data)
	return
}
