// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuApiDao is the data access object for table sys_menu_api.
type SysMenuApiDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SysMenuApiColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuApiColumns defines and stores column names for table sys_menu_api.
type SysMenuApiColumns struct {
	Id        string // id
	MenuId    string // 菜单ID
	ApiId     string // apiId
	IsDeleted string // 是否删除 0未删除 1已删除
	CreatedBy string // 创建人
	CreatedAt string // 创建时间
	DeletedBy string // 删除人
	DeletedAt string // 删除时间
}

// sysMenuApiColumns holds the columns for table sys_menu_api.
var sysMenuApiColumns = SysMenuApiColumns{
	Id:        "id",
	MenuId:    "menu_id",
	ApiId:     "api_id",
	IsDeleted: "is_deleted",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewSysMenuApiDao creates and returns a new DAO object for table data access.
func NewSysMenuApiDao() *SysMenuApiDao {
	return &SysMenuApiDao{
		group:   "default",
		table:   "sys_menu_api",
		columns: sysMenuApiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMenuApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMenuApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMenuApiDao) Columns() SysMenuApiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMenuApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMenuApiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMenuApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
