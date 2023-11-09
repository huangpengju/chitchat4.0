package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"chitchat4.0/pkg/utils/request"
	"chitchat4.0/pkg/utils/set"
)

// 常量
const (
	All = "*"
)

// Scope 表示范围，是自定义 string 类型
type Scope string

const (
	ClusterScope   Scope = "cluster"   // 串范围
	NamespaceScope Scope = "namespace" // 命名空间范围
)

// 角色 结构体
type Role struct {
	ID        uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string `json:"name" gorm:"size:100;not null;unique"`
	Scope     Scope  `json:"scope" gorm:"size:100"`      // Scope 表示范围，string类型
	Namespace string `json:"namespace"  gorm:"size:100"` // 表示命名空间
	Rules     Rules  `json:"rules" gorm:"type:json"`     // Rules 表示规则集合，是切片类型
}

// Operation 表示操作，是自定义 string 类型
type Operation string

// 常量
const (
	AllOperation  Operation = "*"    // 所有操作
	EditOperation Operation = "edit" // 编辑操作
	ViewOperation Operation = "view" // 查看操作
)

// 全局变量
var (
	// 设置编辑操作（创建、删除、更新、补丁，获取，列表）
	EditOperationSet = set.NewString(request.CreateOperation, request.DeleteOperation, request.UpdateOperation, request.PatchOperation, request.GetOperation, request.ListOperation)
	// 设置查看操作（获取、列表）
	ViewOperationSet = set.NewString(request.GetOperation, request.ListOperation)
)

// Contain 控制
func (op Operation) Contain(verb string) bool {
	// 判断 op
	switch op {
	case AllOperation: // 所有操作
		return true
	case EditOperation: // 编辑操作
		return EditOperationSet.Has(verb) // 判断 verb 是否属于创建、删除、更新、补丁，获取，列表之一
	case ViewOperation: // 查看操作
		return ViewOperationSet.Has(verb) // 判断 verb 是否属于获取，列表之一
	default:
		return string(op) == verb // 判断 verb 是否与 string(op) 相等，string(op)用于把op转换为字符串类型
	}
}

// Rule 规则结构体
type Rule struct {
	Resource  string    `json:"resource"`  // 资源
	Operation Operation `json:"operation"` // 操作
}

// Rules 表示规则集合： Rule 切片
type Rules []Rule

func (r *Rules) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to numarshal JSONB value:%v", value)
	}

	result := Rules{}
	err := json.Unmarshal(bytes, &result)
	*r = result
	return err
}

func (r Rules) Value() (driver.Value, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

const (
	ResourceKind = "resource" // 资源种类
	MenuKind     = "menu"     // 菜单种类
)

const (
	ContainerResource = "containers" // 容器资源
	PostResource      = "posts"      // post资源
	UserResource      = "users"      // user资源
	GroupResource     = "groups"     // 组资源
	RoleResource      = "roles"      // Role角色资源
	AuthResource      = "auth"       // 授权资源
	NamespaceResource = "namespaces" // 命名空间资源
)

// Resource 资源结构体
type Resource struct {
	ID    uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name  string `json:"name" gorm:"size:256;not null;unique"`
	Scope Scope  `json:"scope"`
	Kind  string `json:"kind"`
}
