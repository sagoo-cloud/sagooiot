package source

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/source"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var DataSource = cDataSource{}

type cDataSource struct{}

// 添加 api 数据源
func (c *cDataSource) Add(ctx context.Context, req *source.DataSourceApiAddReq) (res *source.DataSourceApiAddRes, err error) {
	_, err = service.DataSource().Add(ctx, req.DataSourceApiAddInput)
	return
}

// 编辑 api 数据源
func (c *cDataSource) Edit(ctx context.Context, req *source.DataSourceApiEditReq) (res *source.DataSourceApiEditRes, err error) {
	err = service.DataSource().Edit(ctx, req.DataSourceApiEditInput)
	return
}

// 获取 api 数据
func (c *cDataSource) GetApiData(ctx context.Context, req *source.DataSourceApiGetReq) (res *source.DataSourceApiGetRes, err error) {
	res = new(source.DataSourceApiGetRes)
	data, err := service.DataSource().GetApiData(ctx, req.SourceId)
	if err != nil {
		return
	}
	if len(data) > 0 {
		res.Data = data[0]
	}
	return
}

// 批量删除数据源
func (c *cDataSource) Del(ctx context.Context, req *source.DataSourceDelReq) (res *source.DataSourceDelRes, err error) {
	err = service.DataSource().Del(ctx, req.Ids)
	return
}

// 搜索数据源
func (c *cDataSource) Search(ctx context.Context, req *source.DataSourceSearchReq) (res *source.DataSourceSearchRes, err error) {
	out, err := service.DataSource().Search(ctx, req.DataSourceSearchInput)
	res = &source.DataSourceSearchRes{
		DataSourceSearchOutput: out,
	}
	return
}

// 数据源列表
func (c *cDataSource) List(ctx context.Context, req *source.DataSourceListReq) (res *source.DataSourceListRes, err error) {
	list, err := service.DataSource().List(ctx)
	res = &source.DataSourceListRes{
		List: list,
	}
	return
}

// 详情
func (c *cDataSource) Detail(ctx context.Context, req *source.DataSourceReq) (res *source.DataSourceRes, err error) {
	res = new(source.DataSourceRes)
	res.Data, err = service.DataSource().Detail(ctx, req.SourceId)
	return
}

// 发布
func (c *cDataSource) Deploy(ctx context.Context, req *source.DataSourceDeployReq) (res *source.DataSourceDeployRes, err error) {
	err = service.DataSource().Deploy(ctx, req.SourceId)
	return
}

// 停用
func (c *cDataSource) Undeploy(ctx context.Context, req *source.DataSourceUndeployReq) (res *source.DataSourceUndeployRes, err error) {
	err = service.DataSource().Undeploy(ctx, req.SourceId)
	return
}

// 获取源数据记录
func (c *cDataSource) GetData(ctx context.Context, req *source.DataSourceDataReq) (res *source.DataSourceDataRes, err error) {
	out, err := service.DataSource().GetData(ctx, req.DataSourceDataInput)
	res = &source.DataSourceDataRes{
		DataSourceDataOutput: out,
	}
	return
}

// 复制数据源
func (c *cDataSource) Copy(ctx context.Context, req *source.DataSourceCopyReq) (res *source.DataSourceCopyRes, err error) {
	err = service.DataSource().CopeSource(ctx, req.SourceId)
	return
}
