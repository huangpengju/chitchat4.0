package main

import (
	"flag"
	"fmt"
	"os"

	"chitchat4.0/pkg/config"
	"chitchat4.0/pkg/server"
	"chitchat4.0/pkg/version"
	"github.com/sirupsen/logrus"
)

// 思路
// 1.查看应用的版本信息
// 2.灵活设置配置文件路径
//
// 方法：Flag 包实现了命令行参数的解析 flag.Type(flag名, 默认值, 帮助信息)*Type
//
// 声明全局变量，定义命令行 flag 参数
var (
	printVersion = flag.Bool("v", false, "打印版本")
	appConfig    = flag.String("config", "config/app.yaml", "应用的配置路径")
)

// @title           ChitChat API
// @version         4.0
// @description     这是 chitchat 服务器 API 文档。
// @description     查看应用版本：项目启动命令后追加-v=true
// @description     指定应用配置路径：项目启动命令后追加-config=配置路径

// @contact.name 作者：黄鹏举
// @contact.url https://huangpengju.github.io/

// @license.name	Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	// 思路
	// 1.解析命令行flag参数
	// 2.打印版本
	// 3.配置日志
	// 4.加载配置
	// 5.初始化服务
	// 6.运行服务

	// 1.
	// flag.Parser() 实现命令行参数的解析
	flag.Parse()
	// 2.
	if *printVersion {
		fmt.Println("这里打印版本")
		version.Print()
		os.Exit(0)
	}
	// 3.
	logger := logrus.StandardLogger()            // 定义一个 Logger 对象
	logger.SetFormatter(&logrus.JSONFormatter{}) // 设置标准 Logger 格式化程序(程序输出 JSON 格式的日志)
	// 4.
	conf, err := config.Parse(*appConfig)
	if err != nil {
		logger.Fatalf("无法分析配置：%v", err)
	}
	// 5.
	s, err := server.New(conf, logger)
	if err != nil {
		logger.Fatalf("初始化服务器失败：%v", err)
	}
	// 6.
	if err := s.Run(); err != nil {
		logger.Fatalf("服务器启动失败：%v", err)
	}
	// 加载配置
	// setting.Init()
	// 初始化数据库
	// models.Init()
	// 普通启动
	// router := routers.InitRouter()
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf("%v:%v", setting.HTTPHost, setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}
