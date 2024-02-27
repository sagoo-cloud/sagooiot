// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAuthorizeDao is the data access object for table sys_authorize.
type SysAuthorizeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysAuthorizeColumns // columns contains all the column names of Table for convenient usage.
}

// SysAuthorizeColumns defines and stores column names for table sys_authorize.
type SysAuthorizeColumns struct {
	Id         string //
	RoleId     string // 角色ID
	ItemsType  string // 项目类型 menu菜单 button按钮 column列表字段 api接口
	ItemsId    string // 项目ID
	IsCheckAll string // 是否全选 1是 0否
	IsDeleted  string // 是否删除 0未删除 1已删除
	CreatedBy  string // 创建人
	CreatedAt  string // 创建时间
	DeletedBy  string // 删除人
	DeletedAt  string // 删除时间
}

// sysAuthorizeColumns holds the columns for table sys_authorize.
var sysAuthorizeColumns = SysAuthorizeColumns{
	Id:         "id",
	RoleId:     "role_id",
	ItemsType:  "items_type",
	ItemsId:    "items_id",
	IsCheckAll: "is_check_all",
	IsDeleted:  "is_deleted",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	DeletedBy:  "deleted_by",
	DeletedAt:  "deleted_at",
}

// NewSysAuthorizeDao creates and returns a new DAO object for table data access.
func NewSysAuthorizeDao() *SysAuthorizeDao {
	return &SysAuthorizeDao{
		group:   "default",
		table:   "sys_authorize",
		columns: sysAuthorizeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAuthorizeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAuthorizeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAuthorizeDao) Columns() SysAuthorizeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAuthorizeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAuthorizeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAuthorizeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
