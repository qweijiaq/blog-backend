package routers

import (
	"backend/api"
)

func (router RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("menus", menuApi.MenuCreateView)
	router.GET("menus", menuApi.MenuListView)
	router.GET("menu_names", menuApi.MenuNameList)
	router.PUT("menus/:id", menuApi.MenuUpdateView)
	router.DELETE("menus", menuApi.MenuRemoveView)
	router.GET("menus/:id", menuApi.MenuDetailView)
}
