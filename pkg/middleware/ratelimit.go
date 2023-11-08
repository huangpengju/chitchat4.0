/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-09-26 18:02:11
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-08 16:28:37
 * @FilePath: \chitchat4.0\pkg\middleware\ratelimit.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middleware

import (
	"net/http"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/utils/ratelimit"
	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware 速度限制中间件
// 函数RateLimitMiddleware接收一个配置数组configs，返回一个gin.HandlerFunc和错误
func RateLimitMiddleware(configs []ratelimit.LimitConfig) (gin.HandlerFunc, error) {

	var limiters []*ratelimit.RateLimiter // 定义限制控制器切片（存放server类型和iP类型的限制）

	// configs= [{server 500 100 1} {ip 50 10 2048}]
	for key := range configs {
		// 生成限制控制器
		limiter, err := ratelimit.NewRateLimiter(&configs[key])
		if err != nil {
			return nil, err
		}
		limiters = append(limiters, limiter)
	}

	return func(c *gin.Context) {
		// 遍历限制控制器，接受请求
		for _, limiter := range limiters {
			if err := limiter.Accept(c); err != nil {
				common.ResponseFailed(c, http.StatusTooManyRequests, err)
				return
			}
		}
		// 执行下一个中间件
		c.Next()
	}, nil
}
