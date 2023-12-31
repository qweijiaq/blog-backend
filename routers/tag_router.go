package routers

import "backend/api"

func (router RouterGroup) TagRouter() {
	TagApi := api.ApiGroupApp.TagApi
	router.GET("tags", TagApi.TagListView)
	router.GET("tag_names", TagApi.TagNameListView)
	router.POST("tags", TagApi.TagCreateView)
	router.PUT("tags/:id", TagApi.TagUpdateView)
	router.DELETE("tags", TagApi.TagRemoveView)
}
