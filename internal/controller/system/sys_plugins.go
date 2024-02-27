package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysPlugins = cSystemSysPlugins{}

type cSystemSysPlugins struct{}

// GetSysPluginsList 获取列表
func (u *cSystemSysPlugins) GetSysPluginsList(ctx context.Context, req *system.GetSysPluginsListReq) (res *system.GetSysPluginsListRes, err error) {
	var inputData = new(model.GetSysPluginsListInput)
	if err = gconv.Scan(req, &inputData); err != nil {
		return
	}
	total, currentPage, dataList, err := service.SysPlugins().GetSysPluginsList(ctx, inputData)
	res = new(system.GetSysPluginsListRes)
	err = gconv.Scan(dataList, &res.Data)
	res.PaginationRes.Total = total
	res.PaginationRes.CurrentPage = currentPage
	return
}

// GetSysPluginsById 获取指定ID数据
func (u *cSystemSysPlugins) GetSysPluginsById(ctx context.Context, req *system.GetSysPluginsByIdReq) (res *system.GetSysPluginsByIdRes, err error) {
	out, err := service.SysPlugins().GetSysPluginsById(ctx, req.Id)
	if err != nil {
		return
	}
	if out != nil {
		res = new(system.GetSysPluginsByIdRes)
		err = gconv.Scan(out, &res)
	}

	return
}

// AddSysPlugins 添加插件
func (u *cSystemSysPlugins) AddSysPlugins(ctx context.Context, req *system.AddSysPluginsReq) (res *system.AddSysPluginsRes, err error) {
	err = service.SysPlugins().AddSysPlugins(ctx, req.File)
	if err != nil {
		return
	}

	return
}

// EditSysPluginsStatus  修改插件的状态
func (u *cSystemSysPlugins) EditSysPluginsStatus(ctx context.Context, req *system.EditSysPluginsStatusReq) (res *system.EditSysPluginsStatusRes, err error) {
	err = service.SysPlugins().EditStatus(ctx, req.Id, req.Status)
	return
}

// EditSysPlugins 添加插件
func (u *cSystemSysPlugins) EditSysPlugins(ctx context.Context, req *system.EditSysPluginsReq) (res *system.EditSysPluginsRes, err error) {
	var input *model.SysPluginsEditInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysPlugins().EditSysPlugins(ctx, input)
	if err != nil {
		return
	}

	return
}

// DelSysPluginsStatus  删除插件
func (u *cSystemSysPlugins) DelSysPluginsStatus(ctx context.Context, req *system.DelSysPluginsStatusReq) (res *system.DelSysPluginsStatusRes, err error) {
	err = service.SysPlugins().DeleteSysPlugins(ctx, req.Ids)
	return
}

// GetSysPluginsTypesAll 获取插件通信方式类型
func (u *cSystemSysPlugins) GetSysPluginsTypesAll(ctx context.Context, req *system.GetSysPluginsTypesAllReq) (res *system.GetSysPluginsTypesAllRes, err error) {
	out, err := service.SysPlugins().GetSysPluginsTypesAll(ctx, req.Types)
	if err != nil {
		return
	}
	var data []*model.SysPluginsInfoRes
	if len(out) > 0 {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &system.GetSysPluginsTypesAllRes{
		Data: data,
	}
	return
}
