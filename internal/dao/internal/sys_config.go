// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigDao is the data access object for table sys_config.
type SysConfigDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SysConfigColumns // columns contains all the column names of Table for convenient usage.
}

// SysConfigColumns defines and stores column names for table sys_config.
type SysConfigColumns struct {
	ConfigId       string // 参数主键
	ModuleClassify string // 所属字典类型数据code
	ConfigName     string // 参数名称
	ConfigKey      string // 参数键名
	ConfigValue    string // 参数键值
	ConfigType     string // 系统内置（1是 2否）
	Remark         string // 备注
	Status         string // 状态 0 停用 1启用
	IsDeleted      string // 是否删除 0未删除 1已删除
	CreatedBy      string // 创建者
	CreatedAt      string // 创建时间
	UpdatedBy      string // 更新者
	UpdatedAt      string // 修改时间
	DeletedBy      string // 删除人
	DeletedAt      string // 删除时间
}

// sysConfigColumns holds the columns for table sys_config.
var sysConfigColumns = SysConfigColumns{
	ConfigId:       "config_id",
	ModuleClassify: "module_classify",
	ConfigName:     "config_name",
	ConfigKey:      "config_key",
	ConfigValue:    "config_value",
	ConfigType:     "config_type",
	Remark:         "remark",
	Status:         "status",
	IsDeleted:      "is_deleted",
	CreatedBy:      "created_by",
	CreatedAt:      "created_at",
	UpdatedBy:      "updated_by",
	UpdatedAt:      "updated_at",
	DeletedBy:      "deleted_by",
	DeletedAt:      "deleted_at",
}

// NewSysConfigDao creates and returns a new DAO object for table data access.
func NewSysConfigDao() *SysConfigDao {
	return &SysConfigDao{
		group:   "default",
		table:   "sys_config",
		columns: sysConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysConfigDao) Columns() SysConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
