// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthProviderDao is the data access object for table oauth_provider.
type OauthProviderDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns OauthProviderColumns // columns contains all the column names of Table for convenient usage.
}

// OauthProviderColumns defines and stores column names for table oauth_provider.
type OauthProviderColumns struct {
	Name      string // 授权对象 qq 或 wechat
	Logo      string // 授权图标地址
	Appid     string // appid
	Appsecret string // appsecret
	Status    string // 状态 0 关闭 1 开启
	CreatedAt string //
	UpdatedAt string //
}

// oauthProviderColumns holds the columns for table oauth_provider.
var oauthProviderColumns = OauthProviderColumns{
	Name:      "name",
	Logo:      "logo",
	Appid:     "appid",
	Appsecret: "appsecret",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewOauthProviderDao creates and returns a new DAO object for table data access.
func NewOauthProviderDao() *OauthProviderDao {
	return &OauthProviderDao{
		group:   "default",
		table:   "oauth_provider",
		columns: oauthProviderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OauthProviderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OauthProviderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OauthProviderDao) Columns() OauthProviderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OauthProviderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OauthProviderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OauthProviderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
