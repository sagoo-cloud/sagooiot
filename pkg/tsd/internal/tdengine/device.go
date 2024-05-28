package tdengine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/service"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/tsd/comm"
	"time"

	"strings"
)

// Initialization 初始化tsd相关的设备数据
func (m *TdEngine) Initialization(ctx context.Context) (err error) {
	// 资源锁
	lockKey := "tdLock:initDb"
	lockVal, err := g.Redis().Do(ctx, "SET", lockKey, gtime.Now().Unix(), "NX", "EX", "3600")
	if err != nil {
		return
	}
	if lockVal.String() != "OK" {
		return
	}
	defer func() {
		_, err = g.Redis().Do(ctx, "DEL", lockKey)
	}()

	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	dbName := g.Cfg().MustGet(context.Background(), "tsd.tdengine.dbName", "sagoo_iot").String()
	_, err = taos.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)

	return
}

// InsertDeviceData 插入设备数据
func (m *TdEngine) InsertDeviceData(deviceKey string, data iotModel.ReportPropertyData, subKey ...string) (err error) {
	if m.db == nil {
		m.connect()
	}

	if len(data) == 0 {
		err = errors.New("数据为空")
		return
	}

	field, value := comm.GetDeviceFieldAndValue(data)
	table := comm.DeviceTableName(deviceKey)
	if len(subKey) > 0 {
		// 子设备
		table = comm.DeviceTableName(subKey[0])
	}
	var baseSQL = "INSERT INTO " + table + " (" + strings.Join(field, ",") + ") VALUES"
	sqlStr := baseSQL + fmt.Sprintf(" (%s)", strings.Join(value, ","))
	_, err = m.db.Exec(sqlStr)
	return
}

// BatchInsertDeviceData 批量插入单设备的数据
func (m *TdEngine) BatchInsertDeviceData(deviceKey string, deviceDataList []iotModel.ReportPropertyData) (resultNum int, err error) {
	if m.db == nil {
		m.connect()
	}
	if len(deviceDataList) == 0 {
		err = errors.New("数据为空")
		return
	}
	table := comm.DeviceTableName(deviceKey)
	keys, fieldList := comm.GetDeviceField(deviceDataList[0])
	var (
		ts         = time.Now().UnixMilli() // Unix 毫秒时间戳
		baseSQL    = "INSERT INTO " + table + " (" + strings.Join(fieldList, ",") + ") VALUES"
		sqlBuilder strings.Builder
		allCount   int
		allTime    int64
	)
	sqlBuilder.WriteString(baseSQL)

	for i, row := range deviceDataList {
		ts++
		value := comm.GetDeviceValue(keys, row) //获取设备数据值
		sqlBuilder.WriteString(fmt.Sprintf(" (%s)", "'"+time.UnixMilli(ts).Format(time.RFC3339Nano)+"',"+strings.Join(value, ",")))
		// 当 SQL 字符串长度超过 15K 或在最后一条数据时执行
		if sqlBuilder.Len() > 15*1024 || i == len(deviceDataList)-1 {
			trimmedSQL := strings.TrimRight(sqlBuilder.String(), " ")
			start := time.Now() // 开始时间
			g.Log().Debug(context.Background(), "====06====BatchInsertDeviceData SQL:", trimmedSQL)
			_, err := m.db.Exec(trimmedSQL)
			if err != nil {
				g.Log().Error(context.Background(), err.Error(), trimmedSQL)
			}
			duration := time.Since(start).Milliseconds() // 执行时间
			executedCount := i + 1 - allCount            // 执行条数
			//fmt.Printf("%d, %dms\n", executedCount, duration)
			allCount += executedCount // 总条数
			allTime += duration       // 总时间
			sqlBuilder.Reset()        // 重置 sqlBuilder
			sqlBuilder.WriteString(baseSQL)
		}
	}
	resultNum = allCount
	return
}

// BatchInsertMultiDeviceData 插入多设备的数据
func (m *TdEngine) BatchInsertMultiDeviceData(multiDeviceDataList map[string][]iotModel.ReportPropertyData) (resultNum int, err error) {
	if m.db == nil {
		m.connect()
	}
	if len(multiDeviceDataList) == 0 {
		err = errors.New("数据为空")
		return
	}

	var (
		ts         = time.Now().UnixMilli() // Unix 毫秒时间戳
		baseSQL    = "INSERT INTO"
		sqlBuilder strings.Builder
		allCount   int
		allTime    int64
	)
	sqlBuilder.WriteString(baseSQL)

	i := 0
	for deviceKey, deviceData := range multiDeviceDataList {
		table := comm.DeviceTableName(deviceKey)
		fieldKeys, fieldList := comm.GetDeviceField(deviceData[0])
		ts++
		sqlBuilder.WriteString(" " + table + " (" + strings.Join(fieldList, ",") + ") VALUES")

		for _, data := range deviceData {
			value := comm.GetDeviceValue(fieldKeys, data)
			sqlBuilder.WriteString(fmt.Sprintf(" (%s)", "'"+time.UnixMilli(ts).Format(time.RFC3339Nano)+"',"+strings.Join(value, ",")))
		}

		// 当 SQL 字符串长度超过 15K 或在最后一条数据时执行
		if sqlBuilder.Len() > 15*1024 || i == len(multiDeviceDataList)-1 {
			trimmedSQL := strings.TrimRight(sqlBuilder.String(), " ")
			start := time.Now() // 开始时间
			//g.Log().Debug(context.Background(), "====06====BatchInsertDeviceData SQL:", trimmedSQL)
			_, err := m.db.Exec(trimmedSQL)
			if err != nil {
				g.Log().Error(context.Background(), err.Error(), trimmedSQL)
			}
			duration := time.Since(start).Milliseconds() // 执行时间
			executedCount := i + 1 - allCount            // 执行条数
			//fmt.Printf("%d, %dms\n", executedCount, duration)
			allCount += executedCount // 总条数
			allTime += duration       // 总时间
			sqlBuilder.Reset()        // 重置 sqlBuilder
			sqlBuilder.WriteString(baseSQL)
		}
		i++
	}
	resultNum = allCount
	//g.Log().Debugf(context.Background(), "Total: %d, Time: %dms\n", allCount, allTime)
	return
}

// WatchDeviceData 监听设备数据日志
func (m *TdEngine) WatchDeviceData(deviceKey string, callback func(data iotModel.ReportPropertyData)) (err error) {

	return
}

// InsertLogData 插入日志数据
func (m *TdEngine) InsertLogData(log iotModel.DeviceLog) (result sql.Result, err error) {
	if m.db == nil {
		m.connect()
	}

	table := comm.DeviceLogTable(log.Device)
	baseSQL := "INSERT INTO %s USING device_log TAGS ('%s') VALUES ('%s', '%s', '%s')"
	sqlStr := fmt.Sprintf(baseSQL, table, log.Device, time.Now().Format(time.RFC3339Nano), log.Type, log.Content)
	_, err = m.db.Exec(sqlStr)

	return
}

// BatchInsertLogData 批量插入日志数据
func (m *TdEngine) BatchInsertLogData(deviceLogList map[string][]iotModel.DeviceLog) (resultNum int, err error) {
	if m.db == nil {
		m.connect()
	}
	if len(deviceLogList) == 0 {
		return
	}
	//g.Log().Debug(context.Background(), "====BatchInsertLogData===接收到   =========", len(deviceLogList))

	var (
		ts         = time.Now().UnixMilli() // Unix 毫秒时间戳
		baseSQL    = "INSERT INTO "
		sqlBuilder strings.Builder
		allCount   int
		allTime    int64
	)
	sqlBuilder.WriteString(baseSQL)

	i := 0
	for k, row := range deviceLogList {
		i++
		table := comm.DeviceLogTable(k)
		tableSql := fmt.Sprintf("%s USING device_log TAGS ('%s') VALUES ", table, k)
		sqlBuilder.WriteString(tableSql)
		ts++
		for _, d := range row {
			sqlBuilder.WriteString(fmt.Sprintf("('%s', '%s', '%s') ", time.UnixMilli(ts).Format(time.RFC3339Nano), d.Type, d.Content))
		}
		// 当 SQL 字符串长度超过 15K 或在最后一条数据时执行
		if sqlBuilder.Len() > 15*1024 || i == len(row) {
			trimmedSQL := strings.TrimRight(sqlBuilder.String(), " ")
			start := time.Now() // 开始时间
			fmt.Println("====写入TD==》》》", sqlBuilder.Len(), len(deviceLogList), trimmedSQL)
			_, err = m.db.Exec(trimmedSQL)
			if err != nil {
				g.Log().Error(context.Background(), err.Error())
			}
			duration := time.Since(start).Milliseconds() // 执行时间
			executedCount := i - allCount                // 执行条数
			//fmt.Printf("%d, %dms\n", executedCount, duration)
			allCount += executedCount // 总条数
			allTime += duration       // 总时间
			sqlBuilder.Reset()        // 重置 sqlBuilder
			sqlBuilder.WriteString(baseSQL)
			resultNum = allCount
		}
	}

	//fmt.Printf("Total: %d, Time: %dms\n", allCount, allTime)
	return
}
