package authorized

import (
	"goframework/configs"
	"goframework/internal/pkg/core"
	"goframework/internal/repository/mysql"
	"goframework/internal/repository/mysql/authorized"
	"goframework/internal/repository/redis"
	"gorm.io/gorm"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	// 先查询 id 是否存在
	authorizedInfo, err := authorized.NewQueryBuilder().
		WhereIsDeleted(mysql.EqualPredicate, -1).
		WhereId(mysql.EqualPredicate, id).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := authorized.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixSignature+authorizedInfo.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}
