package database

import (
	"fmt"

	"chitchat4.0/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgres 接收一个 *config.DBConfig 类型的参数。
// 参数是应用配置 Config 中的子配置 DB ，DB 具有 PostgreSQL 数据库 Host、port 等信息。
// 作用初始化数据库连接
// 返回一个PostgreSQL数据库的 *gorm.DB 结构体 和 error
func NewPostgres(conf *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
