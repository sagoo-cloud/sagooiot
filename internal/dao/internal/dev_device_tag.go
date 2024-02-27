// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevDeviceTagDao is the data access object for table dev_device_tag.
type DevDeviceTagDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns DevDeviceTagColumns // columns contains all the column names of Table for convenient usage.
}

// DevDeviceTagColumns defines and stores column names for table dev_device_tag.
type DevDeviceTagColumns struct {
	Id        string //
	DeptId    string // 部门ID
	DeviceId  string // 设备ID
	DeviceKey string // 设备标识
	Key       string // 标签标识
	Name      string // 标签名称
	Value     string // 标签值
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	DeletedBy string // 删除者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// devDeviceTagColumns holds the columns for table dev_device_tag.
var devDeviceTagColumns = DevDeviceTagColumns{
	Id:        "id",
	DeptId:    "dept_id",
	DeviceId:  "device_id",
	DeviceKey: "device_key",
	Key:       "key",
	Name:      "name",
	Value:     "value",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	DeletedBy: "deleted_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewDevDeviceTagDao creates and returns a new DAO object for table data access.
func NewDevDeviceTagDao() *DevDeviceTagDao {
	return &DevDeviceTagDao{
		group:   "default",
		table:   "dev_device_tag",
		columns: devDeviceTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevDeviceTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevDeviceTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevDeviceTagDao) Columns() DevDeviceTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevDeviceTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevDeviceTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevDeviceTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
