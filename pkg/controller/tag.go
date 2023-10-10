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

func (t *TagController) List(c *gin.Context) {

}

func (t *TagController) Create(c *gin.Context) {

}

func (t *TagController) Get(c *gin.Context) {

}

func (t *TagController) Update(c *gin.Context) {

}

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
