// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMessagereceiveDao is the data access object for table sys_messagereceive.
type SysMessagereceiveDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns SysMessagereceiveColumns // columns contains all the column names of Table for convenient usage.
}

// SysMessagereceiveColumns defines and stores column names for table sys_messagereceive.
type SysMessagereceiveColumns struct {
	Id        string //
	UserId    string // 用户ID
	MessageId string // 消息ID
	IsRead    string // 是否已读 0 未读 1已读
	IsPush    string // 是否已经推送0 否 1是
	IsDeleted string // 是否删除 0未删除 1已删除
	ReadTime  string // 阅读时间
	DeletedAt string // 删除时间
}

// sysMessagereceiveColumns holds the columns for table sys_messagereceive.
var sysMessagereceiveColumns = SysMessagereceiveColumns{
	Id:        "id",
	UserId:    "user_id",
	MessageId: "message_id",
	IsRead:    "is_read",
	IsPush:    "is_push",
	IsDeleted: "is_deleted",
	ReadTime:  "read_time",
	DeletedAt: "deleted_at",
}

// NewSysMessagereceiveDao creates and returns a new DAO object for table data access.
func NewSysMessagereceiveDao() *SysMessagereceiveDao {
	return &SysMessagereceiveDao{
		group:   "default",
		table:   "sys_messagereceive",
		columns: sysMessagereceiveColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMessagereceiveDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMessagereceiveDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMessagereceiveDao) Columns() SysMessagereceiveColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMessagereceiveDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMessagereceiveDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMessagereceiveDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
