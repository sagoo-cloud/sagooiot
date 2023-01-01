package alarm

import (
	"context"
	_ "github.com/sagoo-cloud/sagooiot/internal/logic/notice"
	_ "github.com/sagoo-cloud/sagooiot/internal/logic/product"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestCheck(t *testing.T) {
	productKey := "monipower20221103"
	deviceKey := "t20221333"
	data := map[string]any{
		"ts": gtime.Datetime(),
		"va": 92.12,
	}
	err := service.AlarmRule().Check(context.TODO(), productKey, deviceKey, data)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
}

func TestCache(t *testing.T) {
	rs, err := service.AlarmRule().Cache(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rs)
}
