package middleware

import (
	"net/http"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/utils/ratelimit"
	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware 速度限制中间件
func RateLimitMiddleware(configs []ratelimit.LimitConfig) (gin.HandlerFunc, error) {
	var limiters []*ratelimit.RateLimiter // 定义限制器
	for i := range configs {
		limiter, err := ratelimit.NewRateLimiter(&configs[i])
		if err != nil {
			return nil, err
		}
		limiters = append(limiters, limiter)
	}
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
