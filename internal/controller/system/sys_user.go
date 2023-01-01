package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

// SysUser 用户
var SysUser = cSysUser{}

type cSysUser struct{}

// UserList 用户列表
func (u *cSysUser) UserList(ctx context.Context, req *system.UserListReq) (res *system.UserListRes, err error) {
	//获取所有用户列表
	var input *model.UserListDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, out, err := service.SysUser().UserList(ctx, input)
	if err != nil {
		return
	}
	res = new(system.UserListRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.Data); err != nil {
			return
		}
	}
	return
}

// AddUser 用户添加
func (u *cSysUser) AddUser(ctx context.Context, req *system.AddUserReq) (res *system.AddUserRes, err error) {
	var input *model.AddUserInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysUser().Add(ctx, input)
	return
}

// EditUser 用户编辑
func (u *cSysUser) EditUser(ctx context.Context, req *system.EditUserReq) (res *system.EditUserRes, err error) {
	var input *model.EditUserInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysUser().Edit(ctx, input)
	return
}

// GetUserById 根据ID获取用户信息
func (u *cSysUser) GetUserById(ctx context.Context, req *system.GetUserByIdReq) (res *system.GetUserByIdRes, err error) {
	out, err := service.SysUser().GetUserById(ctx, req.Id)
	if err != nil {
		return
	}
	var userInfoRes *model.UserInfoRes
	if out != nil {
		if err = gconv.Scan(out, &userInfoRes); err != nil {
			return
		}
	}
	res = &system.GetUserByIdRes{
		Data: userInfoRes,
	}
	return
}

// DelUserById 根据ID删除用户
func (u *cSysUser) DelUserById(ctx context.Context, req *system.DeleteUserByIdReq) (res *system.DeleteUserByIdRes, err error) {
	err = service.SysUser().DelInfoById(ctx, req.Id)
	return
}

// ResetPassword 重置密码
func (u *cSysUser) ResetPassword(ctx context.Context, req *system.ResetPasswordReq) (res *system.ResetPasswordRes, err error) {
	err = service.SysUser().ResetPassword(ctx, req.Id, req.UserPassword)
	return
}

// CurrentUser 获取登录用户信息
func (u *cSysUser) CurrentUser(ctx context.Context, req *system.CurrentUserReq) (res *system.CurrentUserRes, err error) {
	userInfoOut, menuTreeOur, err := service.SysUser().CurrentUser(ctx)
	if err != nil {
		return
	}
	var userInfoRes *model.UserInfoRes
	if userInfoOut != nil {
		if err = gconv.Scan(userInfoOut, &userInfoRes); err != nil {
			return
		}
	}
	var userMenuTreeRes []*model.UserMenuTreeRes
	if menuTreeOur != nil {
		if err = gconv.Scan(menuTreeOur, &userMenuTreeRes); err != nil {
			return
		}
	}
	res = &system.CurrentUserRes{
		Info: userInfoRes,
		Data: userMenuTreeRes,
	}
	return
}

// GetParams 获取用户维护相关参数
func (u *cSysUser) GetParams(ctx context.Context, req *system.UserGetParamsReq) (res *system.UserGetParamsRes, err error) {
	res = new(system.UserGetParamsRes)
	res.RoleList, err = service.SysRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Posts, err = service.SysPost().GetUsedPost(ctx)
	return
}

// EditUserStatus 修改用户状态
func (u *cSysUser) EditUserStatus(ctx context.Context, req *system.EditUserStatusReq) (res *system.EditUserStatusRes, err error) {
	err = service.SysUser().EditUserStatus(ctx, req.Id, req.Status)
	return
}

// GetUserAll 所有用户列表
func (u *cSysUser) GetUserAll(ctx context.Context, req *system.GetUserAllReq) (res *system.GetUserAllRes, err error) {
	//获取所有用户列表
	data, err := service.SysUser().GetAll(ctx)
	var userRes []*model.UserRes
	if data != nil {
		if err = gconv.Scan(data, &userRes); err != nil {
			return
		}
	}
	res = &system.GetUserAllRes{
		Data: userRes,
	}
	return
}

// EditUserAvatar 修改用户头像
func (u *cSysUser) EditUserAvatar(ctx context.Context, req *system.EditUserAvatarReq) (res *system.EditUserAvatarRes, err error) {
	err = service.SysUser().EditUserAvatar(ctx, req.Id, req.Avatar)
	return
}
