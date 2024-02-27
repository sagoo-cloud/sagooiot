package product

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestDo(t *testing.T) {
	in := &model.DeviceFunctionInput{
		DeviceKey: "aoxiangTest11",
		FuncKey:   "beginPlay",
		Params: map[string]any{
			"deviceNo":      "设备编号",
			"deviceChannel": "设备通道",
			"in_a":          1,
		},
	}
	out, err := service.DevDeviceFunction().Do(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
