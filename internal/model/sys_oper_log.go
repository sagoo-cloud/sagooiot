package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type SysOperLogDoInput struct {
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
	*PaginationInput
}

type SysOperLogOut struct {
	OperId        uint64      `json:"operId"        description:"日志主键"`
	Title         string      `json:"title"         description:"模块标题"`
	BusinessType  int         `json:"businessType"  description:"业务类型（0其它 1新增 2修改 3删除）"`
	Method        string      `json:"method"        description:"方法名称"`
	RequestMethod string      `json:"requestMethod" description:"请求方式"`
	OperatorType  int         `json:"operatorType"  description:"操作类别（0其它 1后台用户 2手机端用户）"`
	OperName      string      `json:"operName"      description:"操作人员"`
	DeptName      string      `json:"deptName"      description:"部门名称"`
	OperUrl       string      `json:"operUrl"       description:"请求URL"`
	OperIp        string      `json:"operIp"        description:"主机地址"`
	OperLocation  string      `json:"operLocation"  description:"操作地点"`
	OperParam     string      `json:"operParam"     description:"请求参数"`
	JsonResult    string      `json:"jsonResult"    description:"返回参数"`
	Status        int         `json:"status"        description:"操作状态（0正常 1异常）"`
	ErrorMsg      string      `json:"errorMsg"      description:"错误消息"`
	OperTime      *gtime.Time `json:"operTime"      description:"操作时间"`
}

type SysOperLogRes struct {
	OperId        uint64      `json:"operId"        description:"日志主键"`
	Title         string      `json:"title"         description:"模块标题"`
	BusinessType  int         `json:"businessType"  description:"业务类型（0其它 1新增 2修改 3删除）"`
	Method        string      `json:"method"        description:"方法名称"`
	RequestMethod string      `json:"requestMethod" description:"请求方式"`
	OperatorType  int         `json:"operatorType"  description:"操作类别（0其它 1后台用户 2手机端用户）"`
	OperName      string      `json:"operName"      description:"操作人员"`
	DeptName      string      `json:"deptName"      description:"部门名称"`
	OperUrl       string      `json:"operUrl"       description:"请求URL"`
	OperIp        string      `json:"operIp"        description:"主机地址"`
	OperLocation  string      `json:"operLocation"  description:"操作地点"`
	OperParam     string      `json:"operParam"     description:"请求参数"`
	JsonResult    string      `json:"jsonResult"    description:"返回参数"`
	Status        int         `json:"status"        description:"操作状态（0正常 1异常）"`
	ErrorMsg      string      `json:"errorMsg"      description:"错误消息"`
	OperTime      *gtime.Time `json:"operTime"      description:"操作时间"`
}
