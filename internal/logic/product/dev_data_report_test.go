package product

import (
	"context"
	_ "github.com/sagoo-cloud/sagooiot/internal/logic/alarm"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestEvent(t *testing.T) {
	deviceKey := "device1"
	data := map[string]any{
		"aaa": map[string]any{
			"a": 111,
			"b": 222,
		},
	}
	err := service.DevDataReport().Event(context.TODO(), deviceKey, data)
	if err != nil {
		t.Fatal(err)
	}
}
