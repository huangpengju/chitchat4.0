/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-06 15:20:06
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-10 15:27:34
 * @FilePath: \chitchat4.0\pkg\repository\group.go
 * @Description: group 分组仓库，实现接口
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

// group 数据库仓库
type groupRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

/**
 * @description: 返回一个 group 数据库仓库
 * @param {*gorm.DB} db
 * @param {*database.RedisDB} rdb
 * @return {*}
 */
func newGroupRepository(db *gorm.DB, rdb *database.RedisDB) GroupRepository {
	return &groupRepository{
		db:  db,
		rdb: rdb,
	}
}

/**
 * @description: 对 Group 模型进行数据库迁移
 * @return {error}
 */
func (g *groupRepository) Migrate() error {
	return g.db.AutoMigrate(&model.Group{})
}
