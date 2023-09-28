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

func (h *HotSearchController) RegisterRoute(api *gin.RouterGroup) {

}
func (h *HotSearchController) Name() string {
	return "HotSearch"
}
