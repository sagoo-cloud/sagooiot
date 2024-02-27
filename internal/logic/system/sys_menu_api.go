package system

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
)

type sSysMenuApi struct {
}

func sysMenuApiNew() *sSysMenuApi {
	return &sSysMenuApi{}
}

func init() {
	service.RegisterSysMenuApi(sysMenuApiNew())
}

// MenuApiList 根据菜单ID获取API列表
func (s *sSysMenuApi) MenuApiList(ctx context.Context, menuId int) (out []*model.SysApiAllOut, err error) {
	//获取所有的菜单
	menuApiInfo, err := s.GetInfoByMenuId(ctx, menuId)
	if err != nil {
		return
	}
	if menuApiInfo != nil && len(menuApiInfo) > 0 {
		var apiIds []int
		for _, menuApi := range menuApiInfo {
			apiIds = append(apiIds, menuApi.ApiId)
		}
		if apiIds != nil && len(apiIds) > 0 {
			var apiInfos []*entity.SysApi
			apiInfos, err = service.SysApi().GetInfoByIds(ctx, apiIds)
			if err != nil {
				return
			}
			if apiInfos != nil {
				if err = gconv.Scan(apiInfos, &out); err != nil {
					return
				}
			}
		}
	}
	return
}

// GetInfoByIds 根据IDS数组获取菜单信息
func (s *sSysMenuApi) GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysMenuApi, err error) {
	//获取缓存信息
	var tmpData *gvar.Var
	tmpData, err = cache.Instance().Get(ctx, consts.CacheSysMenuApi)
	if err != nil {
		return
	}
	var tmpSysMenuApi []*entity.SysMenuApi
	if tmpData.Val() != nil {
		if err = json.Unmarshal([]byte(tmpData.Val().(string)), &tmpSysMenuApi); err != nil {
			return
		}
		for _, v := range ids {
			for _, tmp := range tmpSysMenuApi {
				if v == int(tmp.Id) {
					data = append(data, tmp)
				}
			}
		}
	}
	if data == nil || len(data) > 0 {
		err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 0,
		}).WhereIn(dao.SysMenuApi.Columns().Id, ids).Scan(&data)
	}
	return
}

// GetInfoByMenuIds 根据菜单ID数组获取菜单信息
func (s *sSysMenuApi) GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuApi, err error) {
	//获取缓存信息
	for _, v := range menuIds {
		var tmpData *gvar.Var
		tmpData, err = cache.Instance().Get(ctx, consts.CacheSysMenuApi+"_"+gconv.String(v))
		if err != nil {
			return
		}
		if tmpData.Val() != nil {
			var sysMenuApi []*entity.SysMenuApi
			err = json.Unmarshal([]byte(tmpData.Val().(string)), &sysMenuApi)
			data = append(data, sysMenuApi...)
		}
	}
	if data == nil || len(data) > 0 {
		err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 0,
		}).WhereIn(dao.SysMenuApi.Columns().MenuId, menuIds).Scan(&data)
	}
	return
}

// GetInfoByApiId 根据接口ID数组获取菜单信息
func (s *sSysMenuApi) GetInfoByApiId(ctx context.Context, apiId int) (data []*entity.SysMenuApi, err error) {
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().IsDeleted: 0,
		dao.SysMenuApi.Columns().ApiId:     apiId,
	}).Scan(&data)
	return
}

// GetAll 获取所有信息
func (s *sSysMenuApi) GetAll(ctx context.Context) (data []*entity.SysMenuApi, err error) {
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

// GetInfoByMenuId 根据菜单ID获取菜单信息
func (s *sSysMenuApi) GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuApi, err error) {
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().IsDeleted: 0,
		dao.SysMenuApi.Columns().MenuId:    menuId,
	}).Scan(&data)
	return
}

// FindParentByChildrenId 根据子节点获取根节点
func FindParentByChildrenId(ctx context.Context, parentId int) *entity.SysApi {
	var api *entity.SysApi

	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: parentId,
	}).Scan(&api)

	if api.ParentId != -1 {
		return FindParentByChildrenId(ctx, api.ParentId)
	}
	return api
}
