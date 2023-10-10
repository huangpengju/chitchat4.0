package controller

import (
	"chitchat4.0/pkg/service"
	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService service.TagService
}

func NewTagController(tagService service.TagService) Controller {
	return &TagController{
		tagService: tagService,
	}
}

// @Summary List user
// @Description 列出tag和存储
// @Produce json
// @Tags user
// @Security JWT
// @Success 200 {object} common.Response{data=model.Users}
// @Router /api/v1/users [get]
func (t *TagController) List(c *gin.Context) {

}

// @Summary Create user
// @Description 创建tag和存储
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param user body model.CreatedUser true "user info"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users [post]
func (t *TagController) Create(c *gin.Context) {

}

// @Summary Get user
// @Description 获取tag和存储
// @Produce json
// @Tags user
// @Security JWT
// @Param id path int true "user id"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [get]
func (t *TagController) Get(c *gin.Context) {

}

// @Summary Update user
// @Description 更新tag和存储
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param id   path      int  true  "user id"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [put]
func (t *TagController) Update(c *gin.Context) {

}

// @Summary Delete user
// @Description 删除tag和存储
// @Produce json
// @Tags user
// @Security JWT
// @Param id path int true "user id"
// @Success 200 {object} common.Response
// @Router /api/v1/users/{id} [delete]
func (t *TagController) Delete(c *gin.Context) {

}

func (t *TagController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/tags", t.List)
	api.POST("/tags", t.Create)
	api.GET("/tags/:id", t.Get)
	api.PUT("/tags:id", t.Update)
	api.DELETE("/tags/:id", t.Delete)
}

func (t *TagController) Name() string {
	return "Tag"
}
