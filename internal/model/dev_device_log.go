package model

type DeviceLogSearchInput struct {
	DeviceKey string   `json:"deviceKey" dc:"设备标识" v:"required#设备标识不能为空"`
	Types     []string `json:"types" dc:"日志类型"`
	PaginationInput
}
type DeviceLogSearchOutput struct {
	List []TdLog `json:"list" dc:"日志类型列表"`
	PaginationOutput
}
