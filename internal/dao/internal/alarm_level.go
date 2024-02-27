// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AlarmLevelDao is the data access object for table alarm_level.
type AlarmLevelDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AlarmLevelColumns // columns contains all the column names of Table for convenient usage.
}

// AlarmLevelColumns defines and stores column names for table alarm_level.
type AlarmLevelColumns struct {
	Level string // 告警级别
	Name  string // 名称
}

// alarmLevelColumns holds the columns for table alarm_level.
var alarmLevelColumns = AlarmLevelColumns{
	Level: "level",
	Name:  "name",
}

// NewAlarmLevelDao creates and returns a new DAO object for table data access.
func NewAlarmLevelDao() *AlarmLevelDao {
	return &AlarmLevelDao{
		group:   "default",
		table:   "alarm_level",
		columns: alarmLevelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AlarmLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AlarmLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AlarmLevelDao) Columns() AlarmLevelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AlarmLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AlarmLevelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AlarmLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
