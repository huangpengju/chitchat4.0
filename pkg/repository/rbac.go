/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-06 15:20:06
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-13 16:46:05
 * @FilePath: \chitchat4.0\pkg\repository\rbac.go
 * @Description: Role-Based Access Control 基于角色的访问控制
 */
package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

// rbac 数据库仓库
type rbacRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

/**
 * @description: newRBACRepository 返回一个RBAC仓库
 * @param {*gorm.DB} db
 * @param {*database.RedisDB} rdb
 * @return {*}
 */
func newRBACRepository(db *gorm.DB, rdb *database.RedisDB) RBACRepository {
	return &rbacRepository{
		db:  db,
		rdb: rdb,
	}
}

/**
 * @description: Create() 实现创建角色
 * @param {*model.Role} role
 * @return {*}
 */
func (rbac *rbacRepository) Create(role *model.Role) (*model.Role, error) {
	err := rbac.db.Create(role).Error
	return role, err
}

/**
 * @description: Migrate 实现数据库model迁移
 * @return {*}
 */
func (rbac *rbacRepository) Migrate() error {
	return rbac.db.AutoMigrate(&model.Role{}, &model.Resource{})
}
