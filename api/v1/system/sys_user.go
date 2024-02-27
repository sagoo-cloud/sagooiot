package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

type UserListReq struct {
	g.Meta   `path:"/user/list" method:"get" summary:"用户列表" tags:"用户管理"`
	KeyWords string `json:"keyWords" description:"关键词(可根据账号或者用户昵称查询)"`
	DeptId   int    `json:"deptId"        description:"部门ID"`
	UserName string `json:"userName"  description:"用户名"`
	Mobile   string `json:"mobile"  description:"手机号"`
	Status   int    `json:"status"  description:"用户状态;0:禁用,1:正常,2:未验证"`
	*common.PaginationReq
}
type UserListRes struct {
	Data []*model.UserListRes
	common.PaginationRes
}

type AddUserReq struct {
	g.Meta       `path:"/user/add" method:"post" summary:"添加用户" tags:"用户管理"`
	UserName     string `json:"userName"      description:"用户名" v:"required#用户名不能为空"`
	UserTypes    string `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile       string `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号" v:"required#手机号不能为空"`
	UserNickname string `json:"userNickname"  description:"用户昵称" v:"required#用户昵称不能为空"`
	Birthday     int    `json:"birthday"      description:"生日"`
	UserPassword string `json:"userPassword"  description:"登录密码;cmf_password加密"`
	UserEmail    string `json:"userEmail"     description:"用户登录邮箱"`
	Sex          int    `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar       string `json:"avatar"        description:"用户头像"`
	DeptId       uint64 `json:"deptId"        description:"部门id" v:"required#部门不能为空"`
	Remark       string `json:"remark"        description:"备注"`
	IsAdmin      int    `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Address      string `json:"address"       description:"联系地址"`
	Describe     string `json:"describe"      description:"描述信息"`
	Status       uint   `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
	RoleIds      []int  `json:"roleIds"      description:"角色ID数组" v:"required#角色不能为空"`
	PostIds      []int  `json:"postIds"      description:"岗位ID数组" v:"required#岗位不能为空"`
}
type AddUserRes struct {
}

type EditUserReq struct {
	g.Meta       `path:"/user/edit" method:"put" summary:"编辑用户" tags:"用户管理"`
	Id           uint64 `json:"id"            description:""`
	UserName     string `json:"userName"      description:"用户名" v:"required#用户名不能为空"`
	UserTypes    string `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile       string `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号" v:"required#手机号不能为空"`
	UserNickname string `json:"userNickname"  description:"用户昵称" v:"required#用户昵称不能为空"`
	Birthday     int    `json:"birthday"      description:"生日"`
	UserEmail    string `json:"userEmail"     description:"用户登录邮箱"`
	Sex          int    `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar       string `json:"avatar"        description:"用户头像"`
	DeptId       uint64 `json:"deptId"        description:"部门id" v:"required#部门不能为空"`
	Remark       string `json:"remark"        description:"备注"`
	IsAdmin      int    `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Address      string `json:"address"       description:"联系地址"`
	Describe     string `json:"describe"      description:"描述信息"`
	Status       uint   `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
	RoleIds      []int  `json:"roleIds"      description:"角色ID数组" v:"required#角色不能为空"`
	PostIds      []int  `json:"postIds"      description:"岗位ID数组" v:"required#岗位不能为空"`
}
type EditUserRes struct {
}

type GetUserByIdReq struct {
	g.Meta `path:"/user/getInfoById" method:"get" summary:"根据ID获取用户" tags:"用户管理"`
	Id     uint `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type GetUserByIdRes struct {
	Data *model.UserInfoRes
}

type DeleteUserByIdReq struct {
	g.Meta `path:"/user/delInfoById" method:"delete" summary:"根据ID删除用户" tags:"用户管理"`
	Id     uint `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type DeleteUserByIdRes struct {
}

type ResetPasswordReq struct {
	g.Meta       `path:"/user/resetPassword" method:"post" summary:"重置用户密码" tags:"用户管理"`
	Id           uint   `json:"id"        description:"ID" v:"required#ID不能为空"`
	UserPassword string `json:"userPassword"  description:"登录密码;cmf_password加密"`
}
type ResetPasswordRes struct {
}

type CurrentUserReq struct {
	g.Meta `path:"/user/currentUser" method:"get" summary:"获取登录用户信息" tags:"用户管理"`
}
type CurrentUserRes struct {
	Info *model.UserInfoRes
	Data []*model.UserMenuTreeRes
}

type UserGetParamsReq struct {
	g.Meta `path:"/user/params" tags:"用户管理" method:"get" summary:"用户维护参数获取"`
}

type UserGetParamsRes struct {
	g.Meta   `mime:"application/json"`
	RoleList []*model.RoleInfoRes   `json:"roleList"`
	Posts    []*model.DetailPostRes `json:"posts"`
}

type EditUserStatusReq struct {
	g.Meta `path:"/user/editStatus" tags:"用户管理" method:"put" summary:"修改用户状态"`
	Id     uint `json:"id"        description:"ID" v:"required#ID不能为空"`
	Status uint `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
}

type EditUserStatusRes struct {
}

type GetUserAllReq struct {
	g.Meta `path:"/user/getAll" method:"get" summary:"所有用户列表" tags:"用户管理"`
}
type GetUserAllRes struct {
	Data []*model.UserRes
}

type EditUserAvatarReq struct {
	g.Meta `path:"/user/editAvatar" tags:"用户管理" method:"put" summary:"修改用户头像"`
	Id     uint   `json:"id"        description:"ID" v:"required#ID不能为空"`
	Avatar string `json:"avatar"        description:"头像" v:"required#头像不能为空"`
}

type EditUserAvatarRes struct {
}

type EditUserInfoReq struct {
	g.Meta       `path:"/user/editUserInfo" tags:"用户管理" method:"put" summary:"修改用户个人资料"`
	Id           uint   `json:"id"        description:"Id" v:"required#ID不能为空"`
	Mobile       string `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname string `json:"userNickname"  description:"用户昵称"`
	Birthday     string `json:"birthday"      description:"生日"`
	UserPassword string `json:"userPassword"  description:"登录密码;cmf_password加密"`
	UserEmail    string `json:"userEmail"     description:"用户登录邮箱"`
	Sex          int    `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar       string `json:"avatar"        description:"用户头像"`
	Address      string `json:"address"       description:"联系地址"`
	Describe     string `json:"describe"      description:"描述信息"`
}

type EditUserInfoRes struct {
}
