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

func (h *HotSearchController) List(c *gin.Context) {

}

func (h *HotSearchController) Create(c *gin.Context) {

}

func (h *HotSearchController) Get(c *gin.Context) {

}

func (h *HotSearchController) Update(c *gin.Context) {

}

func (h *HotSearchController) Delete(c *gin.Context) {

}

func (h *HotSearchController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/hotsearchs", h.List)
	api.POST("/hotsearchs", h.Create)
	api.GET("/hotsearchs/:id", h.Get)
	api.PUT("/hotsearchs/:id", h.Update)
	api.DELETE("/hotsearchs/:id", h.Delete)
}
func (h *HotSearchController) Name() string {
	return "HotSearch"
}
