package proxy

import (
	"context"
	"sagooiot/internal/dao"
	proxyModel "sagooiot/pkg/proxy/model"
)

func GetAlarmLogByDeviceKey(ctx context.Context, deviceKeys []string) (out []*proxyModel.AlarmLog, err error) {
	err = dao.AlarmLog.Ctx(ctx).WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).Scan(&out)
	return
}
