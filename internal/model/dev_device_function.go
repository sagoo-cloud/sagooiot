package model

type DeviceFunctionInput struct {
	DeviceKey string         `json:"deviceKey" dc:"设备标识" v:"required#设备标识不能为空"`
	FuncKey   string         `json:"funcKey" dc:"功能标识" v:"required#功能标识不能为空"`
	Params    map[string]any `json:"params" dc:"功能输入参数"`
}
type DeviceFunctionOutput struct {
	Data map[string]any `json:"data" dc:"功能输出"`
}
