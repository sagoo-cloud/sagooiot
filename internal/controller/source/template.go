package source

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/source"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var DataTemplate = cDataTemplate{}

type cDataTemplate struct{}

// 添加数据模型
func (c *cDataTemplate) Add(ctx context.Context, req *source.DataTemplateAddReq) (res *source.DataTemplateAddRes, err error) {
	_, err = service.DataTemplate().Add(ctx, req.DataTemplateAddInput)
	return
}

// 编辑数据模型
func (c *cDataTemplate) Edit(ctx context.Context, req *source.DataTemplateEditReq) (res *source.DataTemplateEditRes, err error) {
	err = service.DataTemplate().Edit(ctx, req.DataTemplateEditInput)
	return
}

// 批量删除数据模型
func (c *cDataTemplate) Del(ctx context.Context, req *source.DataTemplateDelReq) (res *source.DataTemplateDelRes, err error) {
	err = service.DataTemplate().Del(ctx, req.Ids)
	return
}

// 搜索数据模型
func (c *cDataTemplate) Search(ctx context.Context, req *source.DataTemplateSearchReq) (res *source.DataTemplateSearchRes, err error) {
	out, err := service.DataTemplate().Search(ctx, req.DataTemplateSearchInput)
	res = &source.DataTemplateSearchRes{
		DataTemplateSearchOutput: out,
	}
	return
}

// 已发布数据模型列表
func (c *cDataTemplate) List(ctx context.Context, req *source.DataTemplateListReq) (res *source.DataTemplateListRes, err error) {
	list, err := service.DataTemplate().List(ctx)
	res = &source.DataTemplateListRes{
		List: list,
	}
	return
}

// 详情
func (c *cDataTemplate) Detail(ctx context.Context, req *source.DataTemplateReq) (res *source.DataTemplateRes, err error) {
	res = new(source.DataTemplateRes)
	res.Data, err = service.DataTemplate().Detail(ctx, req.Id)
	return
}

// 发布
func (c *cDataTemplate) Deploy(ctx context.Context, req *source.DataTemplateDeployReq) (res *source.DataTemplateDeployRes, err error) {
	err = service.DataTemplate().Deploy(ctx, req.Id)
	return
}

// 停用
func (c *cDataTemplate) Undeploy(ctx context.Context, req *source.DataTemplateUndeployReq) (res *source.DataTemplateUndeployRes, err error) {
	err = service.DataTemplate().Undeploy(ctx, req.Id)
	return
}

// 获取模型数据
func (c *cDataTemplate) GetData(ctx context.Context, req *source.DataTemplateDataReq) (res *source.DataTemplateDataRes, err error) {
	out, err := service.DataTemplate().GetData(ctx, req.DataTemplateDataInput)
	res = &source.DataTemplateDataRes{
		DataTemplateDataOutput: out,
	}
	return
}

// 复制模型
func (c *cDataTemplate) Copy(ctx context.Context, req *source.DataTemplateCopyReq) (res *source.DataTemplateCopyRes, err error) {
	err = service.DataTemplate().CopeTemplate(ctx, req.Id)
	return
}

// 检测数据模型是否需要设置关联
func (c *cDataTemplate) CheckRelation(ctx context.Context, req *source.DataTemplateCheckRelationReq) (res source.DataTemplateCheckRelationRes, err error) {
	res.Yes, err = service.DataTemplate().CheckRelation(ctx, req.Id)
	return
}

// 设置主源、关联字段
func (c *cDataTemplate) SetRelation(ctx context.Context, req *source.DataTemplateRelationReq) (res *source.DataTemplateRelationRes, err error) {
	err = service.DataTemplate().SetRelation(ctx, req.TemplateDataRelationInput)
	return
}

// 数据源列表
func (c *cDataTemplate) SourceList(ctx context.Context, req *source.TemplateSourceListReq) (res source.TemplateSourceListRes, err error) {
	res.List, err = service.DataTemplate().SourceList(ctx, req.Id)
	return
}
