package models

import (
	"fmt"
	"log"
	"os"
	"strings"

	"chitchat4.0/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func Init() {
	var err error
	str := []string{"host=", setting.DbHost, " port=", setting.DbPort, " user=", setting.DbUser, " password=", setting.DbPassword, " dbname=", setting.DbName, " sslmode=disable"}
	DSN := strings.Join(str, "")
	Db, err = gorm.Open(setting.DbType, DSN)
	if err != nil {
		log.Fatalf("初始化数据库连接失败:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("数据库连接成功")
	// 更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// 设置指定表的前缀
		return setting.DbTablePrefix + defaultTableName
	}
	// 不让表名加s
	Db.SingularTable(true)

	// 启用日志
	Db.LogMode(true)

	// 设置连接池
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	// 自动迁移
	migration()
}
