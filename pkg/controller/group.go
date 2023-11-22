/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-11-14 15:29:14
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-22 12:17:26
 * @FilePath: \chitchat4.0\pkg\controller\group.go
 * @Description:
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package controller

import (
	"fmt"
	"net/http"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/service"
	"chitchat4.0/pkg/utils/trace"
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
	common.TraceStep(c, "开始创建group", trace.Field{Key: "group", Value: group.Name})
	defer common.TraceStep(c, "创建group结束", trace.Field{Key: "group", Value: group.Name})

	group, err := g.groupService.Create(user, group)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, group)
}

// @Summary Get group | 获取 group
// @Description Get group | 通过id查询group
// @Produce json
// @Tags group
// @Security JWT
// @Param id path int true "group id"
// @Success 200 {object} common.Response{data=model.Group}
// @Router /api/v1/groups/{id} [get]
func (g *GroupController) Get(c *gin.Context) {
	group, err := g.groupService.Get(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, group)
}

// @Summary List group | group 列表
// @Description List group | 查询所有group列表
// @Produce json
// @Tags group
// @Security JWT
// @Success 200 {object} common.Response{data=[]model.Group}
// @Router /api/v1/groups [get]
func (g *GroupController) List(c *gin.Context) {
	common.TraceStep(c, "start list group(开始获取组列表)")
	groups, err := g.groupService.List()
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.TraceStep(c, "list group done(group列表结束)")
	common.ResponseSuccess(c, groups)
}

// @Summary Update group | 修改 group
// @Description Update group and storage | 修改group和保存
// @Accept json
// @Produce json
// @Tags group
// @Security JWT
// @Param group body model.UpdatedGroup true "group info"
// @Param id path int true "group id"
// @Success 200 {object} common.Response{data=model.Group}
// @Router /api/v1/groups/{id} [put]
func (g *GroupController) Update(c *gin.Context) {
	user := common.GetUser(c)
	if user == nil {
		common.ResponseFailed(c, http.StatusBadRequest, fmt.Errorf("Update Group:Failed to get user"))
		return
	}

	id := c.Param("id")

	new := new(model.UpdatedGroup)
	if err := c.BindJSON(new); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.TraceStep(c, "start update group", trace.Field{Key: "group", Value: new.Name})
	defer common.TraceStep(c, "update group done", trace.Field{Key: "group", Value: new.Name})

	group, err := g.groupService.Update(id, new.GetGroup(user.ID))
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, group)
}

// @Summary Delete group | 删除group
// @Description Delete group | 删除指定的group
// @Produce json
// @Tags group
// @Security JWT
// @Param id path int true "group id"
// @Success 200 {object} common.Response
// @Router /api/v1/groups/{id} [delete]
func (g *GroupController) Delete(c *gin.Context) {
	user := common.GetUser(c)
	if user == nil {
		common.ResponseFailed(c, http.StatusBadRequest, fmt.Errorf("Delete group:failed to get user"))
		return
	}

	if err := g.groupService.Delete(c.Param("id")); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, nil)
}

// @Summary Group Get users | 获取user集合
// @Description Get users to group| 根据 group 获取user集合
// @Produce json
// @Tags group
// @Security JWT
// @Param id path int true "group id"
// @Success 200 {object} common.Response
// @Router /api/v1/groups/{id}/users [get]
func (g *GroupController) GetUsers(c *gin.Context) {
	users, err := g.groupService.GetUsers(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, users)
}

// @Summary Group Add user | 添加user
// @Description Add user to group | 把user添加到group中
// @Produce json
// @Tags group
// @Security JWT
// @Param id path int true "group id"
// @Param user body model.User true "user info"
// @Success 200 {object} common.Response
// @Router /api/v1/groups/{id}/users [post]
func (g *GroupController) AddUser(c *gin.Context) {
	user := new(model.User)

	if err := c.BindJSON(user); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	if err := g.groupService.AddUser(user, c.Param("id")); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, nil)
}

// @Summary Delete user | 删除user
// @Description Delete user from group | 删除group中的user
// @Produce json
// @Tags group
// @Security JWT
// @Param id path int true "group id"
// @Param uid query int true "user id"
// @Success 200 {object} common.Response
// @Router /api/v1/groups/{id}/users [delete]
func (g *GroupController) DelUser(c *gin.Context) {
	if err := g.groupService.DelUser(c.Param("id"), c.Query("uid")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, nil)
}

// @Summary Add role | 添加角色
// @Description Add role to group | 给 group 添加 role
// @Produce json
// @Tags group
// @Security JWT
// @Param id path int true "group id"
// @param rid path int true "role id"
// @Success 200 {object} common.Response
// @Router /api/v1/groups/{id}/roles/{rid} [post]
func (g *GroupController) AddRole(c *gin.Context) {
	if err := g.groupService.AddRole(c.Param("id"), c.Param("rid")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, nil)

}

/**
 * @description: RegisterRoute() 注册路由
 * @param {*gin.HandlerFunc} api
 * @return {*}
 */
func (g *GroupController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/groups", g.List)
	api.POST("/groups", g.Create)
	api.GET("/groups/:id", g.Get)
	api.PUT("/groups/:id", g.Update)
	api.DELETE("/groups/:id", g.Delete)
	api.GET("/groups/:id/users", g.GetUsers)
	api.POST("/groups/:id/users", g.AddUser)
	api.DELETE("/groups/:id/users", g.DelUser)
	api.POST("/groups/:id/roles/:rid", g.AddRole)
}

/**
 * @description: Name() 返回控制器的名称
 * @return {*}
 */
func (g *GroupController) Name() string {
	return "Group"
}
