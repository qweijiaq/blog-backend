package routers

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/site", settingsApi.SettingsSiteInfoView)
	router.PUT("settings/site", middleware.JwtAdmin(), settingsApi.SettingsSiteUpdateView)
	router.GET("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoView)
	router.PUT("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoUpdateView)
}
