package admin

import (
	"github.com/spf13/cast"
	"goframework/internal/pkg/core"
	"goframework/internal/repository/mysql"
	"strings"
)

type CreateMenuData struct {
	AdminId int32  `form:"admin_id"` // AdminID
	Actions string `form:"actions"`  // 功能权限ID,多个用,分割
}

func (s *service) CreateMenu(ctx core.Context, menuData *CreateMenuData) (err error) {
	qb := admin_menu.NewQueryBuilder()
	qb.WhereAdminId(mysql.EqualPredicate, menuData.AdminId)
	if err = qb.Delete(s.db.GetDbW().WithContext(ctx.RequestContext())); err != nil {
		return
	}

	ActionArr := strings.Split(menuData.Actions, ",")
	for _, v := range ActionArr {
		createModel := admin_menu.NewModel()
		createModel.AdminId = menuData.AdminId
		createModel.MenuId = cast.ToInt32(v)
		createModel.CreatedUser = ctx.SessionUserInfo().UserName

		_, err = createModel.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
		if err != nil {
			return
		}
	}

	return
}
