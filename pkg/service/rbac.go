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
		model.AllOperation,
		model.EditOperation,
		model.ViewOperation,
		request.CreateOperation,
		request.PatchOperation,
		request.UpdateOperation,
		request.GetOperation,
		request.ListOperation,
		request.DeleteOperation,
		"log",
		"exec",
		"proxy",
	}, nil
}
