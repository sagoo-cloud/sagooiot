// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrganizationDao is the data access object for table sys_organization.
type SysOrganizationDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SysOrganizationColumns // columns contains all the column names of Table for convenient usage.
}

// SysOrganizationColumns defines and stores column names for table sys_organization.
type SysOrganizationColumns struct {
	Id        string // 组织ID
	DeptId    string // 部门ID
	ParentId  string // 父组织id
	Ancestors string // 祖级列表
	Name      string // 组织名称
	Number    string // 组织编号
	OrderNum  string // 显示顺序
	Leader    string // 负责人
	Phone     string // 联系电话
	Email     string // 邮箱
	Status    string // 组织状态（0停用 1正常）
	IsDeleted string // 是否删除 0未删除 1已删除
	CreatedAt string // 创建时间
	CreatedBy string // 创建人
	UpdatedBy string // 修改人
	UpdatedAt string // 修改时间
	DeletedBy string // 删除人
	DeletedAt string // 删除时间
}

// sysOrganizationColumns holds the columns for table sys_organization.
var sysOrganizationColumns = SysOrganizationColumns{
	Id:        "id",
	DeptId:    "dept_id",
	ParentId:  "parent_id",
	Ancestors: "ancestors",
	Name:      "name",
	Number:    "number",
	OrderNum:  "order_num",
	Leader:    "leader",
	Phone:     "phone",
	Email:     "email",
	Status:    "status",
	IsDeleted: "is_deleted",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewSysOrganizationDao creates and returns a new DAO object for table data access.
func NewSysOrganizationDao() *SysOrganizationDao {
	return &SysOrganizationDao{
		group:   "default",
		table:   "sys_organization",
		columns: sysOrganizationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysOrganizationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysOrganizationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysOrganizationDao) Columns() SysOrganizationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysOrganizationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysOrganizationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysOrganizationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
