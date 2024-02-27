package model

type AddTagDeviceInput struct {
	DeviceId  uint   `json:"deviceId" dc:"设备ID" v:"required#设备ID不能为空"`
	DeviceKey string `json:"deviceKey" dc:"设备标识" v:"required#设备标识不能为空"`
	Key       string `json:"key" dc:"标签标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入标签标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name      string `json:"name" dc:"标签名称" v:"required#请输入标签名称"`
	Value     string `json:"value" dc:"标签值" v:"required#请输入标签值"`
}

type EditTagDeviceInput struct {
	Id    uint   `json:"id" dc:"标签ID" v:"required#标签ID不能为空"`
	Name  string `json:"name" dc:"标签名称" v:"required#请输入标签名称"`
	Value string `json:"value" dc:"标签值" v:"required#请输入标签值"`
}

type AddTagInput struct {
	Key   string `json:"key" dc:"标签标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入标签标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name  string `json:"name" dc:"标签名称" v:"required#请输入标签名称"`
	Value string `json:"value" dc:"标签值" v:"required#请输入标签值"`
}
