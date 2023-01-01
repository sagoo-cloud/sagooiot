package model

// 绑定业务模型
type DataTemplateBusiAddInput struct {
	DataTemplateId uint64 `json:"dataTemplateId" dc:"数据模型ID"`
	BusiTypes      []int  `json:"busiTypes" dc:"业务单元"`
}
