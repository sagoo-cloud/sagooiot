// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuColumnDao is the data access object for table sys_menu_column.
type SysMenuColumnDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysMenuColumnColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumnColumns defines and stores column names for table sys_menu_column.
type SysMenuColumnColumns struct {
	Id          string //
	ParentId    string // 父ID
	MenuId      string // 菜单ID
	Name        string // 名称
	Code        string // 代表字段
	Description string // 描述
	Status      string // 状态 0 停用 1启用
	IsDeleted   string // 是否删除 0未删除 1已删除
	CreatedBy   string // 创建人
	CreatedAt   string // 创建时间
	UpdatedBy   string // 修改人
	UpdatedAt   string // 更新时间
	DeletedBy   string // 删除人
	DeletedAt   string // 删除时间
}

// sysMenuColumnColumns holds the columns for table sys_menu_column.
var sysMenuColumnColumns = SysMenuColumnColumns{
	Id:          "id",
	ParentId:    "parent_id",
	MenuId:      "menu_id",
	Name:        "name",
	Code:        "code",
	Description: "description",
	Status:      "status",
	IsDeleted:   "is_deleted",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// NewSysMenuColumnDao creates and returns a new DAO object for table data access.
func NewSysMenuColumnDao() *SysMenuColumnDao {
	return &SysMenuColumnDao{
		group:   "default",
		table:   "sys_menu_column",
		columns: sysMenuColumnColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMenuColumnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMenuColumnDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMenuColumnDao) Columns() SysMenuColumnColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMenuColumnDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMenuColumnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMenuColumnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
