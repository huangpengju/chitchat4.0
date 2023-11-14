package service

import (
	"fmt"

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
func NewGroupService(groupRepository repository.GroupRepository, userRepository repository.UserRepository) GroupService {
	return &groupService{
		groupRepository: groupRepository,
		userRepository:  userRepository,
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
