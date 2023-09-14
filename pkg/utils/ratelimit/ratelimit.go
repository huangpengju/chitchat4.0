package ratelimit

type LimitType string

type LimitConfig struct {
	LimitType LimitType `yaml:"limiType"`
	Burst     int       `yaml:"burst"`
	Qps       int       `yaml:"qps"`
	CacheSize int       `yaml:"cacheSize"` // 缓存大小
}
