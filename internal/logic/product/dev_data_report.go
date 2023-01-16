package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sDevDataReport struct{}

func init() {
	service.RegisterDevDataReport(devDataReport())
}

func devDataReport() *sDevDataReport {
	return &sDevDataReport{}
}

//Event 设备事件上报
func (s *sDevDataReport) Event(ctx context.Context, deviceKey string, data map[string]any) error {
	dout, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		return err
	}

	err = service.AlarmRule().Check(ctx, dout.Product.Key, deviceKey, model.AlarmTriggerTypeEvent, data)
	return err
}
