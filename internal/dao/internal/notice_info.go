// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeInfoDao is the data access object for table notice_info.
type NoticeInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns NoticeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// NoticeInfoColumns defines and stores column names for table notice_info.
type NoticeInfoColumns struct {
	Id         string //
	ConfigId   string //
	ComeFrom   string //
	Method     string //
	MsgTitle   string //
	MsgBody    string //
	MsgUrl     string //
	UserIds    string //
	OrgIds     string //
	Totag      string //
	Status     string //
	MethodCron string //
	MethodNum  string //
	CreatedAt  string //
}

// noticeInfoColumns holds the columns for table notice_info.
var noticeInfoColumns = NoticeInfoColumns{
	Id:         "id",
	ConfigId:   "config_id",
	ComeFrom:   "come_from",
	Method:     "method",
	MsgTitle:   "msg_title",
	MsgBody:    "msg_body",
	MsgUrl:     "msg_url",
	UserIds:    "user_ids",
	OrgIds:     "org_ids",
	Totag:      "totag",
	Status:     "status",
	MethodCron: "method_cron",
	MethodNum:  "method_num",
	CreatedAt:  "created_at",
}

// NewNoticeInfoDao creates and returns a new DAO object for table data access.
func NewNoticeInfoDao() *NoticeInfoDao {
	return &NoticeInfoDao{
		group:   "default",
		table:   "notice_info",
		columns: noticeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoticeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoticeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoticeInfoDao) Columns() NoticeInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoticeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoticeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoticeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
