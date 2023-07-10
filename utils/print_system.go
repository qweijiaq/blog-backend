package utils

import "backend/global"

func PrintSystem() {
	ip := global.Config.System.Host
	port := global.Config.System.Port
	if ip == "0.0.0.0" {
		IPList := GetIPList()
		for _, i := range IPList {
			global.Log.Infof("Gin 服务器运行在: http://%s:%d", i, port)
			global.Log.Infof("后端接口文档运行在: http://%s:%d/swagger/index.html#", i, port)
		}
	} else {
		global.Log.Infof("Gin 服务器运行在: http://%s:%d", ip, port)
		global.Log.Infof("后端接口文档运行在: http://%s:%d/swagger/index.html#", ip, port)
	}
}
