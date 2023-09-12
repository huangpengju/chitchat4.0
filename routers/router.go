package routers

import (
	"net/http"

	"chitchat4.0/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 设置应用的模式【debug 开发 release 发布】
	gin.SetMode(setting.RunMode) // 自定义包 setting 中的 RunMode 存放配置文件中 RUN_MODE 的值，现在代替gin.SetMode(gin.DebugMode)

	// 创建不带中间件的路由
	r := gin.New()
	// 挂载中间件
	// 全局中间件 Logger 中间件将日志写入 gin.DefaultWriter
	r.Use(gin.Logger())
	// Recovery 中间件会 revover 恢复 任何 panic
	r.Use(gin.Recovery())

	r.GET("/admin", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "成功了",
		})
	})
	return r
}
