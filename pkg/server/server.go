package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	docs "chitchat4.0/docs"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/config"
	"chitchat4.0/pkg/controller"
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/middleware"
	"chitchat4.0/pkg/repository"
	"chitchat4.0/pkg/service"
	"chitchat4.0/pkg/utils/request"
	"chitchat4.0/pkg/utils/set"
	"chitchat4.0/pkg/version"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	// swagger embed files
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// New 接收两个参数，参数1是配置文件指针 *Config，参数2是日志记录器 *Logger 。
// 作用：返回一个配置好的服务 *Server
func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	// 1. 准备限速的中间件
	rateLimitMiddleware, err := middleware.RateLimitMiddleware(conf.Server.LimitConfig)
	if err != nil {
		return nil, err
	}
	// fmt.Println("logger=", logger)

	db, err := database.NewPostgres(&conf.DB)
	if err != nil {
		return nil, errors.Wrap(err, "db 初始化失败")
	}
	rdb, err := database.NewRedisClient(&conf.Redis)
	if err != nil {
		return nil, errors.Wrap(err, "创建 Reids 客户端失败")
	}
	repository := repository.NewRepository(db, rdb)
	if conf.DB.Migrate {
		if err := repository.Migrate(); err != nil {
			return nil, err
		}
	}

	userService := service.NewUserService(repository.User())
	tagService := service.NewTagService(repository.Tag())
	hotSearchService := service.NewHotSearchService(repository.HotSearch())

	userController := controller.NewUserController(userService)
	tagController := controller.NewTagController(tagService)
	hotSearchController := controller.NewHotSearchController(hotSearchService)

	controllers := []controller.Controller{userController, tagController, hotSearchController}

	gin.SetMode(conf.Server.ENV) // 设置应用的模式(debug|release)

	e := gin.New() // 定义一个 gin 引擎 (不带中间件的路由)
	e.Use(         // 挂载中间件
		rateLimitMiddleware,
		gin.Recovery(), // Recovery 返回一个中间件，可以从任何恐慌中恢复

		middleware.CORSMiddleware(), // CORSMiddleware() 加载cors跨域中间件

		middleware.RequestInfoMiddleware(&request.RequestInfoFactory{APIPrefixes: set.NewString("api")}),

		middleware.LogMiddleware(logger, "/"), // 日志中间件

		middleware.TraceMiddleware(), // 追踪中间件
	)

	// 返回一个服务
	return &Server{
		engine:      e,
		config:      conf,
		logger:      logger,
		controllers: controllers,
	}, nil
}

// Server 自定义一个服务
type Server struct {
	engine *gin.Engine
	config *config.Config
	logger *logrus.Logger

	repository  repository.Repository
	controllers []controller.Controller
}

func (s *Server) Run() error {
	s.initRouter()
	// var err error

	addr := fmt.Sprintf("%s:%d", s.config.Server.Address, s.config.Server.Port)
	s.logger.Infof("启动服务器：%s", addr)
	server := &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}
	server.ListenAndServe()
	return nil
}

func (s *Server) initRouter() {

	root := s.engine

	root.GET("/", common.WrapFunc(s.getRouters))
	root.GET("/index", controller.Index)

	root.GET("/healthz", common.WrapFunc(s.Ping))
	root.GET("/version", common.WrapFunc(version.Get))
	root.GET("/metrics", gin.WrapH(promhttp.Handler()))
	root.Any("/debug/pprof/*any", gin.WrapH(http.DefaultServeMux))

	if gin.Mode() != gin.ReleaseMode {
		docs.SwaggerInfo.BasePath = "/"
		root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	api := root.Group("/api/v1")
	controllers := make([]string, 0, len(s.controllers))
	for _, router := range s.controllers {
		router.RegisterRoute(api)
		controllers = append(controllers, router.Name())
	}
	logrus.Infof("服务器启用控制器：%v", controllers)
}

func (s *Server) getRouters() []string {
	paths := set.NewString()

	for _, r := range s.engine.Routes() {
		if r.Path != "" {
			paths.Insert(r.Path)
		}
	}
	return paths.Slice()
}

type ServerStatus struct {
	Ping         bool `json:"ping"`
	DBRepository bool `json:"dbRepository"`
}

func (s *Server) Ping() *ServerStatus {
	status := &ServerStatus{Ping: true}

	ctx, cannel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cannel()

	if err := s.repository.Ping(ctx); err == nil {
		status.DBRepository = true
	}
	return status
}
