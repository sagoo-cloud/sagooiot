// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthUserDao is the data access object for table oauth_user.
type OauthUserDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns OauthUserColumns // columns contains all the column names of Table for convenient usage.
}

// OauthUserColumns defines and stores column names for table oauth_user.
type OauthUserColumns struct {
	Id        string // 自增id
	Nickname  string // 用户昵称
	AvatarUrl string // 用户授权头像
	Openid    string // 用户授权唯一标识
	Provider  string // 授权对象名 qq 或 wechat
	UserId    string // 用户系统身份 id
	CreatedAt string //
	UpdatedAt string //
}

// oauthUserColumns holds the columns for table oauth_user.
var oauthUserColumns = OauthUserColumns{
	Id:        "id",
	Nickname:  "nickname",
	AvatarUrl: "avatar_url",
	Openid:    "openid",
	Provider:  "provider",
	UserId:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewOauthUserDao creates and returns a new DAO object for table data access.
func NewOauthUserDao() *OauthUserDao {
	return &OauthUserDao{
		group:   "default",
		table:   "oauth_user",
		columns: oauthUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OauthUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OauthUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OauthUserDao) Columns() OauthUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OauthUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OauthUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OauthUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
