package product

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestSet(t *testing.T) {
	in := &model.DevicePropertyInput{
		DeviceKey: "aoxiangTest11",
		Params: map[string]any{
			"a":  9,
			"bb": true,
		},
	}
	out, err := service.DevDeviceProperty().Set(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
