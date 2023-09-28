package controller

import (
	"chitchat4.0/pkg/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) Controller {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {
	// api.GET("/use")
}

func (u *UserController) Name() string {
	return "User"
}
