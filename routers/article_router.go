package routers

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) ArticleRouter() {
	ArticleApi := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), ArticleApi.ArticleCreateView)
	router.PUT("articles", middleware.JwtAuth(), ArticleApi.ArticleUpdateView)
	router.GET("articles", ArticleApi.ArticleListView)
	router.DELETE("articles", middleware.JwtAuth(), ArticleApi.ArticleRemoveView)
	router.GET("articles/detail", ArticleApi.ArticleDetailByTitleView)
	router.GET("articles/:id", ArticleApi.ArticleDetailView)
	router.GET("articles/calendar", ArticleApi.ArticleCalendarView)
	router.GET("articles/tags", ArticleApi.ArticleTagListView)
	router.GET("articles/cats", ArticleApi.ArticleCatView)
	router.POST("articles/collect", middleware.JwtAuth(), ArticleApi.ArticleCollectView)
	router.GET("articles/collect", middleware.JwtAuth(), ArticleApi.ArticleCollListView)
	router.DELETE("articles/collect", middleware.JwtAuth(), ArticleApi.ArticleCollBatchRemoveView)
	router.GET("articles/fulltext", ArticleApi.ArticleFullTextSearchView)
}
