package alarm

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/alarm"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
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
	var reqData = new(model.AlarmLogListInput)
	err = gconv.Scan(req, &reqData)
	out, err := service.AlarmLog().List(ctx, reqData)
	res = &alarm.AlarmLogListRes{
		AlarmLogListOutput: out,
	}
	return
}

func (c *cAlarmLog) Handle(ctx context.Context, req *alarm.AlarmLogHandleReq) (res *alarm.AlarmLogHandleRes, err error) {
	var reqData = new(model.AlarmLogHandleInput)
	err = gconv.Scan(req, &reqData)
	err = service.AlarmLog().Handle(ctx, reqData)
	return
}
