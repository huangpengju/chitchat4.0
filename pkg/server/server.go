package server

import (
	"chitchat4.0/pkg/config"
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Server 自定义一个服务
type Server struct {
	engine *gin.Engine
	config *config.Config
	logger *logrus.Logger
}

// New 接收两个参数，参数1是配置文件指针 *Config，参数2是日志记录器 *Logger 。
// 作用：返回一个配置好的服务 *Server
func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	// 1. 准备限速的中间件
	// middleware.RateLimitMiddleware(conf.Server.LimitConfig)
	// fmt.Println("logger=", logger)
	// os.Exit(0)
	db, err := database.NewPostgres(&conf.DB)
	if err != nil {
		return nil, errors.Wrap(err, "db 初始化失败")
	}

	rdb, err := database.NewRedisClient(&conf.Redis)
	if err != nil {
		return nil, errors.Wrap(err, "创建Reids 客户端失败")
	}

	repository := repository.NewRepository(db, rdb)
	if conf.DB.Migrate {
		if err := repository.Migrate(); err != nil {
			return nil, err
		}
	}

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
