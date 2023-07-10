package routers

import "backend/api"

func (router RouterGroup) ChatRouter() {
	ChatApi := api.ApiGroupApp.ChatApi
	router.GET("chat_group", ChatApi.ChatGroupView)
	router.GET("chat_list", ChatApi.ChatListView)
}
