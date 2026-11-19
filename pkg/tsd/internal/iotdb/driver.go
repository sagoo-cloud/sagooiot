package iotdb

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/apache/iotdb-client-go/client"
)

const driverName = "iotdb"

func init() {
	sql.Register(driverName, iotdbDriver{})
}

type iotdbDriver struct{}

func (iotdbDriver) Open(name string) (driver.Conn, error) {
	config, timeout, err := parseDSN(name)
	if err != nil {
		return nil, err
	}
	session := client.NewSession(config)
	if err := session.Open(false, timeout); err != nil {
		return nil, fmt.Errorf("open IoTDB session: %w", err)
	}
	return &iotdbConn{session: &session}, nil
}

type iotdbConn struct {
	session *client.Session
}

func (c *iotdbConn) Prepare(query string) (driver.Stmt, error) {
	return &iotdbStmt{conn: c, query: query}, nil
}

func (c *iotdbConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return c.Prepare(query)
}

func (c *iotdbConn) Close() error {
	if c.session == nil {
		return nil
	}
	err := c.session.Close()
	c.session = nil
	return err
}

func (c *iotdbConn) Begin() (driver.Tx, error) {
	return nil, errors.New("IoTDB tree sessions do not support transactions")
}

func (c *iotdbConn) Ping(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	rows, err := c.session.ExecuteStatementWithContext(ctx, "SHOW DATABASES")
	if err != nil {
		return err
	}
	return rows.Close()
}

func (c *iotdbConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	statement, err := bindSQL(query, args)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if err := c.session.ExecuteNonQueryStatement(statement); err != nil {
		return nil, err
	}
	return driver.RowsAffected(0), nil
}

func (c *iotdbConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	statement, err := bindSQL(query, args)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	result, err := c.session.ExecuteStatementWithContext(ctx, statement)
	if err != nil {
		return nil, err
	}
	return &iotdbRows{ctx: ctx, result: result, columns: result.GetColumnNames()}, nil
}

type iotdbStmt struct {
	conn  *iotdbConn
	query string
}

func (s *iotdbStmt) Close() error  { return nil }
func (s *iotdbStmt) NumInput() int { return countPlaceholders(s.query) }

func (s *iotdbStmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.conn.ExecContext(context.Background(), s.query, namedValues(args))
}

func (s *iotdbStmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.conn.QueryContext(context.Background(), s.query, namedValues(args))
}

type iotdbRows struct {
	ctx     context.Context
	result  sessionDataSet
	columns []string
	closed  bool
}

type sessionDataSet interface {
	Next() (bool, error)
	GetObjectByIndex(columnIndex int32) (any, error)
	Close() error
}

func (r *iotdbRows) Columns() []string { return r.columns }

func (r *iotdbRows) Close() error {
	if r.closed {
		return nil
	}
	r.closed = true
	return r.result.Close()
}

func (r *iotdbRows) Next(dest []driver.Value) error {
	if err := r.ctx.Err(); err != nil {
		_ = r.Close()
		return err
	}
	next, err := r.result.Next()
	if err != nil {
		_ = r.Close()
		return err
	}
	if !next {
		_ = r.Close()
		return io.EOF
	}
	for index := range r.columns {
		value, err := r.result.GetObjectByIndex(int32(index + 1))
		if err != nil {
			_ = r.Close()
			return err
		}
		dest[index], err = driverValue(value)
		if err != nil {
			_ = r.Close()
			return fmt.Errorf("decode IoTDB column %q: %w", r.columns[index], err)
		}
	}
	return nil
}

func parseDSN(dsn string) (*client.Config, int, error) {
	parsed, err := url.Parse(dsn)
	if err != nil {
		return nil, 0, fmt.Errorf("parse IoTDB DSN: %w", err)
	}
	if parsed.Scheme != driverName || parsed.Hostname() == "" {
		return nil, 0, fmt.Errorf("IoTDB DSN must use iotdb://user:password@host:port")
	}
	port := parsed.Port()
	if port == "" {
		port = "6667"
	}
	user, password := "root", "root"
	if parsed.User != nil {
		user = parsed.User.Username()
		if passwordValue, ok := parsed.User.Password(); ok {
			password = passwordValue
		}
	}
	query := parsed.Query()
	timeout := 10000
	if value := query.Get("connectTimeoutMs"); value != "" {
		timeout, err = strconv.Atoi(value)
		if err != nil || timeout < 0 {
			return nil, 0, fmt.Errorf("invalid connectTimeoutMs %q", value)
		}
	}
	fetchSize := int64(1024)
	if value := query.Get("fetchSize"); value != "" {
		fetchSize, err = strconv.ParseInt(value, 10, 32)
		if err != nil || fetchSize < 1 {
			return nil, 0, fmt.Errorf("invalid fetchSize %q", value)
		}
	}
	return &client.Config{
		Host:            parsed.Hostname(),
		Port:            port,
		UserName:        user,
		Password:        password,
		FetchSize:       int32(fetchSize),
		TimeZone:        query.Get("timeZone"),
		ConnectRetryMax: 1,
	}, timeout, nil
}

func BuildDSN(host string, port int, username, password string) string {
	if port <= 0 {
		port = 6667
	}
	if username == "" {
		username = "root"
	}
	if password == "" {
		password = "root"
	}
	return (&url.URL{
		Scheme: driverName,
		Host:   net.JoinHostPort(host, strconv.Itoa(port)),
		User:   url.UserPassword(username, password),
		RawQuery: url.Values{
			"connectTimeoutMs": {"10000"},
			"timeZone":         {"UTC"},
		}.Encode(),
	}).String()
}

func bindSQL(query string, args []driver.NamedValue) (string, error) {
	var out strings.Builder
	out.Grow(len(query) + len(args)*8)
	argIndex := 0
	quote := byte(0)
	lineComment, blockComment := false, false
	for i := 0; i < len(query); i++ {
		ch := query[i]
		if lineComment {
			out.WriteByte(ch)
			if ch == '\n' {
				lineComment = false
			}
			continue
		}
		if blockComment {
			out.WriteByte(ch)
			if ch == '*' && i+1 < len(query) && query[i+1] == '/' {
				out.WriteByte(query[i+1])
				i++
				blockComment = false
			}
			continue
		}
		if quote != 0 {
			out.WriteByte(ch)
			if ch == quote {
				if i+1 < len(query) && query[i+1] == quote {
					out.WriteByte(query[i+1])
					i++
				} else {
					quote = 0
				}
			}
			continue
		}
		if ch == '-' && i+1 < len(query) && query[i+1] == '-' {
			out.WriteString("--")
			i++
			lineComment = true
			continue
		}
		if ch == '/' && i+1 < len(query) && query[i+1] == '*' {
			out.WriteString("/*")
			i++
			blockComment = true
			continue
		}
		if ch == '\'' || ch == '"' || ch == '`' {
			quote = ch
			out.WriteByte(ch)
			continue
		}
		if ch == '?' {
			if argIndex >= len(args) {
				return "", fmt.Errorf("SQL has more placeholders than arguments")
			}
			literal, err := sqlLiteral(args[argIndex].Value)
			if err != nil {
				return "", err
			}
			out.WriteString(literal)
			argIndex++
			continue
		}
		out.WriteByte(ch)
	}
	if argIndex != len(args) {
		return "", fmt.Errorf("SQL has %d placeholders but received %d arguments", argIndex, len(args))
	}
	return out.String(), nil
}

func countPlaceholders(query string) int {
	count := 0
	quote := byte(0)
	lineComment, blockComment := false, false
	for i := 0; i < len(query); i++ {
		ch := query[i]
		if lineComment {
			if ch == '\n' {
				lineComment = false
			}
			continue
		}
		if blockComment {
			if ch == '*' && i+1 < len(query) && query[i+1] == '/' {
				i++
				blockComment = false
			}
			continue
		}
		if quote != 0 {
			if ch == quote {
				if i+1 < len(query) && query[i+1] == quote {
					i++
				} else {
					quote = 0
				}
			}
			continue
		}
		if ch == '-' && i+1 < len(query) && query[i+1] == '-' {
			i++
			lineComment = true
			continue
		}
		if ch == '/' && i+1 < len(query) && query[i+1] == '*' {
			i++
			blockComment = true
			continue
		}
		if ch == '\'' || ch == '"' || ch == '`' {
			quote = ch
			continue
		}
		if ch == '?' {
			count++
		}
	}
	return count
}

func sqlLiteral(value any) (string, error) {
	switch v := value.(type) {
	case nil:
		return "NULL", nil
	case string:
		return "'" + strings.ReplaceAll(v, "'", "''") + "'", nil
	case []byte:
		return "X'" + hex.EncodeToString(v) + "'", nil
	case bool:
		return strconv.FormatBool(v), nil
	case time.Time:
		return "'" + v.UTC().Format("2006-01-02T15:04:05.999Z07:00") + "'", nil
	case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		return fmt.Sprint(v), nil
	case float64, float32:
		return strconv.FormatFloat(reflectFloat64(v), 'g', -1, 64), nil
	default:
		return "", fmt.Errorf("unsupported IoTDB parameter type %T", value)
	}
}

func reflectFloat64(value any) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}

func namedValues(values []driver.Value) []driver.NamedValue {
	result := make([]driver.NamedValue, len(values))
	for i, value := range values {
		result[i] = driver.NamedValue{Ordinal: i + 1, Value: value}
	}
	return result
}

func driverValue(value any) (driver.Value, error) {
	switch v := value.(type) {
	case nil, int64, float64, bool, []byte, string, time.Time:
		return v, nil
	case *client.Binary:
		if v == nil {
			return nil, nil
		}
		return append([]byte(nil), v.GetValues()...), nil
	case int32:
		return int64(v), nil
	case float32:
		return float64(v), nil
	default:
		return nil, fmt.Errorf("unsupported IoTDB result type %T", value)
	}
}

var _ driver.Driver = iotdbDriver{}
var _ driver.Conn = (*iotdbConn)(nil)
var _ driver.ConnPrepareContext = (*iotdbConn)(nil)
var _ driver.ExecerContext = (*iotdbConn)(nil)
var _ driver.QueryerContext = (*iotdbConn)(nil)
var _ driver.Pinger = (*iotdbConn)(nil)
var _ driver.Stmt = (*iotdbStmt)(nil)
var _ driver.Rows = (*iotdbRows)(nil)
