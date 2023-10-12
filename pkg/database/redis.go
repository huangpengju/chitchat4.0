package database

import (
	"context"
	"errors"
	"fmt"

	"chitchat4.0/pkg/config"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	RedisDisableError = errors.New("redis disable")
)

type RedisDB struct {
	enable bool
	*redis.Client
}

// NewRedisClient 接收一个 *config.RedisConfig 类型的参数。
// 参数是应用配置 Config 中的子配置 Redis ，Redis 具有 Redis 数据库 Host、port 等。
// 作用初始化 Redis 客户端
// 返回一个Redis客户端的 *RedisDB 结构体 和 error
func NewRedisClient(conf *config.RedisConfig) (*RedisDB, error) {
	if !conf.Enable {
		logrus.Info("redis 禁用")
		return &RedisDB{}, nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       0, // 数据库从0开始
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return &RedisDB{
		enable: true,
		Client: rdb,
	}, nil
}

// HSet
// 参数 key：users:id，
// 参数 field：id，
// 参数 val：user
func (rdb *RedisDB) HSet(key, field string, val interface{}) error {
	if !rdb.enable {
		return nil
	}
	// 设置 key-value 键值对
	return rdb.Client.HSet(context.Background(), key, field, val).Err()
}

// HGet
func (rdb *RedisDB) HGet(key, field string, obj interface{}) error {
	if !rdb.enable {
		return RedisDisableError
	}

	return rdb.Client.HGet(context.Background(), key, field).Scan(obj)
}
