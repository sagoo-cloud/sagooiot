// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NetworkServerDao is the data access object for table network_server.
type NetworkServerDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns NetworkServerColumns // columns contains all the column names of Table for convenient usage.
}

// NetworkServerColumns defines and stores column names for table network_server.
type NetworkServerColumns struct {
	Id            string //
	DeptId        string // 部门ID
	Name          string //
	Types         string // tcp/udp
	Addr          string //
	Register      string // 注册包
	Heartbeat     string // 心跳包
	Protocol      string // 协议
	Devices       string // 默认设备
	Status        string //
	CreatedAt     string //
	UpdatedAt     string //
	CreateBy      string //
	Remark        string // 备注
	IsTls         string // 开启TLS:1=是，0=否
	AuthType      string // 认证方式（1=Basic，2=AccessToken，3=证书）
	AuthUser      string // 认证用户
	AuthPasswd    string // 认证密码
	AccessToken   string // AccessToken
	CertificateId string // 证书ID
	Stick         string // 粘包处理方式
}

// networkServerColumns holds the columns for table network_server.
var networkServerColumns = NetworkServerColumns{
	Id:            "id",
	DeptId:        "dept_id",
	Name:          "name",
	Types:         "types",
	Addr:          "addr",
	Register:      "register",
	Heartbeat:     "heartbeat",
	Protocol:      "protocol",
	Devices:       "devices",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	CreateBy:      "create_by",
	Remark:        "remark",
	IsTls:         "is_tls",
	AuthType:      "auth_type",
	AuthUser:      "auth_user",
	AuthPasswd:    "auth_passwd",
	AccessToken:   "access_token",
	CertificateId: "certificate_id",
	Stick:         "stick",
}

// NewNetworkServerDao creates and returns a new DAO object for table data access.
func NewNetworkServerDao() *NetworkServerDao {
	return &NetworkServerDao{
		group:   "default",
		table:   "network_server",
		columns: networkServerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetworkServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetworkServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetworkServerDao) Columns() NetworkServerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NetworkServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetworkServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetworkServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
