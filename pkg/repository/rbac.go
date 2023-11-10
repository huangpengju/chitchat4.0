/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-06 15:20:06
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-10 14:11:45
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

// newRBACRepository
//
//	@param db
//	@param rdb
//	@return RBACRepository
func newRBACRepository(db *gorm.DB, rdb *database.RedisDB) RBACRepository {
	return &rbacRepository{
		db:  db,
		rdb: rdb,
	}
}

func (rbac *rbacRepository) Migrate() error {
	return rbac.db.AutoMigrate(&model.Role{}, &model.Resource{})
}
