package iotdb

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/tsd/comm"
)

type memoryStore struct {
	mu     sync.Mutex
	schema map[string]string
	rows   map[string][]map[string]any
}

type memoryConnector struct{ store *memoryStore }

func (c memoryConnector) Connect(context.Context) (driver.Conn, error) {
	return &memoryConn{store: c.store}, nil
}

func (c memoryConnector) Driver() driver.Driver { return memoryDriver{store: c.store} }

type memoryDriver struct{ store *memoryStore }

func (d memoryDriver) Open(string) (driver.Conn, error) {
	return &memoryConn{store: d.store}, nil
}

type memoryConn struct{ store *memoryStore }

func (c *memoryConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("prepared statements are not supported by the test driver")
}

func (*memoryConn) Close() error { return nil }

func (*memoryConn) Begin() (driver.Tx, error) {
	return nil, errors.New("transactions are not supported by the test driver")
}

func (c *memoryConn) ExecContext(_ context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	bound, err := bindSQL(query, args)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(bound, "CREATE TIMESERIES ") {
		parts := strings.Fields(bound)
		if len(parts) != 7 || parts[3] != "WITH" {
			return nil, fmt.Errorf("invalid create timeseries query %q", bound)
		}
		c.store.mu.Lock()
		c.store.schema[parts[2]] = strings.TrimSuffix(strings.TrimPrefix(parts[4], "DATATYPE="), ",")
		c.store.mu.Unlock()
		return driver.RowsAffected(0), nil
	}
	if strings.HasPrefix(bound, "INSERT INTO ") {
		return c.insert(bound)
	}
	return nil, fmt.Errorf("unsupported test-driver exec: %s", bound)
}

func (c *memoryConn) insert(query string) (driver.Result, error) {
	rest := strings.TrimPrefix(query, "INSERT INTO ")
	openColumns := strings.Index(rest, " (")
	closeColumns := strings.Index(rest, ") VALUES (")
	if openColumns < 0 || closeColumns < openColumns {
		return nil, fmt.Errorf("invalid insert query %q", query)
	}
	path := rest[:openColumns]
	columns := strings.Split(rest[openColumns+2:closeColumns], ", ")
	valuesText := strings.TrimSuffix(rest[closeColumns+len(") VALUES ("):], ")")
	values, err := splitSQLValues(valuesText)
	if err != nil || len(values) != len(columns) {
		return nil, fmt.Errorf("parse insert values: %v", err)
	}
	row := make(map[string]any, len(columns))
	for index, column := range columns {
		row[column] = parseSQLValue(values[index])
	}
	c.store.mu.Lock()
	c.store.rows[path] = append(c.store.rows[path], row)
	c.store.mu.Unlock()
	return driver.RowsAffected(1), nil
}

func (c *memoryConn) QueryContext(_ context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	bound, err := bindSQL(query, args)
	if err != nil {
		return nil, err
	}
	selectIndex := strings.Index(strings.ToLower(bound), "select ")
	fromIndex := strings.Index(strings.ToLower(bound), " from ")
	if selectIndex < 0 || fromIndex < 0 {
		return nil, fmt.Errorf("unsupported test-driver query: %s", bound)
	}
	measurement := strings.TrimSpace(bound[selectIndex+len("select ") : fromIndex])
	remaining := strings.TrimSpace(bound[fromIndex+len(" from "):])
	path := strings.Fields(remaining)[0]
	c.store.mu.Lock()
	defer c.store.mu.Unlock()
	if _, exists := c.store.schema[path+"."+measurement]; !exists {
		return nil, fmt.Errorf("timeseries %s.%s does not exist", path, measurement)
	}
	rows := c.store.rows[path]
	for index := len(rows) - 1; index >= 0; index-- {
		value, exists := rows[index][measurement]
		if !exists || value == nil {
			continue
		}
		timestamp := rows[index]["time"]
		return &memoryRows{
			columns: []string{"Time", path + "." + measurement},
			values:  []driver.Value{timestamp.(driver.Value), value.(driver.Value)},
		}, nil
	}
	return &memoryRows{columns: []string{"Time", measurement}}, nil
}

type memoryRows struct {
	columns []string
	values  []driver.Value
	done    bool
}

func (r *memoryRows) Columns() []string { return r.columns }
func (*memoryRows) Close() error        { return nil }
func (r *memoryRows) Next(dest []driver.Value) error {
	if r.done || len(r.values) == 0 {
		return io.EOF
	}
	r.done = true
	copy(dest, r.values)
	return nil
}

func splitSQLValues(values string) ([]string, error) {
	var result []string
	start := 0
	quoted := false
	for index := 0; index < len(values); index++ {
		if values[index] == '\'' {
			if quoted && index+1 < len(values) && values[index+1] == '\'' {
				index++
				continue
			}
			quoted = !quoted
		} else if values[index] == ',' && !quoted {
			result = append(result, strings.TrimSpace(values[start:index]))
			start = index + 1
		}
	}
	if quoted {
		return nil, errors.New("unterminated string literal")
	}
	return append(result, strings.TrimSpace(values[start:])), nil
}

func parseSQLValue(value string) any {
	if len(value) >= 2 && value[0] == '\'' && value[len(value)-1] == '\'' {
		return strings.ReplaceAll(value[1:len(value)-1], "''", "'")
	}
	if integer, err := strconv.ParseInt(value, 10, 64); err == nil {
		return integer
	}
	if decimal, err := strconv.ParseFloat(value, 64); err == nil {
		return decimal
	}
	return value
}

func TestPlatformQueryReadsSchemaWrittenDeviceData(t *testing.T) {
	store := &memoryStore{schema: make(map[string]string), rows: make(map[string][]map[string]any)}
	db := &Database{Option: comm.Option{Database: "root.sagoo_iot"}}
	db.db = sql.OpenDB(memoryConnector{store: store})
	defer db.Close()

	ctx := context.Background()
	path := db.devicePath("sensor-1")
	if err := db.ensureTimeseries(ctx, path, "p_temperature", "DOUBLE"); err != nil {
		t.Fatal(err)
	}
	if err := db.InsertDeviceData("gateway", iotModel.ReportPropertyData{"temperature": {Value: 21.5, CreateTime: 1763460000}}, "sensor-1"); err != nil {
		t.Fatal(err)
	}

	result, err := db.GetTableDataOne(ctx, "SELECT p_temperature FROM device_sensor_1 WHERE p_temperature IS NOT NULL ORDER BY time DESC LIMIT 1")
	if err != nil {
		t.Fatal(err)
	}
	if result["temperature"].Float64() != 21.5 {
		t.Fatalf("platform query result = %#v, want temperature 21.5", result)
	}
	reported, err := db.GetTableDataOne(ctx, "SELECT p_temperature_time FROM device_sensor_1")
	if err != nil || reported["temperature_time"].Int64() != 1763460000000 || store.schema[path+".p_temperature_time"] != "INT64" {
		t.Fatalf("report time read back = %v, schema = %v, error = %v", reported, store.schema, err)
	}
	if got, want := qualifyFromTables("SELECT p_temperature FROM device_sensor_1", db.rootPath()), "SELECT p_temperature FROM root.sagoo_iot.device_sensor_1"; got != want {
		t.Fatalf("qualified query = %q, want %q", got, want)
	}
}

func TestQualifyFromTablesSkipsStringsAndComments(t *testing.T) {
	query := "SELECT 'FROM imaginary' AS text FROM device_1 -- FROM comment\n/* FROM block */"
	want := "SELECT 'FROM imaginary' AS text FROM root.sagoo_iot.device_1 -- FROM comment\n/* FROM block */"
	if got := qualifyFromTables(query, "root.sagoo_iot"); got != want {
		t.Fatalf("qualified query = %q, want %q", got, want)
	}
}

var _ driver.Connector = memoryConnector{}
var _ driver.Driver = memoryDriver{}
var _ driver.Conn = (*memoryConn)(nil)
var _ driver.ExecerContext = (*memoryConn)(nil)
var _ driver.QueryerContext = (*memoryConn)(nil)
var _ driver.Rows = (*memoryRows)(nil)

func TestDeviceWindowArgumentsAreCheckedBeforeQuery(t *testing.T) {
	db := new(Database)
	start := time.Date(2026, 11, 18, 0, 0, 0, 0, time.UTC)
	_, _, err := db.DeviceWindows(context.Background(), "meter", []string{"p_t"}, 0, 1, start, start.Add(time.Hour), 1, 10)
	if err == nil {
		t.Fatal("zero-width window accepted")
	}
}

func TestDeviceWindowStepsFollowPlatformUnits(t *testing.T) {
	for unit, want := range map[uint]int64{1: 1000, 2: 60_000, 3: 3_600_000, 4: 86_400_000} {
		got, err := deviceWindowStep(1, unit)
		if err != nil || got != want {
			t.Errorf("deviceWindowStep(1, %d) = %d, %v; want %d", unit, got, err, want)
		}
	}
	for _, unit := range []uint{0, 5} {
		if _, err := deviceWindowStep(1, unit); err == nil {
			t.Errorf("deviceWindowStep accepted unsupported unit %d", unit)
		}
	}
	if _, err := deviceWindowStep(0, 1); err == nil {
		t.Fatal("deviceWindowStep accepted a zero interval")
	}
	if _, err := deviceWindowStep(uint(math.MaxInt64/86_400_000+1), 4); err == nil {
		t.Fatal("deviceWindowStep accepted an overflowing day interval")
	}
}

func TestWindowOriginFloorsPreEpochTimestamps(t *testing.T) {
	for timestamp, want := range map[int64]int64{-1001: -2000, -1000: -1000, -1: -1000, 0: 0, 1001: 1000} {
		if got := windowOrigin(timestamp, 1000); got != want {
			t.Errorf("windowOrigin(%d, 1000) = %d, want %d", timestamp, got, want)
		}
	}
}
