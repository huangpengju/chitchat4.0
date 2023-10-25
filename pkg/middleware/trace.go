package middleware

import (
	"time"

	"chitchat4.0/pkg/common"
	utiltrace "chitchat4.0/pkg/utils/trace"
	"github.com/bombsimon/logrusr/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TraceMiddleware 追踪中间件
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 返回一个 Trace 结构体
		// Trace 跟踪一组“步骤”，并允许我们记录一个特定的步骤，如果它花费的时间超过了它在总允许时间中的份额
		trace := utiltrace.New("Handler", // Handler 处理者

			// logrus.StandardLogger() 是 Logrus 库中的一个函数，它返回一个预配置的 *logrus.Logger 实例，这个实例通常用于在 Go 程序中进行日志记录。
			logrusr.New(logrus.StandardLogger()),                    // New将返回一个新日志。从logrus创建的记录器。FieldLogger。
			utiltrace.Field{Key: "method", Value: c.Request.Method}, // {method "GET"}，
			utiltrace.Field{Key: "path", Value: c.Request.URL.Path}, // {path "/api/usrs/"}
		)

		// 在Go语言中，defer语句用于延迟（推迟）一个函数的执行，直到包含它的函数返回之前才会被执行
		// defer语句通常用于在函数返回之前执行一些清理操作，例如关闭文件、释放资源或打印日志等。
		defer trace.LogIfLong(100 * time.Millisecond) // 100 * time.Millisecond 100毫秒

		// 把trace 存储到Context.keys
		common.SetTrace(c, trace) // Context.Keys是专门用于每个请求context的键/值对。

		c.Next()
	}
}
