package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var SysMenu = cMenu{}

type cMenu struct{}

func (a *cMenu) MenuTree(ctx context.Context, req *system.MenuDoReq) (res *system.MenuDoRes, err error) {
	//获取所有的菜单
	menuInfo, err := service.SysMenu().GetTree(ctx, req.Title, req.Status)
	if err != nil {
		return nil, err
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
func (a *cMenu) AddMenu(ctx context.Context, req *system.AddMenuReq) (res *system.AddMenuRes, err error) {
	var input *model.AddMenuInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenu().Add(ctx, input)
	return
}

// DetailMenu 获取菜单详情
func (a *cMenu) DetailMenu(ctx context.Context, req *system.DetailMenuReq) (res *system.DetailMenuRes, err error) {
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
func (a *cMenu) EditMenu(ctx context.Context, req *system.EditMenuReq) (res *system.EditMenuRes, err error) {
	var input *model.EditMenuInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenu().Edit(ctx, input)
	return
}

// DelMenu 根据ID删除菜单
func (a *cMenu) DelMenu(ctx context.Context, req *system.DelMenuReq) (res *system.DelMenuRes, err error) {
	err = service.SysMenu().Del(ctx, req.Id)
	return
}

// MenuButtonTree 获取菜单按钮树结构列表
func (a *cMenu) MenuButtonTree(ctx context.Context, req *system.MenuButtonDoReq) (res *system.MenuButtonDoRes, err error) {
	//获取所有的菜单
	menuButtonInfo, err := service.SysMenuButton().GetList(ctx, req.Status, req.Name, req.MenuId)
	if err != nil {
		return nil, err
	}
	var parentNodeRes []*model.UserMenuButtonRes
	if menuButtonInfo != nil {
		//获取所有的根节点
		for _, v := range menuButtonInfo {
			var parentNode *model.UserMenuButtonRes
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeRes = append(parentNodeRes, parentNode)
			}
		}
	}
	treeData := buttonTree(parentNodeRes, menuButtonInfo)
	res = &system.MenuButtonDoRes{
		Data: treeData,
	}
	return
}

// buttonTree MenuButtonTree 生成菜单按钮树结构
func buttonTree(parentNodeRes []*model.UserMenuButtonRes, data []model.UserMenuButtonRes) (dataTree []*model.UserMenuButtonRes) {
	//循环所有一级菜单
	for k, v := range parentNodeRes {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.UserMenuButtonRes
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeRes[k].Children = append(parentNodeRes[k].Children, node)
			}
		}
		buttonTree(v.Children, data)
	}
	return parentNodeRes
}

// AddMenuButton AddMenu 添加菜单按钮
func (a *cMenu) AddMenuButton(ctx context.Context, req *system.AddMenuButtonReq) (res *system.AddMenuButtonRes, err error) {
	var input *model.AddMenuButtonInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuButton().Add(ctx, input)
	return
}

// DetailMenuButton DetailMenu 获取菜单按钮详情
func (a *cMenu) DetailMenuButton(ctx context.Context, req *system.DetailMenuButtonReq) (res *system.DetailMenuButtonRes, err error) {
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
func (a *cMenu) EditMenuButton(ctx context.Context, req *system.EditMenuButtonReq) (res *system.EditMenuButtonRes, err error) {
	var input *model.EditMenuButtonInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuButton().Edit(ctx, input)
	return
}

// DelMenuButton  根据ID删除菜单按钮
func (a *cMenu) DelMenuButton(ctx context.Context, req *system.DelMenuButtonReq) (res *system.DelMenuButtonRes, err error) {
	err = service.SysMenuButton().Del(ctx, req.Id)
	return
}

// EditMenuButtonStatus  编辑菜单按钮状态
func (a *cMenu) EditMenuButtonStatus(ctx context.Context, req *system.EditMenuButtonStatusReq) (res *system.EditMenuButtonStatusRes, err error) {
	err = service.SysMenuButton().EditStatus(ctx, req.Id, req.MenuId, req.Status)
	return
}

// MenuColumnTree 获取菜单列表树结构列表
func (a *cMenu) MenuColumnTree(ctx context.Context, req *system.MenuColumnDoReq) (res *system.MenuColumnDoRes, err error) {
	var input *model.MenuColumnDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	//获取所有的菜单
	menuColumnInfo, err := service.SysMenuColumn().GetList(ctx, input)
	if err != nil {
		return nil, err
	}
	var parentNodeRes []*model.UserMenuColumnRes
	if menuColumnInfo != nil {
		//获取所有的根节点
		for _, v := range menuColumnInfo {
			var parentNode *model.UserMenuColumnRes
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeRes = append(parentNodeRes, parentNode)
			}
		}
	}
	treeData := ColumnTree(parentNodeRes, menuColumnInfo)
	res = &system.MenuColumnDoRes{
		Data: treeData,
	}
	return
}

// ColumnTree MenuColumnTree 生成菜单列表树结构
func ColumnTree(parentNodeRes []*model.UserMenuColumnRes, data []model.UserMenuColumnRes) (dataTree []*model.UserMenuColumnRes) {
	//循环所有一级菜单
	for k, v := range parentNodeRes {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.UserMenuColumnRes
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeRes[k].Children = append(parentNodeRes[k].Children, node)
			}
		}
		ColumnTree(v.Children, data)
	}
	return parentNodeRes
}

// AddMenuColumn AddMenu 添加菜单列表
func (a *cMenu) AddMenuColumn(ctx context.Context, req *system.AddMenuColumnReq) (res *system.AddMenuColumnRes, err error) {
	var input *model.AddMenuColumnInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuColumn().Add(ctx, input)
	return
}

// DetailMenuColumn DetailMenu 获取菜单列表详情
func (a *cMenu) DetailMenuColumn(ctx context.Context, req *system.DetailMenuColumnReq) (res *system.DetailMenuColumnRes, err error) {
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
func (a *cMenu) EditMenuColumn(ctx context.Context, req *system.EditMenuColumnReq) (res *system.EditMenuColumnRes, err error) {
	var input *model.EditMenuColumnInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysMenuColumn().Edit(ctx, input)
	return
}

// DelMenuColumn DelMenu 根据ID删除菜单列表
func (a *cMenu) DelMenuColumn(ctx context.Context, req *system.DelMenuColumnReq) (res *system.DelMenuColumnRes, err error) {
	err = service.SysMenuColumn().Del(ctx, req.Id)
	return
}

// EditMenuColumnStatus  编辑菜单列表状态
func (a *cMenu) EditMenuColumnStatus(ctx context.Context, req *system.EditMenuColumnStatusReq) (res *system.EditMenuColumnStatusRes, err error) {
	err = service.SysMenuColumn().EditStatus(ctx, req.Id, req.MenuId, req.Status)
	return
}
