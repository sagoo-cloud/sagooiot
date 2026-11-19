package iotdb

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
)

// DeviceWindows streams non-empty Tree windows to count sparse history without
// retaining every window in memory. Only the requested page is returned.
func (m *Database) DeviceWindows(ctx context.Context, key string, columns []string, interval, unit uint, start, end time.Time, page, size int) (gdb.Result, int, error) {
	step, err := deviceWindowStep(interval, unit)
	if err != nil {
		return nil, 0, err
	}
	if page < 1 || size < 1 || page-1 > math.MaxInt/size || len(columns) == 0 {
		return nil, 0, fmt.Errorf("invalid device aggregation page or properties")
	}
	path := m.devicePath(key)
	var values, nonempty []string
	for _, col := range columns {
		if col != iotdbSegment(col) || !strings.HasPrefix(col, "p_") {
			return nil, 0, fmt.Errorf("invalid device aggregation property %q", col)
		}
		values = append(values, "mode("+col+") as "+col)
		nonempty = append(nonempty, "count("+col+") > 0")
	}
	if start.IsZero() && end.IsZero() {
		// The platform omits a range when requesting all history. Fetch indexed
		// boundary rows to supply the explicit range required by Tree GROUP BY.
		series := strings.Join(columns, ", ")
		for _, boundary := range []struct {
			order  string
			target *time.Time
		}{{"ASC", &start}, {"DESC", &end}} {
			row, err := m.GetTableDataOne(ctx, "SELECT "+series+" FROM "+path+" ORDER BY TIME "+boundary.order+" LIMIT 1")
			if err != nil || row == nil {
				return nil, 0, err
			}
			*boundary.target = row["Time"].Time()
		}
		end = end.Add(time.Millisecond)
	}
	if start.IsZero() || end.IsZero() || !start.Before(end) {
		return nil, 0, fmt.Errorf("invalid device aggregation time range")
	}
	first, last := start.UnixMilli(), end.UnixMilli()
	// Keep bucket starts stable across requests, including pre-epoch reports.
	origin := windowOrigin(first, step)
	query := fmt.Sprintf("SELECT %s FROM %s WHERE time >= ? AND time < ? GROUP BY ([?, ?), %dms) HAVING %s ORDER BY TIME DESC", strings.Join(values, ", "), path, step, strings.Join(nonempty, " OR "))
	rows, err := m.queryContext(ctx, query, first, last, origin, last)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	offset, total := (page-1)*size, 0
	var result gdb.Result
	for rows.Next() {
		total++
		if total <= offset || len(result) == size {
			continue
		}
		record, err := readPropertyRecord(rows)
		if err != nil {
			return nil, 0, err
		}
		windowStart := record["Time"].Time()
		delete(record, "Time")
		record["_wstart"] = gvar.New(windowStart)
		record["_wend"] = gvar.New(time.UnixMilli(windowStart.UnixMilli() + step))
		result = append(result, record)
	}
	return result, total, rows.Err()
}

func deviceWindowStep(interval, unit uint) (int64, error) {
	units := [...]int64{1000, 60_000, 3_600_000, 86_400_000}
	if unit == 0 || unit > uint(len(units)) || interval == 0 || uint64(interval) > uint64(math.MaxInt64/units[unit-1]) {
		return 0, fmt.Errorf("invalid device aggregation interval or unit")
	}
	return int64(interval) * units[unit-1], nil
}

func windowOrigin(timestamp, step int64) int64 {
	origin := timestamp - timestamp%step
	if origin > timestamp {
		origin -= step
	}
	return origin
}

// Count log messages from exactly the requested devices in [start, end).
func (m *Database) CountLogs(ctx context.Context, deviceKeys []string, start, end time.Time) (int, error) {
	if len(deviceKeys) == 0 {
		return 0, nil
	}
	paths := make([]string, len(deviceKeys))
	for i, key := range deviceKeys {
		paths[i] = m.logDevicePath(key)
	}
	row, err := m.GetTableDataOne(ctx, "SELECT count(content) FROM "+strings.Join(paths, ", ")+" WHERE time >= ? AND time < ? GROUP BY LEVEL=0", start.UnixMilli(), end.UnixMilli())
	if err != nil {
		return 0, err
	}
	count := 0
	for key, value := range row {
		if key != "Time" {
			count += value.Int()
		}
	}
	return count, nil
}
