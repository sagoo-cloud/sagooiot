package product

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DevicePropertyReq struct {
	g.Meta `path:"/property/set" method:"post" summary:"设备属性设置" tags:"设备"`
	*model.DevicePropertyInput
}
type DevicePropertyRes struct {
	*model.DevicePropertyOutput
}
