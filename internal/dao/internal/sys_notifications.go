// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysNotificationsDao is the data access object for table sys_notifications.
type SysNotificationsDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns SysNotificationsColumns // columns contains all the column names of Table for convenient usage.
}

// SysNotificationsColumns defines and stores column names for table sys_notifications.
type SysNotificationsColumns struct {
	Id        string //
	Title     string // 标题
	Doc       string // 描述
	Source    string // 消息来源
	Types     string // 类型
	CreatedAt string // 发送时间
	Status    string // 0，未读，1，已读
}

// sysNotificationsColumns holds the columns for table sys_notifications.
var sysNotificationsColumns = SysNotificationsColumns{
	Id:        "id",
	Title:     "title",
	Doc:       "doc",
	Source:    "source",
	Types:     "types",
	CreatedAt: "created_at",
	Status:    "status",
}

// NewSysNotificationsDao creates and returns a new DAO object for table data access.
func NewSysNotificationsDao() *SysNotificationsDao {
	return &SysNotificationsDao{
		group:   "default",
		table:   "sys_notifications",
		columns: sysNotificationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysNotificationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysNotificationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysNotificationsDao) Columns() SysNotificationsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysNotificationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysNotificationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysNotificationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
