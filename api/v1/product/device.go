package product

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
)

type DetailDeviceReq struct {
	g.Meta    `path:"/device/detail" method:"get" summary:"设备详情" tags:"设备"`
	DeviceKey string `json:"deviceKey" dc:"设备标识deviceKey" v:"required#设备标识不能为空"`
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
	g.Meta     `path:"/device/list" method:"get" summary:"已发布产品设备列表" tags:"设备"`
	ProductKey string `json:"productKey" dc:"产品Key"`
	KeyWord    string `json:"keyWord" dc:"搜索设备的关键词"`
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

type UpdateDeviceExtensionInfoReq struct {
	g.Meta `path:"/device/update" method:"put" summary:"更新设备信息" tags:"设备"`
	*model.DeviceExtensionInfoInput
}
type UpdateDeviceExtensionInfoRes struct{}

type UpdateDeviceExtendReq struct {
	g.Meta `path:"/device/extend/update" method:"put" summary:"更新设备扩展信息" tags:"设备"`
	*model.DeviceExtendInput
}
type UpdateDeviceExtendRes struct{}

type DelDeviceReq struct {
	g.Meta `path:"/device/del" method:"delete" summary:"删除设备" tags:"设备"`
	Keys   []string `json:"keys" dc:"设备Keys" v:"required#设备ID不能为空"`
}
type DelDeviceRes struct{}

type DeployDeviceReq struct {
	g.Meta    `path:"/device/deploy" method:"post" summary:"启用设备" tags:"设备"`
	DeviceKey string `json:"deviceKey" dc:"设备标识deviceKey" v:"required#设备标识不能为空"`
}
type DeployDeviceRes struct{}

type UndeployDeviceReq struct {
	g.Meta    `path:"/device/undeploy" method:"post" summary:"禁用设备" tags:"设备"`
	DeviceKey string `json:"deviceKey" dc:"设备标识deviceKey" v:"required#设备标识不能为空"`
}
type UndeployDeviceRes struct{}

type OnlineDeviceReq struct {
	g.Meta    `path:"/device/online" method:"post" summary:"上线设备" tags:"设备"`
	DeviceKey string `json:"deviceKey" dc:"设备标识deviceKey" v:"required#设备标识不能为空"`
}
type OnlineDeviceRes struct{}

type OfflineDeviceReq struct {
	g.Meta    `path:"/device/offline" method:"post" summary:"下线设备" tags:"设备"`
	DeviceKey string `json:"deviceKey" dc:"设备标识deviceKey" v:"required#设备标识不能为空"`
}
type OfflineDeviceRes struct{}

type DeviceRunStatusReq struct {
	g.Meta    `path:"/device/run_status" method:"get" summary:"运行状态" tags:"设备"`
	DeviceKey string `json:"device_key" dc:"设备ID" v:"required#设备device_key不能为空"`
}
type DeviceRunStatusRes struct {
	*model.DeviceRunStatusOutput
}

type DeviceGetLatestPropertyReq struct {
	g.Meta    `path:"/device/get_latest_property" method:"get" summary:"获取设备最新的属性值" tags:"设备"`
	DeviceKey string `json:"deviceKey" dc:"设备标识deviceKey" v:"required#设备标识不能为空"`
}
type DeviceGetLatestPropertyRes struct {
	List []model.DeviceLatestProperty `json:"list" dc:"设备最新的属性值"`
}

type DeviceGetPropertyReq struct {
	g.Meta `path:"/device/property/get" method:"get" summary:"获取指定属性值" tags:"设备"`
	*model.DeviceGetPropertyInput
}
type DeviceGetPropertyRes struct {
	*model.DevicePropertiy
}

type DeviceGetPropertyListReq struct {
	g.Meta      `path:"/device/property/list" method:"get" summary:"属性详情列表" tags:"设备"`
	DeviceKey   string `json:"deviceKey" dc:"设备标识"`
	PropertyKey string `json:"propertyKey" dc:"属性标识" v:"required#属性标识不能为空"`
	common.PaginationReq
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

type DeviceBindReq struct {
	g.Meta `path:"/device/bind_sub" method:"post" summary:"绑定子设备" tags:"设备"`
	*model.DeviceBindInput
}
type DeviceBindRes struct{}

type DeviceUnBindReq struct {
	g.Meta `path:"/device/unbind_sub" method:"post" summary:"解绑子设备" tags:"设备"`
	*model.DeviceBindInput
}
type DeviceUnBindRes struct{}

type BindListReq struct {
	g.Meta `path:"/device/bind_list" method:"get" summary:"已绑定子设备列表" tags:"设备"`
	*model.DeviceBindListInput
}
type BindListRes struct {
	*model.DeviceBindListOutput
}

type ListForSubReq struct {
	g.Meta `path:"/device/sub_list" method:"get" summary:"子设备列表" tags:"设备"`
	*model.ListForSubInput
}
type ListForSubRes struct {
	*model.ListDeviceForPageOutput
}

type DelSubDeviceReq struct {
	g.Meta    `path:"/device/del_sub" method:"delete" summary:"删除子设备" tags:"设备"`
	DeviceKey string `json:"deviceKey" dc:"子设备标识deviceKey" v:"required#子设备标识不能为空"`
}
type DelSubDeviceRes struct{}

type ImportDevicesReq struct {
	g.Meta     `path:"/device/import" method:"post" summary:"导入设备" tags:"设备"`
	File       *ghttp.UploadFile `json:"file" type:"file" dc:"上传文件" v:"required#请上传文件"`
	ProductKey string            `json:"productKey" dc:"产品Key" v:"required#产品key不能为空"`
}
type ImportDevicesRes struct {
	Success    int      `json:"success" dc:"导入成功设备数"`
	Fail       int      `json:"fail" dc:"导入失败数"`
	DevicesKey []string `json:"deviceKey" dc:"失败的设备标识"`
}

type ExportDevicesReq struct {
	g.Meta     `path:"/device/export" method:"get" summary:"导出设备" tags:"设备"`
	ProductKey string `json:"productKey" dc:"产品key" v:"required#产品key不能为空"`
}
type ExportDevicesRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type SetDeviceStatusReq struct {
	g.Meta `path:"/device/setDeviceStatus" method:"post" summary:"批量启用/禁用设备" tags:"设备"`
	Keys   []string `json:"ids" dc:"设备keys" v:"array#设备keys为数组"`
	Status int      `json:"status"  dc:"0禁用，1启用" v:"required#status不能为空"`
}
type SetDeviceStatusRes struct {
}

type DeviceDataListReq struct {
	g.Meta `path:"/device/data/list" method:"get" summary:"获取设备属性数据列表" tags:"设备"`
	*model.DeviceDataListInput
}
type DeviceDataListRes struct {
	*model.DeviceDataListOutput
}
