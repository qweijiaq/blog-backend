package routers

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) CommentRouter() {
	CommentApi := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), CommentApi.CommentCreateView)
	router.GET("comments", middleware.JwtAuth(), CommentApi.CommentListView)
	router.GET("comments/:id", middleware.JwtAuth(), CommentApi.CommentDigg)
	router.DELETE("comments/:id", middleware.JwtAuth(), CommentApi.CommentRemoveView)
}
