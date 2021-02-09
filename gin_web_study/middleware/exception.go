package middleware

import (
	"fmt"
	"gin_web_study/config/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

//全局异常处理中间件
func Exception(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 将异常写入日志
			global.Log.Error(fmt.Sprintf("[Exception]未知异常: %v\n堆栈信息: %v", err, string(debug.Stack())))
			// 服务器异常
			// 以json方式写入响应
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errorToString(err),
			})
			return
		}
	}()
	c.Next()
}
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}