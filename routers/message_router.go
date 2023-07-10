package routers

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) MessageRouter() {
	MsgApi := api.ApiGroupApp.MessageApi
	router.POST("messages", middleware.JwtAuth(), MsgApi.MessageCreateView)
	router.GET("messages", middleware.JwtAuth(), MsgApi.MessageListView)
	router.GET("messages_all", middleware.JwtAuth(), MsgApi.MessageListAllView)
	router.POST("message_detail", middleware.JwtAuth(), MsgApi.MessageDetailView)
}
