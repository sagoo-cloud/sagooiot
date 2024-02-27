// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevDeviceTreeInfoDao is the data access object for table dev_device_tree_info.
type DevDeviceTreeInfoDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns DevDeviceTreeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// DevDeviceTreeInfoColumns defines and stores column names for table dev_device_tree_info.
type DevDeviceTreeInfoColumns struct {
	Id        string //
	DeptId    string // 部门ID
	Name      string // 名称
	Code      string // 编码
	DeviceKey string // 设备标识
	Company   string // 所属公司
	Area      string // 区域
	Address   string // 地址
	Lng       string // 经度
	Lat       string // 纬度
	Contact   string // 联系人
	Phone     string // 联系电话
	StartDate string // 服务周期：开始日期
	EndDate   string // 服务周期：截止日期
	Image     string // 图片
	Duration  string // 时间窗口值
	TimeUnit  string // 时间单位：1=秒，2=分钟，3=小时，4=天
	Template  string // 页面模板，默认：default
	Category  string // 分类
	Types     string // 类型
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	DeletedBy string // 删除者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// devDeviceTreeInfoColumns holds the columns for table dev_device_tree_info.
var devDeviceTreeInfoColumns = DevDeviceTreeInfoColumns{
	Id:        "id",
	DeptId:    "dept_id",
	Name:      "name",
	Code:      "code",
	DeviceKey: "device_key",
	Company:   "company",
	Area:      "area",
	Address:   "address",
	Lng:       "lng",
	Lat:       "lat",
	Contact:   "contact",
	Phone:     "phone",
	StartDate: "start_date",
	EndDate:   "end_date",
	Image:     "image",
	Duration:  "duration",
	TimeUnit:  "time_unit",
	Template:  "template",
	Category:  "category",
	Types:     "types",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	DeletedBy: "deleted_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewDevDeviceTreeInfoDao creates and returns a new DAO object for table data access.
func NewDevDeviceTreeInfoDao() *DevDeviceTreeInfoDao {
	return &DevDeviceTreeInfoDao{
		group:   "default",
		table:   "dev_device_tree_info",
		columns: devDeviceTreeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevDeviceTreeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevDeviceTreeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevDeviceTreeInfoDao) Columns() DevDeviceTreeInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevDeviceTreeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevDeviceTreeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevDeviceTreeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
