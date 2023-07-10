package routers

import "backend/api"

func (router RouterGroup) DiggRouter() {
	diggApi := api.ApiGroupApp.DiggApi
	router.POST("digg/article", diggApi.DiggArticleView)
}
