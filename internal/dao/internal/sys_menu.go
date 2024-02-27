// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for table sys_menu.
type SysMenuDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumns defines and stores column names for table sys_menu.
type SysMenuColumns struct {
	Id         string //
	ParentId   string // 父ID
	Name       string // 规则名称
	Title      string // 菜单名称
	Icon       string // 图标
	Condition  string // 条件
	Remark     string // 备注
	MenuType   string // 类型 0目录 1菜单 2按钮
	Weigh      string // 权重
	IsHide     string // 显示状态
	Path       string // 路由地址
	Component  string // 组件路径
	IsLink     string // 是否外链 1是 0否
	ModuleType string // 所属模块 system 运维 company企业
	ModelId    string // 模型ID
	IsIframe   string // 是否内嵌iframe
	IsCached   string // 是否缓存
	Redirect   string // 路由重定向地址
	IsAffix    string // 是否固定
	LinkUrl    string // 链接地址
	Status     string // 状态 0 停用 1启用
	IsDeleted  string // 是否删除 0未删除 1已删除
	CreatedBy  string // 创建人
	CreatedAt  string // 创建时间
	UpdatedBy  string // 修改人
	UpdatedAt  string // 更新时间
	DeletedBy  string // 删除人
	DeletedAt  string // 删除时间
}

// sysMenuColumns holds the columns for table sys_menu.
var sysMenuColumns = SysMenuColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Name:       "name",
	Title:      "title",
	Icon:       "icon",
	Condition:  "condition",
	Remark:     "remark",
	MenuType:   "menu_type",
	Weigh:      "weigh",
	IsHide:     "is_hide",
	Path:       "path",
	Component:  "component",
	IsLink:     "is_link",
	ModuleType: "module_type",
	ModelId:    "model_id",
	IsIframe:   "is_iframe",
	IsCached:   "is_cached",
	Redirect:   "redirect",
	IsAffix:    "is_affix",
	LinkUrl:    "link_url",
	Status:     "status",
	IsDeleted:  "is_deleted",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
	DeletedBy:  "deleted_by",
	DeletedAt:  "deleted_at",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		group:   "default",
		table:   "sys_menu",
		columns: sysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
