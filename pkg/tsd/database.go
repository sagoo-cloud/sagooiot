package tsd

import (
	"context"
	"database/sql"
	"sagooiot/pkg/iotModel"

	"github.com/gogf/gf/v2/database/gdb"
)

// Database 接口，所有数据库都应该实现这个接口
type Database interface {
	Close()
	Query(sql string) (rows *sql.Rows, err error)
	Count(table string) (int, error)
	InsertLogData(log iotModel.DeviceLog) (result sql.Result, err error)
	BatchInsertLogData(deviceLogList map[string][]iotModel.DeviceLog) (resultNum int, err error)
	InsertDeviceData(deviceKey string, data iotModel.ReportPropertyData, subKey ...string) (err error)
	BatchInsertDeviceData(deviceKey string, deviceDataList []iotModel.ReportPropertyData) (resultNum int, err error)
	BatchInsertMultiDeviceData(multiDeviceDataList map[string][]iotModel.ReportPropertyData) (resultNum int, err error)
	WatchDeviceData(deviceKey string, callback func(data iotModel.ReportPropertyData)) (err error)

	//获取所有数据库名称
	GetAllDatabaseName(ctx context.Context) (names []string, err error)
	//获取指定的数据库下所有的表
	GetTableListByDatabase(ctx context.Context, dbName string) (tableList []iotModel.TsdTables, err error)
	//获取指定数据表结构信息
	GetTableInfo(ctx context.Context, tableName string) (table []*iotModel.TsdTableInfo, err error)
	//获取指定数据表数据信息
	GetTableData(ctx context.Context, tableName string) (table *iotModel.TsdTableDataInfo, err error)
	//获取超级表的单条数据
	GetTableDataOne(ctx context.Context, sqlStr string, args ...any) (rs gdb.Record, err error)
	//获取超级表的多条数据
	GetTableDataAll(ctx context.Context, sqlStr string, args ...any) (rs gdb.Result, err error)
}
