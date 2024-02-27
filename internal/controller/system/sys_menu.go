package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysMenu = cMenu{}

type cMenu struct{}

func (s *cMenu) MenuTree(ctx context.Context, req *system.MenuDoReq) (res *system.MenuDoRes, err error) {
	//获取所有的菜单
	menuInfo, err := service.SysMenu().GetTree(ctx, req.Title, req.Status)
	if err != nil {
		return
	}
	var treeData []*model.SysMenuRes
	if menuInfo != nil {
		if err = gconv.Scan(menuInfo, &treeData); err != nil {
			return
		}
	}
	res = &system.MenuDoRes{
		Data: treeData,
	}
	return
}

// AddMenu 添加菜单
func (s *cMenu) AddMenu(ctx context.Context, req *system.AddMenuReq) (res *system.AddMenuRes, err error) {
	var input *model.AddMenuInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenu().Add(ctx, input)
	return
}

// DetailMenu 获取菜单详情
func (s *cMenu) DetailMenu(ctx context.Context, req *system.DetailMenuReq) (res *system.DetailMenuRes, err error) {
	data, err := service.SysMenu().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *model.DetailMenuRes
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &system.DetailMenuRes{
			Data: detailRes,
		}
	}
	return
}

// EditMenu 编辑菜单
func (s *cMenu) EditMenu(ctx context.Context, req *system.EditMenuReq) (res *system.EditMenuRes, err error) {
	var input *model.EditMenuInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenu().Edit(ctx, input)
	return
}

// DelMenu 根据ID删除菜单
func (s *cMenu) DelMenu(ctx context.Context, req *system.DelMenuReq) (res *system.DelMenuRes, err error) {
	err = service.SysMenu().Del(ctx, req.Id)
	return
}
