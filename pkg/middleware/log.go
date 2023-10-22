package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	hostname, _ = os.Hostname() // Hostname返回内核报告的主机名。
)

// LogMiddleware 日志中间件
func LogMiddleware(logger *logrus.Logger, pathPrefix ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path // 访问的URL(用于客户端请求)路径。 比如：/api/v1/users/7
		logged := len(pathPrefix) == 0
		for _, prefix := range pathPrefix {
			if strings.HasPrefix(path, prefix) {
				logged = true
				break
			}
		}
		if !logged {
			return
		}

		start := time.Now() // 返回当前本地时间

		defer func() {
			latency := time.Since(start)             // 传输时间
			statusCode := c.Writer.Status()          // Status返回当前请求的HTTP响应状态码。
			clientIP := c.ClientIP()                 // ClientIP实现了一个最佳努力算法来返回真实的客户端IP
			clientUserAgent := c.Request.UserAgent() // 如果在请求中发送，UserAgent返回客户端的用户代理。

			entry := logger.WithFields(logrus.Fields{
				"hostname":   hostname,         // 主机名
				"path":       path,             // 访问的URL(用于客户端请求)路径。 比如：/api/v1/users/7
				"method":     c.Request.Method, // Method指定HTTP方法(GET、POST、PUT等)。对于客户端请求，空字符串表示GET。
				"statusCode": statusCode,       // Status返回当前请求的HTTP响应状态码。
				"clientIP":   clientIP,         // ClientIP实现了一个最佳努力算法来返回真实的客户端IP
				"userAgent":  clientUserAgent,  // 如果在请求中发送，UserAgent返回客户端的用户代理。
			})

			if len(c.Errors) > 0 {
				entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				msg := fmt.Sprintf("[%s %s] %d %v", c.Request.Method, c.Request.URL, statusCode, latency)
				if statusCode >= http.StatusInternalServerError {
					entry.Error(msg)
				} else if statusCode >= http.StatusBadRequest {
					entry.Warn(msg)
				} else {
					entry.Info(msg)
				}
			}
		}()

		c.Next()
	}
}
