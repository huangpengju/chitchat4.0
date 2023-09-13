package models

import (
	"chitchat4.0/pkg/setting"
)

func migration() {
	// 创建表的时候使用 Set 设置一些额外的表属性，比如：数据库引擎、字符集
	if setting.DbType == "mysql" {
		Db.Set("gorm:table_options", "ENGINE=InnoDB charset=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(&Tag{}).AutoMigrate(&HotList{})
		// 给 hot_list 表，设置外键
		Db.Model(&HotList{}).AddForeignKey("tag_id", "cc_tag(id)", "CASCADE", "CASCADE")
	} else {
		Db.AutoMigrate(&Tag{}).AutoMigrate(&HotList{})
	}

}
