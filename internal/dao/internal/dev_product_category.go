// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevProductCategoryDao is the data access object for table dev_product_category.
type DevProductCategoryDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns DevProductCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// DevProductCategoryColumns defines and stores column names for table dev_product_category.
type DevProductCategoryColumns struct {
	Id        string //
	DeptId    string // 部门ID
	ParentId  string // 父ID
	Key       string // 分类标识
	Name      string // 分类名称
	Sort      string // 排序
	Desc      string // 描述
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	DeletedBy string // 删除者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// devProductCategoryColumns holds the columns for table dev_product_category.
var devProductCategoryColumns = DevProductCategoryColumns{
	Id:        "id",
	DeptId:    "dept_id",
	ParentId:  "parent_id",
	Key:       "key",
	Name:      "name",
	Sort:      "sort",
	Desc:      "desc",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	DeletedBy: "deleted_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewDevProductCategoryDao creates and returns a new DAO object for table data access.
func NewDevProductCategoryDao() *DevProductCategoryDao {
	return &DevProductCategoryDao{
		group:   "default",
		table:   "dev_product_category",
		columns: devProductCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevProductCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevProductCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevProductCategoryDao) Columns() DevProductCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevProductCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevProductCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevProductCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
