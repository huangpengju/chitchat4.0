/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-09-27 14:50:58
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-13 15:33:52
 * @FilePath: \chitchat4.0\pkg\repository\interface.go
 * @Description: 接口
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package repository

import (
	"context"

	"chitchat4.0/pkg/model"
)

type Repository interface {
	User() UserRepository   // 实现
	Group() GroupRepository //
	RBAC() RBACRepository
	Tag() TagRepository
	HotSearch() HotSearchRepository

	Ping(ctx context.Context) error

	Migrant
}

type Migrant interface {
	Migrate() error
}

// User 用户接口
type UserRepository interface {
	GetUserByID(uint) (*model.User, error)
	GetUserByName(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Migrate() error
}

// 分组
type GroupRepository interface {
	Migrate() error
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

type RBACRepository interface {
	List() ([]model.Role, error)
	ListResources() ([]model.Resource, error)
	Create(role *model.Role) (*model.Role, error)
	GetRoleByID(id int) (*model.Role, error)
	Update(role *model.Role) (*model.Role, error)
	Delete(id uint) error
	Migrate() error
}
