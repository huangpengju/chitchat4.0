package server

import (
	"os"

	"chitchat4.0/pkg/config"
	"chitchat4.0/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Server 自定义一个服务
type Server struct {
	engine *gin.Engine
	config *config.Config
	logger *logrus.Logger
}

func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	// 1. 准备限速的中间件
	middleware.RateLimitMiddleware(conf.Server.LimitConfig)
	// fmt.Println("logger=", logger)
	os.Exit(0)

	gin.SetMode(conf.Server.ENV) // 设置应用的模式(debug|release)

	e := gin.New() // 定义一个 gin 引擎 (不带中间件的路由)
	e.Use(         // 挂载中间件
		gin.Recovery(), // Recovery 返回一个中间件，可以从任何恐慌中恢复
	)

	// 返回一个服务
	return &Server{
		engine: e,
		config: conf,
		logger: logger,
	}, nil
}

func (s *Server) Run() error {

	var err error
	return err
}
