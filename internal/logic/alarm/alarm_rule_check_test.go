package alarm

import (
	"context"
	"sagooiot/internal/consts"
	_ "sagooiot/internal/logic/context"
	_ "sagooiot/internal/logic/notice"
	_ "sagooiot/internal/logic/product"
	_ "sagooiot/internal/logic/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestCheck(t *testing.T) {
	productKey := "monipower20221103"
	deviceKey := "t20221333"
	// param := model.ReportPropertyData{
	// 	"va": {92.12, gtime.Now().Unix()},
	// }
	param := model.ReportStatusData{
		Status:     "online",
		CreateTime: gtime.Now().Unix(),
	}
	err := service.AlarmRule().Check(context.TODO(), productKey, deviceKey, consts.AlarmTriggerTypeOnline, param)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
}
