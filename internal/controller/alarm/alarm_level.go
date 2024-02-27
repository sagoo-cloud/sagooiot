package alarm

import (
	"context"
	"sagooiot/api/v1/alarm"
	"sagooiot/internal/service"
)

var AlarmLevel = cAlarmLevel{}

type cAlarmLevel struct{}

func (c *cAlarmLevel) All(ctx context.Context, req *alarm.AlarmLevelReq) (res *alarm.AlarmLevelRes, err error) {
	list, err := service.AlarmLevel().All(ctx)
	if err != nil || list == nil {
		return
	}
	res = new(alarm.AlarmLevelRes)
	res.AlarmLevelListOutput = list
	return
}

func (c *cAlarmLevel) Edit(ctx context.Context, req *alarm.AlarmLevelEditReq) (res *alarm.AlarmLevelEditRes, err error) {
	err = service.AlarmLevel().Edit(ctx, req.List)
	return
}
