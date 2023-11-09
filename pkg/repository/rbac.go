/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-06 15:20:06
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-09 16:56:11
 * @FilePath: \chitchat4.0\pkg\repository\rbac.go
 * @Description: Role-Based Access Control 基于角色的访问控制 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

type rbacRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

// newRBACRepository 
//  @param db 
//  @param rdb 
//  @return RBACRepository 
func newRBACRepository(db *gorm.DB, rdb *database.RedisDB) RBACRepository {
	return &rbacRepository{
		db:  db,
		rdb: rdb,
	}
}

func (rbac *rbacRepository) Migrate() error {
	return rbac.db.AutoMigrate(&model.Role{}, &model.Resource{})
}
