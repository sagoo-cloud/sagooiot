// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysApiDao is the data access object for table sys_api.
type SysApiDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SysApiColumns // columns contains all the column names of Table for convenient usage.
}

// SysApiColumns defines and stores column names for table sys_api.
type SysApiColumns struct {
	Id        string //
	ParentId  string //
	Name      string // 名称
	Types     string // 1 分类 2接口
	ApiTypes  string // 数据字典维护
	Method    string // 请求方式(数据字典维护)
	Address   string // 接口地址
	Remark    string // 备注
	Status    string // 状态 0 停用 1启用
	Sort      string // 排序
	IsDeleted string // 是否删除 0未删除 1已删除
	CreatedBy string // 创建者
	CreatedAt string // 创建时间
	UpdatedBy string // 更新者
	UpdatedAt string // 修改时间
	DeletedBy string // 删除人
	DeletedAt string // 删除时间
}

// sysApiColumns holds the columns for table sys_api.
var sysApiColumns = SysApiColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	Types:     "types",
	ApiTypes:  "api_types",
	Method:    "method",
	Address:   "address",
	Remark:    "remark",
	Status:    "status",
	Sort:      "sort",
	IsDeleted: "is_deleted",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewSysApiDao creates and returns a new DAO object for table data access.
func NewSysApiDao() *SysApiDao {
	return &SysApiDao{
		group:   "default",
		table:   "sys_api",
		columns: sysApiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysApiDao) Columns() SysApiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysApiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
