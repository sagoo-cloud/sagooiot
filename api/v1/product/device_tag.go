package product

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AddTagDeviceReq struct {
	g.Meta `path:"/device/tag/add" method:"post" summary:"标签添加" tags:"设备"`
	*model.AddTagDeviceInput
}
type AddTagDeviceRes struct{}

type EditTagDeviceReq struct {
	g.Meta `path:"/device/tag/edit" method:"put" summary:"标签编辑" tags:"设备"`
	*model.EditTagDeviceInput
}
type EditTagDeviceRes struct{}

type DelTagDeviceReq struct {
	g.Meta `path:"/device/tag/del" method:"delete" summary:"标签删除" tags:"设备"`
	Id     uint `json:"id" dc:"标签ID" v:"required#标签ID不能为空"`
}
type DelTagDeviceRes struct{}
