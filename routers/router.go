package routers

import (
	"net/http"

	docs "chitchat4.0/docs"
	"chitchat4.0/pkg/setting"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

	// Swagger API 文档的路由

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/admin", Helloworld)
	return r
}

// @BasePath /admin

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /admin [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
