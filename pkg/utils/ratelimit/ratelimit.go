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
	keyFunc            func(*gin.Context) string         // 匿名函数，参数是 c *gin.Context包含了大量的关于请求和响应的信息 ,返回值类型 string
	cache              *lru.Cache[string, *rate.Limiter] // LRU 最近最少使用（固定大小的缓存）
	rateLimiterFactory func() *rate.Limiter              // 限制器控制事件允许发生的频率。
}

// NewRateLimiter 创建一个新的速率限制器
func NewRateLimiter(conf *LimitConfig) (*RateLimiter, error) {
	if conf == nil {
		return nil, errors.New("无效的 config")
	}
	// 使用 Validate 验证 conf 中指定字段的值，
	// 判断 QPS、Burst 的值不能为0
	// QPS 不能大于B urtst，
	// CacheSize 不能为0
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	// c *gin.Context 参数包含了大量的关于请求和响应的信息，可以用于获取客户端传递过来的请求参数、设置响应头等操作。
	// keyFunc 是函数类型
	var keyFunc func(*gin.Context) string

	// witch 判断 conf.LimitType 是 server 还是 ip
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

	// LRU算法即最近最少使用算法，它可以通过记录缓存内每个元素的使用情况来淘汰最近最少使用的元素，以达到缓存利用效率的最大化。）
	// 创建一个新的 LRU，同时指定大小（固定大小的缓存）
	// 参数：rate.Limiter 是一种限流工具（限制器）
	c, err := lru.New[string, *rate.Limiter](conf.CacheSize)
	if err != nil {
		return nil, err
	}

	// rateLimiterFactory 限制器控制事件允许发生的频率(QPS的最大速率，Burst个token的爆发)。
	rateLimiterFactory := func() *rate.Limiter {
		// NewLimiter 返回一个新的Limiter，（创建了一个rate.Limiter对象）
		// NewLimiter(r Limit, b int) *Limiter
		// NewLimiter返回一个新的Limiter，允许事件的最高速率为r，最多b个token的爆发。

		// 它允许事件的最高速率为conf.QPS（平均每秒处理的请求数量），
		// 并允许最多conf.Burst(在指定时间内允许的最大请求数量，用于处理突发流量)个令牌的爆发。
		// rate.Limit 定义了某些事件的最大频率。
		// 限制表示为每秒的事件数。
		// 0 Limit不允许任何事件。
		// rate.Limit(conf.QPS) 将 conf.QPS 转换为 Limit 类型
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
		return fmt.Errorf("LimitConfig 限制配置中 Burst and QPS 不能为0")
	}
	if c.QPS > c.Burst {
		return fmt.Errorf("LimitConfig中QPS（平均每秒处理的请求数量）(%d) 必须小于 Burst(最大请求数量)(%d)", c.QPS, c.Burst)
	}
	if c.CacheSize == 0 {
		c.CacheSize = defaultCacheSize // defaultCacheSize 2048 默认缓存大小
	}
	return nil
}

// Accept 接受
func (rl *RateLimiter) Accept(c *gin.Context) error {
	key := rl.keyFunc(c)
	limiter := rl.get(key)

	// Allow 允许报告事件现在是否可能发生。。
	if !limiter.Allow() {
		return fmt.Errorf("键 %v 在 %s 上达到极限", key, rl.limitType)
	}
	return nil
}

// get
func (rl *RateLimiter) get(key string) *rate.Limiter {
	value, found := rl.cache.Get(key)
	if !found {
		new := rl.rateLimiterFactory()
		// Add向缓存中添加一个值。如果发生了驱逐，则返回true。
		rl.cache.Add(key, new)
		return new
	}
	return value
}
