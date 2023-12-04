/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-09-27 14:50:58
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-22 13:55:41
 * @FilePath: \chitchat4.0\pkg\repository\interface.go
 * @Description: 接口
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package repository

import (
	"context"

	"chitchat4.0/pkg/model"
	"gorm.io/gorm/clause"
)

type Repository interface {
	User() UserRepository   // 实现
	Group() GroupRepository //
	RBAC() RBACRepository   //
	Close() error           // -

	// Tag() TagRepository
	// HotSearch() HotSearchRepository

	Ping(ctx context.Context) error

	Init() error // -

	Migrant
}

type Migrant interface {
	Migrate() error
}

// User 用户接口13
type UserRepository interface {
	GetUserByID(uint) (*model.User, error)                        // 实现通过id获取user
	GetUserByAuthID(authType, authID string) (*model.User, error) // 实现通过授权类型和授权ID获取userId，进一步通过userId获取user
	GetUserByName(string) (*model.User, error)                    // 实现通过name获取user
	List() (model.Users, error)                                   // 获取user列表
	Create(*model.User) (*model.User, error)                      // 创建user
	Update(*model.User) (*model.User, error)                      // 修改user
	Delete(*model.User) error                                     // 删除user

	GetGroups(*model.User) ([]model.Group, error)     // 获取user的全部group
	AddRole(role *model.Role, user *model.User) error // 给user添加role
	DelRole(role *model.Role, user *model.User) error // 删除user的role

	AddAuthInfo(authInfo *model.AuthInfo) error // 添加授权信息
	DelAuthInfo(authInfo *model.AuthInfo) error // 删除授权信息

	Migrate() error // 自动迁移
}

// 分组14-14
type GroupRepository interface {
	GetGroupByID(uint) (*model.Group, error)     // 实现通过id获取group
	GetGroupByName(string) (*model.Group, error) // 实现通过name获取group

	List() ([]model.Group, error)                           // 获取group列表
	Create(*model.User, *model.Group) (*model.Group, error) // 创建group

	Update(*model.Group) (*model.Group, error)          // 修改group
	Delete(uint) error                                  // 删除group
	GetUsers(*model.Group) (model.Users, error)         // 获取group下的全部user
	AddUser(user *model.User, group *model.Group) error // 给group添加user
	DelUser(user *model.User, group *model.Group) error // 删除group下的user
	AddRole(role *model.Role, group *model.Group) error // 给group添加role
	DelRole(role *model.Role, group *model.Group) error // 删除group对应的role

	RoleBinding(role *model.Role, group *model.Group) error // 创建默认group时，绑定role
	Migrate() error

	CreateGroups(groups []model.Group, conds ...clause.Expression) error // -
}

// Tag 标签接口
type TagRepository interface {
	List() ([]model.Tag, error)
	Create(*model.User, *model.Tag) (*model.Tag, error)
	Migrate() error
}

// HotSearchRepository 热搜列表仓库接口
type HotSearchRepository interface {
	List() ([]model.HotSearch, error)
	Create(*model.Tag, *model.HotSearch) (*model.HotSearch, error)
	Migrate() error
}

// 12-7
type RBACRepository interface {
	List() ([]model.Role, error)                  // 获取role列表
	ListResources() ([]model.Resource, error)     // resource 列表
	Create(role *model.Role) (*model.Role, error) // 创建role
	GetRoleByID(id int) (*model.Role, error)      // 通过id获取role
	Update(role *model.Role) (*model.Role, error) // 修改role
	Delete(id uint) error                         // 删除role
	Migrate() error                               // 自动迁移

	CreateResource(resource *model.Resource) (*model.Resource, error)
	CreateResources(resource []model.Resource, conds ...clause.Expression) error
	GetResource(id int) (*model.Resource, error)
	GetRoleByName(name string) (*model.Role, error)
	DeleteResource(id uint) error
}
