package routers

import "backend/api"

func (router RouterGroup) CatRouter() {
	CatApi := api.ApiGroupApp.CatApi
	router.GET("cats", CatApi.CatListView)
	router.GET("cat_names", CatApi.CatNameListView)
	router.POST("cats", CatApi.CatCreateView)
	router.PUT("cats/:id", CatApi.CatUpdateView)
	router.DELETE("cats", CatApi.CatRemoveView)
}
