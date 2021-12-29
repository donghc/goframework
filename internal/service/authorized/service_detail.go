package authorized

import (
	"goframework/internal/pkg/core"
	"goframework/internal/repository/mysql"
	"goframework/internal/repository/mysql/authorized"
)

func (s *service) Detail(ctx core.Context, id int32) (info *authorized.Authorized, err error) {
	qb := authorized.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	qb.WhereId(mysql.EqualPredicate, id)

	info, err = qb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
