package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
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
	data, err := service.SysPlugins().GetSysPluginsById(ctx, req.Id)
	res = new(system.GetSysPluginsByIdRes)
	err = gconv.Scan(data, &res)
	return
}

// EditSysPluginsStatus  修改插件的状态
func (a *cMenu) EditSysPluginsStatus(ctx context.Context, req *system.EditSysPluginsStatusReq) (res *system.EditSysPluginsStatusRes, err error) {
	err = service.SysPlugins().EditStatus(ctx, req.Id, req.Status)
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
