// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevDeviceTreeDao is the data access object for table dev_device_tree.
type DevDeviceTreeDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns DevDeviceTreeColumns // columns contains all the column names of Table for convenient usage.
}

// DevDeviceTreeColumns defines and stores column names for table dev_device_tree.
type DevDeviceTreeColumns struct {
	Id           string //
	InfoId       string // 设备树信息ID
	ParentInfoId string // 父ID
}

// devDeviceTreeColumns holds the columns for table dev_device_tree.
var devDeviceTreeColumns = DevDeviceTreeColumns{
	Id:           "id",
	InfoId:       "info_id",
	ParentInfoId: "parent_info_id",
}

// NewDevDeviceTreeDao creates and returns a new DAO object for table data access.
func NewDevDeviceTreeDao() *DevDeviceTreeDao {
	return &DevDeviceTreeDao{
		group:   "default",
		table:   "dev_device_tree",
		columns: devDeviceTreeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevDeviceTreeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevDeviceTreeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevDeviceTreeDao) Columns() DevDeviceTreeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevDeviceTreeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevDeviceTreeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevDeviceTreeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
