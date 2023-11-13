package service

import (
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
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

/**
 * @description: Create() 创建角色的服务
 * @param {*model.Role} role
 * @return {*}
 */
func (rbac *rbacService) Create(role *model.Role) (*model.Role, error) {
	return rbac.rbacRepository.Create(role)
}
