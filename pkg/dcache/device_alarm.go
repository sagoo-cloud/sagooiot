package dcache

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/pkg/cache"
)

// GetDeviceAlarm 基于产品key获取设备告警规则
func GetDeviceAlarm(ctx context.Context, productKey string) (out []model.AlarmRuleOutput) {
	data, err := cache.Instance().Get(ctx, consts.DeviceAlarmRulePrefix+productKey)
	if err != nil || data.Val() == nil {
		return
	}
	if err = gconv.Scan(data, &out); err != nil {
		return
	}

	return
}

// SetDeviceAlarmRule 基于产品key设置设备告警规则
func SetDeviceAlarmRule(ctx context.Context, productKey string, data []model.AlarmRuleOutput) (err error) {
	if data == nil {
		return
	}
	err = cache.Instance().Set(ctx, consts.DeviceAlarmRulePrefix+productKey, data, 0)
	return
}
