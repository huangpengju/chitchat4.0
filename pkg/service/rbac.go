/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-13 15:40:47
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-14 15:09:16
 * @FilePath: \chitchat4.0\pkg\service\rbac.go
 * @Description:
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package service

import (
	"strconv"

	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
	"chitchat4.0/pkg/utils/request"
)

/**
 * @description: rbacService 服务结构体
 * 封装 RBAC 仓库
 */
type rbacService struct {
	rbacRepository repository.RBACRepository
}

/**
 * @description: NewRBACService 创建一个RBAC服务，该服务包需要RBAC仓库
 * @param {repository.RBACRepository} rbacRepository
 * @return {*}
 */
func NewRBACService(rbacRepository repository.RBACRepository) RBACService {
	return &rbacService{
		rbacRepository: rbacRepository,
	}
}

func (rbac *rbacService) List() ([]model.Role, error) {
	return rbac.rbacRepository.List()
}

/**
 * @description: Create() 创建角色的服务
 * @param {*model.Role} role
 * @return {*}
 */
func (rbac *rbacService) Create(role *model.Role) (*model.Role, error) {
	return rbac.rbacRepository.Create(role)
}

/**
 * @description: Get() 获取角色的服务
 * @param {string} id
 * @return {*}
 */
func (rbac *rbacService) Get(id string) (*model.Role, error) {
	rid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return rbac.rbacRepository.GetRoleByID(rid)
}

func (rbac *rbacService) Update(id string, role *model.Role) (*model.Role, error) {
	rid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	role.ID = uint(rid)
	return rbac.rbacRepository.Update(role)
}

func (rbac *rbacService) Delete(id string) error {
	rid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return rbac.rbacRepository.Delete(uint(rid))
}

func (rbac *rbacService) ListResources() ([]model.Resource, error) {
	return rbac.rbacRepository.ListResources()
}

func (rbac *rbacService) ListOperations() ([]model.Operation, error) {
	return []model.Operation{
		model.AllOperation,      // 所有操作
		model.EditOperation,     // 编辑操作
		model.ViewOperation,     // 查看操作
		request.CreateOperation, // create创建操作
		request.PatchOperation,  // patch更新局部操作
		request.UpdateOperation, // update更新全部操作
		request.GetOperation,    // get获取单个（详情）操作
		request.ListOperation,   // list获取列表操作
		request.DeleteOperation, // delete删除操作
		"log",                   // 日志
		"exec",                  // 执行
		"proxy",                 // 代理
	}, nil
}
