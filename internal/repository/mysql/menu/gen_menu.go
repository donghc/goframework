///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package menu

import (
	"fmt"
	"goframework/internal/repository/mysql"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Menu {
	return new(Menu)
}

func NewQueryBuilder() *menuQueryBuilder {
	return new(menuQueryBuilder)
}

func (t *Menu) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type menuQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *menuQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *menuQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Menu{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *menuQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Menu{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *menuQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Menu{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *menuQueryBuilder) First(db *gorm.DB) (*Menu, error) {
	ret := &Menu{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *menuQueryBuilder) QueryOne(db *gorm.DB) (*Menu, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *menuQueryBuilder) QueryAll(db *gorm.DB) ([]*Menu, error) {
	var ret []*Menu
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *menuQueryBuilder) Limit(limit int) *menuQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *menuQueryBuilder) Offset(offset int) *menuQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *menuQueryBuilder) WhereId(p mysql.Predicate, value int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIdIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIdNotIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderById(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *menuQueryBuilder) WherePid(p mysql.Predicate, value int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pid", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WherePidIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pid", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WherePidNotIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByPid(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "pid "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereName(p mysql.Predicate, value string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereNameIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereNameNotIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByName(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereLink(p mysql.Predicate, value string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "link", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereLinkIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "link", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereLinkNotIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "link", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByLink(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "link "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereIcon(p mysql.Predicate, value string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIconIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIconNotIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByIcon(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "icon "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereLevel(p mysql.Predicate, value int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "level", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereLevelIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "level", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereLevelNotIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "level", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByLevel(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "level "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereSort(p mysql.Predicate, value int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereSortIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereSortNotIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderBySort(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereIsUsed(p mysql.Predicate, value int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIsUsedIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIsUsedNotIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByIsUsed(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_used "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIsDeletedIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereIsDeletedNotIn(value []int32) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByIsDeleted(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereCreatedAtIn(value []time.Time) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByCreatedAt(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereCreatedUserIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereCreatedUserNotIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByCreatedUser(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereUpdatedAtIn(value []time.Time) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByUpdatedAt(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *menuQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereUpdatedUserIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) WhereUpdatedUserNotIn(value []string) *menuQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *menuQueryBuilder) OrderByUpdatedUser(asc bool) *menuQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}