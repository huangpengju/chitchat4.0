/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-10-07 14:40:19
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-13 16:01:16
 * @FilePath: \chitchat4.0\pkg\controller\interface.go
 * @Description: 控制器接口
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package controller

import "github.com/gin-gonic/gin"

/**
 * @description:Controller 接口，需要实现两个方法；
 * Name()返回控制器名称；
 * RegisterRoute()注册路由；
 *
 */
type Controller interface {
	Name() string                   // 名称
	RegisterRoute(*gin.RouterGroup) // 注册路由
}
