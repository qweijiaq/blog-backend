package data_api

import (
	"backend/global"
	"backend/models"
	"backend/models/res"
	"time"

	"github.com/gin-gonic/gin"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
}

func (DataApi) SevenLogin(c *gin.Context) {
	var loginDateCount, signDateCount []DateCount

	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDateCount)
	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&signDateCount)
	var loginDateCountMap = map[string]int{}
	var signDateCountMap = map[string]int{}
	var loginCountList, signCountList []int
	now := time.Now()
	for _, i2 := range loginDateCount {
		loginDateCountMap[i2.Date] = i2.Count
	}
	for _, i2 := range signDateCount {
		signDateCountMap[i2.Date] = i2.Count
	}
	var dateList []string
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		loginCount := loginDateCountMap[day]
		signCount := signDateCountMap[day]
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginCount)
		signCountList = append(signCountList, signCount)
	}

	res.OkWithData(DateCountResponse{
		DateList:  dateList,
		LoginData: loginCountList,
		SignData:  signCountList,
	}, c)

}
