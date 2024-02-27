// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeLogDao is the data access object for table notice_log.
type NoticeLogDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns NoticeLogColumns // columns contains all the column names of Table for convenient usage.
}

// NoticeLogColumns defines and stores column names for table notice_log.
type NoticeLogColumns struct {
	Id          string //
	DeptId      string // 部门ID
	SendGateway string // 通知渠道
	TemplateId  string // 通知模板ID
	Addressee   string // 收信人列表
	Title       string // 通知标题
	Content     string // 通知内容
	Status      string // 发送状态：0=失败，1=成功
	FailMsg     string // 失败信息
	SendTime    string // 发送时间
}

// noticeLogColumns holds the columns for table notice_log.
var noticeLogColumns = NoticeLogColumns{
	Id:          "id",
	DeptId:      "dept_id",
	SendGateway: "send_gateway",
	TemplateId:  "template_id",
	Addressee:   "addressee",
	Title:       "title",
	Content:     "content",
	Status:      "status",
	FailMsg:     "fail_msg",
	SendTime:    "send_time",
}

// NewNoticeLogDao creates and returns a new DAO object for table data access.
func NewNoticeLogDao() *NoticeLogDao {
	return &NoticeLogDao{
		group:   "default",
		table:   "notice_log",
		columns: noticeLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoticeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoticeLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoticeLogDao) Columns() NoticeLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoticeLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoticeLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoticeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
