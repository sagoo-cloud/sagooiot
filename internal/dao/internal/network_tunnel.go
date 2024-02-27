// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NetworkTunnelDao is the data access object for table network_tunnel.
type NetworkTunnelDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns NetworkTunnelColumns // columns contains all the column names of Table for convenient usage.
}

// NetworkTunnelColumns defines and stores column names for table network_tunnel.
type NetworkTunnelColumns struct {
	Id        string //
	DeptId    string // 部门ID
	ServerId  string // 服务ID
	Name      string //
	Types     string //
	Addr      string //
	Remote    string //
	Retry     string // 断线重连
	Heartbeat string // 心跳包
	Serial    string // 串口参数
	Protoccol string // 适配协议
	DeviceKey string // 设备标识
	Status    string //
	Last      string //
	CreatedAt string //
	UpdatedAt string //
	Remark    string // 备注
}

// networkTunnelColumns holds the columns for table network_tunnel.
var networkTunnelColumns = NetworkTunnelColumns{
	Id:        "id",
	DeptId:    "dept_id",
	ServerId:  "server_id",
	Name:      "name",
	Types:     "types",
	Addr:      "addr",
	Remote:    "remote",
	Retry:     "retry",
	Heartbeat: "heartbeat",
	Serial:    "serial",
	Protoccol: "protoccol",
	DeviceKey: "device_key",
	Status:    "status",
	Last:      "last",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Remark:    "remark",
}

// NewNetworkTunnelDao creates and returns a new DAO object for table data access.
func NewNetworkTunnelDao() *NetworkTunnelDao {
	return &NetworkTunnelDao{
		group:   "default",
		table:   "network_tunnel",
		columns: networkTunnelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetworkTunnelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetworkTunnelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetworkTunnelDao) Columns() NetworkTunnelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NetworkTunnelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetworkTunnelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetworkTunnelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
