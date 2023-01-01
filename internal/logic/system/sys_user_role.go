package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sSysUserRole struct {
}

func init() {
	service.RegisterSysUserRole(sysUserRoleNew())
}

func sysUserRoleNew() *sSysUserRole {
	return &sSysUserRole{}
}

//GetInfoByUserId 根据用户ID获取信息
func (s *sSysUserRole) GetInfoByUserId(ctx context.Context, userId int) (data []*entity.SysUserRole, err error) {
	var userRole []*entity.SysUserRole
	err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, userId).Scan(&userRole)
	if userRole != nil {
		for _, role := range userRole {
			//判断角色是否为已启动并未删除状态
			num, _ := dao.SysRole.Ctx(ctx).Where(g.Map{
				dao.SysRole.Columns().Id:        role.RoleId,
				dao.SysRole.Columns().Status:    1,
				dao.SysRole.Columns().IsDeleted: 0,
			}).Count()

			if num > 0 {
				data = append(data, role)
			}
		}
	}
	return
}
