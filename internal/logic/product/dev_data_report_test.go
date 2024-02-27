package product

import (
	"context"
	_ "sagooiot/internal/logic/alarm"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestEvent(t *testing.T) {
	deviceKey := "t20221333"
	data := model.ReportEventData{
		Key: "aaa",
		Param: model.ReportEventParam{
			Value:      map[string]any{"a": 111, "b": 222},
			CreateTime: time.Now().Unix(),
		},
	}
	err := service.DevDataReport().Event(context.TODO(), deviceKey, data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProperty(t *testing.T) {
	deviceKey := "aoxiang_d_gw"
	subKey := "aoxiang_d_sub"
	data := model.ReportPropertyData{
		"b": model.ReportPropertyNode{
			Value:      111,
			CreateTime: time.Now().Unix(),
		},
	}
	err := service.DevDataReport().Property(context.TODO(), deviceKey, data, subKey)
	if err != nil {
		t.Fatal(err)
	}
}
