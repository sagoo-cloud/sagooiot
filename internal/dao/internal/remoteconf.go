// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RemoteconfDao is the data access object for table remoteconf.
type RemoteconfDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns RemoteconfColumns // columns contains all the column names of Table for convenient usage.
}

// RemoteconfColumns defines and stores column names for table remoteconf.
type RemoteconfColumns struct {
	Id              string // 配置ID
	ConfigName      string // 配置名称
	ConfigFormat    string // 配置格式，json等
	ConfigContent   string // 配置内容
	ConfigSize      string // 配置文件大小（按字节计算）
	ProductKey      string // 产品key
	Scope           string // 配置范围：产品=product 设备=device
	Status          string // 状态： 0=停用 1=启用
	ContainedOssUrl string // 包含OssURL
	OssPath         string // Oss文件位置
	OssUrl          string // Oss链接
	Sign            string // 签名
	SignMethod      string // 签名方式，sha256等
	GmtCreate       string // 创建时间
	UtcCreate       string // UTC格式的创建时间
}

// remoteconfColumns holds the columns for table remoteconf.
var remoteconfColumns = RemoteconfColumns{
	Id:              "id",
	ConfigName:      "config_name",
	ConfigFormat:    "config_format",
	ConfigContent:   "config_content",
	ConfigSize:      "config_size",
	ProductKey:      "product_key",
	Scope:           "scope",
	Status:          "status",
	ContainedOssUrl: "contained_oss_url",
	OssPath:         "oss_path",
	OssUrl:          "oss_url",
	Sign:            "sign",
	SignMethod:      "sign_method",
	GmtCreate:       "gmt_create",
	UtcCreate:       "utc_create",
}

// NewRemoteconfDao creates and returns a new DAO object for table data access.
func NewRemoteconfDao() *RemoteconfDao {
	return &RemoteconfDao{
		group:   "default",
		table:   "remoteconf",
		columns: remoteconfColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RemoteconfDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RemoteconfDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RemoteconfDao) Columns() RemoteconfColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RemoteconfDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RemoteconfDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RemoteconfDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
