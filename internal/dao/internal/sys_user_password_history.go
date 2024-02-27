// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserPasswordHistoryDao is the data access object for table sys_user_password_history.
type SysUserPasswordHistoryDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns SysUserPasswordHistoryColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserPasswordHistoryColumns defines and stores column names for table sys_user_password_history.
type SysUserPasswordHistoryColumns struct {
	Id             string //
	UserId         string // 用户ID
	BeforePassword string // 变更之前密码
	AfterPassword  string // 变更之后密码
	ChangeTime     string // 变更时间
	CreatedAt      string //
	CreatedBy      string //
}

// sysUserPasswordHistoryColumns holds the columns for table sys_user_password_history.
var sysUserPasswordHistoryColumns = SysUserPasswordHistoryColumns{
	Id:             "id",
	UserId:         "user_id",
	BeforePassword: "before_password",
	AfterPassword:  "after_password",
	ChangeTime:     "change_time",
	CreatedAt:      "created_at",
	CreatedBy:      "created_by",
}

// NewSysUserPasswordHistoryDao creates and returns a new DAO object for table data access.
func NewSysUserPasswordHistoryDao() *SysUserPasswordHistoryDao {
	return &SysUserPasswordHistoryDao{
		group:   "default",
		table:   "sys_user_password_history",
		columns: sysUserPasswordHistoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserPasswordHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserPasswordHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysUserPasswordHistoryDao) Columns() SysUserPasswordHistoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserPasswordHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserPasswordHistoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserPasswordHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
