package model

import "github.com/gogf/gf/v2/os/gtime"

type LoginLogParams struct {
	Status    int
	Username  string
	Ip        string
	UserAgent string
	Msg       string
	Module    string
}

type SysLoginLogInput struct {
	LoginName     string      `json:"loginName"     description:"登录账号"`
	Ipaddr        string      `json:"ipaddr"        description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" description:"登录地点"`
	Browser       string      `json:"browser"       description:"浏览器类型"`
	Os            string      `json:"os"            description:"操作系统"`
	Status        int         `json:"status"        description:"登录状态（0失败 1成功）"`
	Msg           string      `json:"msg"           description:"提示消息"`
	LoginTime     *gtime.Time `json:"loginTime"     description:"登录时间"`
	Module        string      `json:"module"        description:"登录模块"`
	PaginationInput
}

type SysLoginLogListOut struct {
	Data []*SysLoginLogOut
	PaginationOutput
}

type SysLoginLogOut struct {
	InfoId        int64       `json:"infoId"        description:"访问ID"`
	LoginName     string      `json:"loginName"     description:"登录账号"`
	Ipaddr        string      `json:"ipaddr"        description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" description:"登录地点"`
	Browser       string      `json:"browser"       description:"浏览器类型"`
	Os            string      `json:"os"            description:"操作系统"`
	Status        int         `json:"status"        description:"登录状态（0失败 1成功）"`
	Msg           string      `json:"msg"           description:"提示消息"`
	LoginTime     *gtime.Time `json:"loginTime"     description:"登录时间"`
	Module        string      `json:"module"        description:"登录模块"`
}

type SysLoginLogExportOut struct {
	InfoId        int64       `json:"infoId"        description:"访问ID"`
	LoginName     string      `json:"loginName"     description:"登录账号"`
	Ipaddr        string      `json:"ipaddr"        description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" description:"登录地点"`
	Browser       string      `json:"browser"       description:"浏览器类型"`
	Os            string      `json:"os"            description:"操作系统"`
	Status        string      `json:"status"        description:"登录状态"`
	Msg           string      `json:"msg"           description:"提示消息"`
	LoginTime     *gtime.Time `json:"loginTime"     description:"登录时间"`
	Module        string      `json:"module"        description:"登录模块"`
}
