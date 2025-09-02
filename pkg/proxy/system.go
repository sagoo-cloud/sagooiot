package proxy

import (
	"context"
	"database/sql"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	proxyModel "sagooiot/pkg/proxy/model"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetUserInfo 获取当前登录用户信息
func GetUserInfo(ctx context.Context) (userInfo *model.ContextUser, err error) {
	userInfo = service.Context().GetLoginUser(ctx)
	return
}

func GetUserAll(ctx context.Context) (data []*proxyModel.UserInfoOut, err error) {
	//获取所有用户列表
	sysUserInfo, err := service.SysUser().GetAll(ctx)
	if err != nil {
		return
	}
	if len(sysUserInfo) > 0 {
		err = gconv.Scan(sysUserInfo, &data)
	}
	return
}

// GetDeptInfoById 根据部门id获取部门信息
func GetDeptInfoById(ctx context.Context, deptId int64) (out *proxyModel.SysDeptOut, err error) {
	//获取部门名称
	deptOut, _ := service.SysDept().Detail(ctx, deptId)
	if deptOut != nil {
		if err = gconv.Scan(deptOut, &out); err != nil {
			return
		}
	}
	return
}

func GetDeptAll(ctx context.Context) (out []*proxyModel.SysDeptOut, err error) {
	//获取部门名称
	deptOut, _ := service.SysDept().GetAll(ctx)
	if deptOut != nil {
		if err = gconv.Scan(deptOut, &out); err != nil {
			return
		}
	}
	return
}

// GetDeptInfosByParentId 根据父ID获取子部门信息
func GetDeptInfosByParentId(ctx context.Context, parentId int) (data []*proxyModel.SysDeptOut, err error) {
	all, err := service.SysDept().GetFromCache(ctx)
	if err != nil {
		return
	}
	if len(all) == 0 {
		return
	}
	out := service.SysDept().FindSonByParentId(all, int64(parentId))
	if len(out) > 0 {
		err = gconv.Scan(out, &data)
	}
	return
}

// GetUserInfoById 根据用户id获取用户信息
func GetUserInfoById(ctx context.Context, userId uint) (out *proxyModel.UserInfoOut, err error) {
	userOut, _ := service.SysUser().GetUserById(ctx, userId)
	if userOut != nil {
		if err = gconv.Scan(userOut, &out); err != nil {
			return
		}
	}
	return
}

// GetFromCache 获取部门信息并更新缓存
func GetFromCache(ctx context.Context) (list []*proxyModel.SysDeptOut, err error) {
	all, err := service.SysDept().GetFromCache(ctx)
	if err != nil {
		return
	}
	if len(all) > 0 {
		err = gconv.Scan(all, &list)
	}
	return
}

// GetSequences 获取主键ID
func GetSequences(ctx context.Context, result sql.Result, tableName string, primaryKey string) (lastInsertId int64, err error) {
	lastInsertId, err = service.Sequences().GetSequences(ctx, result, tableName, primaryKey)
	return
}
