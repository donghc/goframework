package router

import (
	"go.uber.org/zap"
	"goframework/configs"
	"goframework/internal/alert"
	"goframework/internal/metrics"
	"goframework/internal/pkg/core"
	"goframework/internal/repository/mysql"
	"goframework/internal/repository/redis"
	"goframework/pkg/errors"
	"goframework/pkg/file"
)

type resource struct {
	mux    core.Mux
	logger *zap.Logger
	db     mysql.Repo
	cache  redis.Repo
}

type Server struct {
	Mux   core.Mux
	Db    mysql.Repo
	Cache redis.Repo
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger
	openBrowserUri := configs.ProjectName + configs.ProjectPort

	_, ok := file.IsExists(configs.ProjectInstallMark)
	{
		if !ok { //未安装
			openBrowserUri += "/install"
		} else { //已安装
			//初始化DB
			dbRepo, err := mysql.New()
			if err != nil {
				logger.Fatal("new db err :", zap.Error(err))
			}
			r.db = dbRepo

			//初始化 redis
			redisRepo, err := redis.New()
			if err != nil {
				logger.Fatal("new redis err", zap.Error(err))
			}
			r.cache = redisRepo
		}
	}

	mux, err := core.New(logger,
		core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithAlertNotify(alert.NotifyHandler(logger)),
		core.WithRecordMetrics(metrics.RecordHandler(logger)),
	)

	if err != nil {
		panic(err)
	}

	r.mux = mux
	// 设置 Render 路由
	setRenderRouter(r)
	// 设置 API 路由
	setApiRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache

	return s, nil
}
