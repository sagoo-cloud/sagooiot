package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	systemV1 "sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysOrganization = cOrganization{}

type cOrganization struct{}

func (a *cOrganization) OrganizationTree(ctx context.Context, req *systemV1.OrganizationDoReq) (res *systemV1.OrganizationDoRes, err error) {
	//获取所有的组织
	out, err := service.SysOrganization().GetTree(ctx, req.Name, req.Status)
	if err != nil {
		return nil, err
	}
	var treeData []*model.OrganizationRes
	if out != nil {
		if err = gconv.Scan(out, &treeData); err != nil {
			return
		}
	}
	res = &systemV1.OrganizationDoRes{
		Data: treeData,
	}
	return
}

// AddOrganization 添加组织
func (a *cOrganization) AddOrganization(ctx context.Context, req *systemV1.AddOrganizationReq) (res *systemV1.AddOrganizationRes, err error) {
	var input *model.AddOrganizationInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysOrganization().Add(ctx, input)
	return
}

// EditOrganization 编辑组织
func (a *cOrganization) EditOrganization(ctx context.Context, req *systemV1.EditOrganizationReq) (res *systemV1.EditOrganizationRes, err error) {
	var input *model.EditOrganizationInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysOrganization().Edit(ctx, input)
	return
}

// DetailOrganization 获取组织详情
func (a *cOrganization) DetailOrganization(ctx context.Context, req *systemV1.DetailOrganizationReq) (res *systemV1.DetailOrganizationRes, err error) {
	data, err := service.SysOrganization().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *model.DetailOrganizationRes
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &systemV1.DetailOrganizationRes{
			Data: detailRes,
		}
	}
	return
}

// DelOrganization 根据ID删除组织
func (a *cOrganization) DelOrganization(ctx context.Context, req *systemV1.DelOrganizationReq) (res *systemV1.DelOrganizationRes, err error) {
	err = service.SysOrganization().Del(ctx, req.Id)
	return
}

// GetCount 获取组织数量
func (a *cOrganization) GetCount(ctx context.Context, req *systemV1.GetOrganizationCountReq) (res *systemV1.GetOrganizationCountRes, err error) {
	count, err := service.SysOrganization().Count(ctx)
	if err != nil {
		return
	}
	res = &systemV1.GetOrganizationCountRes{
		Count: count,
	}
	return
}
