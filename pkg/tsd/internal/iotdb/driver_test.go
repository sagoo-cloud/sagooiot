package iotdb

import (
	"context"
	"database/sql/driver"
	"io"
	"math"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/apache/iotdb-client-go/client"
)

type stubSessionDataSet struct {
	indexes []int32
	values  []any
	rowRead bool
}

func (s *stubSessionDataSet) Next() (bool, error) {
	if s.rowRead {
		return false, nil
	}
	s.rowRead = true
	return true, nil
}

func (s *stubSessionDataSet) GetObjectByIndex(index int32) (any, error) {
	s.indexes = append(s.indexes, index)
	return s.values[index-1], nil
}

func (*stubSessionDataSet) Close() error { return nil }

func TestRowsUseOneBasedSessionIndexesAndDecodeText(t *testing.T) {
	dataSet := &stubSessionDataSet{values: []any{int64(123), client.NewBinary([]byte("hello"))}}
	rows := &iotdbRows{ctx: context.Background(), result: dataSet, columns: []string{"Time", "value"}}
	dest := make([]driver.Value, 2)
	if err := rows.Next(dest); err != nil {
		t.Fatal(err)
	}
	if len(dataSet.indexes) != 2 || dataSet.indexes[0] != 1 || dataSet.indexes[1] != 2 {
		t.Fatalf("session indexes = %v, want [1 2]", dataSet.indexes)
	}
	if dest[0] != int64(123) || string(dest[1].([]byte)) != "hello" {
		t.Fatalf("decoded row = %#v", dest)
	}
	if err := rows.Next(dest); err != io.EOF {
		t.Fatalf("Next() = %v, want EOF", err)
	}
}

func TestBuildAndParseDSN(t *testing.T) {
	dsn := BuildDSN("::1", 6667, "iot user", "p@ss:/word")
	config, timeout, err := parseDSN(dsn)
	if err != nil {
		t.Fatal(err)
	}
	if config.Host != "::1" || config.Port != "6667" {
		t.Fatalf("unexpected endpoint: %s:%s", config.Host, config.Port)
	}
	if config.UserName != "iot user" || config.Password != "p@ss:/word" {
		t.Fatalf("credentials did not round trip")
	}
	if timeout != 10000 || config.TimeZone != "UTC" {
		t.Fatalf("unexpected defaults: timeout=%d timezone=%q", timeout, config.TimeZone)
	}
}

func TestParseDSNDefaultsAndValidation(t *testing.T) {
	config, timeout, err := parseDSN("iotdb://localhost")
	if err != nil {
		t.Fatal(err)
	}
	if config.Port != "6667" || config.UserName != "root" || config.Password != "root" || timeout != 10000 {
		t.Fatalf("unexpected DSN defaults: %#v timeout=%d", config, timeout)
	}
	if _, _, err := parseDSN("mysql://localhost"); err == nil {
		t.Fatal("expected non-IoTDB DSN to fail")
	}
	if _, _, err := parseDSN("iotdb://localhost?fetchSize=0"); err == nil {
		t.Fatal("expected invalid fetchSize to fail")
	}
}

func TestBindSQLRespectsQuotedTextAndComments(t *testing.T) {
	query := "INSERT INTO root.sg.d1 (time, s1, s2) VALUES (?, ?, '?') -- ?\n/* ? */"
	result, err := bindSQL(query, []driver.NamedValue{
		{Ordinal: 1, Value: int64(42)},
		{Ordinal: 2, Value: "O'Reilly"},
	})
	if err != nil {
		t.Fatal(err)
	}
	want := "INSERT INTO root.sg.d1 (time, s1, s2) VALUES (42, 'O''Reilly', '?') -- ?\n/* ? */"
	if result != want {
		t.Fatalf("unexpected bound SQL:\n got: %s\nwant: %s", result, want)
	}
	if count := countPlaceholders(query); count != 2 {
		t.Fatalf("countPlaceholders() = %d, want 2", count)
	}
}

func TestBindSQLRejectsMismatchedAndUnsupportedArguments(t *testing.T) {
	if _, err := bindSQL("SELECT 1", []driver.NamedValue{{Ordinal: 1, Value: int64(1)}}); err == nil {
		t.Fatal("expected extra argument to fail")
	}
	if _, err := bindSQL("SELECT ?", nil); err == nil {
		t.Fatal("expected missing argument to fail")
	}
	if _, err := bindSQL("SELECT ?", []driver.NamedValue{{Ordinal: 1, Value: map[string]int{"x": 1}}}); err == nil {
		t.Fatal("expected unsupported argument type to fail")
	}
}

func TestSQLLiteralFormatsSupportedValues(t *testing.T) {
	when := time.Date(2026, 11, 18, 10, 11, 12, 0, time.FixedZone("UTC+8", 8*60*60))
	tests := []struct {
		value any
		want  string
	}{
		{value: nil, want: "NULL"},
		{value: "text", want: "'text'"},
		{value: []byte{0xab, 0xcd}, want: "X'abcd'"},
		{value: true, want: "true"},
		{value: int64(9), want: "9"},
		{value: float64(1.25), want: "1.25"},
		{value: when, want: "'2026-11-18T02:11:12Z'"},
	}
	for _, test := range tests {
		got, err := sqlLiteral(test.value)
		if err != nil {
			t.Fatalf("sqlLiteral(%T): %v", test.value, err)
		}
		if got != test.want {
			t.Errorf("sqlLiteral(%T) = %q, want %q", test.value, got, test.want)
		}
	}
}

func TestDriverValueRejectsUnsupportedObject(t *testing.T) {
	if _, err := driverValue(struct{}{}); err == nil {
		t.Fatal("expected unsupported result type to fail")
	}
	if got, err := driverValue(int32(4)); err != nil || got != int64(4) {
		t.Fatalf("int32 conversion = %v, %v", got, err)
	}
	if got, err := driverValue(client.NewBinary([]byte("value"))); err != nil || string(got.([]byte)) != "value" {
		t.Fatalf("IoTDB Binary conversion = %v, %v", got, err)
	}
}

func TestBuildDSNURLSafety(t *testing.T) {
	parsed, err := url.Parse(BuildDSN("iotdb.example", 0, "root", "root"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(parsed.String(), "iotdb://root:root@iotdb.example:6667") {
		t.Fatalf("unexpected DSN: %s", parsed)
	}
	if value, err := sqlLiteral(math.NaN()); err != nil || value != "NaN" {
		t.Fatalf("float SQL literal = %q, %v", value, err)
	}
}
