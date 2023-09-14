package config

import (
	"os"

	"chitchat4.0/pkg/utils/ratelimit"
	"gopkg.in/yaml.v3"
)

// Config 应用的全部配置
type Config struct {
	Server ServerConfig `yaml:"server"` // 服务相关配置
	DB     DBConfig     `yaml:"db"`     // 数据库相关配置
}

// ServerConfig 服务配置
type ServerConfig struct {
	ENV                    string                  `yaml:"env"`                    // 项目运行环境
	Address                string                  `yaml:"address"`                // 主机地址
	Port                   int                     `yaml:"port"`                   // 端口
	GracefulShutdownPeriod int                     `yaml:"gracefulShutdownPeriod"` // 正常停机时间
	LimitConfig            []ratelimit.LimitConfig `yaml:"rateLimis"`              // 等级界限
	JWTSecret              string                  `yaml:"jwtSecret"`              // jsonWebToken
}

// DBConfig 数据库配置
type DBConfig struct {
	Host     string `yaml:"host"`     // 数据库主机
	Port     int    `yaml:"port"`     // 数据库端口
	Name     string `yaml:"name"`     // 库名称
	User     string `yaml:"user"`     // 用户
	Password string `yaml:"password"` // 密码
	Migrate  bool   `yaml:"migrate"`  // 迁移（是否自动迁移）
}

// Parse 根据传入的路径分析配置信息，返回应用的配置 Config 和 error
func Parse(appConfig string) (*Config, error) {
	config := &Config{} // 定义一个空的配置

	file, err := os.Open(appConfig) // 读取配置文件
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// NewDecoder 创建一个新的解码器
	// Decode 给 config 结构填充相应的数据
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
