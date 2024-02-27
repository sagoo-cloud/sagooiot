package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysMenuColumn = cMenuColumn{}

type cMenuColumn struct{}

// MenuColumnTree 获取菜单列表树结构列表
func (c *cMenuColumn) MenuColumnTree(ctx context.Context, req *system.MenuColumnDoReq) (res *system.MenuColumnDoRes, err error) {
	var input *model.MenuColumnDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	out, err := service.SysMenuColumn().GetList(ctx, input)
	if err != nil {
		return
	}
	var data []*model.UserMenuColumnRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &system.MenuColumnDoRes{
		Data: data,
	}
	return
}

// AddMenuColumn AddMenu 添加菜单列表
func (c *cMenuColumn) AddMenuColumn(ctx context.Context, req *system.AddMenuColumnReq) (res *system.AddMenuColumnRes, err error) {
	var input *model.AddMenuColumnInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuColumn().Add(ctx, input)
	return
}

// DetailMenuColumn DetailMenu 获取菜单列表详情
func (c *cMenuColumn) DetailMenuColumn(ctx context.Context, req *system.DetailMenuColumnReq) (res *system.DetailMenuColumnRes, err error) {
	data, err := service.SysMenuColumn().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *model.DetailMenuColumnRes
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &system.DetailMenuColumnRes{
			Data: detailRes,
		}
	}
	return
}

// EditMenuColumn EditMenu 编辑菜单列表
func (c *cMenuColumn) EditMenuColumn(ctx context.Context, req *system.EditMenuColumnReq) (res *system.EditMenuColumnRes, err error) {
	var input *model.EditMenuColumnInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuColumn().Edit(ctx, input)
	return
}

// DelMenuColumn DelMenu 根据ID删除菜单列表
func (c *cMenuColumn) DelMenuColumn(ctx context.Context, req *system.DelMenuColumnReq) (res *system.DelMenuColumnRes, err error) {
	err = service.SysMenuColumn().Del(ctx, req.Id)
	return
}

// EditMenuColumnStatus  编辑菜单列表状态
func (c *cMenuColumn) EditMenuColumnStatus(ctx context.Context, req *system.EditMenuColumnStatusReq) (res *system.EditMenuColumnStatusRes, err error) {
	err = service.SysMenuColumn().EditStatus(ctx, req.Id, req.MenuId, req.Status)
	return
}
