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
func (t *TagController) RegisterRoute(api *gin.RouterGroup) {

}

func (t *TagController) Name() string {
	return "Tag"
}
