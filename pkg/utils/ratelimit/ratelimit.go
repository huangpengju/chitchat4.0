package ratelimit

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru/v2"
	"golang.org/x/time/rate"
)

const defaultCacheSize = 2048 // 默认缓存大小
type LimitType string

const (
	ServerLimitType LimitType = "server" //限制类型server
	IPLimitType     LimitType = "ip"     //限制类型ip
)

// LimitConfig 限制配置(存放配置文件中的数据)
type LimitConfig struct {
	LimitType LimitType `yaml:"limitType"` // 速率限制的类型
	Burst     int       `yaml:"burst"`     // 在指定时间内允许的最大请求数量，用于处理突发流量
	QPS       int       `yaml:"qps"`       // 每秒请求数（平均每秒处理的请求数量）
	CacheSize int       `yaml:"cacheSize"` // 用于存储限制的缓存大小（用于优化性能，降低查询频率）
}

// RateLimiter 速度限制结构
type RateLimiter struct {
	limitType          LimitType                         // 限制的类型 server | ip
	keyFunc            func(*gin.Context) string         // // 匿名函数，参数是 c *gin.Context包含了大量的关于请求和响应的信息 ,返回值类型 string
	cache              *lru.Cache[string, *rate.Limiter] // LRU 缓存
	rateLimiterFactory func() *rate.Limiter              // 限制器控制事件允许发生的频率。
}

func NewRateLimiter(conf *LimitConfig) (*RateLimiter, error) {
	if conf == nil {
		return nil, errors.New("无效的 config")
	}
	// 使用 Validate 验证 conf 中指定字段的值
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	// c *gin.Context 参数包含了大量的关于请求和响应的信息，可以用于获取客户端传递过来的请求参数、设置响应头等操作。
	var keyFunc func(*gin.Context) string
	// conf.LimitType 可能是 server 也可能是 ip
	// 常量 ServerLimitType = server
	// 常量 IPLimitType = ip
	switch conf.LimitType {
	case ServerLimitType:
		keyFunc = func(c *gin.Context) string {
			return ""
		}
	case IPLimitType:
		keyFunc = func(c *gin.Context) string {
			return c.ClientIP() // 返回客户端IP
		}
	default:
		return nil, fmt.Errorf("不确定限制的类型（server或ip） %s", conf.LimitType)
	}
	// fmt.Printf("keyFun=%T\n", keyFunc) // keyFun=func(*gin.Context) string

	// 创建给定大小的LRU
	c, err := lru.New[string, *rate.Limiter](conf.CacheSize)
	if err != nil {
		return nil, err
	}

	rateLimiterFactory := func() *rate.Limiter {
		// NewLimiter返回一个新的Limiter，
		// 它允许事件的最高速率为conf.QPS（平均每秒处理的请求数量），
		// 并允许最多conf.Burst(在指定时间内允许的最大请求数量，用于处理突发流量)个令牌的爆发。
		return rate.NewLimiter(rate.Limit(conf.QPS), conf.Burst)
	}

	return &RateLimiter{
		limitType:          conf.LimitType,     // 限制的类型
		keyFunc:            keyFunc,            // 匿名函数，参数是 c *gin.Context ,返回值类型 string
		cache:              c,                  // 创建给定大小的LRU
		rateLimiterFactory: rateLimiterFactory, // 限制器控制事件允许发生的频率。
	}, nil
}

// Validate 用于验证 Server 中的 rateLimits（数据在LimitConfig结构体中），
// PQS和Burst 不能为0；
// PQS必须小于Burst；
// CacheSize为0时，使用 defaultCacheSize 2048。
func (c *LimitConfig) Validate() error {
	if c.QPS == 0 || c.Burst == 0 {
		return fmt.Errorf("LimitConfig 限制配置中 Burst and QPS 不能为空")
	}
	if c.QPS > c.Burst {
		return fmt.Errorf("LimitConfig中QPS（平均每秒处理的请求数量）(%d) 必须小于 Burst(最大请求数量)(%d)", c.QPS, c.Burst)
	}
	if c.CacheSize == 0 {
		c.CacheSize = defaultCacheSize // defaultCacheSize 2048 默认缓存大小
	}
	return nil
}

func (rl *RateLimiter) Accept(c *gin.Context) error {
	key := rl.keyFunc(c)
	limiter := rl.get(key)

	if !limiter.Allow() {
		return fmt.Errorf("键 %v 在 %s 上达到极限", key, rl.limitType)
	}
	return nil
}
func (rl *RateLimiter) get(key string) *rate.Limiter {
	value, found := rl.cache.Get(key)
	if !found {
		new := rl.rateLimiterFactory()
		rl.cache.Add(key, new)
		return new
	}
	return value
}
