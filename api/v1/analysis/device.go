package analysis

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

// DeviceDataTotalCountReq 设备消息总数统计（当年、当月、当日）
type DeviceDataTotalCountReq struct {
	g.Meta   `path:"/deviceDataTotalCount" method:"get" summary:"设备消息总数统计（当年、当月、当日）" tags:"IOT数据分析"`
	DateType string `json:"dateType" v:"required#日期类型不能为空" dc:"日期类型：year 年，month 月，day 日"`
}
type DeviceDataTotalCountRes struct {
	Number int64 `json:"number" dc:"消息总数"`
}

// DeviceOnlineOfflineCountReq 设备在线离线统计
type DeviceOnlineOfflineCountReq struct {
	g.Meta `path:"/deviceOnlineOfflineCount" method:"get" summary:"设备在线离线统计" tags:"IOT数据分析"`
}
type DeviceOnlineOfflineCountRes struct {
	Data model.DeviceOnlineOfflineCount
}

// DeviceDataCountReq 按年度每月设备消息统计
type DeviceDataCountReq struct {
	g.Meta   `path:"/deviceDataCount" method:"get" summary:"按年度统计1-12月份设备消息统计" tags:"IOT数据分析"`
	DateType string `json:"dateType" v:"required#日期类型不能为空" dc:"日期类型：year 年，month 月，day 日"`
	//Date     string `json:"date" v:"" dc:"日期：年-yyyy"`
}
type DeviceDataCountRes struct {
	Data []model.CountData `json:"data" dc:"1-12月份设备消息量统计数据"`
}
