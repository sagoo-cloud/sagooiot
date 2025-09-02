package model

import "github.com/gogf/gf/v2/os/gtime"

type UserInfoOut struct {
	Id            uint64      `json:"id"            description:""`
	UserName      string      `json:"userName"      description:"用户名"`
	UserTypes     string      `json:"userTypes"     description:"系统 system 企业 company"`
	Mobile        string      `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string      `json:"userNickname"  description:"用户昵称"`
	Birthday      *gtime.Time `json:"birthday"      description:"生日"`
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
	CreatedBy     uint        `json:"createdBy"      description:"创建者"`
	IsDeleted     int         `json:"isDeleted"     orm:"is_deleted"      description:"是否删除 0未删除 1已删除"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建日期"`
	UpdatedBy     uint        `json:"updatedBy"      description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"修改日期"`
	RoleIds       []int       `json:"roleIds"      description:"角色ID数组" v:"required#角色不能为空"`
	PostIds       []int       `json:"postIds"      description:"岗位ID数组" v:"required#岗位不能为空"`
}

type LoginUserOut struct {
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	Avatar       string `orm:"avatar" json:"avatar"`                 //头像
}
