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

// @Summary Create rbac role
// @Description Create rbac role | 创建 rbac 的角色
// @Accept json
// @Produce json
// @Tags rbac
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
	api.POST("/roles", rbac.Create) // rbac.Create 处理程序函数，开始创建角色
}
