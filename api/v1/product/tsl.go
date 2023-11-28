package product

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据类型

type DateTypeReq struct {
	g.Meta `path:"/tsl/data_type" method:"get" summary:"数据类型" tags:"物模型"`
}
type DateTypeRes struct {
	DataType *model.DataTypeOutput `json:"dataType" dc:"数据类型"`
}

// 属性

type ListTSLPropertyReq struct {
	g.Meta `path:"/tsl/property/list" method:"get" summary:"属性列表" tags:"物模型"`
	*model.ListTSLPropertyInput
}
type ListTSLPropertyRes struct {
	*model.ListTSLPropertyOutput
}

type AllTSLPropertyReq struct {
	g.Meta `path:"/tsl/property/all" method:"get" summary:"所有属性列表" tags:"物模型"`
	Key    string `json:"key" dc:"产品标识" v:"required#产品标识不能为空"`
}
type AllTSLPropertyRes struct {
	Data []model.TSLProperty
}

type AddTSLPropertyReq struct {
	g.Meta `path:"/tsl/property/add" method:"post" summary:"属性添加" tags:"物模型"`
	*model.TSLPropertyInput
}
type AddTSLPropertyRes struct{}

type EditTSLPropertyReq struct {
	g.Meta `path:"/tsl/property/edit" method:"put" summary:"属性编辑" tags:"物模型"`
	*model.TSLPropertyInput
}
type EditTSLPropertyRes struct{}

type DelTSLPropertyReq struct {
	g.Meta `path:"/tsl/property/del" method:"delete" summary:"属性删除" tags:"物模型"`
	*model.DelTSLPropertyInput
}
type DelTSLPropertyRes struct{}

// 功能

type ListTSLFunctionReq struct {
	g.Meta `path:"/tsl/function/list" method:"get" summary:"功能列表" tags:"物模型"`
	*model.ListTSLFunctionInput
}
type ListTSLFunctionRes struct {
	*model.ListTSLFunctionOutput
}

type AddTSLFunctionReq struct {
	g.Meta `path:"/tsl/function/add" method:"post" summary:"功能添加" tags:"物模型"`
	*model.TSLFunctionAddInput
}
type AddTSLFunctionRes struct{}

type EditTSLFunctionReq struct {
	g.Meta `path:"/tsl/function/edit" method:"put" summary:"功能编辑" tags:"物模型"`
	*model.TSLFunctionAddInput
}
type EditTSLFunctionRes struct{}

type DelTSLFunctionReq struct {
	g.Meta `path:"/tsl/function/del" method:"delete" summary:"功能删除" tags:"物模型"`
	*model.DelTSLFunctionInput
}
type DelTSLFunctionRes struct{}

// 事件

type ListTSLEventReq struct {
	g.Meta `path:"/tsl/event/list" method:"get" summary:"事件列表" tags:"物模型"`
	*model.ListTSLEventInput
}
type ListTSLEventRes struct {
	*model.ListTSLEventOutput
}

type AddTSLEventReq struct {
	g.Meta `path:"/tsl/event/add" method:"post" summary:"事件添加" tags:"物模型"`
	*model.TSLEventInput
}
type AddTSLEventRes struct{}

type EditTSLEventReq struct {
	g.Meta `path:"/tsl/event/edit" method:"put" summary:"事件编辑" tags:"物模型"`
	*model.TSLEventInput
}
type EditTSLEventRes struct{}

type DelTSLEventReq struct {
	g.Meta `path:"/tsl/event/del" method:"delete" summary:"事件删除" tags:"物模型"`
	*model.DelTSLEventInput
}
type DelTSLEventRes struct{}

// 标签

type ListTSLTagReq struct {
	g.Meta `path:"/tsl/tag/list" method:"get" summary:"标签列表" tags:"物模型"`
	*model.ListTSLTagInput
}
type ListTSLTagRes struct {
	*model.ListTSLTagOutput
}

type AddTSLTagReq struct {
	g.Meta `path:"/tsl/tag/add" method:"post" summary:"标签添加" tags:"物模型"`
	*model.TSLTagInput
}
type AddTSLTagRes struct{}

type EditTSLTagReq struct {
	g.Meta `path:"/tsl/tag/edit" method:"put" summary:"标签编辑" tags:"物模型"`
	*model.TSLTagInput
}
type EditTSLTagRes struct{}

type DelTSLTagReq struct {
	g.Meta `path:"/tsl/tag/del" method:"delete" summary:"标签删除" tags:"物模型"`
	*model.DelTSLTagInput
}
type DelTSLTagRes struct{}

type AllTSLFunctionReq struct {
	g.Meta           `path:"/tsl/function/all" method:"get" summary:"所有功能列表" tags:"物模型"`
	Key              string `json:"key" dc:"产品标识" v:"required#产品标识不能为空"`
	InputsValueTypes string `json:"inputsValueTypes" dc:"参数值类型"`
}
type AllTSLFunctionRes struct {
	Data []model.TSLFunction
}
