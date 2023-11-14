package controller

import (
	"net/http"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/service"
	"github.com/gin-gonic/gin"
)

/**
 * @description: RBAC控制器
 * 封装 RBAC 服务
 */
type RBACController struct {
	rbacService service.RBACService
}

/**
 * @description: NewRbacController 返回一个接口类型，接口值是RBAC控制器
 * @param {service.RBACService} rbacService
 * @return {*}
 */
func NewRbacController(rbacService service.RBACService) Controller {
	return &RBACController{rbacService: rbacService}
}

// @Summary List rbac role
// @Description List rbac role | rbac 角色列表
// @Product json
// @Tags rbac
// @Security JWT
// @Success 200 {object} common.Response{data=[]model.Role}
// @Router /api/v1/roles [get]
func (rbac *RBACController) List(c *gin.Context) {
	roles, err := rbac.rbacService.List()
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, roles)
}

// @Summary Create rbac role
// @Description Create rbac role | 创建 rbac 的角色
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param role body model.Role true "rbac role info"
// @Success 200 {object} common.Response
// @Router /api/v1/roles [post]
func (rbac *RBACController) Create(c *gin.Context) {
	role := &model.Role{}
	if err := c.BindJSON(role); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	role, err := rbac.rbacService.Create(role) // 调用 RBAC 服务
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, role)

}

// @Summary Get role
// @Description Get role  | 获取一个 rbac 的角色
// @Produce json
// @Tags rbac
// @Security JWT
// @Param id path int true "role id"
// @Success 200 {object} common.Response{data=model.Role}
// @Router /api/v1/roles/{id} [get]
func (rbac *RBACController) Get(c *gin.Context) {

	role, err := rbac.rbacService.Get(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, role)
}

// @Summary Update rbac role
// @Description Update rbac role | rbac 修改角色
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param role body model.Role true "rbac role info"
// @Success 200 {object} common.Response{data=model.Role}
// @Param id path int true "role id"
// @Router /api/v1/roles/{id} [put]
func (rbac *RBACController) Update(c *gin.Context) {
	role := &model.Role{}
	if err := c.BindJSON(role); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	role, err := rbac.rbacService.Update(id, role)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, role)
}

// @Summary Delete role
// @Description Delete role | 删除角色
// @Produce json
// @Tags rbac
// @Security JWT
// @Param id path int true "role id"
// @Success 200 {object} common.Response
// @Router /api/v1/roles/{id} [delete]
func (rbac *RBACController) Delete(c *gin.Context) {
	if err := rbac.rbacService.Delete(c.Param("id")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, nil)
}

// @Summary List resources
// @Description List resources | 资源列表
// @Produce json
// @Tags rbac
// @Security JWT
// @Success 200 {object} common.Response{data=[]model.Resource}
// @Router /api/v1/resources [get]
func (rbac *RBACController) ListResources(c *gin.Context) {
	data, err := rbac.rbacService.ListResources()
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
	}
	common.ResponseSuccess(c, data)
}

// @Summary List operations
// @Description List operations | 操作列表
// @Produce json
// @Tags rbac
// @Security JWT
// @Success 200 {object} common.Response{data=[]model.Operation}
// @Router /api/v1/operations [get]
func (rbac *RBACController) ListOperations(c *gin.Context) {
	data, err := rbac.rbacService.ListOperations()
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, data)
}

/**
 * @description:Name() 返回控制器名称
 * @return {*}
 */
func (rbac *RBACController) Name() string {
	return "RBAC"
}

/**
 * @description: RegisterRoute 实现Controller接口中的方法，注册路由
 * @param {*gin.RouterGroup} api
 * @return {*}
 */
func (rbac *RBACController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/roles", rbac.List)                // rbac.List 获取role列表
	api.POST("/roles", rbac.Create)             // rbac.Create 处理程序函数，开始创建角色
	api.GET("/roles/:id", rbac.Get)             // rbac.Get 获取指定id的 roles
	api.PUT("/roles/:id", rbac.Update)          // rbac.Update 更新role
	api.DELETE("/roles/:id", rbac.Delete)       // rbac.Delete 删除role
	api.GET("/resources", rbac.ListResources)   // rbac.ListResources 资源列表
	api.GET("/operations", rbac.ListOperations) // rbac.ListOperations 操作列表
}
