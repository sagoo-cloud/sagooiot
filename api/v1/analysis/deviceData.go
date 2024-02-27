package analysis

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

// DeviceDataReq 按年度每月设备消息统计
type DeviceDataReq struct {
	g.Meta    `path:"/deviceData" method:"get" summary:"设备最近的数据" tags:"IOT数据分析"`
	DeviceKey string `json:"deviceKey" v:"required#设备key不能为空" dc:"设备key"`
	common.PaginationReq
}
type DeviceDataRes struct {
	Data []interface{}
	common.PaginationRes
}

// DeviceDataForProductByLatestReq 按产品查询设备最近的数据
type DeviceDataForProductByLatestReq struct {
	g.Meta     `path:"/deviceDataForProductByLatest" method:"get" summary:"按产品查询设备最近的数据" tags:"IOT数据分析"`
	ProductKey string `json:"productKey" v:"required#产品key不能为空" dc:"产品key"`
}

type DeviceDataForProductByLatestRes struct {
	Data []model.DeviceDataRes
}

type DeviceDataForTsdReq struct {
	g.Meta `path:"/deviceDataForTsd" method:"get" summary:"设备的时序数据" tags:"IOT数据分析"`
	common.PaginationReq
}
type DeviceDataForTsdRes struct {
	Data []interface{}
	common.PaginationRes
}

// DeviceAlarmLogDataReq 设备告警日志数据请求
type DeviceAlarmLogDataReq struct {
	g.Meta `path:"/deviceAlarmLogData" method:"get" summary:"设备告警日志数据" tags:"IOT数据分析"`
	common.PaginationReq
}
type DeviceAlarmLogDataRes struct {
	Data interface{}
	common.PaginationRes
}
