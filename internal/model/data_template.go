package model

import (
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	DataTemplateStatusOff int = iota // 数据模型未发布
	DataTemplateStatusOn             // 数据模型已发布
)

type DataTemplate struct {
	*entity.DataTemplate

	// 绑定的业务
	DataTemplateBusi []DataTemplateBusi `json:"dataTemplateBusi" orm:"with:data_template_id=id" dc:"绑定的业务单元"`
	BusiTypes        []int              `json:"busiTypes" dc:"业务单元"`
}

// 绑定业务模型
type DataTemplateBusi struct {
	DataTemplateId uint64 `json:"dataTemplateId" dc:"数据模型ID"`
	BusiTypes      int    `json:"busiTypes" dc:"业务单元"`
}

// 搜索数据模型
type DataTemplateSearchInput struct {
	Key  string `json:"key" dc:"数据模型标识"`
	Name string `json:"name" dc:"数据模型名称"`
	PaginationInput
}
type DataTemplateSearchOutput struct {
	List []DataTemplate `json:"list" dc:"数据模型列表"`
	PaginationOutput
}

// 添加数据模型
type DataTemplateAddInput struct {
	Key  string `json:"key" dc:"数据模型标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入数据模型标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name string `json:"name" dc:"数据模型名称" v:"required#请输入数据模型名称"`
	Desc string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`

	// 数据更新间隔，cron 格式
	CronExpression string `json:"cronExpression" dc:"任务执行表达式" v:"required#任务表达式不能为空"`
}

// 编辑数据模型
type DataTemplateEditInput struct {
	Id   uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
	Name string `json:"name" dc:"数据模型名称" v:"required#请输入数据模型名称"`
	Desc string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Key  string `json:"key" dc:"数据模型标识" v:"regex:^[A-Za-z_]+[\\w]*$#标识由字母、数字和下划线组成,且不能以数字开头"`

	// 数据更新间隔，cron 格式
	CronExpression string `json:"cronExpression" dc:"任务执行表达式" v:"required#任务表达式不能为空"`

	// 绑定业务
	BusiTypes []int `json:"busiTypes" dc:"业务单元"`
}

// 数据模型
type DataTemplateOutput struct {
	*entity.DataTemplate
}

// 数据模型获取数据的内网方法列表，供大屏使用
type AllTemplateOut struct {
	Id   uint64 `json:"id" dc:"数据模型ID"`
	Name string `json:"name" dc:"数据模型名称"`
	Path string `json:"path" dc:"接口地址"`
}

type TemplateDataAllInput struct {
	Id    uint64                 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
	Param map[string]interface{} `json:"param" dc:"搜索哪些字段的数据"`
}
type TemplateDataAllOutput struct {
	List g.List `json:"data" dc:"模型数据记录"`
}

type TemplateDataLastInput struct {
	Id    uint64                 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
	Param map[string]interface{} `json:"param" dc:"搜索哪些字段的数据"`
}
type TemplateDataLastOutput struct {
	Data g.Map `json:"data" dc:"模型数据记录"`
}

// 数据模型设置主源、关联字段
type TemplateDataRelationInput struct {
	Id            uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
	MainSourceId  uint64 `json:"mainSourceId" dc:"主数据源" v:"required#主数据源ID不能为空"`
	SourceNodeKey string `json:"sourceNodeKey" dc:"关联节点" v:"required#关联节点标识不能为空"`
}

type DataTemplateDataInput struct {
	Id    uint64                 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
	Param map[string]interface{} `json:"param" dc:"搜索哪些字段的数据"`
	PaginationInput
}
type DataTemplateDataOutput struct {
	List string `json:"data" dc:"模型数据记录"`
	PaginationOutput
}
