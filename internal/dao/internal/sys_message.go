// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMessageDao is the data access object for table sys_message.
type SysMessageDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SysMessageColumns // columns contains all the column names of Table for convenient usage.
}

// SysMessageColumns defines and stores column names for table sys_message.
type SysMessageColumns struct {
	Id        string //
	Title     string // 标题
	Types     string // 字典表
	Scope     string // 消息范围
	Content   string // 内容
	IsDeleted string // 是否删除 0未删除 1已删除
	CreatedBy string // 创建者
	CreatedAt string // 创建日期
	DeletedBy string // 删除人
	DeletedAt string // 删除时间
}

// sysMessageColumns holds the columns for table sys_message.
var sysMessageColumns = SysMessageColumns{
	Id:        "id",
	Title:     "title",
	Types:     "types",
	Scope:     "scope",
	Content:   "content",
	IsDeleted: "is_deleted",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewSysMessageDao creates and returns a new DAO object for table data access.
func NewSysMessageDao() *SysMessageDao {
	return &SysMessageDao{
		group:   "default",
		table:   "sys_message",
		columns: sysMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMessageDao) Columns() SysMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
