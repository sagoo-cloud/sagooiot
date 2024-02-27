package system

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SysLoginLogDoReq struct {
	g.Meta        `path:"/login/log/list"       tags:"访问日志管理" method:"get" summary:"访问日志列表"`
	LoginName     string `p:"login_name"        description:"登录账号"`
	Ipaddr        string `p:"ipaddr"            description:"登录IP地址"`
	LoginLocation string `p:"login_location"    description:"登录地点"`
	Browser       string `p:"browser"           description:"浏览器类型"`
	Os            string `p:"os"                description:"操作系统"`
	Status        int    `p:"status"            description:"登录状态（0失败 1成功）"`
	*common.PaginationReq
}
type SysLoginLogDoRes struct {
	Data []*model.SysLoginLogOut
	common.PaginationRes
}

type SysLoginLogDoExportReq struct {
	g.Meta        `path:"/login/log/export"       tags:"访问日志管理" method:"get" summary:"访问日志导出"`
	LoginName     string `p:"login_name"        description:"登录账号"`
	Ipaddr        string `p:"ipaddr"            description:"登录IP地址"`
	LoginLocation string `p:"login_location"    description:"登录地点"`
	Browser       string `p:"browser"           description:"浏览器类型"`
	Os            string `p:"os"                description:"操作系统"`
	Status        int    `p:"status"            description:"登录状态（0失败 1成功）"`
	common.PaginationReq
}
type SysLoginLogDoExportRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type DetailSysLoginLogReq struct {
	g.Meta `path:"/login/log/detail" tags:"访问日志管理"   method:"get" summary:"根据ID获取访问日志详情"`
	InfoId int `json:"infoId"     description:"访问日志ID"  v:"required#日志ID不能为空"`
}
type DetailSysLoginLogRes struct {
	Data *entity.SysLoginLog
}

type DelSysLoginLogReq struct {
	g.Meta  `path:"/login/log/del" method:"delete" summary:"根据ID删除访问日志" tags:"访问日志管理"`
	InfoIds []int `json:"infoIds" description:"访问日志ID"  v:"required#ID不能为空"`
}
type DelSysLoginLogRes struct {
}
