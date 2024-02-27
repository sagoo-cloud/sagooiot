package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysMenuButton = cMenuButton{}

type cMenuButton struct{}

// MenuButtonTree 获取菜单按钮树结构列表
func (c *cMenuButton) MenuButtonTree(ctx context.Context, req *system.MenuButtonDoReq) (res *system.MenuButtonDoRes, err error) {
	out, err := service.SysMenuButton().GetList(ctx, req.Status, req.Name, req.MenuId)
	if err != nil {
		return
	}
	var data []*model.UserMenuButtonRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &system.MenuButtonDoRes{
		Data: data,
	}
	return
}

// AddMenuButton AddMenu 添加菜单按钮
func (c *cMenuButton) AddMenuButton(ctx context.Context, req *system.AddMenuButtonReq) (res *system.AddMenuButtonRes, err error) {
	var input *model.AddMenuButtonInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuButton().Add(ctx, input)
	return
}

// DetailMenuButton DetailMenu 获取菜单按钮详情
func (c *cMenuButton) DetailMenuButton(ctx context.Context, req *system.DetailMenuButtonReq) (res *system.DetailMenuButtonRes, err error) {
	data, err := service.SysMenuButton().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *model.DetailMenuButtonRes
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &system.DetailMenuButtonRes{
			Data: detailRes,
		}
	}
	return
}

// EditMenuButton  编辑菜单按钮
func (c *cMenuButton) EditMenuButton(ctx context.Context, req *system.EditMenuButtonReq) (res *system.EditMenuButtonRes, err error) {
	var input *model.EditMenuButtonInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuButton().Edit(ctx, input)
	return
}

// DelMenuButton  根据ID删除菜单按钮
func (c *cMenuButton) DelMenuButton(ctx context.Context, req *system.DelMenuButtonReq) (res *system.DelMenuButtonRes, err error) {
	err = service.SysMenuButton().Del(ctx, req.Id)
	return
}

// EditMenuButtonStatus  编辑菜单按钮状态
func (c *cMenuButton) EditMenuButtonStatus(ctx context.Context, req *system.EditMenuButtonStatusReq) (res *system.EditMenuButtonStatusRes, err error) {
	err = service.SysMenuButton().EditStatus(ctx, req.Id, req.MenuId, req.Status)
	return
}
