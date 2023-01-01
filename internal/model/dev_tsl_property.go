package model

// 添加、编辑属性
type TSLPropertyInput struct {
	ProductId uint `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	TSLProperty
}

// 删除属性
type DelTSLPropertyInput struct {
	ProductId uint   `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	Key       string `json:"key" dc:"属性标识" v:"required#属性标识不能为空"`
}

type ListTSLPropertyInput struct {
	ProductId uint   `json:"productId" dc:"产品ID" v:"required#产品ID不能为空"`
	Name      string `json:"name" dc:"属性名称"`
	DateType  string `json:"dateType" dc:"数据类型"`
	PaginationInput
}
type ListTSLPropertyOutput struct {
	Data []TSLProperty
	PaginationOutput
}
