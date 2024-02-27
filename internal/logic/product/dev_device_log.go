package product

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/dcache"
)

type sDevDeviceLog struct{}

func init() {
	service.RegisterDevDeviceLog(devDeviceLog())
}

func devDeviceLog() *sDevDeviceLog {
	return &sDevDeviceLog{}
}

// LogType 日志类型
func (s *sDevDeviceLog) LogType(ctx context.Context) (list []string) {
	list = consts.GetTopicTypes()
	return
}

// Search 日志搜索
func (s *sDevDeviceLog) Search(ctx context.Context, in *model.DeviceLogSearchInput) (out *model.DeviceLogSearchOutput, err error) {
	out = new(model.DeviceLogSearchOutput)

	result, total, currentPage, err := dcache.GetDataByPage(ctx, in.DeviceKey, in.PageNum, in.PageSize, in.Types, in.DateRange)
	if err != nil {
		return
	}

	out.Total = total
	out.CurrentPage = currentPage

	var logs []model.TdLog
	if err := gconv.Scan(result, &logs); err != nil {
		return nil, err
	}
	out.List = logs

	//var whereOr []string
	//for _, v := range in.Types {
	//	whereOr = append(whereOr, "type='"+v+"'")
	//}
	//
	//where := ""
	//if len(whereOr) > 0 {
	//	where = " and (" + strings.Join(whereOr, " or ") + ") "
	//}
	//
	//if len(in.DateRange) > 0 {
	//	where += " and (ts >= '" + in.DateRange[0] + " 00:00:00" + "' and ts <= '" + in.DateRange[1] + " 23:59:59" + "') "
	//}
	//
	//// TDengine
	//sql := "select count(*) as num from device_log where device='?'" + where
	//rs, err := service.TdEngine().GetOne(ctx, sql, in.DeviceKey)
	//if err != nil {
	//	return
	//}
	//out.Total = rs["num"].Int()
	//out.CurrentPage = in.PageNum
	//
	//sql = "select * from device_log where device='?'" + where + " order by ts desc limit ?, ?"
	//out.List, err = service.TdLogTable().GetAll(ctx, sql, in.DeviceKey, (in.PageNum-1)*in.PageSize, in.PageSize)
	return
}
