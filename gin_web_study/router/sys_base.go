package router

import (
	"gin_web_study/api"
	"gin_web_study/middleware"
	"github.com/gin-gonic/gin"
)
//公共路由,任何人可访问
func InitPublicRouter(r *gin.Engine) (R gin.IRoutes) {
	{
		r.POST("get", api.GetLeaves)
		r.GET("save", api.SaveTest)
	}
	return r
}

func InitBaseRouter(r *gin.Engine) (R gin.IRoutes) {
	r.Use(middleware.LoginAuth())
	{
	}
	return r
}

