package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
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

//GetInfoByIds 根据IDS数组获取菜单信息
func (s *sSysMenuApi) GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysMenuApi, err error) {
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysMenuApi.Columns().Id, ids).Scan(&data)
	return
}

//GetInfoByMenuIds 根据菜单ID数组获取菜单信息
func (s *sSysMenuApi) GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuApi, err error) {
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysMenuApi.Columns().MenuId, menuIds).Scan(&data)
	return
}

//GetInfoByApiId 根据接口ID数组获取菜单信息
func (s *sSysMenuApi) GetInfoByApiId(ctx context.Context, apiId int) (data []*entity.SysMenuApi, err error) {
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().IsDeleted: 0,
		dao.SysMenuApi.Columns().ApiId:     apiId,
	}).Scan(&data)
	return
}
