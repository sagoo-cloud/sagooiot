// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysJobDao is the data access object for table sys_job.
type SysJobDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SysJobColumns // columns contains all the column names of Table for convenient usage.
}

// SysJobColumns defines and stores column names for table sys_job.
type SysJobColumns struct {
	JobId          string // 任务ID
	JobName        string // 任务名称
	JobParams      string // 参数
	JobGroup       string // 任务组名
	InvokeTarget   string // 调用目标字符串
	CronExpression string // cron执行表达式
	MisfirePolicy  string // 计划执行策略（1多次执行 2执行一次）
	Concurrent     string // 是否并发执行（0允许 1禁止）
	Status         string // 状态（0正常 1暂停）
	CreatedBy      string // 创建者
	UpdatedBy      string // 更新者
	Remark         string // 备注信息
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// sysJobColumns holds the columns for table sys_job.
var sysJobColumns = SysJobColumns{
	JobId:          "job_id",
	JobName:        "job_name",
	JobParams:      "job_params",
	JobGroup:       "job_group",
	InvokeTarget:   "invoke_target",
	CronExpression: "cron_expression",
	MisfirePolicy:  "misfire_policy",
	Concurrent:     "concurrent",
	Status:         "status",
	CreatedBy:      "created_by",
	UpdatedBy:      "updated_by",
	Remark:         "remark",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewSysJobDao creates and returns a new DAO object for table data access.
func NewSysJobDao() *SysJobDao {
	return &SysJobDao{
		group:   "default",
		table:   "sys_job",
		columns: sysJobColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysJobDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysJobDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysJobDao) Columns() SysJobColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysJobDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysJobDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysJobDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
