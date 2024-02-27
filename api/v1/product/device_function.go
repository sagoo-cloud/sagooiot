package product

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DeviceFunctionReq struct {
	g.Meta `path:"/function/do" method:"post" summary:"设备功能执行" tags:"设备"`
	*model.DeviceFunctionInput
}
type DeviceFunctionRes struct {
	*model.DeviceFunctionOutput
}
