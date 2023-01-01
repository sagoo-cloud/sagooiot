package product

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceReq struct {
	g.Meta `path:"/device/get" method:"get" summary:"设备详情" tags:"设备"`
	Key    string `json:"key" dc:"设备标识" v:"required#设备标识不能为空"`
}
type GetDeviceRes struct {
	Data *model.DeviceOutput `json:"data" dc:"设备详情"`
}

type DetailDeviceReq struct {
	g.Meta `path:"/device/detail" method:"get" summary:"设备详情" tags:"设备"`
	Id     uint `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
}
type DetailDeviceRes struct {
	Data *model.DeviceOutput `json:"data" dc:"设备详情"`
}

type ListDeviceForPageReq struct {
	g.Meta `path:"/device/page_list" method:"get" summary:"设备搜索列表（分页）" tags:"设备"`
	*model.ListDeviceForPageInput
}
type ListDeviceForPageRes struct {
	*model.ListDeviceForPageOutput
}

type ListDeviceReq struct {
	g.Meta `path:"/device/list" method:"get" summary:"设备列表" tags:"设备"`
	*model.ListDeviceInput
}
type ListDeviceRes struct {
	Device []*model.DeviceOutput `json:"device" dc:"设备列表"`
}

type AddDeviceReq struct {
	g.Meta `path:"/device/add" method:"post" summary:"添加设备" tags:"设备"`
	*model.AddDeviceInput
}
type AddDeviceRes struct{}

type EditDeviceReq struct {
	g.Meta `path:"/device/edit" method:"put" summary:"编辑设备" tags:"设备"`
	*model.EditDeviceInput
}
type EditDeviceRes struct{}

type DelDeviceReq struct {
	g.Meta `path:"/device/del" method:"delete" summary:"删除设备" tags:"设备"`
	Ids    []uint `json:"ids" dc:"设备Ids" v:"required#设备ID不能为空"`
}
type DelDeviceRes struct{}

type DeployDeviceReq struct {
	g.Meta `path:"/device/deploy" method:"post" summary:"启用设备" tags:"设备"`
	Id     uint `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
}
type DeployDeviceRes struct{}

type UndeployDeviceReq struct {
	g.Meta `path:"/device/undeploy" method:"post" summary:"禁用设备" tags:"设备"`
	Id     uint `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
}
type UndeployDeviceRes struct{}

type OnlineDeviceReq struct {
	g.Meta `path:"/device/online" method:"post" summary:"上线设备" tags:"设备"`
	Key    string `json:"key" dc:"设备标识" v:"required#设备标识不能为空"`
}
type OnlineDeviceRes struct{}

type OfflineDeviceReq struct {
	g.Meta `path:"/device/offline" method:"post" summary:"下线设备" tags:"设备"`
	Key    string `json:"key" dc:"设备标识" v:"required#设备标识不能为空"`
}
type OfflineDeviceRes struct{}

type DeviceRunStatusReq struct {
	g.Meta `path:"/device/run_status" method:"get" summary:"运行状态" tags:"设备"`
	Id     uint `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
}
type DeviceRunStatusRes struct {
	*model.DeviceRunStatusOutput
}

type DeviceGetPropertyReq struct {
	g.Meta `path:"/device/property/get" method:"get" summary:"获取指定属性值" tags:"设备"`
	*model.DeviceGetPropertyInput
}
type DeviceGetPropertyRes struct {
	*model.DevicePropertiy
}

type DeviceGetPropertyListReq struct {
	g.Meta `path:"/device/property/list" method:"get" summary:"属性详情列表" tags:"设备"`
	*model.DeviceGetPropertyListInput
}
type DeviceGetPropertyListRes struct {
	*model.DeviceGetPropertyListOutput
}

type DeviceStatisticsReq struct {
	g.Meta `path:"/device/statistics" method:"get" summary:"设备相关统计" tags:"设备"`
}
type DeviceStatisticsRes struct {
	DeviceTotal model.DeviceTotalOutput `json:"deviceTotal" dc:"设备相关统计"`
}

type DeviceStatisticsForMonthsReq struct {
	g.Meta `path:"/device/statistics/months" method:"get" summary:"设备消息量、告警量月度统计" tags:"设备"`
}
type DeviceStatisticsForMonthsRes struct {
	MsgTotal   map[int]int `json:"msgTotal" dc:"设备消息量月度统计"`
	AlarmTotal map[int]int `json:"alarmTotal" dc:"设备告警量月度统计"`
}
