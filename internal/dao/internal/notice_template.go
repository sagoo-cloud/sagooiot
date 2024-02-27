// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeTemplateDao is the data access object for table notice_template.
type NoticeTemplateDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns NoticeTemplateColumns // columns contains all the column names of Table for convenient usage.
}

// NoticeTemplateColumns defines and stores column names for table notice_template.
type NoticeTemplateColumns struct {
	Id          string //
	DeptId      string // 部门ID
	ConfigId    string //
	SendGateway string //
	Code        string //
	Title       string //
	Content     string //
	CreatedAt   string //
}

// noticeTemplateColumns holds the columns for table notice_template.
var noticeTemplateColumns = NoticeTemplateColumns{
	Id:          "id",
	DeptId:      "dept_id",
	ConfigId:    "config_id",
	SendGateway: "send_gateway",
	Code:        "code",
	Title:       "title",
	Content:     "content",
	CreatedAt:   "created_at",
}

// NewNoticeTemplateDao creates and returns a new DAO object for table data access.
func NewNoticeTemplateDao() *NoticeTemplateDao {
	return &NoticeTemplateDao{
		group:   "default",
		table:   "notice_template",
		columns: noticeTemplateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoticeTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoticeTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoticeTemplateDao) Columns() NoticeTemplateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoticeTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoticeTemplateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoticeTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
