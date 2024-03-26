package product

import (
	"context"
	_ "sagooiot/internal/logic/alarm"
	_ "sagooiot/internal/logic/common"
	_ "sagooiot/internal/logic/context"
	_ "sagooiot/internal/logic/middleware"
	_ "sagooiot/internal/logic/network"
	_ "sagooiot/internal/logic/notice"
	_ "sagooiot/internal/logic/system"
	_ "sagooiot/internal/logic/tdengine"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func TestRunStatus(t *testing.T) {
	out, err := service.DevDevice().RunStatus(context.TODO(), "t20221222")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestCheckBind(t *testing.T) {
	in := &model.CheckBindInput{
		GatewayKey: "aoxiang_d_gw",
		SubKey:     "aoxiang_d_sub",
	}
	yes, err := service.DevDevice().CheckBind(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(yes)
}

func TestAuthInfo(t *testing.T) {
	in := &model.AuthInfoInput{
		DeviceKey:  "aoxiangTest1",
		ProductKey: "",
	}
	out, err := service.DevDevice().AuthInfo(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestGetData(t *testing.T) {
	in := &model.DeviceGetDataInput{
		DeviceKey:   "t20221222",
		PropertyKey: "va",
		DateRange:   []string{"2023-05-30 00:00:00", "2023-05-30 00:00:30"},
	}
	out, err := service.DevDevice().GetData(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(out)
}

func TestGetPropertyList(t *testing.T) {
	in := &model.DeviceGetPropertyListInput{
		DeviceKey:       "t20221222",
		PropertyKey:     "va",
		PaginationInput: model.PaginationInput{PageNum: 1, PageSize: 10},
	}
	out, err := service.DevDevice().GetPropertyList(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(out)
}

func TestGetLatestProperty(t *testing.T) {
	list, err := service.DevDevice().GetLatestProperty(context.TODO(), "aoxiang925d")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(list)
}

// TestGetDeviceOnlineTimeOut 测试获取设备在线超时时间
func TestGetDeviceOnlineTimeOut(t *testing.T) {
	timeOut := service.DevDevice().GetDeviceOnlineTimeOut(context.TODO(), "t202210000")
	g.Dump(timeOut)
}

func TestCacheDeviceDetailList(t *testing.T) {
	err := service.DevDevice().CacheDeviceDetailList(context.Background())
	if err != nil {
		t.Fatal(err)
	}

}
