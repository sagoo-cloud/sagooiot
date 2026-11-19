package iotdb

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/apache/iotdb-client-go/client"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/tsd/comm"
)

type Database struct {
	Option comm.Option

	mu sync.Mutex
	db *sql.DB
}

func (m *Database) connect() (*sql.DB, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.db != nil {
		return m.db, nil
	}
	host := m.Option.Host
	if host == "" {
		host = "127.0.0.1"
	}
	conn, err := sql.Open(driverName, BuildDSN(host, m.Option.Port, m.Option.Username, m.Option.Password))
	if err != nil {
		return nil, fmt.Errorf("open IoTDB database: %w", err)
	}
	if m.Option.MaxIdleConns > 0 {
		conn.SetMaxIdleConns(m.Option.MaxIdleConns)
	}
	if m.Option.MaxOpenConns > 0 {
		conn.SetMaxOpenConns(m.Option.MaxOpenConns)
	}
	m.db = conn
	return m.db, nil
}

func (m *Database) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.db == nil {
		return
	}
	_ = m.db.Close()
	m.db = nil
}

func (m *Database) Query(query string) (*sql.Rows, error) {
	db, err := m.connect()
	if err != nil {
		return nil, err
	}
	return db.QueryContext(context.Background(), qualifyFromTables(query, m.rootPath()))
}

func (m *Database) Count(path string) (int, error) {
	db, err := m.connect()
	if err != nil {
		return 0, err
	}
	var count int
	queryPath := safePath(path, m.rootPath())
	if !strings.HasPrefix(queryPath, m.rootPath()+".") {
		queryPath = m.rootPath() + "." + queryPath
	}
	query := "SELECT count(*) FROM " + queryPath
	if err := db.QueryRow(query).Scan(&count); err != nil {
		return 0, fmt.Errorf("count IoTDB path %q: %w", path, err)
	}
	return count, nil
}

func (m *Database) InsertLogData(log iotModel.DeviceLog) (sql.Result, error) {
	path := m.logDevicePath(log.Device)
	if err := m.ensureTimeseries(context.Background(), path, "log_type", "TEXT"); err != nil {
		return nil, err
	}
	if err := m.ensureTimeseries(context.Background(), path, "content", "TEXT"); err != nil {
		return nil, err
	}
	db, err := m.connect()
	if err != nil {
		return nil, err
	}
	timestamp := time.Now().UnixMilli()
	if log.Ts != nil {
		timestamp = log.Ts.UnixMilli()
	}
	query := fmt.Sprintf("INSERT INTO %s(time, log_type, content) VALUES (?, ?, ?)", path)
	return db.Exec(query, timestamp, log.Type, log.Content)
}

func (m *Database) BatchInsertLogData(records map[string][]iotModel.DeviceLog) (int, error) {
	count := 0
	for device, logs := range records {
		for _, log := range logs {
			if log.Device == "" {
				log.Device = device
			}
			if _, err := m.InsertLogData(log); err != nil {
				return count, err
			}
			count++
		}
	}
	return count, nil
}

func (m *Database) InsertDeviceData(deviceKey string, data iotModel.ReportPropertyData, subKey ...string) error {
	if len(data) == 0 {
		return fmt.Errorf("device data is empty")
	}
	if len(subKey) > 0 && subKey[0] != "" {
		deviceKey = subKey[0]
	}
	values := make(map[string]any, len(data)*2)
	for key, node := range data {
		field := comm.PropertyColumn(key, true)
		values[field] = node.Value
		values[field+"_time"] = comm.IoTDBReportMillis(node.CreateTime)
	}
	return m.insertRow(context.Background(), m.devicePath(deviceKey), time.Now().UnixMilli(), values, nil)
}

func (m *Database) BatchInsertDeviceData(deviceKey string, data []iotModel.ReportPropertyData) (int, error) {
	count := 0
	for _, row := range data {
		if err := m.InsertDeviceData(deviceKey, row); err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (m *Database) BatchInsertMultiDeviceData(records map[string][]iotModel.ReportPropertyData) (int, error) {
	count := 0
	for deviceKey, rows := range records {
		inserted, err := m.BatchInsertDeviceData(deviceKey, rows)
		count += inserted
		if err != nil {
			return count, err
		}
	}
	return count, nil
}

func (m *Database) WatchDeviceData(string, func(iotModel.ReportPropertyData)) error {
	return fmt.Errorf("IoTDB change subscription is not implemented by the session query API")
}

func (m *Database) GetAllDatabaseName(ctx context.Context) ([]string, error) {
	db, err := m.connect()
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var result []string
	for rows.Next() {
		values := make([]any, len(columns))
		scan := make([]any, len(columns))
		for i := range values {
			scan[i] = &values[i]
		}
		if err := rows.Scan(scan...); err != nil {
			return nil, err
		}
		if len(values) > 0 && values[0] != nil {
			result = append(result, comm.IoTDBMetadataText(values[0]))
		}
	}
	return result, rows.Err()
}

func (m *Database) GetTableListByDatabase(ctx context.Context, database string) ([]iotModel.TsdTables, error) {
	db, err := m.connect()
	if err != nil {
		return nil, err
	}
	root := m.rootPath()
	if database != "" && strings.HasPrefix(database, "root.") {
		root = safePath(database, root)
	}
	rows, err := db.QueryContext(ctx, "SHOW DEVICES "+root+".**")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var result []iotModel.TsdTables
	for rows.Next() {
		values, err := scanRow(rows, len(columns))
		if err != nil {
			return nil, err
		}
		if len(values) == 0 {
			continue
		}
		path := comm.IoTDBMetadataText(values[0])
		parts := strings.Split(path, ".")
		result = append(result, iotModel.TsdTables{
			TableName:  path,
			DbName:     root,
			StableName: strings.Join(parts[:max(0, len(parts)-1)], "."),
		})
	}
	return result, rows.Err()
}

func (m *Database) GetTableInfo(ctx context.Context, path string) ([]*iotModel.TsdTableInfo, error) {
	db, err := m.connect()
	if err != nil {
		return nil, err
	}
	queryPath := safePath(path, m.rootPath())
	if !strings.HasPrefix(queryPath, m.rootPath()+".") {
		queryPath = m.rootPath() + "." + queryPath
	}
	rows, err := db.QueryContext(ctx, "SHOW TIMESERIES "+queryPath+".**")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var result []*iotModel.TsdTableInfo
	for rows.Next() {
		values, err := scanRow(rows, len(columns))
		if err != nil {
			return nil, err
		}
		if len(values) < 4 {
			continue
		}
		result = append(result, &iotModel.TsdTableInfo{
			Field:  comm.IoTDBMetadataText(values[0]),
			Type:   comm.IoTDBMetadataText(values[3]),
			Length: 0,
		})
	}
	return result, rows.Err()
}

func (m *Database) GetTableData(ctx context.Context, path string) (*iotModel.TsdTableDataInfo, error) {
	db, err := m.connect()
	if err != nil {
		return nil, err
	}
	queryPath := safePath(path, m.rootPath())
	if !strings.HasPrefix(queryPath, m.rootPath()+".") {
		queryPath = m.rootPath() + "." + queryPath
	}
	rows, err := db.QueryContext(ctx, "SELECT * FROM "+queryPath+" LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	result := &iotModel.TsdTableDataInfo{Filed: append([]string(nil), columns...)}
	for rows.Next() {
		values, err := scanRow(rows, len(columns))
		if err != nil {
			return nil, err
		}
		row := make(map[string]interface{}, len(columns))
		for i, column := range columns {
			row[column] = values[i]
		}
		result.Info = append(result.Info, row)
	}
	return result, rows.Err()
}

func (m *Database) GetTableDataOne(ctx context.Context, query string, args ...any) (gdb.Record, error) {
	rows, err := m.queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, rows.Err()
	}
	return readPropertyRecord(rows)
}

func (m *Database) GetTableDataAll(ctx context.Context, query string, args ...any) (gdb.Result, error) {
	rows, err := m.queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result gdb.Result
	for rows.Next() {
		record, err := readPropertyRecord(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, rows.Err()
}

func (m *Database) queryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	db, err := m.connect()
	if err != nil {
		return nil, err
	}
	return db.QueryContext(ctx, qualifyFromTables(query, m.rootPath()), args...)
}

func qualifyFromTables(query, root string) string {
	var out strings.Builder
	out.Grow(len(query) + len(root))
	quote := byte(0)
	lineComment, blockComment := false, false
	for i := 0; i < len(query); {
		current := query[i]
		if lineComment {
			out.WriteByte(current)
			i++
			if current == '\n' {
				lineComment = false
			}
			continue
		}
		if blockComment {
			out.WriteByte(current)
			i++
			if current == '*' && i < len(query) && query[i] == '/' {
				out.WriteByte(query[i])
				i++
				blockComment = false
			}
			continue
		}
		if quote != 0 {
			out.WriteByte(current)
			i++
			if current == quote {
				if i < len(query) && query[i] == quote {
					out.WriteByte(query[i])
					i++
				} else {
					quote = 0
				}
			}
			continue
		}
		if current == '-' && i+1 < len(query) && query[i+1] == '-' {
			out.WriteString("--")
			i += 2
			lineComment = true
			continue
		}
		if current == '/' && i+1 < len(query) && query[i+1] == '*' {
			out.WriteString("/*")
			i += 2
			blockComment = true
			continue
		}
		if current == '\'' || current == '"' || current == '`' {
			quote = current
			out.WriteByte(current)
			i++
			continue
		}
		if isSQLWordStart(current) {
			end := i + 1
			for end < len(query) && isSQLWordPart(query[end]) {
				end++
			}
			word := query[i:end]
			out.WriteString(word)
			i = end
			if !strings.EqualFold(word, "FROM") {
				continue
			}
			spaceStart := i
			for i < len(query) && (query[i] == ' ' || query[i] == '\t' || query[i] == '\r' || query[i] == '\n') {
				i++
			}
			out.WriteString(query[spaceStart:i])
			tableStart := i
			for i < len(query) && isSQLPathPart(query[i]) {
				i++
			}
			if tableStart == i {
				continue
			}
			path := query[tableStart:i]
			if strings.HasPrefix(strings.ToLower(path), "device_") && !strings.Contains(path, ".") {
				path = root + "." + path
			}
			out.WriteString(path)
			continue
		}
		out.WriteByte(current)
		i++
	}
	return out.String()
}

func isSQLWordStart(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value == '_'
}

func isSQLWordPart(value byte) bool {
	return isSQLWordStart(value) || value >= '0' && value <= '9' || value == '$'
}

func isSQLPathPart(value byte) bool {
	return isSQLWordPart(value) || value == '.' || value == '*'
}

func readPropertyRecord(rows *sql.Rows) (gdb.Record, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values, err := scanRow(rows, len(columns))
	if err != nil {
		return nil, err
	}
	result := make(gdb.Record, len(columns))
	for i, column := range columns {
		key := column
		if strings.HasPrefix(key, "root.") {
			key = key[strings.LastIndexByte(key, '.')+1:]
		}
		key = strings.TrimPrefix(key, comm.TdPropertyPrefix)
		result[key] = gvar.New(values[i])
	}
	return result, nil
}

func (m *Database) insertRow(ctx context.Context, path string, timestamp int64, values map[string]any, declaredTypes map[string]string) error {
	if len(values) == 0 {
		return fmt.Errorf("device data is empty")
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	// Keep a deterministic column order so equivalent batches generate stable SQL.
	sort.Strings(keys)
	db, err := m.connect()
	if err != nil {
		return err
	}
	for _, key := range keys {
		dataType := ""
		if declaredTypes != nil {
			dataType = declaredTypes[key]
		}
		if dataType == "" {
			dataType = inferType(values[key])
		}
		if err := createTimeseries(ctx, db, path+"."+iotdbSegment(key), dataType); err != nil {
			return err
		}
	}
	columns := make([]string, 0, len(keys)+1)
	columns = append(columns, "time")
	arguments := make([]any, 0, len(keys)+1)
	arguments = append(arguments, timestamp)
	placeholders := make([]string, 0, len(keys)+1)
	placeholders = append(placeholders, "?")
	for _, key := range keys {
		columns = append(columns, iotdbSegment(key))
		placeholders = append(placeholders, "?")
		arguments = append(arguments, values[key])
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", path, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err = db.ExecContext(ctx, query, arguments...)
	return err
}

func (m *Database) ensureTimeseries(ctx context.Context, path, measurement, dataType string) error {
	db, err := m.connect()
	if err != nil {
		return err
	}
	return createTimeseries(ctx, db, path+"."+iotdbSegment(measurement), dataType)
}

func createTimeseries(ctx context.Context, db *sql.DB, path, dataType string) error {
	err := comm.CreateIoTDBSchema(ctx, db, fmt.Sprintf("CREATE TIMESERIES %s WITH DATATYPE=%s, ENCODING=%s, COMPRESSOR=SNAPPY", path, dataType, encodingFor(dataType)), client.TimeseriesAlreadyExist)
	if err != nil {
		return fmt.Errorf("create IoTDB timeseries %s: %w", path, err)
	}
	return nil
}

func (m *Database) rootPath() string {
	return comm.IoTDBRoot(m.Option.Database)
}

func (m *Database) devicePath(device string) string {
	return comm.IoTDBDevicePath(m.rootPath(), device)
}

func (m *Database) logDevicePath(device string) string {
	return m.rootPath() + ".logs.device_" + iotdbSegment(device)
}

func safePath(value, fallback string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback
	}
	parts := strings.Split(value, ".")
	for i := range parts {
		parts[i] = iotdbSegment(parts[i])
	}
	return strings.Join(parts, ".")
}

func iotdbSegment(value string) string {
	return comm.IoTDBNode(value)
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

func encodingFor(dataType string) string {
	switch dataType {
	case "BOOLEAN":
		return "RLE"
	case "INT32", "INT64", "TIMESTAMP":
		return "TS_2DIFF"
	case "FLOAT", "DOUBLE":
		return "GORILLA"
	default:
		return "PLAIN"
	}
}

func scanRow(rows *sql.Rows, columnCount int) ([]any, error) {
	values := make([]any, columnCount)
	arguments := make([]any, columnCount)
	for i := range values {
		arguments[i] = &values[i]
	}
	if err := rows.Scan(arguments...); err != nil {
		return nil, err
	}
	return values, nil
}

var _ interface {
	Close()
	Query(string) (*sql.Rows, error)
	Count(string) (int, error)
	InsertLogData(iotModel.DeviceLog) (sql.Result, error)
	BatchInsertLogData(map[string][]iotModel.DeviceLog) (int, error)
	InsertDeviceData(string, iotModel.ReportPropertyData, ...string) error
	BatchInsertDeviceData(string, []iotModel.ReportPropertyData) (int, error)
	BatchInsertMultiDeviceData(map[string][]iotModel.ReportPropertyData) (int, error)
	WatchDeviceData(string, func(iotModel.ReportPropertyData)) error
	GetAllDatabaseName(context.Context) ([]string, error)
	GetTableListByDatabase(context.Context, string) ([]iotModel.TsdTables, error)
	GetTableInfo(context.Context, string) ([]*iotModel.TsdTableInfo, error)
	GetTableData(context.Context, string) (*iotModel.TsdTableDataInfo, error)
	GetTableDataOne(context.Context, string, ...any) (gdb.Record, error)
	GetTableDataAll(context.Context, string, ...any) (gdb.Result, error)
} = (*Database)(nil)
