package system

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SysOperLogDoReq struct {
	g.Meta        `path:"/oper/log/list"       tags:"操作日志管理" method:"get" summary:"操作日志列表"`
	Title         string `p:"title"            description:"模块标题"`
	BusinessType  string `p:"business_type"    description:"业务类型（0其它 1新增 2修改 3删除）"`
	Method        string `p:"method"           description:"方法名称"`
	RequestMethod string `p:"request_method"   description:"请求方式"`
	OperatorType  string `p:"operator_type"    description:"操作类别（0其它 1后台用户 2手机端用户）"`
	OperName      string `p:"oper_name"        description:"操作人员"`
	DeptName      string `p:"dept_name"        description:"部门名称"`
	OperUrl       string `p:"oper_url"         description:"请求URL"`
	OperIp        string `p:"oper_ip"          description:"主机地址"`
	OperLocation  string `p:"oper_location"    description:"操作地点"`
	Status        int    `p:"status"           description:"状态:-1为全部,0为正常,1为停用"`
	*common.PaginationReq
}
type SysOperLogDoRes struct {
	Data []*model.SysOperLogRes
	common.PaginationRes
}

/**
type AddSysOperLogReq struct {
	g.Meta `path:"/oper/log/add"     tags:"操作日志管理" method:"post" summary:"添加操作日志"`
	*entity.SysOperLog
}
type AddSysOperLogRes struct {
}
*/

type DetailSysOperLogReq struct {
	g.Meta `path:"/oper/log/detail" tags:"操作日志管理"   method:"get" summary:"根据ID获取操作日志详情"`
	OperId int `json:"operId"     description:"操作日志ID"  v:"required#日志ID不能为空"`
}
type DetailSysOperLogRes struct {
	Data *entity.SysOperLog
}

type DelSysOperLogReq struct {
	g.Meta  `path:"/oper/log/del" method:"delete" summary:"根据ID删除操作日志" tags:"操作日志管理"`
	OperIds []int `json:"operIds" description:"操作日志ID"  v:"required#ID不能为空"`
}
type DelSysOperLogRes struct {
}
