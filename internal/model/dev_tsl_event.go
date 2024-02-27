package model

// TSLEventAddInput 事件：添加、编辑
type TSLEventAddInput struct {
	ProductKey string `json:"productKey" dc:"产品Key" v:"required#产品Key不能为空"`
	TSLEvent
}

// DelTSLEventInput 事件：删除
type DelTSLEventInput struct {
	ProductKey string `json:"productKey" dc:"产品Key" v:"required#产品Key不能为空"`
	Key        string `json:"key" dc:"事件标识" v:"required#事件标识不能为空"`
}

type ListTSLEventInput struct {
	ProductKey string `json:"productKey" dc:"产品Key" v:"required#产品Key不能为空"`
	PaginationInput
}
type ListTSLEventOutput struct {
	Data []TSLEvent
	PaginationOutput
}
