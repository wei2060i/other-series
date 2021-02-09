package initialize

import (
	"gin_web_study/config/global"
	"gin_web_study/middleware"
	"gin_web_study/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	// r := gin.Default()
	// 创建不带中间件的路由:
	r := gin.New()
	//添加全局异常处理中间件
	r.Use(middleware.Exception)
	// ping
	//apiGroup.GET("/ping", api.Ping)
	router.InitPublicRouter(r)                       // 注册公共路由
	global.Log.Info("初始化路由完成")
	return r
}

