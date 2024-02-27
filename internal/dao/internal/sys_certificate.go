// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysCertificateDao is the data access object for table sys_certificate.
type SysCertificateDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysCertificateColumns // columns contains all the column names of Table for convenient usage.
}

// SysCertificateColumns defines and stores column names for table sys_certificate.
type SysCertificateColumns struct {
	Id                string //
	DeptId            string // 部门ID
	Name              string // 名称
	Standard          string // 证书标准
	FileContent       string // 证书文件内容
	PublicKeyContent  string // 证书公钥内容
	PrivateKeyContent string // 证书私钥内容
	Description       string // 说明
	Status            string // 状态  0未启用  1启用
	IsDeleted         string // 是否删除 0未删除 1已删除
	CreatedBy         string // 创建者
	CreatedAt         string // 创建日期
	UpdatedBy         string // 修改人
	UpdatedAt         string // 更新时间
	DeletedBy         string // 删除人
	DeletedAt         string // 删除时间
}

// sysCertificateColumns holds the columns for table sys_certificate.
var sysCertificateColumns = SysCertificateColumns{
	Id:                "id",
	DeptId:            "dept_id",
	Name:              "name",
	Standard:          "standard",
	FileContent:       "file_content",
	PublicKeyContent:  "public_key_content",
	PrivateKeyContent: "private_key_content",
	Description:       "description",
	Status:            "status",
	IsDeleted:         "is_deleted",
	CreatedBy:         "created_by",
	CreatedAt:         "created_at",
	UpdatedBy:         "updated_by",
	UpdatedAt:         "updated_at",
	DeletedBy:         "deleted_by",
	DeletedAt:         "deleted_at",
}

// NewSysCertificateDao creates and returns a new DAO object for table data access.
func NewSysCertificateDao() *SysCertificateDao {
	return &SysCertificateDao{
		group:   "default",
		table:   "sys_certificate",
		columns: sysCertificateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysCertificateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysCertificateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysCertificateDao) Columns() SysCertificateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysCertificateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysCertificateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysCertificateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
