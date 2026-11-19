package iotdb

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/tsd/comm"
)

// Run against an isolated local Apache IoTDB using SAGOO_IOTDB_TEST_PORT.
func TestLiveTreeDeviceQueries(t *testing.T) {
	portString := os.Getenv("SAGOO_IOTDB_TEST_PORT")
	if portString == "" {
		t.Skip("SAGOO_IOTDB_TEST_PORT is unset")
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	root := fmt.Sprintf("root.sagoo_test_%d", time.Now().UnixNano())
	m := &Database{Option: comm.Option{Host: "127.0.0.1", Port: port, Database: root, Username: "root", Password: "root"}}
	defer m.Close()
	db, err := m.connect()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "CREATE DATABASE "+root); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if _, err := db.ExecContext(ctx, "DELETE DATABASE "+root); err != nil {
			t.Error(err)
		}
	}()
	start := time.Now().UTC().Truncate(24 * time.Hour)
	end := start.AddDate(0, 0, 1)
	for i, seconds := range []int64{1, 2, 3, 61} {
		value := []int64{21, 22, 21, 25}[i]
		ms := start.UnixMilli() + seconds*1000
		if err := m.insertRow(ctx, m.devicePath("meter-1"), ms, map[string]any{"p_temp": value, "p_temp_time": ms}, nil); err != nil {
			t.Fatal(err)
		}
	}
	latest, err := m.GetTableDataOne(ctx, "SELECT p_temp FROM device_meter_1 WHERE p_temp IS NOT NULL ORDER BY p_temp_time DESC LIMIT 1")
	if err != nil || latest["temp"].Int() != 25 {
		t.Fatalf("latest = %v (%v)", latest, err)
	}
	series, err := m.GetTableDataAll(ctx, "SELECT p_temp_time AS ts, p_temp FROM device_meter_1 WHERE p_temp_time >= ? AND p_temp_time < ? ORDER BY p_temp_time ASC", start.UnixMilli()+2000, start.UnixMilli()+61000)
	if err != nil || len(series) != 2 || series[0]["temp"].Int() != 22 || series[1]["ts"].Int64() != start.UnixMilli()+3000 {
		t.Fatalf("reported range = %v (%v)", series, err)
	}
	for _, ranges := range [][2]time.Time{{start, end}, {time.Time{}, time.Time{}}} {
		windows, total, err := m.DeviceWindows(ctx, "meter-1", []string{"p_temp"}, 1, 2, ranges[0], ranges[1], 2, 1)
		if err != nil || total != 2 || len(windows) != 1 || windows[0]["temp"].Int() != 21 || !windows[0]["_wstart"].Time().Equal(start) || !windows[0]["_wend"].Time().Equal(start.Add(time.Minute)) {
			t.Fatalf("mode windows = %v, total = %v (%v)", windows, total, err)
		}
	}
	if err := m.InsertDeviceData("gateway", iotModel.ReportPropertyData{"humidity": {Value: 47.5, CreateTime: start.Unix()}}, "child-1"); err != nil {
		t.Fatal(err)
	}
	child, err := m.GetTableDataOne(ctx, "SELECT p_humidity, p_humidity_time FROM device_child_1")
	if err != nil || child["humidity"].Float64() != 47.5 || child["humidity_time"].Int64() != start.UnixMilli() {
		t.Fatalf("child data = %v (%v)", child, err)
	}
	for _, key := range []string{"meter-1", "child-1", "unselected"} {
		if _, err := m.InsertLogData(iotModel.DeviceLog{Device: key, Type: "report", Content: "ready", Ts: gtime.New(start)}); err != nil {
			t.Fatal(err)
		}
	}
	logs, err := m.CountLogs(ctx, []string{"meter-1", "child-1"}, start, end)
	if err != nil || logs != 2 {
		t.Fatalf("log count = %v (%v)", logs, err)
	}
}
