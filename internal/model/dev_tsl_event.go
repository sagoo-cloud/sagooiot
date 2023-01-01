package model

// 事件：添加、编辑
type TSLEventInput struct {
	ProductId uint `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	TSLEvent
}

// 事件：删除
type DelTSLEventInput struct {
	ProductId uint   `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	Key       string `json:"key" dc:"事件标识" v:"required#事件标识不能为空"`
}

type ListTSLEventInput struct {
	ProductId uint `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	PaginationInput
}
type ListTSLEventOutput struct {
	Data []TSLEvent
	PaginationOutput
}
