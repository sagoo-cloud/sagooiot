package tsd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/pkg/iotModel"
	"testing"
	"time"

	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
)

func TestGetTableListByDatabase(t *testing.T) {
	db := DB()
	tableList, err := db.GetTableListByDatabase(context.Background(), "sagoo_iot")
	if err != nil {
		t.Error(err)
	}
	t.Log(tableList)

}

func TestGetDatabaseInstance(t *testing.T) {
	db := DB()
	rows, err := db.Query("select * from device_log LIMIT 10")
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var log iotModel.DeviceLog
		err = rows.Scan(&log.Ts, &log.Type, &log.Content, &log.Device)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(log)
	}

	num, err := DB().Count("device_log")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("device 日志数：", num)

}

func TestCount(t *testing.T) {
	num, err := DB().Count("device_log")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(num)
}

// 插入单条日志数据
func TestInsertLog(t *testing.T) {
	logData := iotModel.DeviceLog{
		Device:  "yxjc012311290004983",
		Type:    "属性上报",
		Content: `{"device_id":"yxjc012311290004983","return_time":"2022-11-10 10:49:33","property_99":2,"property_98":2,"property_97":3,"property_96":4,"property_95":2}`,
	}
	_, err := DB().InsertLogData(logData)
	if err != nil {
		t.Fatal(err)
	}
}

// 批量插入日志数据
func TestBatchInsertLogData(t *testing.T) {
	var logList []iotModel.DeviceLog
	for i := 0; i < 10; i++ {
		logData := iotModel.DeviceLog{
			Device:  "yxjc012311290004983",
			Type:    "属性上报",
			Content: `{"device_id":"yxjc012311290004983","return_time":"2022-11-10 10:49:33","property_99":2,"property_98":2,"property_97":3,"property_96":4,"property_95":2}`,
		}
		logList = append(logList, logData)
	}

	var upData = make(map[string][]iotModel.DeviceLog)
	upData["yxjc012311290004983"] = logList
	num, err := DB().BatchInsertLogData(upData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("插入数据条数：", num)
}

// 获取设备数据
func TestGetDeviceData(t *testing.T) {
	db := DB()

	sqlStr := fmt.Sprintf("select * from device_%s", "t202200001")
	rows, err := db.Query(sqlStr)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	var rs gdb.Result

	for rows.Next() {
		values := make([]any, len(columns))
		for i := range values {
			values[i] = new(any)
		}

		err = rows.Scan(values...)
		if err != nil {
			t.Error(err)
		}

		m := make(gdb.Record, len(columns))
		for i, c := range columns {
			// 去除前缀
			if c[:2] == consts.TdPropertyPrefix {
				c = c[2:]
			}
			m[c] = toTime(gvar.New(values[i]))
		}
		rs = append(rs, m)
	}
	t.Log(rs)
}

// Time REST连接时区处理
func toTime(v *g.Var) (rs *g.Var) {
	driver := g.Cfg().MustGet(context.TODO(), "tdengine.type")
	if driver.String() == "taosRestful" {
		if t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", v.String()); err == nil {
			rs = gvar.New(t.Local().Format("2006-01-02 15:04:05"))
			return
		}
	}

	rs = v
	return
}
