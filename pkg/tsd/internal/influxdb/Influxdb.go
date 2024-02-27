package influxdb

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/tsd/comm"
)

type Influxdb struct {
	Option   comm.Option
	Database string
}

type ReportReq struct {
	Timestamp  int64 //
	Metric     string
	Dimensions map[string]string
	Value      float64
}

func (m *Influxdb) client() {

	return
}

func (m *Influxdb) Close() {
	return
}
func (m *Influxdb) Query(sql string) (rows *sql.Rows, err error) {

	return
}
func (m *Influxdb) Count(table string) (int, error) {

	return 50, nil
}
func (m *Influxdb) InsertLogData(log iotModel.DeviceLog) (result sql.Result, err error) {

	return

}
func (m *Influxdb) BatchInsertLogData(ddeviceLogList map[string][]iotModel.DeviceLog) (resultNum int, err error) {

	return
}

// GetAllDatabaseName 获取所有数据库名称
func (m *Influxdb) GetAllDatabaseName(ctx context.Context) (names []string, err error) {

	return
}

// GetTableListByDatabase 获取指定的数据库下所有的表
func (m *Influxdb) GetTableListByDatabase(ctx context.Context, dbName string) (tableList []iotModel.TsdTables, err error) {

	return
}

// GetTableInfo 获取指定数据表结构信息
func (m *Influxdb) GetTableInfo(ctx context.Context, tableName string) (table []*iotModel.TsdTableInfo, err error) {
	return
}

// GetTableData 获取指定数据表数据信息
func (m *Influxdb) GetTableData(ctx context.Context, tableName string) (table *iotModel.TsdTableDataInfo, err error) {

	return
}

// GetTableDataOne 获取超级表的单条数据
func (m *Influxdb) GetTableDataOne(ctx context.Context, sqlStr string, args ...any) (rs gdb.Record, err error) {

	return
}

// GetTableDataAll 获取超级表的多条数据
func (m *Influxdb) GetTableDataAll(ctx context.Context, sqlStr string, args ...any) (rs gdb.Result, err error) {

	return
}
