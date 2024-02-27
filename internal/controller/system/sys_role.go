package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

// SysRole 角色
var SysRole = cSysRole{}

type cSysRole struct{}

// RoleTree 角色树状列表
func (a *cSysRole) RoleTree(ctx context.Context, req *system.RoleTreeReq) (res *system.RoleTreeRes, err error) {
	//获取所有的角色
	out, err := service.SysRole().GetTree(ctx, req.Name, req.Status)
	if err != nil {
		return nil, err
	}
	var treeData []*model.RoleTreeRes
	if out != nil {
		if err = gconv.Scan(out, &treeData); err != nil {
			return
		}
	}
	res = &system.RoleTreeRes{
		Data: treeData,
	}
	return
}

// AddRole 添加
func (a *cSysRole) AddRole(ctx context.Context, req *system.AddRoleReq) (res *system.AddRoleRes, err error) {
	var input *model.AddRoleInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysRole().Add(ctx, input)
	return
}

// EditRole 编辑
func (a *cSysRole) EditRole(ctx context.Context, req *system.EditRoleReq) (res *system.EditRoleRes, err error) {
	var input *model.EditRoleInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysRole().Edit(ctx, input)
	return
}

// GetRoleById 根据ID获取角色信息
func (a *cSysRole) GetRoleById(ctx context.Context, req *system.GetRoleByIdReq) (res *system.GetRoleByIdRes, err error) {
	data, err := service.SysRole().GetInfoById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var roleInfoRes *model.RoleInfoRes
		if err = gconv.Scan(data, &roleInfoRes); err != nil {
			return nil, err
		}
		if roleInfoRes.DataScope == 2 {
			//获取部门ID
			roleDeptInfo, _ := service.SysRoleDept().GetInfoByRoleId(ctx, int(roleInfoRes.Id))
			var deptIds []int64
			if roleDeptInfo != nil {
				for _, roleDept := range roleDeptInfo {
					deptIds = append(deptIds, roleDept.DeptId)
				}
			}
			roleInfoRes.DeptIds = append(roleInfoRes.DeptIds, deptIds...)
		}
		res = &system.GetRoleByIdRes{
			Data: roleInfoRes,
		}
	}
	return
}

// DelRoleById 根据ID删除角色
func (a *cSysRole) DelRoleById(ctx context.Context, req *system.DeleteRoleByIdReq) (res *system.DeleteRoleByIdRes, err error) {
	err = service.SysRole().DelInfoById(ctx, req.Id)
	return
}

// DataScope 角色数据权限授权
func (a *cSysRole) DataScope(ctx context.Context, req *system.DataScopeReq) (res *system.DataScopeRes, err error) {
	err = service.SysRole().DataScope(ctx, req.Id, req.DataScope, req.DeptIds)
	return
}

// GetAuthorizeById 根据ID获取权限信息
func (c *cSysAuthorize) GetAuthorizeById(ctx context.Context, req *system.GetAuthorizeByIdReq) (res *system.GetAuthorizeByIdRes, err error) {
	menuIds, menuButtonIds, menuColumnIds, menuApiIds, err := service.SysRole().GetAuthorizeById(ctx, req.Id)
	if err != nil {
		return
	}

	res = &system.GetAuthorizeByIdRes{
		MenuIds:   menuIds,
		ButtonIds: menuButtonIds,
		ColumnIds: menuColumnIds,
		ApiIds:    menuApiIds,
	}
	return

}
