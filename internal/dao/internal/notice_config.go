// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeConfigDao is the data access object for table notice_config.
type NoticeConfigDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NoticeConfigColumns // columns contains all the column names of Table for convenient usage.
}

// NoticeConfigColumns defines and stores column names for table notice_config.
type NoticeConfigColumns struct {
	Id          string //
	DeptId      string // 部门ID
	Title       string //
	SendGateway string //
	Types       string //
	CreatedAt   string //
}

// noticeConfigColumns holds the columns for table notice_config.
var noticeConfigColumns = NoticeConfigColumns{
	Id:          "id",
	DeptId:      "dept_id",
	Title:       "title",
	SendGateway: "send_gateway",
	Types:       "types",
	CreatedAt:   "created_at",
}

// NewNoticeConfigDao creates and returns a new DAO object for table data access.
func NewNoticeConfigDao() *NoticeConfigDao {
	return &NoticeConfigDao{
		group:   "default",
		table:   "notice_config",
		columns: noticeConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoticeConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoticeConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoticeConfigDao) Columns() NoticeConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoticeConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoticeConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoticeConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
