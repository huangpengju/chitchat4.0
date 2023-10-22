package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	// cors.New函数创建了一个CORS中间件实例
	return cors.New(cors.Config{
		// 当函数返回 true 时，表示允许该来源的跨源请求；返回 false 则表示不允许。如果你希望对允许访问的来源有更精细的控制，你可以在函数中添加适当的逻辑。例如，你可以检查 origin 是否匹配某个特定的模式，或者检查请求的其他头或参数。
		AllowOriginFunc: func(origin string) bool {
			return true
		}, // AllowOriginFunc是一个用于验证起源的自定义函数。它将origin原点作为参数，如果允许则返回true，否则返回false。如果设置了这个选项，AllowOrigins的内容将被忽略。
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST", "OPTIONS"},          // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},             // 允许暴露的
		AllowCredentials: true,                                                                  // AllowCredentials指示请求是否可以包含用户凭据，如cookie、HTTP身份验证或客户端SSL证书。
		MaxAge:           12 * time.Hour,                                                        // 表示预检请求的结果可以缓存多长时间(以秒精度计算)
		AllowWebSockets:  true,                                                                  // 允许使用WebSocket协议
	})
}
