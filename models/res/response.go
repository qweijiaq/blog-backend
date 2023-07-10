package res

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ListResponse[T any] struct {
	List  T     `json:"list"`
	Count int64 `json:"count"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(200, Response{Code: code, Data: data, Msg: msg})
}

func Ok(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func OkContext(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "成功", c)
}

func OkWithList(list any, count int64, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func Fail(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]any{}, message, c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, msg, c)
}

// func OkWithList[T any](List []T, count any, c *gin.Context) {
// 	if len(List) == 0 {
// 		List = []T{}
// 	}
// 	Result(SUCCESS, ListResponse[T]{
// 		List: List,
// 		Count: count,
// 	}, "成功", c)
// }
