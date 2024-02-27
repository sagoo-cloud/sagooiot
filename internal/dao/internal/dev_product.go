// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevProductDao is the data access object for table dev_product.
type DevProductDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns DevProductColumns // columns contains all the column names of Table for convenient usage.
}

// DevProductColumns defines and stores column names for table dev_product.
type DevProductColumns struct {
	Id                string //
	DeptId            string // 部门ID
	Key               string // 产品标识
	Name              string // 产品名称
	CategoryId        string // 所属品类
	MessageProtocol   string // 消息协议
	TransportProtocol string // 传输协议: MQTT,COAP,UDP
	ProtocolId        string // 协议id
	DeviceType        string // 设备类型: 网关，设备，子设备
	Desc              string // 描述
	Icon              string // 图片地址
	Metadata          string // 物模型
	MetadataTable     string // 是否生成物模型表：0=否，1=是
	Policy            string // 采集策略
	Status            string // 发布状态：0=未发布，1=已发布
	AuthType          string // 认证方式（1=Basic，2=AccessToken，3=证书）
	AuthUser          string // 认证用户
	AuthPasswd        string // 认证密码
	AccessToken       string // AccessToken
	CertificateId     string // 证书ID
	ScriptInfo        string // 脚本信息
	CreatedBy         string // 创建者
	UpdatedBy         string // 更新者
	DeletedBy         string // 删除者
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
	DeletedAt         string // 删除时间
}

// devProductColumns holds the columns for table dev_product.
var devProductColumns = DevProductColumns{
	Id:                "id",
	DeptId:            "dept_id",
	Key:               "key",
	Name:              "name",
	CategoryId:        "category_id",
	MessageProtocol:   "message_protocol",
	TransportProtocol: "transport_protocol",
	ProtocolId:        "protocol_id",
	DeviceType:        "device_type",
	Desc:              "desc",
	Icon:              "icon",
	Metadata:          "metadata",
	MetadataTable:     "metadata_table",
	Policy:            "policy",
	Status:            "status",
	AuthType:          "auth_type",
	AuthUser:          "auth_user",
	AuthPasswd:        "auth_passwd",
	AccessToken:       "access_token",
	CertificateId:     "certificate_id",
	ScriptInfo:        "script_info",
	CreatedBy:         "created_by",
	UpdatedBy:         "updated_by",
	DeletedBy:         "deleted_by",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
}

// NewDevProductDao creates and returns a new DAO object for table data access.
func NewDevProductDao() *DevProductDao {
	return &DevProductDao{
		group:   "default",
		table:   "dev_product",
		columns: devProductColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevProductDao) Columns() DevProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
