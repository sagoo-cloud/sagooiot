package thing

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/thing"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var DataOverview = cDataOverview{}

type cDataOverview struct{}

// 物联概览
func (a *cDataOverview) ThingOverview(ctx context.Context, req *thing.ThingOverviewReq) (res thing.ThingOverviewRes, err error) {
	res.Overview, err = service.DevDevice().Total(ctx)
	if err != nil {
		return
	}
	res.Device.MsgTotal, err = service.DevDevice().TotalForMonths(ctx)
	if err != nil {
		return
	}
	res.Device.AlarmTotal, err = service.DevDevice().AlarmTotalForMonths(ctx)
	if err != nil {
		return
	}
	res.AlarmLevel, err = service.AlarmLog().TotalForLevel(ctx)

	return
}
