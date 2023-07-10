package routers

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("mJAxmalxcoHsa73dNWwc8"))

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", userApi.EmailLogin)
	router.GET("login", userApi.QQLoginView)
	router.POST("users", middleware.JwtAdmin(), userApi.UserCreateView)
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("user_pwd", middleware.JwtAuth(), userApi.UserUpdatePwdView)
	router.GET("logout", middleware.JwtAuth(), userApi.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), userApi.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
}
