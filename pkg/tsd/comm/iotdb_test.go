package comm

import "testing"

func TestIoTDBPathsEncodeUntrustedSegments(t *testing.T) {
	root := IoTDBRoot("tenant-a.sg")
	if want := "root.tenant_x2d_a.sg"; root != want {
		t.Fatalf("IoTDBRoot() = %q, want %q", root, want)
	}
	if got, want := IoTDBDevicePath(root, "gateway.1/child"), "root.tenant_x2d_a.sg.device_gateway_x2e_1_x2f_child"; got != want {
		t.Fatalf("IoTDBDevicePath() = %q, want %q", got, want)
	}
	if got := PropertyColumn("Temperature", true); got != "p_temperature" {
		t.Fatalf("PropertyColumn() = %q, want p_temperature", got)
	}
}
