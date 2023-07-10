package routers

import "backend/api"

func (router RouterGroup) NewsRouter() {
	NewsApi := api.ApiGroupApp.NewsApi
	router.POST("news", NewsApi.NewListView)
}
