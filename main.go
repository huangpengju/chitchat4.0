package main

import (
	"fmt"
	"net/http"

	"chitchat4.0/models"
	"chitchat4.0/pkg/setting"
	"chitchat4.0/routers"
)

func main() {
	// 加载配置
	setting.Init()
	// 初始化数据库
	models.Init()
	// 普通启动
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", setting.HTTPHost, setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
