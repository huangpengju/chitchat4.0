/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-06 15:20:06
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-22 14:00:51
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
	"gorm.io/gorm/clause"
)

var (
	groupUpdateFields = []string{"Name", "Describe", "Roles", "UpdaterId"}
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
 * @description: Create()实现 group的创建
 * @param {*model.User} user
 * @param {*model.Group} group
 * @return {*}
 */
func (g *groupRepository) Create(user *model.User, group *model.Group) (*model.Group, error) {
	group.CreatorId = user.ID
	group.Users = []model.User{*user}
	err := g.db.Create(group).Error
	return group, err
}

/**
 * @description: RoleBinding()实现group与role的绑定
 * @param {*model.Role} role
 * @param {*model.Group} group
 * @return {*}
 */
func (g *groupRepository) RoleBinding(role *model.Role, group *model.Group) error {
	return g.db.Model(group).Association("Roles").Append(role)
}

/**
 * @description: 对 Group 模型进行数据库迁移
 * @return {error}
 */
func (g *groupRepository) Migrate() error {
	return g.db.AutoMigrate(&model.Group{})
}

/**
 * @description: GetGroupByID()实现id查询group，以及与之联系的User和Role
 * @param {uint} id
 * @return {*}
 */
func (g *groupRepository) GetGroupByID(id uint) (*model.Group, error) {
	group := new(model.Group)
	if err := g.db.Preload("Users").Preload("Roles").First(group, id).Error; err != nil {
		return nil, err
	}
	return group, nil
}

/**
 * @description: List() 实现获取Group列表
 * @return {*}
 */
func (g *groupRepository) List() ([]model.Group, error) {
	groups := make([]model.Group, 0)
	if err := g.db.Order("name").Preload("Roles").Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *groupRepository) Update(group *model.Group) (*model.Group, error) {
	err := g.db.Model(group).Select(groupUpdateFields).Updates(group).Error
	return group, err
}

func (g *groupRepository) Delete(id uint) error {
	return g.db.Delete(&model.Group{}, id).Error
}

func (g *groupRepository) GetUsers(group *model.Group) (model.Users, error) {
	users := make(model.Users, 0)
	err := g.db.Model(group).Association(model.UserAssociation).Find(&users)
	return users, err
}

func (g *groupRepository) AddUser(user *model.User, group *model.Group) error {
	return g.db.Model(group).Association(model.UserAssociation).Append(user)
}

func (g *groupRepository) DelUser(user *model.User, group *model.Group) error {
	return g.db.Model(group).Association(model.UserAssociation).Delete(user)
}

func (g *groupRepository) AddRole(role *model.Role, group *model.Group) error {
	var err error
	if group.ID == 0 {
		group, err = g.GetGroupByName(group.Name)
	}
	if err != nil {
		return err
	}
	return g.db.Model(group).Association("Roles").Append(role)
}

func (g *groupRepository) GetGroupByName(name string) (*model.Group, error) {
	group := new(model.Group)
	if err := g.db.Preload("Users").Preload("Roles").Where("name = ?", name).First(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (g *groupRepository) DelRole(role *model.Role, group *model.Group) error {
	var err error
	if group.ID == 0 {
		group, err = g.GetGroupByName(group.Name)
	}
	if err != nil {
		return err
	}
	return g.db.Model(group).Association("Roles").Delete(role)
}

// 在repository仓库Init时调用
func (g *groupRepository) CreateGroups(groups []model.Group, conds ...clause.Expression) error {
	return g.db.Clauses(conds...).Create(groups).Error
}
