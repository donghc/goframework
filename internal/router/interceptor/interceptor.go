package interceptor

import (
	"go.uber.org/zap"
	"goframework/internal/pkg/core"
	"goframework/internal/proposal"
	"goframework/internal/repository/mysql"
	"goframework/internal/repository/redis"
	"goframework/internal/service/admin"
	"goframework/internal/service/authorized"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	i()

	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	// CheckRBAC 验证 RBAC 权限是否合法
	CheckRBAC() core.HandlerFunc

	// CheckSignature 验证签名是否合法，对用签名算法 pkg/signature
	CheckSignature() core.HandlerFunc
}

type interceptor struct {
	logger            *zap.Logger
	cache             redis.Repo
	db                mysql.Repo
	authorizedService authorized.Service
	adminService      admin.Service
}

func New(logger *zap.Logger, cache redis.Repo, db mysql.Repo) Interceptor {
	return &interceptor{
		logger:            logger,
		cache:             cache,
		db:                db,
		authorizedService: authorized.New(db, cache),
		adminService:      admin.New(db, cache),
	}
}

func (i *interceptor) i() {}
