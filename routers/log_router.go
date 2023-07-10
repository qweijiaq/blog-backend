package routers

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) LogRouter() {
	LogApi := api.ApiGroupApp.LogApi
	router.GET("logs", LogApi.LogListView)
	router.DELETE("logs", middleware.JwtAdmin(), LogApi.LogRemoveView)
}
