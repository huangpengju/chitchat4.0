/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-06 15:20:06
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-14 16:32:51
 * @FilePath: \chitchat4.0\pkg\model\group.go
 * @Description: group 分组的 model
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package model

const (
	RootGroup            = "root"                   // 根用户组
	AuthenticatedGroup   = "system:authenticated"   // 系统已认证
	UnAuthenticatedGroup = "system:unauthenticated" // 系统未认证
	SystemGroup          = "system"                 // 系统
	CustomGroup          = "custom"                 // 自定义
)

type Group struct {
	ID        uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string `json:"name" gorm:"size:100;not null;unique"`
	Kind      string `json:"kind" gorm:"size:100"`                // 种类
	Describe  string `json:"describe" gorm:"size:1024;"`          // 描述
	CreatorId uint   `json:"creatorId"`                           // 创作者Id
	UpdaterId uint   `json:"updaterId"`                           // 更新 Id
	Users     []User `json:"users" gorm:"many2many:user_groups;"` // 用户集合
	Roles     []Role `json:"roles" gorm:"many2many:group_roles;"` // 角色组集合

	BaseModel
}

// CreatedGroup 创建分组结构体
type CreatedGroup struct {
	Name      string `json:"name"`
	Describe  string `json:"describe"`  // 描述
	CreatorId uint   `json:"creatorId"` // 创建者ID
}

/**
 * @description: 接收者CreatedGroup，返回Group分组Name、描述、创建者Id
 * @param {uint} uid
 * @return {Group}
 */
func (g *CreatedGroup) GetGroup(uid uint) *Group {
	return &Group{
		Name:      g.Name,
		Describe:  g.Describe,
		CreatorId: g.CreatorId, // uid
	}
}

// UpdatedGroup 修改分组结构体
type UpdatedGroup struct {
	Name      string `json:"name"`
	Describe  string `json:"describe"`
	UpdaterId uint   `json:"updaterId"`
}

/**
 * @description: 接收者UpdatedGroup，返回Group分组Name、描述、更新Id
 * @param {uint} uid
 * @return {*}
 */
func (g *UpdatedGroup) GetGroup(uid uint) *Group {
	return &Group{
		Name:      g.Name,
		Describe:  g.Describe,
		UpdaterId: g.UpdaterId,
	}
}
