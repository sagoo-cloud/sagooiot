package alarm

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/alarm"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var AlarmLog = cAlarmLog{}

type cAlarmLog struct{}

func (c *cAlarmLog) Detail(ctx context.Context, req *alarm.AlarmLogDetailReq) (res *alarm.AlarmLogDetailRes, err error) {
	out, err := service.AlarmLog().Detail(ctx, req.Id)
	if err != nil || out == nil {
		return
	}
	res = new(alarm.AlarmLogDetailRes)
	res.Data = out
	return
}

func (c *cAlarmLog) List(ctx context.Context, req *alarm.AlarmLogListReq) (res *alarm.AlarmLogListRes, err error) {
	out, err := service.AlarmLog().List(ctx, req.AlarmLogListInput)
	res = &alarm.AlarmLogListRes{
		AlarmLogListOutput: out,
	}
	return
}

func (c *cAlarmLog) Handle(ctx context.Context, req *alarm.AlarmLogHandleReq) (res *alarm.AlarmLogHandleRes, err error) {
	err = service.AlarmLog().Handle(ctx, &req.AlarmLogHandleInput)
	return
}
