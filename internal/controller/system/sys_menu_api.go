package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysMenuApi = cMenuApi{}

type cMenuApi struct{}

// MenuApiTree 获取菜单API树结构列表
func (c *cMenuApi) MenuApiTree(ctx context.Context, req *system.MenuApiDoReq) (res *system.MenuApiDoRes, err error) {
	out, err := service.SysMenuApi().MenuApiList(ctx, req.MenuId)
	if err != nil {
		return
	}
	var data []*model.SysApiAllRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &system.MenuApiDoRes{
		Data: data,
	}
	return
}

// AddMenuApi 绑定菜单和API关联关系
func (c *cMenuApi) AddMenuApi(ctx context.Context, req *system.AddMenuApiReq) (res *system.AddMenuApiRes, err error) {
	err = service.SysApi().AddMenuApi(ctx, "menu", req.ApiIds, []int{req.MenuId})
	if err != nil {
		return
	}
	return
}
