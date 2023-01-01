package model

import (
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

// 添加节点
type DataTemplateNodeAddInput struct {
	Tid      uint64 `json:"tid" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
	From     int    `json:"from" dc:"字段生成方式:1=自动生成,2=数据源" v:"required|in:1,2#请选择字段生成方式|未知方式，请正确选择"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required-if:from,2#数据源ID不能为空"`
	NodeId   uint64 `json:"nodeId" dc:"数据节点ID" v:"required-if:from,2#数据节点ID不能为空"`
	Key      string `json:"key" dc:"模型节点标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入模型节点标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name     string `json:"name" dc:"模型节点名称" v:"required#请输入模型节点名称"`
	DataType string `json:"dataType" dc:"数据类型" v:"required#请选择数据类型"`
	Default  string `json:"default" dc:"默认值"`
	Method   string `json:"method" dc:"数值类型，取值方式：max、min、avg"`
	IsPk     int    `json:"isPk" dc:"是否主键"`
	Desc     string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`

	IsSorting int `json:"isSorting" dc:"是否参与排序:0=否,1=是" v:"required|in:0,1#请选择是否参与排序|请正确选择"`
	IsDesc    int `json:"isDesc" dc:"排序方式:1=倒序,2=正序" v:"required-if:isSorting,1#请选择排序方式"`
}

// 编辑节点
type DataTemplateNodeEditInput struct {
	Id       uint64 `json:"id" dc:"模型节点ID" v:"required#模型节点ID不能为空"`
	From     int    `json:"from" dc:"字段生成方式:1=自动生成,2=数据源" v:"required|in:1,2#请选择字段生成方式|未知方式，请正确选择"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required-if:from,2#数据源ID不能为空"`
	NodeId   uint64 `json:"nodeId" dc:"数据节点ID" v:"required-if:from,2#数据节点ID不能为空"`
	Name     string `json:"name" dc:"模型节点名称" v:"required#请输入模型节点名称"`
	Default  string `json:"default" dc:"默认值"`
	Desc     string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`

	IsSorting int `json:"isSorting" dc:"是否参与排序:0=否,1=是" v:"required|in:0,1#请选择是否参与排序|请正确选择"`
	IsDesc    int `json:"isDesc" dc:"排序方式:1=倒序,2=正序" v:"required-if:isSorting,1#请选择排序方式"`
}

// 数据模型节点
type DataTemplateNodeOutput struct {
	*entity.DataTemplateNode
	IsSorting int `json:"isSorting" dc:"是否参与排序:0=否,1=是"`
	IsDesc    int `json:"isDesc" dc:"排序方式:1=倒序,2=正序"`

	Source *WithSource `json:"source" orm:"with:source_id, where:source_id>0" dc:"数据源"`
	Node   *WithNode   `json:"node" orm:"with:node_id, where:node_id>0" dc:"数据源节点"`
}

type WithSource struct {
	gmeta.Meta `orm:"table:data_source"`
	SourceId   uint64 `json:"sourceId" dc:"数据源ID"`
	Key        string `json:"key" dc:"数据源标识"`
	Name       string `json:"name" dc:"数据源名称"`
	From       int    `json:"from" dc:"数据来源:1=api导入,2=数据库,3=文件,4=设备"`
}
type WithNode struct {
	gmeta.Meta `orm:"table:data_node"`
	NodeId     uint64 `json:"nodeId" dc:"数据节点ID"`
	Key        string `json:"key" dc:"数据节点标识"`
	Name       string `json:"name" dc:"数据节点名称"`
}
