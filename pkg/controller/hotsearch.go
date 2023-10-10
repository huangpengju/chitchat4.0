package controller

import (
	"chitchat4.0/pkg/service"
	"github.com/gin-gonic/gin"
)

type HotSearchController struct {
	hotSearchService service.HotSearchService
}

func NewHotSearchController(hotSearchService service.HotSearchService) Controller {
	return &HotSearchController{
		hotSearchService: hotSearchService,
	}
}

// @Summary List user
// @Description 列出热搜和存储
// @Produce json
// @Tags user
// @Security JWT
// @Success 200 {object} common.Response{data=model.Users}
// @Router /api/v1/users [get]
func (h *HotSearchController) List(c *gin.Context) {

}

// @Summary Create user
// @Description 创建热搜和存储
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param user body model.CreatedUser true "user info"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users [post]
func (h *HotSearchController) Create(c *gin.Context) {

}

// @Summary Get user
// @Description 获取热搜和存储
// @Produce json
// @Tags user
// @Security JWT
// @Param id path int true "user id"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [get]
func (h *HotSearchController) Get(c *gin.Context) {

}

// @Summary Update user
// @Description 更新热搜和存储
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param id   path      int  true  "user id"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [put]
func (h *HotSearchController) Update(c *gin.Context) {

}

// @Summary Delete user
// @Description 删除热搜和存储
// @Produce json
// @Tags user
// @Security JWT
// @Param id path int true "user id"
// @Success 200 {object} common.Response
// @Router /api/v1/users/{id} [delete]
func (h *HotSearchController) Delete(c *gin.Context) {

}

func (h *HotSearchController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/hotsearchs", h.List)
	api.POST("/hotsearchs", h.Create)
	api.GET("/hotsearchs/:id", h.Get)
	api.PUT("/hotsearchs:id", h.Update)
	api.DELETE("/hotsearchs/:id", h.Delete)
}
func (h *HotSearchController) Name() string {
	return "HotSearch"
}
