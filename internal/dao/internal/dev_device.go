// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevDeviceDao is the data access object for table dev_device.
type DevDeviceDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns DevDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// DevDeviceColumns defines and stores column names for table dev_device.
type DevDeviceColumns struct {
	Id             string //
	DeptId         string // 部门ID
	Key            string // 设备标识
	Name           string // 设备名称
	ProductKey     string // 所属产品KEY
	Desc           string // 描述
	MetadataTable  string // 是否生成物模型子表：0=否，1=是
	Status         string // 状态：0=未启用,1=离线,2=在线
	OnlineTimeout  string // 设备在线超时设置，单位：秒
	RegistryTime   string // 激活时间
	LastOnlineTime string // 最后上线时间
	Version        string // 固件版本号
	TunnelId       string // tunnelId
	Lng            string // 经度
	Lat            string // 纬度
	AuthType       string // 认证方式（1=Basic，2=AccessToken，3=证书）
	AuthUser       string // 认证用户
	AuthPasswd     string // 认证密码
	AccessToken    string // AccessToken
	CertificateId  string // 证书ID
	CreatedBy      string // 创建者
	UpdatedBy      string // 更新者
	DeletedBy      string // 删除者
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// devDeviceColumns holds the columns for table dev_device.
var devDeviceColumns = DevDeviceColumns{
	Id:             "id",
	DeptId:         "dept_id",
	Key:            "key",
	Name:           "name",
	ProductKey:     "product_key",
	Desc:           "desc",
	MetadataTable:  "metadata_table",
	Status:         "status",
	OnlineTimeout:  "online_timeout",
	RegistryTime:   "registry_time",
	LastOnlineTime: "last_online_time",
	Version:        "version",
	TunnelId:       "tunnel_id",
	Lng:            "lng",
	Lat:            "lat",
	AuthType:       "auth_type",
	AuthUser:       "auth_user",
	AuthPasswd:     "auth_passwd",
	AccessToken:    "access_token",
	CertificateId:  "certificate_id",
	CreatedBy:      "created_by",
	UpdatedBy:      "updated_by",
	DeletedBy:      "deleted_by",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewDevDeviceDao creates and returns a new DAO object for table data access.
func NewDevDeviceDao() *DevDeviceDao {
	return &DevDeviceDao{
		group:   "default",
		table:   "dev_device",
		columns: devDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevDeviceDao) Columns() DevDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
