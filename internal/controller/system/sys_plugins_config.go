package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var SysPluginsConfig = cSystemPlugins_config{}

type cSystemPlugins_config struct{}

// GetPluginsConfigList 获取列表
func (u *cSystemPlugins_config) GetPluginsConfigList(ctx context.Context, req *system.GetPluginsConfigListReq) (res *system.GetPluginsConfigListRes, err error) {
	var inputData = new(model.GetPluginsConfigListInput)
	if err = gconv.Scan(req, &inputData); err != nil {
		return
	}
	total, currentPage, dataList, err := service.SystemPluginsConfig().GetPluginsConfigList(ctx, inputData)
	res = new(system.GetPluginsConfigListRes)
	err = gconv.Scan(dataList, &res.Data)
	res.PaginationRes.Total = total
	res.PaginationRes.CurrentPage = currentPage
	return
}

// GetPluginsConfigById 获取指定ID数据
func (u *cSystemPlugins_config) GetPluginsConfigById(ctx context.Context, req *system.GetPluginsConfigByIdReq) (res *system.GetPluginsConfigByIdRes, err error) {
	data, err := service.SystemPluginsConfig().GetPluginsConfigById(ctx, req.Id)
	res = new(system.GetPluginsConfigByIdRes)
	err = gconv.Scan(data, &res)
	return
}

// GetPluginsConfigByName 获取指定类型与名称的插件配置数据
func (u *cSystemPlugins_config) GetPluginsConfigByName(ctx context.Context, req *system.GetPluginsConfigByNameReq) (res *system.GetPluginsConfigByNameRes, err error) {
	data, err := service.SystemPluginsConfig().GetPluginsConfigByName(ctx, req.Type, req.Name)
	res = new(system.GetPluginsConfigByNameRes)
	err = gconv.Scan(data, &res)
	return
}

// AddPluginsConfig 添加数据
func (u *cSystemPlugins_config) AddPluginsConfig(ctx context.Context, req *system.AddPluginsConfigReq) (res *system.AddPluginsConfigRes, err error) {
	var data = model.PluginsConfigAddInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	err = service.SystemPluginsConfig().AddPluginsConfig(ctx, data)
	return
}

// EditPluginsConfig 修改数据
func (u *cSystemPlugins_config) EditPluginsConfig(ctx context.Context, req *system.EditPluginsConfigReq) (res *system.EditPluginsConfigRes, err error) {
	var data = model.PluginsConfigEditInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	err = service.SystemPluginsConfig().EditPluginsConfig(ctx, data)
	return
}

// SavePluginsConfig 修改数据
func (u *cSystemPlugins_config) SavePluginsConfig(ctx context.Context, req *system.SavePluginsConfigReq) (res *system.SavePluginsConfigRes, err error) {
	var data = model.PluginsConfigAddInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	err = service.SystemPluginsConfig().SavePluginsConfig(ctx, data)
	return
}

// DeletePluginsConfig 删除数据
func (u *cSystemPlugins_config) DeletePluginsConfig(ctx context.Context, req *system.DeletePluginsConfigReq) (res *system.DeletePluginsConfigRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.SystemPluginsConfig().DeletePluginsConfig(ctx, req.Ids)
	return
}
