package product

import (
	"reflect"
	"testing"
	"time"
)

func TestDeviceCalendarDateBounds(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 11, 1, 15, 0, 0, 0, loc)
	want := time.Date(2026, 11, 1, 0, 0, 0, 0, loc)
	for _, dates := range [][]string{nil, {"2026-11-01", "2026-11-01"}} {
		start, end, err := deviceDateBounds(dates, now)
		if err != nil || !reflect.DeepEqual(propertyTimeArgs(true, start, end), []any{want.UnixMilli(), want.Add(25 * time.Hour).UnixMilli()}) {
			t.Fatalf("dates %v => %v to %v (%v)", dates, start, end, err)
		}
	}
	start, end, err := deviceDateBounds([]string{"2026-11-01T10:05:30.123Z", "2026-11-01T10:05:30.123Z"}, now)
	if err != nil || end.Sub(start) != time.Millisecond {
		t.Fatalf("inclusive timestamp = %v to %v (%v)", start, end, err)
	}
	for _, dates := range [][]string{{"2026-11-01"}, {"2026-11-02", "2026-11-01"}, {"bad", "2026-11-01"}} {
		if _, _, err := deviceDateBounds(dates, now); err == nil {
			t.Errorf("accepted invalid range %v", dates)
		}
	}
}

func TestPropertyTimeArgsKeepInclusiveTimestampPrecision(t *testing.T) {
	start, end := time.Date(2026, 11, 18, 10, 11, 12, 123_000_000, time.UTC), time.Date(2026, 11, 18, 10, 11, 12, 124_000_000, time.UTC)
	if got, want := propertyTimeArgs(true, start, end), []any{start.UnixMilli(), end.UnixMilli()}; !reflect.DeepEqual(got, want) {
		t.Fatalf("tree time args = %v, want %v", got, want)
	}
	if got, want := propertyTimeArgs(false, start, end), []any{"2026-11-18 10:11:12.123", "2026-11-18 10:11:12.124"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("SQL time args = %v, want %v", got, want)
	}
}

func TestUnsafeTimeseriesIdentifiers(t *testing.T) {
	if validateTsdIdentifiers("device_meter_1", "p_temperature") != nil {
		t.Fatal("valid identifiers were rejected")
	}
	for _, invalid := range []string{"p_a/b", "device_a.b", "p_a b", "p_a'", "p_a*"} {
		if validateTsdIdentifiers(invalid) == nil {
			t.Errorf("accepted identifier %q", invalid)
		}
	}
}
