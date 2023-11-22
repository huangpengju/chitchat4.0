/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-09-28 16:29:32
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-21 16:18:49
 * @FilePath: \chitchat4.0\pkg\service\interface.go
 * @Description:
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package service

import "chitchat4.0/pkg/model"

type UserService interface {
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Get(string) (*model.User, error)
	Update(string, *model.User) (*model.User, error)
	// Delete(string) error
	Validate(*model.User) error
	Auth(*model.AuthUser) (*model.User, error)
	Default(*model.User)
}

type GroupService interface {
	List() ([]model.Group, error)
	Create(*model.User, *model.Group) (*model.Group, error)
	Get(string) (*model.Group, error)
	Update(string, *model.Group) (*model.Group, error)
	Delete(string) error
	GetUsers(string) (model.Users, error)
	AddUser(user *model.User, gid string) error
	DelUser(gid, uid string) error
	AddRole(id, rid string) error
}

type TagService interface {
	List() ([]model.Tag, error)
	Create(*model.User, *model.Tag) (*model.Tag, error)
	// Get(string) (*model.Tag, error)
	// Update(string, *model.Tag) (*model.Tag, error)
	// Delete(string) error
	// Validate(*model.Tag) error
}

type HotSearchService interface {
	List() ([]model.HotSearch, error)
	Create(*model.Tag, *model.HotSearch) (*model.HotSearch, error)
	// Get(string) (*model.HotSearch, error)
	// Update(string, *model.HotSearch) (*model.HotSearch, error)
	// Delete(string) error
	// Validate(*model.HotSearch) error
}

/**
 * @description: RBACService 基于角色访问控制的服务
 *
 */
type RBACService interface {
	List() ([]model.Role, error)
	Create(role *model.Role) (*model.Role, error)
	Get(id string) (*model.Role, error)
	Update(id string, role *model.Role) (*model.Role, error)
	Delete(id string) error
	ListResources() ([]model.Resource, error)
	ListOperations() ([]model.Operation, error)
}
