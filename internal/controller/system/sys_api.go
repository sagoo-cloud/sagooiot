package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var SysApi = cApi{}

type cApi struct{}

// GetApiAll 获取所有接口
func (a *cApi) GetApiAll(ctx context.Context, req *system.GetApiAllReq) (res *system.GetApiAllRes, err error) {
	apiInfo, err := service.SysApi().GetApiAll(ctx)
	if err != nil {
		return
	}
	var apiInfoRes []*model.SysApiAllRes
	if apiInfo != nil {
		if err = gconv.Scan(apiInfo, &apiInfoRes); err != nil {
			return
		}
	}
	res = &system.GetApiAllRes{
		Data: apiInfoRes,
	}
	return
}

// GetApiTree 获取接口树状结构
func (a *cApi) GetApiTree(ctx context.Context, req *system.GetApiTreeReq) (res *system.GetApiTreeRes, err error) {
	out, err := service.SysApi().GetApiTree(ctx, req.Name, req.Address, req.Status, req.Types)
	if err != nil {
		return nil, err
	}
	var treeData []*model.SysApiTreeRes
	if out != nil {
		if err = gconv.Scan(out, &treeData); err != nil {
			return
		}
	}
	res = &system.GetApiTreeRes{
		Info: treeData,
	}
	return
}

// AddApi 添加Api列表
func (a *cApi) AddApi(ctx context.Context, req *system.AddApiReq) (res *system.AddApiRes, err error) {
	var input *model.AddApiInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysApi().Add(ctx, input)
	return
}

// DetailApi 获取Api列表详情
func (a *cApi) DetailApi(ctx context.Context, req *system.DetailApiReq) (res *system.DetailApiRes, err error) {
	out, err := service.SysApi().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if out != nil {
		var detailRes *model.SysApiRes
		if err = gconv.Scan(out, &detailRes); err != nil {
			return nil, err
		}
		res = &system.DetailApiRes{
			Data: detailRes,
		}
	}
	return
}

// EditApi 编辑Api
func (a *cApi) EditApi(ctx context.Context, req *system.EditApiReq) (res *system.EditApiRes, err error) {
	var input *model.EditApiInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysApi().Edit(ctx, input)
	return
}

// DelApi 根据ID删除Api
func (a *cApi) DelApi(ctx context.Context, req *system.DelApiReq) (res *system.DelApiRes, err error) {
	err = service.SysApi().Del(ctx, req.Id)
	return
}

// EditApiStatus  编辑API状态
func (a *cApi) EditApiStatus(ctx context.Context, req *system.EditApiStatusReq) (res *system.EditApiStatusRes, err error) {
	err = service.SysApi().EditStatus(ctx, req.Id, req.Status)
	return
}
