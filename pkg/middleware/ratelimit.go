package middleware

import (
	"net/http"
	"os"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/utils/ratelimit"
	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware 速度限制中间件
func RateLimitMiddleware(configs []ratelimit.LimitConfig) (gin.HandlerFunc, error) {

	var limiters []*ratelimit.RateLimiter // 定义限制控制器切片（存放server类型和iP类型的限制）

	// configs= [{server 500 100 1} {ip 50 10 2048}]
	for i := range configs {
		// 生成限制控制器
		limiter, err := ratelimit.NewRateLimiter(&configs[i])
		if err != nil {
			return nil, err
		}
		limiters = append(limiters, limiter)
	}
	os.Exit(0)

	return func(c *gin.Context) {
		for _, limiter := range limiters {
			if err := limiter.Accept(c); err != nil {
				common.ResponseFailed(c, http.StatusTooManyRequests, err)
				return
			}
		}
		c.Next()
	}, nil
}
