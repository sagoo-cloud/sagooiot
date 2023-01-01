package model

import "github.com/sagoo-cloud/sagooiot/internal/model/entity"

// 添加节点
type DataNodeAddInput struct {
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
	Key      string `json:"key" dc:"数据节点标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入数据节点标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name     string `json:"name" dc:"数据节点名称" v:"required#请输入数据节点名称"`
	DataType string `json:"dataType" dc:"数据类型" v:"required#请选择数据类型"`
	Value    string `json:"value" dc:"取值项" v:"required#请输入取值项"`
	IsPk     int    `json:"isPk" dc:"是否主键"`

	Rule []DataSourceRule `json:"rule" dc:"规则配置"`
}

// 编辑节点
type DataNodeEditInput struct {
	NodeId uint64 `json:"nodeId" dc:"数据节点ID" v:"required#数据节点ID不能为空"`
	Name   string `json:"name" dc:"数据节点名称" v:"required#请输入数据节点名称"`
	Value  string `json:"value" dc:"取值项" v:"required#请输入取值项"`
}

// 数据节点
type DataNodeOutput struct {
	*entity.DataNode

	NodeRule []*DataSourceRule `json:"nodeRule" dc:"数据节点规则配置"`
}
