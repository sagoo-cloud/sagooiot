// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPluginsDao is the data access object for table sys_plugins.
type SysPluginsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SysPluginsColumns // columns contains all the column names of Table for convenient usage.
}

// SysPluginsColumns defines and stores column names for table sys_plugins.
type SysPluginsColumns struct {
	Id                    string // ID
	DeptId                string // 部门ID
	Types                 string // 插件与SagooIOT的通信方式
	HandleType            string // 功能类型
	Name                  string // 名称
	Title                 string // 标题
	Description           string // 介绍
	Version               string // 版本
	Author                string // 作者
	Icon                  string // 插件图标
	Link                  string // 插件的网址。指向插件的 github 链接。值应为一个可访问的网址
	Command               string // 插件的运行指令
	Args                  string // 插件的指令参数
	Status                string // 状态  0未启用  1启用
	FrontendUi            string // 是否有插件页面
	FrontendUrl           string // 插件页面地址
	FrontendConfiguration string // 是否显示配置页面
	StartTime             string // 启动时间
	IsDeleted             string // 是否删除 0未删除 1已删除
	CreatedBy             string // 创建者
	CreatedAt             string // 创建日期
	UpdatedBy             string // 修改人
	UpdatedAt             string // 更新时间
	DeletedBy             string // 删除人
	DeletedAt             string // 删除时间
}

// sysPluginsColumns holds the columns for table sys_plugins.
var sysPluginsColumns = SysPluginsColumns{
	Id:                    "id",
	DeptId:                "dept_id",
	Types:                 "types",
	HandleType:            "handle_type",
	Name:                  "name",
	Title:                 "title",
	Description:           "description",
	Version:               "version",
	Author:                "author",
	Icon:                  "icon",
	Link:                  "link",
	Command:               "command",
	Args:                  "args",
	Status:                "status",
	FrontendUi:            "frontend_ui",
	FrontendUrl:           "frontend_url",
	FrontendConfiguration: "frontend_configuration",
	StartTime:             "start_time",
	IsDeleted:             "is_deleted",
	CreatedBy:             "created_by",
	CreatedAt:             "created_at",
	UpdatedBy:             "updated_by",
	UpdatedAt:             "updated_at",
	DeletedBy:             "deleted_by",
	DeletedAt:             "deleted_at",
}

// NewSysPluginsDao creates and returns a new DAO object for table data access.
func NewSysPluginsDao() *SysPluginsDao {
	return &SysPluginsDao{
		group:   "default",
		table:   "sys_plugins",
		columns: sysPluginsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPluginsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPluginsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysPluginsDao) Columns() SysPluginsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPluginsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPluginsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysPluginsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
