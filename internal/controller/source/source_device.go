package source

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/source"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

// 添加 设备 数据源
func (c *cDataSource) AddDevice(ctx context.Context, req *source.DataSourceDeviceAddReq) (res *source.DataSourceDeviceAddRes, err error) {
	_, err = service.DataSource().AddDevice(ctx, req.DataSourceDeviceAddInput)
	return
}

// 编辑 设备 数据源
func (c *cDataSource) EditDevice(ctx context.Context, req *source.DataSourceDeviceEditReq) (res *source.DataSourceDeviceEditRes, err error) {
	err = service.DataSource().EditDevice(ctx, req.DataSourceDeviceEditInput)
	return
}

// 获取 设备 数据
func (c *cDataSource) GetDeviceData(ctx context.Context, req *source.DataSourceDeviceGetReq) (res *source.DataSourceDeviceGetRes, err error) {
	res = new(source.DataSourceDeviceGetRes)
	res.Data, err = service.DataSource().GetDeviceData(ctx, req.SourceId)
	return
}
