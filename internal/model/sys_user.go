package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type UserListDoInput struct {
	KeyWords string `json:"keyWords" description:"关键词(可根据账号或者用户昵称查询)"`
	DeptId   int    `json:"deptId"        description:"部门ID"`
	UserName string `json:"userName"  description:"用户名"`
	Mobile   string `json:"mobile"  description:"手机号"`
	Status   int    `json:"status"  description:"用户状态;0:禁用,1:正常,2:未验证"`
	*PaginationInput
}

// LoginUserRes 登录返回
type LoginUserRes struct {
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	Avatar       string `orm:"avatar" json:"avatar"`                 //头像
}

type LoginUserOut struct {
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	Avatar       string `orm:"avatar" json:"avatar"`                 //头像
}

type UserListOut struct {
	Id            uint64         `json:"id"            description:""`
	UserName      string         `json:"userName"      description:"用户名"`
	UserTypes     string         `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile        string         `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string         `json:"userNickname"  description:"用户昵称"`
	Birthday      int            `json:"birthday"      description:"生日"`
	UserEmail     string         `json:"userEmail"     description:"用户登录邮箱"`
	Sex           int            `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar        string         `json:"avatar"        description:"用户头像"`
	DeptId        int64          `json:"deptId"        description:"部门id"`
	Remark        string         `json:"remark"        description:"备注"`
	IsAdmin       int            `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Address       string         `json:"address"       description:"联系地址"`
	Describe      string         `json:"describe"      description:"描述信息"`
	LastLoginIp   string         `json:"lastLoginIp"   description:"最后登录ip"`
	LastLoginTime *gtime.Time    `json:"lastLoginTime" description:"最后登录时间"`
	Status        uint           `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
	CreateBy      uint           `json:"createBy"      description:"创建者"`
	CreatedAt     *gtime.Time    `json:"createdAt"     description:"创建日期"`
	UpdateBy      uint           `json:"updateBy"      description:"更新者"`
	UpdatedAt     *gtime.Time    `json:"updatedAt"     description:"修改日期"`
	Dept          *DetailDeptRes `json:"dept"     description:"部门信息"`
	RolesNames    string         `json:"rolesNames"     description:"角色信息"`
}

type UserListRes struct {
	Id            uint64         `json:"id"            description:""`
	UserName      string         `json:"userName"      description:"用户名"`
	UserTypes     string         `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile        string         `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string         `json:"userNickname"  description:"用户昵称"`
	Birthday      int            `json:"birthday"      description:"生日"`
	UserEmail     string         `json:"userEmail"     description:"用户登录邮箱"`
	Sex           int            `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar        string         `json:"avatar"        description:"用户头像"`
	DeptId        int64          `json:"deptId"        description:"部门id"`
	Remark        string         `json:"remark"        description:"备注"`
	IsAdmin       int            `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Address       string         `json:"address"       description:"联系地址"`
	Describe      string         `json:"describe"      description:"描述信息"`
	LastLoginIp   string         `json:"lastLoginIp"   description:"最后登录ip"`
	LastLoginTime *gtime.Time    `json:"lastLoginTime" description:"最后登录时间"`
	Status        uint           `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
	CreateBy      uint           `json:"createBy"      description:"创建者"`
	CreatedAt     *gtime.Time    `json:"createdAt"     description:"创建日期"`
	UpdateBy      uint           `json:"updateBy"      description:"更新者"`
	UpdatedAt     *gtime.Time    `json:"updatedAt"     description:"修改日期"`
	Dept          *DetailDeptRes `json:"dept"     description:"部门信息"`
	RolesNames    string         `json:"rolesNames"     description:"角色信息"`
}

type UserRes struct {
	Id            uint64         `json:"id"            description:""`
	UserName      string         `json:"userName"      description:"用户名"`
	UserTypes     string         `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile        string         `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string         `json:"userNickname"  description:"用户昵称"`
	Birthday      int            `json:"birthday"      description:"生日"`
	UserEmail     string         `json:"userEmail"     description:"用户登录邮箱"`
	Sex           int            `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar        string         `json:"avatar"        description:"用户头像"`
	DeptId        int64          `json:"deptId"        description:"部门id"`
	Remark        string         `json:"remark"        description:"备注"`
	IsAdmin       int            `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Address       string         `json:"address"       description:"联系地址"`
	Describe      string         `json:"describe"      description:"描述信息"`
	LastLoginIp   string         `json:"lastLoginIp"   description:"最后登录ip"`
	LastLoginTime *gtime.Time    `json:"lastLoginTime" description:"最后登录时间"`
	Status        uint           `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
	CreateBy      uint           `json:"createBy"      description:"创建者"`
	CreatedAt     *gtime.Time    `json:"createdAt"     description:"创建日期"`
	UpdateBy      uint           `json:"updateBy"      description:"更新者"`
	UpdatedAt     *gtime.Time    `json:"updatedAt"     description:"修改日期"`
	Dept          *DetailDeptRes `json:"dept"     description:"部门信息"`
	RolesNames    string         `json:"rolesNames"     description:"角色信息"`
}

type AddUserInput struct {
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

type EditUserInput struct {
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

type UserInfoRes struct {
	Id            uint64      `json:"id"            description:""`
	UserName      string      `json:"userName"      description:"用户名"`
	UserTypes     string      `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile        string      `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string      `json:"userNickname"  description:"用户昵称"`
	Birthday      int         `json:"birthday"      description:"生日"`
	UserEmail     string      `json:"userEmail"     description:"用户登录邮箱"`
	Sex           int         `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar        string      `json:"avatar"        description:"用户头像"`
	DeptId        uint64      `json:"deptId"        description:"部门id"`
	Remark        string      `json:"remark"        description:"备注"`
	IsAdmin       int         `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Address       string      `json:"address"       description:"联系地址"`
	Describe      string      `json:"describe"      description:"描述信息"`
	LastLoginIp   string      `json:"lastLoginIp"   description:"最后登录ip"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" description:"最后登录时间"`
	Status        uint        `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
	CreateBy      uint        `json:"createBy"      description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建日期"`
	UpdateBy      uint        `json:"updateBy"      description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"修改日期"`
	RoleIds       []int       `json:"roleIds"      description:"角色ID数组" v:"required#角色不能为空"`
	PostIds       []int       `json:"postIds"      description:"岗位ID数组" v:"required#岗位不能为空"`
}
type UserInfoOut struct {
	Id            uint64      `json:"id"            description:""`
	UserName      string      `json:"userName"      description:"用户名"`
	UserTypes     string      `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile        string      `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string      `json:"userNickname"  description:"用户昵称"`
	Birthday      int         `json:"birthday"      description:"生日"`
	UserEmail     string      `json:"userEmail"     description:"用户登录邮箱"`
	Sex           int         `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar        string      `json:"avatar"        description:"用户头像"`
	DeptId        uint64      `json:"deptId"        description:"部门id"`
	Remark        string      `json:"remark"        description:"备注"`
	IsAdmin       int         `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Address       string      `json:"address"       description:"联系地址"`
	Describe      string      `json:"describe"      description:"描述信息"`
	LastLoginIp   string      `json:"lastLoginIp"   description:"最后登录ip"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" description:"最后登录时间"`
	Status        uint        `json:"status"        description:"用户状态;0:禁用,1:正常,2:未验证"`
	CreateBy      uint        `json:"createBy"      description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建日期"`
	UpdateBy      uint        `json:"updateBy"      description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"修改日期"`
	RoleIds       []int       `json:"roleIds"      description:"角色ID数组" v:"required#角色不能为空"`
	PostIds       []int       `json:"postIds"      description:"岗位ID数组" v:"required#岗位不能为空"`
}
