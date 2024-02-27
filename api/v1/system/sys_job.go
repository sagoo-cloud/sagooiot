package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

type GetJobListReq struct {
	g.Meta   `path:"/job/list" method:"get" summary:"获取任务列表" tags:"定时任务管理"`
	JobName  string `json:"jobName" description:"任务名称"`
	JobGroup string `json:"jobGroup" description:"任务组名"`
	Status   string `json:"status" description:"状态（0正常 1暂停）"`
	*common.PaginationReq
}
type GetJobListRes struct {
	Data []*model.SysJobRes
	common.PaginationRes
}

type GetJobFunListReq struct {
	g.Meta `path:"/job/fun_list" method:"get" summary:"获取任务可用方法列表" tags:"定时任务管理"`
}
type GetJobFunListRes struct {
	Data []*model.SysJobFunListOut
}

type AddJobReq struct {
	g.Meta         `path:"/job/add" method:"post" summary:"添加定时任务" tags:"定时任务管理"`
	JobName        string `json:"jobName"  description:"任务名称" v:"required#任务名称不能来空"`
	JobParams      string `json:"jobParams"  description:"任务参数"`
	JobGroup       string `json:"jobGroup" description:"分组"`
	InvokeTarget   string `json:"invokeTarget" description:"执行方法" v:"required#执行方法不能为空"`
	CronExpression string `json:"cronExpression" description:"任务执行表达式" v:"required#任务表达式不能为空"`
	MisfirePolicy  int    `json:"misfirePolicy"`
	Concurrent     int    `json:"concurrent" `
	Status         int    `json:"status" description:"状态" v:"required#状态（0正常 1暂停）不能为空"`
	Remark         string `json:"remark" `
	CreateBy       uint64
}
type AddJobRes struct {
}

type EditJobReq struct {
	g.Meta         `path:"/job/edit" method:"put" summary:"编辑定时任务" tags:"定时任务管理"`
	JobId          int64  `json:"job_id" v:"min:1#任务id不能为空"`
	JobName        string `json:"jobName"  description:"任务名称" v:"required#任务名称不能来空"`
	JobParams      string `json:"jobParams"  description:"任务参数"`
	JobGroup       string `json:"jobGroup" description:"分组"`
	InvokeTarget   string `json:"invokeTarget" description:"执行方法" v:"required#执行方法不能为空"`
	CronExpression string `json:"cronExpression" description:"任务执行表达式" v:"required#任务表达式不能为空"`
	MisfirePolicy  int    `json:"misfirePolicy"`
	Concurrent     int    `json:"concurrent" `
	Status         int    `json:"status" description:"状态" v:"required#状态（0正常 1暂停）不能为空"`
	Remark         string `json:"remark" `
	CreateBy       uint64
	UpdateBy       uint64
}
type EditJobRes struct {
}

type DeleteJobByIdReq struct {
	g.Meta `path:"/job/delJobById" method:"delete" summary:"根据ID删除任务" tags:"定时任务管理"`
	Id     []int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type DeleteJobByIdRes struct {
}

type GetJobByIdReq struct {
	g.Meta `path:"/job/getJobById" method:"get" summary:"根据ID获取任务" tags:"定时任务管理"`
	Id     int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type GetJobByIdRes struct {
	Data *model.SysJobRes
}

type StartJobByIdReq struct {
	g.Meta `path:"/job/start" method:"put" summary:"开始一个任务" tags:"定时任务管理"`
	Id     int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type StartJobByIdRes struct {
}

type StopJobByIdReq struct {
	g.Meta `path:"/job/stop" method:"put" summary:"结束一个任务" tags:"定时任务管理"`
	Id     int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type StopJobByIdRes struct {
}

type RunJobByIdReq struct {
	g.Meta `path:"/job/run" method:"put" summary:"执行一个任务" tags:"定时任务管理"`
	Id     int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type RunJobByIdRes struct {
}
