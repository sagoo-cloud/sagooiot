package product

import (
	"fmt"
	"time"
)

// Date-only upper bounds include the entire local calendar day, including DST days.
// Explicit timestamps retain the API's inclusive millisecond endpoint.
func deviceDateBounds(dateRange []string, now time.Time) (time.Time, time.Time, error) {
	loc := now.Location()
	if len(dateRange) == 0 {
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		return start, start.AddDate(0, 0, 1), nil
	}
	if len(dateRange) != 2 {
		return time.Time{}, time.Time{}, fmt.Errorf("device date range requires two endpoints")
	}
	var bounds [2]time.Time
	for i, text := range dateRange {
		for _, layout := range []string{time.DateOnly, time.DateTime, time.RFC3339Nano} {
			parsed, err := time.ParseInLocation(layout, text, loc)
			if err != nil {
				continue
			}
			if i == 1 {
				if layout == time.DateOnly {
					parsed = parsed.AddDate(0, 0, 1)
				} else {
					parsed = parsed.Add(time.Millisecond)
				}
			}
			bounds[i] = parsed
			break
		}
		if bounds[i].IsZero() {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid device date %q", text)
		}
	}
	if !bounds[0].Before(bounds[1]) {
		return time.Time{}, time.Time{}, fmt.Errorf("device date range end precedes start")
	}
	return bounds[0], bounds[1], nil
}

func propertyTimeArgs(tree bool, start, end time.Time) []any {
	if tree {
		return []any{start.UnixMilli(), end.UnixMilli()}
	}
	const layout = "2006-01-02 15:04:05.000"
	return []any{start.Format(layout), end.Format(layout)}
}
