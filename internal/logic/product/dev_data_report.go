package product

import (
	"context"
	"errors"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/dcache"
)

type sDevDataReport struct{}

func init() {
	service.RegisterDevDataReport(devDataReport())
}

func devDataReport() *sDevDataReport {
	return &sDevDataReport{}
}

// Event 设备事件上报
func (s *sDevDataReport) Event(ctx context.Context, deviceKey string, data model.ReportEventData, subKey ...string) error {
	device, err := dcache.GetDeviceDetailInfo(deviceKey)
	if err != nil {
		return err
	}
	if device == nil {
		return errors.New("未发现设备")
	}

	err = service.AlarmRule().Check(ctx, device.Product.Key, deviceKey, consts.AlarmTriggerTypeEvent, data, subKey...)

	return err
}

// Property 设备属性上报
func (s *sDevDataReport) Property(ctx context.Context, deviceKey string, data model.ReportPropertyData, subKey ...string) error {
	//logC, err := json.Marshal(data)
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//}
	//g.Log().Debugf(ctx, "dev_data_report: deveceKey(%s), subKey(%s), data(%s)", deviceKey, strings.Join(subKey, ","), logC)

	//if len(subKey) > 0 {
	//	skey := subKey[0]
	//	service.DevDevice().UpdateStatus(ctx, skey)
	//} else {
	//	service.DevDevice().UpdateStatus(ctx, deviceKey)
	//}

	if err := service.TSLTable().Insert(ctx, deviceKey, data, subKey...); err != nil {
		return err
	}

	return nil
}
