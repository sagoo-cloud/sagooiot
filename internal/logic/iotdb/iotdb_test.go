package iotdb

import "testing"

func TestTreePathAndTypeMapping(t *testing.T) {
	if got, want := nodeName("Sensor-A 01"), "sensor_x2d_a_x20_01"; got != want {
		t.Fatalf("nodeName() = %q, want %q", got, want)
	}
	if got, want := devicePath("device-2"), "root.sagoo_iot.device_device_2"; got != want {
		t.Fatalf("devicePath() = %q, want %q", got, want)
	}
	types := map[string]string{
		"int": "INT32", "long": "INT64", "float": "FLOAT", "double": "DOUBLE",
		"boolean": "BOOLEAN", "date": "DATE", "string": "TEXT", "unknown": "TEXT",
	}
	for input, want := range types {
		if got := tslType(input); got != want {
			t.Errorf("tslType(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestParseTagValue(t *testing.T) {
	tests := []struct {
		typeName string
		value    string
		want     any
	}{
		{"int", "17", int32(17)},
		{"long", "9000000000", int64(9000000000)},
		{"float", "1.5", float32(1.5)},
		{"double", "2.25", float64(2.25)},
		{"boolean", "true", true},
		{"string", "zone-a", "zone-a"},
	}
	for _, test := range tests {
		got, err := parseTagValue(test.value, test.typeName)
		if err != nil {
			t.Errorf("parseTagValue(%q, %q): %v", test.value, test.typeName, err)
			continue
		}
		if got != test.want {
			t.Errorf("parseTagValue(%q, %q) = %#v, want %#v", test.value, test.typeName, got, test.want)
		}
	}
	if _, err := parseTagValue("not-an-int", "int"); err == nil {
		t.Fatal("expected invalid integer tag to fail")
	}
}
