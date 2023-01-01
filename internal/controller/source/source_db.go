package source

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/source"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

// 添加 数据库 数据源
func (c *cDataSource) AddDb(ctx context.Context, req *source.DataSourceDbAddReq) (res *source.DataSourceDbAddRes, err error) {
	_, err = service.DataSource().AddDb(ctx, req.DataSourceDbAddInput)
	return
}

// 编辑 数据库 数据源
func (c *cDataSource) EditDb(ctx context.Context, req *source.DataSourceDbEditReq) (res *source.DataSourceDbEditRes, err error) {
	err = service.DataSource().EditDb(ctx, req.DataSourceDbEditInput)
	return
}

// 获取 数据库 数据
func (c *cDataSource) GetDbData(ctx context.Context, req *source.DataSourceDbGetReq) (res *source.DataSourceDbGetRes, err error) {
	res = new(source.DataSourceDbGetRes)
	res.Data, err = service.DataSource().GetDbData(ctx, req.SourceId)
	return
}

// 获取 数据表 字段
func (c *cDataSource) GetDbFields(ctx context.Context, req *source.DataSourceDbFieldsReq) (res *source.DataSourceDbFieldsRes, err error) {
	res = new(source.DataSourceDbFieldsRes)
	res.Data, err = service.DataSource().GetDbFields(ctx, req.SourceId)
	return
}
