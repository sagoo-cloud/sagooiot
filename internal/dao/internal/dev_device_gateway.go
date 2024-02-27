// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevDeviceGatewayDao is the data access object for table dev_device_gateway.
type DevDeviceGatewayDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns DevDeviceGatewayColumns // columns contains all the column names of Table for convenient usage.
}

// DevDeviceGatewayColumns defines and stores column names for table dev_device_gateway.
type DevDeviceGatewayColumns struct {
	Id         string //
	GatewayKey string // 网关标识
	SubKey     string // 子设备标识
	CreatedBy  string // 创建者
	UpdatedBy  string // 更新者
	DeletedBy  string // 删除者
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// devDeviceGatewayColumns holds the columns for table dev_device_gateway.
var devDeviceGatewayColumns = DevDeviceGatewayColumns{
	Id:         "id",
	GatewayKey: "gateway_key",
	SubKey:     "sub_key",
	CreatedBy:  "created_by",
	UpdatedBy:  "updated_by",
	DeletedBy:  "deleted_by",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewDevDeviceGatewayDao creates and returns a new DAO object for table data access.
func NewDevDeviceGatewayDao() *DevDeviceGatewayDao {
	return &DevDeviceGatewayDao{
		group:   "default",
		table:   "dev_device_gateway",
		columns: devDeviceGatewayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevDeviceGatewayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevDeviceGatewayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevDeviceGatewayDao) Columns() DevDeviceGatewayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevDeviceGatewayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevDeviceGatewayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevDeviceGatewayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
