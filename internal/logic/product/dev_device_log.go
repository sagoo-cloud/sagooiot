package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"strings"
)

type sDevDeviceLog struct{}

func init() {
	service.RegisterDevDeviceLog(devDeviceLog())
}

func devDeviceLog() *sDevDeviceLog {
	return &sDevDeviceLog{}
}

// 日志类型
func (s *sDevDeviceLog) LogType(ctx context.Context) (list []string) {
	list = consts.GetTopicTypes()
	return
}

// 日志搜索
func (s *sDevDeviceLog) Search(ctx context.Context, in *model.DeviceLogSearchInput) (out *model.DeviceLogSearchOutput, err error) {
	out = new(model.DeviceLogSearchOutput)

	var whereOr []string
	for _, v := range in.Types {
		whereOr = append(whereOr, "type='"+v+"'")
	}

	where := ""
	if len(whereOr) > 0 {
		where = " and (" + strings.Join(whereOr, " or ") + ") "
	}

	if len(in.DateRange) > 0 {
		where += " and (ts >= '" + in.DateRange[0] + "' and ts <= '" + in.DateRange[1] + "') "
	}

	// TDengine
	sql := "select count(*) as num from device_log where device='?'" + where
	rs, err := service.TdEngine().GetOne(ctx, sql, in.DeviceKey)
	if err != nil {
		return
	}
	out.Total = rs["num"].Int()
	out.CurrentPage = in.PageNum

	sql = "select * from device_log where device='?'" + where + " order by ts desc limit ?, ?"
	out.List, err = service.TdLogTable().GetAll(ctx, sql, in.DeviceKey, (in.PageNum-1)*in.PageSize, in.PageSize)
	return
}
