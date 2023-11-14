package controller

import (
	"fmt"
	"net/http"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/service"
	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupService service.GroupService
}

/**
 * @description: NewGroupController() 返回一个Group控制器
 * @param {service.GroupService} groupService
 * @return {*}
 */
func NewGroupController(groupService service.GroupService) Controller {
	return &GroupController{
		groupService: groupService,
	}
}

// @Summary Create group | 创建 group
// @Description Create group and stroage | 创建 group 和 stroage 存储
// @Accept json
// @produce json
// @Tags group
// @Security JWT
// @Param group body model.CreatedGroup true "group info"
// @Success 200 {object} common.Response{data=model.Group}
// @Router /api/v1/groups [post]
func (g *GroupController) Create(c *gin.Context) {
	user := common.GetUser(c)
	if user == nil {
		common.ResponseFailed(c, http.StatusBadRequest, fmt.Errorf("Create Group 获取User失败"))
	}
	createdGroup := new(model.CreatedGroup)
	if err := c.BindJSON(createdGroup); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	group := createdGroup.GetGroup(user.ID)
	// common.TraceStep(c, "开始创建group", trace.Field{Key: "group", Value: group.Name})
	// defer common.TraceStep(c, "创建group结束", trace.Field{Key: "group", Value: group.Name})

	group, err := g.groupService.Create(user, group)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, group)
}

/**
 * @description: RegisterRoute() 注册路由
 * @param {*gin.HandlerFunc} api
 * @return {*}
 */
func (g *GroupController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/groups", g.Create)
}

/**
 * @description: Name() 返回控制器的名称
 * @return {*}
 */
func (g *GroupController) Name() string {
	return "Group"
}
