package api

import (
	"backend/api/article_api"
	"backend/api/category_api"
	"backend/api/chat_api"
	"backend/api/comment_api"
	"backend/api/data_api"
	"backend/api/digg_api"
	"backend/api/images_api"
	"backend/api/log_api"
	"backend/api/menu_api"
	"backend/api/message_api"
	"backend/api/news_api"
	"backend/api/settings_api"
	"backend/api/tag_api"
	"backend/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
	CatApi      category_api.CatApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
	NewsApi     news_api.NewApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	DataApi     data_api.DataApi
}

var ApiGroupApp = new(ApiGroup)
