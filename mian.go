package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"goframework/configs"
	"goframework/internal/router"
	"goframework/pkg/env"
	"goframework/pkg/logger"
	"goframework/pkg/shutdown"
	"goframework/pkg/timeutil"
	"net/http"
	"time"
)

func main() {
	//初始化 access log,
	accessLogger, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
	}()

	//初始化 http 服务
	s, err := router.NewHTTPServer(accessLogger)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	//优雅关闭服务
	shutdown.NewHook().Close(
		//关闭http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}

		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}

				if err := s.Db.DbRClose(); err != nil {
					accessLogger.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},
	)

}
