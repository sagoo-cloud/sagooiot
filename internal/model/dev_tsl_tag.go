package model

// TSLTagInput 添加、编辑标签
type TSLTagInput struct {
	ProductKey string `json:"productKey" dc:"产品标识key" v:"required#产品KEY不能为空"`
	TSLTag
}

// DelTSLTagInput 删除标签
type DelTSLTagInput struct {
	ProductKey string `json:"productKey" dc:"产品标识key" v:"required#产品KEY不能为空"`
	Key        string `json:"key" dc:"标签标识" v:"required#标签标识不能为空"`
}

type ListTSLTagInput struct {
	ProductKey string `json:"productKey" dc:"产品标识key" v:"required#产品KEY不能为空"`
	PaginationInput
}
type ListTSLTagOutput struct {
	Data []TSLTag
	PaginationOutput
}
