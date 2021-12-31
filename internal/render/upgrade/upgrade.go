package upgrade

import (
	"go.uber.org/zap"
	"goframework/internal/repository/mysql"
	"goframework/internal/repository/redis"
)

type handler struct {
	db     mysql.Repo
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}
