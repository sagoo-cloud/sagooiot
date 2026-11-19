package comm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/apache/iotdb-client-go/client"
)

func IoTDBNode(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "_empty"
	}
	var out strings.Builder
	for _, r := range strings.ToLower(value) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '_' {
			out.WriteRune(r)
		} else {
			fmt.Fprintf(&out, "_x%x_", r)
		}
	}
	name := out.String()
	if name[0] >= '0' && name[0] <= '9' {
		return "n_" + name
	}
	return name
}

func IoTDBRoot(value string) string {
	value = strings.Trim(value, ". ")
	if value == "" {
		return "root.sagoo_iot"
	}
	parts := strings.Split(value, ".")
	for i := range parts {
		parts[i] = IoTDBNode(parts[i])
	}
	if parts[0] != "root" {
		parts = append([]string{"root"}, parts...)
	}
	return strings.Join(parts, ".")
}

func IoTDBDevicePath(root, key string) string {
	return root + "." + IoTDBNode(DeviceTableName(key))
}

func PropertyColumn(key string, tree bool) string {
	if tree {
		return TdPropertyPrefix + strings.ToLower(key)
	}
	return TsdColumnName(key)
}

// The Go session returns binary-encoded text for SHOW metadata columns.
func IoTDBMetadataText(value any) string {
	if text, ok := value.([]byte); ok {
		return string(text)
	}
	return fmt.Sprint(value)
}

// The device protocol sends milliseconds, but the platform also produces Unix seconds.
func IoTDBReportMillis(timestamp int64) int64 {
	if timestamp > -10_000_000_000 && timestamp < 10_000_000_000 {
		return timestamp * 1000
	}
	return timestamp
}

// Tree DDL has no IF NOT EXISTS clause. Accept only the server's exact duplicate code.
func CreateIoTDBSchema(ctx context.Context, db *sql.DB, statement string, duplicate int32) error {
	_, err := db.ExecContext(ctx, statement)
	var server *client.ExecutionError
	if errors.As(err, &server) && server.Code == duplicate {
		return nil
	}
	return err
}
