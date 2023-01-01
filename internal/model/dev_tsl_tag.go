package model

// 添加、编辑标签
type TSLTagInput struct {
	ProductId uint `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	TSLTag
}

// 删除标签
type DelTSLTagInput struct {
	ProductId uint   `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	Key       string `json:"key" dc:"标签标识" v:"required#标签标识不能为空"`
}

type ListTSLTagInput struct {
	ProductId uint `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	PaginationInput
}
type ListTSLTagOutput struct {
	Data []TSLTag
	PaginationOutput
}
