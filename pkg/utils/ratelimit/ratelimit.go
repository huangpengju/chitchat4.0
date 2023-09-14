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

// LimitConfig 限制配置
type LimitConfig struct {
	LimitType LimitType `yaml:"limiType"`  // 速率限制的类型
	Burst     int       `yaml:"burst"`     // 在指定时间内允许的最大请求数量，用于处理突发流量
	QPS       int       `yaml:"qps"`       // 每秒请求数（平均每秒处理的请求数量）
	CacheSize int       `yaml:"cacheSize"` // 用于存储限制的缓存大小（用于优化性能，降低查询频率）
}

// RateLimiter 速度限制结构
type RateLimiter struct {
	limitType          LimitType
	keyFunc            func(*gin.Context) string
	cache              *lru.Cache[string, *rate.Limiter] // LRU 缓存
	rateLimiterFactory func() *rate.Limiter
}

func NewRateLimiter(conf *LimitConfig) (*RateLimiter, error) {
	if conf == nil {
		return nil, errors.New("无效的 config")
	}
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	var keyFunc func(*gin.Context) string
	switch conf.LimitType {
	case ServerLimitType:
		keyFunc = func(c *gin.Context) string {
			return ""
		}
	case IPLimitType:
		keyFunc = func(c *gin.Context) string {
			return c.ClientIP()
		}
	default:
		return nil, fmt.Errorf("不知道限制的类型 %s", conf.LimitType)
	}
	c, err := lru.New[string, *rate.Limiter](conf.CacheSize)
	if err != nil {
		return nil, err
	}
	rateLimiterFactory := func() *rate.Limiter {
		return rate.NewLimiter(rate.Limit(conf.QPS), conf.Burst)
	}
	return &RateLimiter{
		limitType:          conf.LimitType,
		keyFunc:            keyFunc,
		cache:              c,
		rateLimiterFactory: rateLimiterFactory,
	}, nil
}

// Validate 用于验证 LimitConfig 中的PQS和Burst
func (c *LimitConfig) Validate() error {
	if c.QPS == 0 || c.Burst == 0 {
		return fmt.Errorf("LimitConfig 限制配置中 Burst and QPS 不能为空")
	}
	if c.QPS > c.Burst {
		return fmt.Errorf("LimitConfig中QPS(%d) 必须小于 Burst(%d)", c.QPS, c.Burst)
	}
	if c.CacheSize == 0 {
		c.CacheSize = defaultCacheSize
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
