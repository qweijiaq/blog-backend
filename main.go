package main

import (
	"backend/core"
	_ "backend/docs"
	"backend/flag"
	"backend/global"
	"backend/routers"
	"backend/utils"
)

// @title backend API 文档
// @version 1.0
// @description blog_backend API 文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 连接 Redis
	global.Redis = core.ConnectRedis()
	// 连接 ES
	global.ESClient = core.EsConnect()

	core.InitAddrDB()
	defer global.AddrDB.Close()

	option := flag.Parse()
	//fmt.Println(option)
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	utils.PrintSystem()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
