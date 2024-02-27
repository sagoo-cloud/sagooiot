// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPluginsConfigDao is the data access object for table sys_plugins_config.
type SysPluginsConfigDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns SysPluginsConfigColumns // columns contains all the column names of Table for convenient usage.
}

// SysPluginsConfigColumns defines and stores column names for table sys_plugins_config.
type SysPluginsConfigColumns struct {
	Id    string //
	Type  string // 插件类型
	Name  string // 插件名称
	Value string // 配置内容
	Doc   string // 配置说明
}

// sysPluginsConfigColumns holds the columns for table sys_plugins_config.
var sysPluginsConfigColumns = SysPluginsConfigColumns{
	Id:    "id",
	Type:  "type",
	Name:  "name",
	Value: "value",
	Doc:   "doc",
}

// NewSysPluginsConfigDao creates and returns a new DAO object for table data access.
func NewSysPluginsConfigDao() *SysPluginsConfigDao {
	return &SysPluginsConfigDao{
		group:   "default",
		table:   "sys_plugins_config",
		columns: sysPluginsConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPluginsConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPluginsConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysPluginsConfigDao) Columns() SysPluginsConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPluginsConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPluginsConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysPluginsConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
