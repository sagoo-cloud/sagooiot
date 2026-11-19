package iotdb

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/apache/iotdb-client-go/client"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/dcache"
	_ "sagooiot/pkg/tsd"
	"sagooiot/pkg/tsd/comm"
)

const (
	iotdbDatabaseType = "IoTDB"
	iotdbDriverName   = "iotdb"
)

type engine struct {
	mu sync.Mutex
	db *sql.DB
}

type tslTable struct{}
type logTable struct{}

func init() {
	if strings.EqualFold(configString(context.Background(), "tsd.database", "TdEngine"), iotdbDatabaseType) {
		registerServices()
	}
}

func registerServices() {
	service.RegisterTdEngine(&engine{})
	service.RegisterTSLTable(&tslTable{})
	service.RegisterTdLogTable(&logTable{})
}

func configString(ctx context.Context, key, fallback string) string {
	value, err := g.Cfg().Get(ctx, key, fallback)
	if err != nil || value == nil {
		return fallback
	}
	return value.String()
}

func (e *engine) GetConn(ctx context.Context, _ string) (*sql.DB, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.db != nil {
		return e.db, nil
	}
	host := g.Cfg().MustGet(ctx, "tsd.iotdb.host", "127.0.0.1").String()
	port := g.Cfg().MustGet(ctx, "tsd.iotdb.port", 6667).Int()
	username := g.Cfg().MustGet(ctx, "tsd.iotdb.username", "root").String()
	password := g.Cfg().MustGet(ctx, "tsd.iotdb.password", "root").String()
	if host == "" {
		host = "127.0.0.1"
	}
	if port <= 0 {
		port = 6667
	}
	dsn := (&url.URL{
		Scheme: iotdbDriverName,
		Host:   net.JoinHostPort(host, strconv.Itoa(port)),
		User:   url.UserPassword(username, password),
		RawQuery: url.Values{
			"connectTimeoutMs": {"10000"},
			"timeZone":         {"UTC"},
		}.Encode(),
	}).String()
	db, err := sql.Open(iotdbDriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("open IoTDB: %w", err)
	}
	maxIdle := g.Cfg().MustGet(ctx, "tsd.iotdb.maxIdleConns", 8).Int()
	maxOpen := g.Cfg().MustGet(ctx, "tsd.iotdb.maxOpenConns", 32).Int()
	if maxIdle > 0 {
		db.SetMaxIdleConns(maxIdle)
	}
	if maxOpen > 0 {
		db.SetMaxOpenConns(maxOpen)
	}
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("connect to IoTDB: %w", err)
	}
	e.db = db
	return db, nil
}

func (e *engine) Time(value *g.Var) *g.Var { return value }

func (e *engine) ClearLogByDays(ctx context.Context, days int) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	cutoff := gtime.Now().AddDate(0, 0, -days).UnixMilli()
	_, err = db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s.logs.** WHERE time < ?", rootPath()), cutoff)
	return err
}

func (t *tslTable) CreateDatabase(ctx context.Context) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	return comm.CreateIoTDBSchema(ctx, db, "CREATE DATABASE "+rootPath(), client.DatabaseAlreadyExists)
}

func (t *tslTable) CreateStable(ctx context.Context, tsl *model.TSL) error {
	if tsl == nil || tsl.Key == "" {
		return errors.New("IoTDB product schema requires a product key")
	}
	if err := t.CreateDatabase(ctx); err != nil {
		return err
	}
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	base := schemaProductPath(tsl.Key)
	for _, property := range tsl.Properties {
		field := comm.PropertyColumn(property.Key, true)
		if err := createSeries(ctx, db, base+"."+nodeName(field), tslType(property.ValueType.Type)); err != nil {
			return err
		}
		if err := createSeries(ctx, db, base+"."+nodeName(field+"_time"), "INT64"); err != nil {
			return err
		}
	}
	for _, tag := range tsl.Tags {
		field := "t_" + tag.Key
		if err := createSeries(ctx, db, base+".tags."+nodeName(field), tslType(tag.ValueType.Type)); err != nil {
			return err
		}
	}
	return nil
}

func (t *tslTable) CreateTable(ctx context.Context, stable, table string) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	schema := schemaProductPath(stable)
	rows, err := db.QueryContext(ctx, "SHOW TIMESERIES "+schema+".**")
	if err != nil {
		return fmt.Errorf("read IoTDB product schema %q: %w", stable, err)
	}
	series, err := readSeries(rows)
	if err != nil {
		return err
	}
	if len(series) == 0 {
		return fmt.Errorf("IoTDB schema for product %q has no measurements", stable)
	}
	path := devicePath(table)
	for _, item := range series {
		measurement := item.name[strings.LastIndex(item.name, ".")+1:]
		if err := createSeries(ctx, db, path+"."+measurement, item.dataType); err != nil {
			return err
		}
	}
	return nil
}

func (t *tslTable) Insert(ctx context.Context, deviceKey string, data model.ReportPropertyData, subKey ...string) error {
	if len(data) == 0 {
		return errors.New("reported IoTDB property data is empty")
	}
	target := deviceKey
	if len(subKey) > 0 && subKey[0] != "" {
		target = subKey[0]
	}
	device, err := dcache.GetDeviceDetailInfo(target)
	if err != nil || device == nil || device.Product == nil {
		device, err = service.DevDevice().Detail(ctx, target)
		if err != nil {
			return fmt.Errorf("resolve IoTDB device product for %q: %w", target, err)
		}
	}
	if device == nil || device.Product == nil {
		return fmt.Errorf("IoTDB device %q has no product metadata", target)
	}
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return fmt.Errorf("get IoTDB connection: %w", err)
	}
	path := devicePath(target)
	values := make(map[string]any, len(data)*2)
	declaredTypes := make(map[string]string, len(data)*2)
	for key, node := range data {
		field := comm.PropertyColumn(key, true)
		values[field] = normalizeValue(node.Value)
		values[field+"_time"] = comm.IoTDBReportMillis(node.CreateTime)
		if device.TSL != nil {
			for _, property := range device.TSL.Properties {
				if property.Key == key {
					declaredTypes[field] = tslType(property.ValueType.Type)
					break
				}
			}
		}
		declaredTypes[field+"_time"] = "INT64"
	}
	if device.TSL != nil {
		for _, tag := range device.TSL.Tags {
			for _, tagValue := range device.Tags {
				if tagValue != nil && tagValue.Key == tag.Key {
					field := "t_" + tag.Key
					value, err := parseTagValue(tagValue.Value, tag.ValueType.Type)
					if err != nil {
						return fmt.Errorf("parse IoTDB tag %q: %w", tag.Key, err)
					}
					values[field] = value
					declaredTypes[field] = tslType(tag.ValueType.Type)
					break
				}
			}
		}
	}
	return insertRow(ctx, db, path, time.Now().UnixMilli(), values, declaredTypes)
}

func (t *tslTable) DropStable(ctx context.Context, stable string) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	paths := []string{schemaProductPath(stable) + ".**"}
	devices, err := productDeviceKeys(ctx, stable)
	if err != nil {
		return err
	}
	for _, device := range devices {
		paths = append(paths, devicePath(device)+".**")
	}
	for _, path := range paths {
		if _, err := db.ExecContext(ctx, "DELETE TIMESERIES "+path); err != nil {
			return err
		}
	}
	return nil
}

func (t *tslTable) DropTable(ctx context.Context, table string) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	return deleteTimeseries(ctx, db, devicePath(table)+".**")
}

func (t *tslTable) CheckStable(ctx context.Context, stable string) (bool, error) {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return false, err
	}
	// The platform passes the formatted table name to its existence checks.
	rows, err := db.QueryContext(ctx, "SHOW TIMESERIES "+rootPath()+".schemas."+nodeName(stable)+".**")
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), rows.Err()
}

func (t *tslTable) CheckTable(ctx context.Context, table string) (bool, error) {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return false, err
	}
	rows, err := db.QueryContext(ctx, "SHOW TIMESERIES "+rootPath()+"."+nodeName(table)+".**")
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), rows.Err()
}

func (t *tslTable) AddDatabaseField(ctx context.Context, product, field, dataType string, _ int) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	measurement := "p_" + nodeName(field)
	typeName := tslType(dataType)
	if err := createSeries(ctx, db, schemaProductPath(product)+"."+measurement, typeName); err != nil {
		return err
	}
	if err := createSeries(ctx, db, schemaProductPath(product)+"."+measurement+"_time", "INT64"); err != nil {
		return err
	}
	devices, err := productDeviceKeys(ctx, product)
	if err != nil {
		return err
	}
	for _, device := range devices {
		if err := createSeries(ctx, db, devicePath(device)+"."+measurement, typeName); err != nil {
			return err
		}
		if err := createSeries(ctx, db, devicePath(device)+"."+measurement+"_time", "INT64"); err != nil {
			return err
		}
	}
	return nil
}

func (t *tslTable) DelDatabaseField(ctx context.Context, product, field string) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	measurement := "p_" + nodeName(field)
	devices, err := productDeviceKeys(ctx, product)
	if err != nil {
		return err
	}
	paths := []string{schemaProductPath(product) + "." + measurement, schemaProductPath(product) + "." + measurement + "_time"}
	for _, device := range devices {
		path := devicePath(device)
		paths = append(paths, path+"."+measurement, path+"."+measurement+"_time")
	}
	return deleteTimeseries(ctx, db, paths...)
}

func (t *tslTable) ModifyDatabaseField(ctx context.Context, product, field, dataType string, _ int) error {
	// IoTDB Tree schema has no field-length property. Type changes are handled by the product schema migration path.
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	rows, err := db.QueryContext(ctx, "SHOW TIMESERIES "+schemaProductPath(product)+".p_"+nodeName(field))
	if err != nil {
		return err
	}
	series, readErr := readSeries(rows)
	if readErr != nil {
		return readErr
	}
	if len(series) == 0 {
		return fmt.Errorf("IoTDB property %q does not exist for product %q", field, product)
	}
	if series[0].dataType != tslType(dataType) {
		return fmt.Errorf("IoTDB cannot change a timeseries data type in place (%s to %s); migrate the product schema and stored data", series[0].dataType, tslType(dataType))
	}
	return nil
}

func (t *tslTable) AddTag(ctx context.Context, product, tag, dataType string, _ int) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	typeName := tslType(dataType)
	if err := createSeries(ctx, db, schemaProductPath(product)+".tags.t_"+nodeName(tag), typeName); err != nil {
		return err
	}
	devices, err := productDeviceKeys(ctx, product)
	if err != nil {
		return err
	}
	for _, device := range devices {
		if err := createSeries(ctx, db, devicePath(device)+".t_"+nodeName(tag), typeName); err != nil {
			return err
		}
	}
	return nil
}

func (t *tslTable) DelTag(ctx context.Context, product, tag string) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	paths := []string{schemaProductPath(product) + ".tags.t_" + nodeName(tag)}
	devices, err := productDeviceKeys(ctx, product)
	if err != nil {
		return err
	}
	for _, device := range devices {
		paths = append(paths, devicePath(device)+".t_"+nodeName(tag))
	}
	return deleteTimeseries(ctx, db, paths...)
}

func (t *tslTable) ModifyTag(ctx context.Context, product, tag, dataType string, _ int) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	rows, err := db.QueryContext(ctx, "SHOW TIMESERIES "+schemaProductPath(product)+".tags.t_"+nodeName(tag))
	if err != nil {
		return err
	}
	series, err := readSeries(rows)
	if err != nil {
		return err
	}
	if len(series) == 0 {
		return fmt.Errorf("IoTDB tag %q does not exist for product %q", tag, product)
	}
	if series[0].dataType != tslType(dataType) {
		return fmt.Errorf("IoTDB cannot change a timeseries data type in place (%s to %s); migrate the product schema and stored data", series[0].dataType, tslType(dataType))
	}
	return nil
}

func (t *logTable) CreateStable(ctx context.Context) error {
	return t.CreateDatabase(ctx)
}

func (t *logTable) CreateDatabase(ctx context.Context) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	return comm.CreateIoTDBSchema(ctx, db, "CREATE DATABASE "+rootPath(), client.DatabaseAlreadyExists)
}

func (t *logTable) Insert(ctx context.Context, log *model.TdLogAddInput) error {
	if log == nil {
		return errors.New("IoTDB log cannot be nil")
	}
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	path := logDevicePath(log.Device)
	if err := createSeries(ctx, db, path+".log_type", "TEXT"); err != nil {
		return err
	}
	if err := createSeries(ctx, db, path+".content", "TEXT"); err != nil {
		return err
	}
	timestamp := time.Now().UnixMilli()
	if log.Ts != nil {
		timestamp = log.Ts.UnixMilli()
	}
	_, err = db.ExecContext(ctx, "INSERT INTO "+path+" (time, log_type, content) VALUES (?, ?, ?)", timestamp, log.Type, log.Content)
	return err
}

func (t *logTable) Clear(ctx context.Context) error {
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return err
	}
	cutoff := time.Now().Add(-7 * 24 * time.Hour).UnixMilli()
	_, err = db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s.logs.** WHERE time < ?", rootPath()), cutoff)
	return err
}

func (t *logTable) GetAll(ctx context.Context, _ string, args ...any) ([]model.TdLog, error) {
	if len(args) == 0 {
		return nil, errors.New("IoTDB device log query requires a device key")
	}
	device := fmt.Sprint(args[0])
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		return nil, err
	}
	query := "SELECT log_type, content FROM " + logDevicePath(device) + " ORDER BY TIME DESC LIMIT 100"
	if len(args) >= 3 {
		query = "SELECT log_type, content FROM " + logDevicePath(device) + " ORDER BY TIME DESC LIMIT ? OFFSET ?"
	}
	var rows *sql.Rows
	if len(args) >= 3 {
		rows, err = db.QueryContext(ctx, query, args[2], args[1])
	} else {
		rows, err = db.QueryContext(ctx, query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []model.TdLog
	for rows.Next() {
		var item model.TdLog
		var ts time.Time
		if err := rows.Scan(&ts, &item.Type, &item.Content); err != nil {
			return nil, err
		}
		item.Ts = gtime.New(ts)
		item.Device = device
		result = append(result, item)
	}
	return result, rows.Err()
}

type seriesInfo struct {
	name     string
	dataType string
}

func insertRow(ctx context.Context, db *sql.DB, path string, timestamp int64, values map[string]any, declaredTypes map[string]string) error {
	if db == nil {
		return errors.New("IoTDB connection is not initialized")
	}
	if len(values) == 0 {
		return errors.New("IoTDB property data is empty")
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		dataType := declaredTypes[key]
		if dataType == "" {
			dataType = inferType(values[key])
		}
		if err := createSeries(ctx, db, path+"."+nodeName(key), dataType); err != nil {
			return err
		}
	}
	columns := []string{"time"}
	placeholders := []string{"?"}
	args := []any{timestamp}
	for _, key := range keys {
		columns = append(columns, nodeName(key))
		placeholders = append(placeholders, "?")
		args = append(args, values[key])
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", path, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err := db.ExecContext(ctx, query, args...)
	return err
}

func createSeries(ctx context.Context, db *sql.DB, path, dataType string) error {
	err := comm.CreateIoTDBSchema(ctx, db, fmt.Sprintf("CREATE TIMESERIES %s WITH DATATYPE=%s, ENCODING=%s, COMPRESSOR=SNAPPY", path, dataType, encodingFor(dataType)), client.TimeseriesAlreadyExist)
	if err != nil {
		return fmt.Errorf("create IoTDB timeseries %s: %w", path, err)
	}
	return nil
}

func deleteTimeseries(ctx context.Context, db *sql.DB, paths ...string) error {
	for _, path := range paths {
		if _, err := db.ExecContext(ctx, "DELETE TIMESERIES "+path); err != nil {
			return err
		}
	}
	return nil
}

func readSeries(rows *sql.Rows) ([]seriesInfo, error) {
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	nameIndex, typeIndex := 0, 3
	for index, column := range columns {
		switch strings.ToLower(strings.ReplaceAll(column, "_", "")) {
		case "timeseries", "path", "measurement":
			nameIndex = index
		case "datatype", "type":
			typeIndex = index
		}
	}
	var result []seriesInfo
	for rows.Next() {
		values := make([]any, len(columns))
		args := make([]any, len(columns))
		for index := range values {
			args[index] = &values[index]
		}
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		if nameIndex >= len(values) || typeIndex >= len(values) {
			continue
		}
		result = append(result, seriesInfo{name: comm.IoTDBMetadataText(values[nameIndex]), dataType: strings.ToUpper(comm.IoTDBMetadataText(values[typeIndex]))})
	}
	return result, rows.Err()
}

func rootPath() string {
	return comm.IoTDBRoot(configString(context.Background(), "tsd.iotdb.database", "root.sagoo_iot"))
}

func schemaProductPath(product string) string {
	return rootPath() + ".schemas." + productPath(product)
}

func devicePath(device string) string {
	return comm.IoTDBDevicePath(rootPath(), device)
}

func productDeviceKeys(ctx context.Context, product string) ([]string, error) {
	devices, err := service.DevDevice().GetAllForProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(devices))
	for _, device := range devices {
		if device != nil && device.Key != "" {
			keys = append(keys, device.Key)
		}
	}
	return keys, nil
}

func logDevicePath(device string) string {
	return rootPath() + ".logs.device_" + nodeName(device)
}

func productPath(product string) string { return nodeName(comm.ProductTableName(product)) }

func nodeName(value string) string {
	return comm.IoTDBNode(value)
}

func tslType(valueType string) string {
	switch strings.ToLower(valueType) {
	case "int":
		return "INT32"
	case "long":
		return "INT64"
	case "float":
		return "FLOAT"
	case "double":
		return "DOUBLE"
	case "boolean", "bool":
		return "BOOLEAN"
	case "timestamp":
		return "TIMESTAMP"
	case "date":
		return "DATE"
	default:
		return "TEXT"
	}
}

func encodingFor(dataType string) string {
	switch dataType {
	case "BOOLEAN":
		return "RLE"
	case "INT32", "INT64", "TIMESTAMP", "DATE":
		return "TS_2DIFF"
	case "FLOAT", "DOUBLE":
		return "GORILLA"
	default:
		return "PLAIN"
	}
}

func inferType(value any) string {
	if value == nil {
		return "TEXT"
	}
	switch value.(type) {
	case bool:
		return "BOOLEAN"
	case int, int8, int16, int32, uint, uint8, uint16, uint32:
		return "INT32"
	case int64, uint64:
		return "INT64"
	case float32:
		return "FLOAT"
	case float64:
		return "DOUBLE"
	case time.Time, *gtime.Time:
		return "TIMESTAMP"
	default:
		return "TEXT"
	}
}

func normalizeValue(value any) any {
	switch value.(type) {
	case map[string]any, []any:
		encoded, _ := json.Marshal(value)
		return string(encoded)
	default:
		return value
	}
}

func parseTagValue(value, dataType string) (any, error) {
	switch strings.ToLower(dataType) {
	case "int":
		parsed, err := strconv.ParseInt(value, 10, 32)
		return int32(parsed), err
	case "long":
		return strconv.ParseInt(value, 10, 64)
	case "float":
		parsed, err := strconv.ParseFloat(value, 32)
		return float32(parsed), err
	case "double":
		return strconv.ParseFloat(value, 64)
	case "boolean", "bool":
		return strconv.ParseBool(value)
	default:
		return value, nil
	}
}

var _ service.ITdEngine = (*engine)(nil)
var _ service.ITSLTable = (*tslTable)(nil)
var _ service.ITdLogTable = (*logTable)(nil)
