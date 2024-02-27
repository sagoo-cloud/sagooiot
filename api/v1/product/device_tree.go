package product

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DeviceTreeListReq struct {
	g.Meta `path:"/device_tree/list" method:"get" summary:"设备树列表" tags:"设备树"`
}
type DeviceTreeListRes struct {
	List []*model.DeviceTreeListOutput `json:"list" dc:"设备树列表"`
}

type DeviceTreeChangeReq struct {
	g.Meta       `path:"/device_tree/change" method:"post" summary:"更换上下级" tags:"设备树"`
	InfoId       int `json:"infoId" dc:"信息ID" v:"required#信息ID不能为空"`
	ParentInfoId int `json:"parentInfoId" dc:"所属信息ID" v:"required#所属信息ID不能为空"`
}
type DeviceTreeChangeRes struct{}

type DetailDeviceTreeInfoReq struct {
	g.Meta `path:"/device_tree/info/detail" method:"get" summary:"信息详情" tags:"设备树"`
	InfoId int `json:"infoId" dc:"信息ID" v:"required#信息ID不能为空"`
}
type DetailDeviceTreeInfoRes struct {
	Data *model.DetailDeviceTreeInfoOutput `json:"data" dc:"信息详情"`
}

type AddDeviceTreeInfoReq struct {
	g.Meta `path:"/device_tree/info/add" method:"post" summary:"添加设备树基本信息" tags:"设备树"`
	*model.AddDeviceTreeInfoInput
}
type AddDeviceTreeInfoRes struct{}

type EditDeviceTreeInfoReq struct {
	g.Meta `path:"/device_tree/info/edit" method:"put" summary:"编辑设备树基本信息" tags:"设备树"`
	*model.EditDeviceTreeInfoInput
}
type EditDeviceTreeInfoRes struct{}

type DelDeviceTreeInfoReq struct {
	g.Meta `path:"/device_tree/info/del" method:"delete" summary:"删除设备树基本信息" tags:"设备树"`
	Id     int `json:"id" dc:"信息ID" v:"required#信息ID不能为空"`
}
type DelDeviceTreeInfoRes struct{}
