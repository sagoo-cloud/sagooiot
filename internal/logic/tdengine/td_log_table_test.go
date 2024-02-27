package tdengine

import (
	"context"
	_ "sagooiot/internal/logic/product"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestInsertLog(t *testing.T) {
	logData := &model.TdLogAddInput{
		Ts:      gtime.Now(),
		Device:  "k213213",
		Type:    "属性上报",
		Content: `{"device_id":"k213213","return_time":"2022-11-10 10:49:33","property_99":2,"property_98":2,"property_97":3,"property_96":4,"property_95":2}`,
	}
	err := service.TdLogTable().Insert(context.TODO(), logData)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClear(t *testing.T) {
	err := service.TdLogTable().Clear(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
}
