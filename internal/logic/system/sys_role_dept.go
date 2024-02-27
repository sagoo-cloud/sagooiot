package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/dao"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
)

type sSysRoleDept struct {
}

func init() {
	service.RegisterSysRoleDept(sysRoleDeptNew())
}

func sysRoleDeptNew() *sSysRoleDept {
	return &sSysRoleDept{}
}

// GetInfoByRoleId 根据角色ID获取信息
func (s *sSysRoleDept) GetInfoByRoleId(ctx context.Context, roleId int) (data []*entity.SysRoleDept, err error) {
	var roleDepts []*entity.SysRoleDept
	err = dao.SysRoleDept.Ctx(ctx).Where(dao.SysRoleDept.Columns().RoleId, roleId).Scan(&roleDepts)
	if roleDepts != nil {
		for _, roleDept := range roleDepts {
			//判断角色是否为已启动并未删除状态
			num, _ := dao.SysDept.Ctx(ctx).Where(g.Map{
				dao.SysDept.Columns().DeptId:    roleDept.DeptId,
				dao.SysDept.Columns().Status:    1,
				dao.SysDept.Columns().IsDeleted: 0,
			}).Count()

			if num > 0 {
				data = append(data, roleDept)
			}
		}
	}
	return
}
