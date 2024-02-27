package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/dao"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
)

type sSysUserRole struct {
}

func init() {
	service.RegisterSysUserRole(sysUserRoleNew())
}

func sysUserRoleNew() *sSysUserRole {
	return &sSysUserRole{}
}

// GetInfoByUserId 根据用户ID获取信息
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

// BindUserAndRole 添加用户与角色绑定关系
func (s *sSysUserRole) BindUserAndRole(ctx context.Context, userId int, roleIds []int) (err error) {
	if len(roleIds) > 0 {
		//删除原有用户与角色绑定管理
		_, err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, userId).Delete()
		if err != nil {
			return gerror.New("删除用户与角色绑定关系失败")
		}

		var sysUserRoles []*entity.SysUserRole
		for _, roleId := range roleIds {
			/*var sysUserRole *entity.SysUserRole
			err = dao.SysUserRole.Ctx(ctx).Where(g.Map{
				dao.SysUserRole.Columns().UserId: userId,
				dao.SysUserRole.Columns().RoleId: roleId,
			}).Scan(&sysUserRole)*/

			//添加用户与角色绑定管理
			var sysUserRole = new(entity.SysUserRole)
			sysUserRole.UserId = userId
			sysUserRole.RoleId = roleId
			sysUserRoles = append(sysUserRoles, sysUserRole)
		}
		_, err = dao.SysUserRole.Ctx(ctx).Data(sysUserRoles).Insert()
		if err != nil {
			return gerror.New("绑定角色失败")
		}
		//删除缓存
		_, err = gcache.Remove(ctx, "RoleListAtName"+gconv.String(userId))
		if err != nil {
			return err
		}
	}
	return
}
