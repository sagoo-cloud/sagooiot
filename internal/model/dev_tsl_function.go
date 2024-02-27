package model

// 功能：添加、编辑
type TSLFunctionAddInput struct {
	ProductKey string `json:"productKey" dc:"产品Key" v:"required#产品Key不能为空"`
	TSLFunction
}

// 功能：删除
type DelTSLFunctionInput struct {
	ProductKey string `json:"productKey" dc:"产品Key" v:"required#产品Key不能为空"`
	Key        string `json:"key" dc:"功能标识" v:"required#功能标识不能为空"`
}

type ListTSLFunctionInput struct {
	ProductKey string `json:"productKey" dc:"产品Key" v:"required#产品Key不能为空"`
	PaginationInput
}
type ListTSLFunctionOutput struct {
	Data []TSLFunction
	PaginationOutput
}
