/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-14 15:30:32
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-22 14:03:15
 * @FilePath: \chitchat4.0\pkg\service\group.go
 * @Description:
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package service

import (
	"fmt"
	"strconv"

	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
)

type groupService struct {
	userRepository  repository.UserRepository
	groupRepository repository.GroupRepository
	rbacRepository  repository.RBACRepository
}

/**
 * @description: NewGroupService() 返回一个 group服务
 * @param {repository.GroupRepository} groupRepository
 * @param {repository.UserRepository} userRepository
 * @return {*}
 */
func NewGroupService(groupRepository repository.GroupRepository, userRepository repository.UserRepository, rbacRepository repository.RBACRepository) GroupService {
	return &groupService{
		groupRepository: groupRepository,
		userRepository:  userRepository,
		rbacRepository:  rbacRepository,
	}
}

/**
 * @description: Create() 创建group服务，创建成功后与角色进行绑定
 * @param {*model.User} user
 * @param {*model.Group} group
 * @return {*}
 */
func (g *groupService) Create(user *model.User, group *model.Group) (*model.Group, error) {
	group, err := g.groupRepository.Create(user, group)
	if err != nil {
		return nil, err
	}

	if err := g.createDefaultRoles(group); err != nil {
		return nil, err
	}
	return group, nil
}

/**
 * @description: createDefaultRoles 指定（新建）三个角色，并与group绑定
 * @param {*model.Group} group
 * @return {*}
 */
func (g *groupService) createDefaultRoles(group *model.Group) error {
	roles := []model.Role{
		{
			Name:      fmt.Sprintf("ns-%s-%s", group.Name, "admin"),
			Scope:     model.NamespaceScope,
			Namespace: group.Name,
			Rules: []model.Rule{
				{
					Resource:  model.All,
					Operation: model.All,
				},
			},
		},
		{
			Name:      fmt.Sprintf("ns-%s-%s", group.Name, "edit"),
			Scope:     model.NamespaceScope,
			Namespace: group.Name,
			Rules: []model.Rule{
				{
					Resource:  model.All,
					Operation: model.EditOperation,
				},
			},
		},
		{
			Name:      fmt.Sprintf("ns-%s-%s", group.Name, "view"),
			Scope:     model.NamespaceScope,
			Namespace: group.Name,
			Rules: []model.Rule{
				{
					Resource:  model.All,
					Operation: model.ViewOperation,
				},
			},
		},
	}
	for i := range roles {
		if _, err := g.rbacRepository.Create(&roles[i]); err != nil {
			return err
		}
	}

	return g.groupRepository.RoleBinding(&roles[0], group)
}

/**
 * @description: Get() 通过id查询group服务
 * @param {string} id
 * @return {*}
 */
func (g *groupService) Get(id string) (*model.Group, error) {
	gid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return g.groupRepository.GetGroupByID(uint(gid))
}

/**
 * @description: List()查询所有group的服务
 * @return {*}
 */
func (g *groupService) List() ([]model.Group, error) {
	return g.groupRepository.List()
}

func (g *groupService) Update(id string, group *model.Group) (*model.Group, error) {
	gid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	group.ID = uint(gid)
	return g.groupRepository.Update(group)
}

func (g *groupService) Delete(id string) error {
	gid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return g.groupRepository.Delete(uint(gid))
}

func (g *groupService) GetUsers(id string) (model.Users, error) {
	gid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return g.groupRepository.GetUsers(&model.Group{ID: uint(gid)})
}

func (g *groupService) AddUser(user *model.User, id string) error {
	// var err error
	if user.ID == 0 {
		return fmt.Errorf("%v", "Group AddUser:invaild user info")
	}
	gid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return g.groupRepository.AddUser(user, &model.Group{ID: uint(gid)})
}

func (g *groupService) DelUser(gid, uid string) error {
	groupId, err := strconv.Atoi(gid)
	if err != nil {
		return err
	}

	userId, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}

	user := new(model.User)
	user.ID = uint(userId)
	if user.ID == 0 {
		return fmt.Errorf("%v", "Group DelUser:invaild user info")
	}

	return g.groupRepository.DelUser(user, &model.Group{ID: uint(groupId)})
}

func (g *groupService) AddRole(id, rid string) error {
	gid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	roleId, err := strconv.Atoi(rid)
	if err != nil {
		return err
	}

	return g.groupRepository.AddRole(&model.Role{ID: uint(roleId)}, &model.Group{ID: uint(gid)})
}

func (g *groupService) DelRole(id, rid string) error {
	gid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	roleId, err := strconv.Atoi(rid)
	if err != nil {
		return err
	}

	return g.groupRepository.DelRole(&model.Role{ID: uint(roleId)}, &model.Group{ID: uint(gid)})
}
