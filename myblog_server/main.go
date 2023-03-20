package main

import (
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("myblog_server运行在：%s", addr)
	router.Run(addr)
}
