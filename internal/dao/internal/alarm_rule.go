// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AlarmRuleDao is the data access object for table alarm_rule.
type AlarmRuleDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AlarmRuleColumns // columns contains all the column names of Table for convenient usage.
}

// AlarmRuleColumns defines and stores column names for table alarm_rule.
type AlarmRuleColumns struct {
	Id               string //
	DeptId           string // 部门ID
	Name             string // 告警规则名称
	Level            string // 告警级别，默认：4（一般）
	ProductKey       string // 产品标识
	DeviceKey        string // 设备标识
	TriggerMode      string // 触发方式：1=设备触发，2=定时触发
	TriggerType      string // 触发类型：1=上线，2=离线，3=属性上报, 4=事件上报
	EventKey         string // 事件标识
	TriggerCondition string // 触发条件
	Action           string // 执行动作
	Status           string // 状态：0=未启用，1=已启用
	CreatedBy        string // 创建者
	UpdatedBy        string // 更新者
	DeletedBy        string // 删除者
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 删除时间
}

// alarmRuleColumns holds the columns for table alarm_rule.
var alarmRuleColumns = AlarmRuleColumns{
	Id:               "id",
	DeptId:           "dept_id",
	Name:             "name",
	Level:            "level",
	ProductKey:       "product_key",
	DeviceKey:        "device_key",
	TriggerMode:      "trigger_mode",
	TriggerType:      "trigger_type",
	EventKey:         "event_key",
	TriggerCondition: "trigger_condition",
	Action:           "action",
	Status:           "status",
	CreatedBy:        "created_by",
	UpdatedBy:        "updated_by",
	DeletedBy:        "deleted_by",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewAlarmRuleDao creates and returns a new DAO object for table data access.
func NewAlarmRuleDao() *AlarmRuleDao {
	return &AlarmRuleDao{
		group:   "default",
		table:   "alarm_rule",
		columns: alarmRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AlarmRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AlarmRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AlarmRuleDao) Columns() AlarmRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AlarmRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AlarmRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AlarmRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
