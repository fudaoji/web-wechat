package main

import (
	"web-wechat/core"

	"github.com/gin-gonic/gin"

	. "web-wechat/db"
	"web-wechat/global"
	"web-wechat/logger"
	"web-wechat/middleware"

	//"web-wechat/oss"
	"web-wechat/route"
)

// 程序启动入口
func main() {
	// 初始化日志
	logger.InitLogger()
	// 初始化Gin
	app := gin.Default()

	// 定义全局异常处理
	app.NoRoute(core.NotFoundErrorHandler())
	// AppKey预检
	app.Use(middleware.CheckAppKeyExistMiddleware(), middleware.CheckAppKeyIsLoggedInMiddleware())
	// 初始化路由
	route.InitRoute(app)

	// 初始化WechatBotMap
	global.InitWechatBotsMap()

	//读取配置
	core.InitConfig()

	// 初始化OSS
	//oss.InitOssConnHandle()

	// 初始化MongoDB
	//InitMongoConnHandle()

	// 初始化Redis连接
	InitRedisConnHandle()

	// 初始化Redis里登录的数据
	global.InitBotWithStart()

	// 定时更新 Bot 的热登录数据
	global.UpdateHotLoginData()

	// 保活
	global.KeepAliveHandle()

	// 监听端口
	_ = app.Run(":8889")
}
