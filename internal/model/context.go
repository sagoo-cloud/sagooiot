package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id             int    // 用户ID
	UserName       string // 用户账号
	Nickname       string // 用户名称
	Avatar         string // 用户
	IsAdmin        bool   // 是否是管理员
	DeptId         int    // 部门ID
	RequestWay     string // 请求方式
	ChildrenDeptId []int  //子部门ID
}
