package system

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sSysMenuApi struct {
}

func sysMenuApiNew() *sSysMenuApi {
	return &sSysMenuApi{}
}

func init() {
	service.RegisterSysMenuApi(sysMenuApiNew())
}

// GetInfoByIds 根据IDS数组获取菜单信息
func (s *sSysMenuApi) GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysMenuApi, err error) {
	cache := common.Cache()
	//获取缓存信息
	var tmpData *gvar.Var
	tmpData = cache.Get(ctx, consts.CacheSysMenuApi)
	if err != nil {
		return
	}
	var tmpSysMenuApi []*entity.SysMenuApi
	if tmpData.Val() != nil {
		json.Unmarshal([]byte(tmpData.Val().(string)), &tmpSysMenuApi)
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
	cache := common.Cache()
	//获取缓存信息
	for _, v := range menuIds {
		var tmpData *gvar.Var
		tmpData = cache.Get(ctx, consts.CacheSysMenuApi+"_"+gconv.String(v))
		if err != nil {
			return
		}
		if tmpData.Val() != nil {
			var sysMenuApi []*entity.SysMenuApi
			json.Unmarshal([]byte(tmpData.Val().(string)), &sysMenuApi)
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
