package analysis

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

// DeviceIndicatorTrendReq 设备指标趋势
type DeviceIndicatorTrendReq struct {
	g.Meta     `path:"/deviceIndicatorTrend" method:"get" summary:"设备指标趋势" tags:"IOT数据分析"`
	ProductKey string `json:"productKey" v:"required#产品key不能为空"`
	DeviceKey  string `json:"deviceKey" v:"required#设备key不能为空"`
	Properties string `json:"properties" v:"required#设备属性不能为空"`
	*common.PaginationReq
}

type DeviceIndicatorTrendRes struct {
	Data []*model.DeviceIndicatorTrendRes
}

// DeviceIndicatorPolymerizeReq 设备指标聚合
type DeviceIndicatorPolymerizeReq struct {
	g.Meta     `path:"/deviceIndicatorPolymerize" method:"get" summary:"设备指标聚合" tags:"IOT数据分析"`
	DateType   string `json:"dateType" v:"required#日期类型不能为空" dc:"日期类型：1 yyyy-MM-dd HH:mm 5分钟，2 一小时 yyyy-MM-dd HH ，3 一天 yyyy-MM-dd；对应时间范围为 一周，一个月和一年"`
	ProductKey string `json:"productKey" v:"required#产品key不能为空"`
	DeviceKey  string `json:"deviceKey" v:"required#设备key不能为空"`
	Properties string `json:"properties" v:"required#设备属性不能为空"`
	*common.PaginationReq
}

type DeviceIndicatorPolymerizeRes struct {
	Data []*model.DeviceIndicatorPolymerizeRes
}
