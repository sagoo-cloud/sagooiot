package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysJob = cSysJob{}

type cSysJob struct{}

// List 任务列表
func (a *cSysJob) List(ctx context.Context, req *system.GetJobListReq) (res *system.GetJobListRes, err error) {

	var input *model.GetJobListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	total, out, err := service.SysJob().JobList(ctx, input)

	if err != nil {
		return
	}
	res = new(system.GetJobListRes)
	res.Total = total
	res.CurrentPage = req.PageNum

	if out != nil {
		if err = gconv.Scan(out, &res.Data); err != nil {
			return
		}
	}

	return
}

// FunList 获取任务可用方法名列表
func (a *cSysJob) FunList(ctx context.Context, req *system.GetJobFunListReq) (res *system.GetJobFunListRes, err error) {
	resData, err := service.SysJob().GetJobFuns(ctx)
	res = &system.GetJobFunListRes{
		Data: resData,
	}
	return
}

func (a *cSysJob) Add(ctx context.Context, req *system.AddJobReq) (res *system.AddJobRes, err error) {

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if loginUserId == 0 {
		err = gerror.New("未登录或TOKEN失效,请重新登录")
		return
	}

	req.CreateBy = gconv.Uint64(loginUserId)
	var input *model.SysJobAddInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	err = service.SysJob().AddJob(ctx, input)
	return
}

func (a *cSysJob) Get(ctx context.Context, req *system.GetJobByIdReq) (res *system.GetJobByIdRes, err error) {
	job, err := service.SysJob().GetJobInfoById(ctx, req.Id)
	if err != nil {
		return
	}
	var jobRes *model.SysJobRes
	if job != nil {
		if err = gconv.Scan(job, &jobRes); err != nil {
			return
		}
	}
	res = &system.GetJobByIdRes{
		Data: jobRes,
	}
	return
}

func (a *cSysJob) Edit(ctx context.Context, req *system.EditJobReq) (res *system.EditJobRes, err error) {

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if loginUserId == 0 {
		err = gerror.New("未登录或TOKEN失效,请重新登录")
		return
	}

	req.UpdateBy = gconv.Uint64(loginUserId) //获取登陆用户id

	var input *model.SysJobEditInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	err = service.SysJob().EditJob(ctx, input)
	return
}

// Start 启动任务
func (a *cSysJob) Start(ctx context.Context, req *system.StartJobByIdReq) (res *system.StartJobByIdRes, err error) {

	job, err := service.SysJob().GetJobInfoById(ctx, req.Id)
	err = service.SysJob().JobStart(ctx, job)
	return
}

// Stop 停止任务
func (a *cSysJob) Stop(ctx context.Context, req *system.StopJobByIdReq) (res *system.StopJobByIdRes, err error) {

	job, err := service.SysJob().GetJobInfoById(ctx, req.Id)
	err = service.SysJob().JobStop(ctx, job)
	return
}

// Run 执行任务
func (a *cSysJob) Run(ctx context.Context, req *system.RunJobByIdReq) (res *system.RunJobByIdRes, err error) {

	job, err := service.SysJob().GetJobInfoById(ctx, req.Id)
	err = service.SysJob().JobRun(ctx, job)
	return
}

// Delete 删除任务
func (a *cSysJob) Delete(ctx context.Context, req *system.DeleteJobByIdReq) (res *system.DeleteJobByIdRes, err error) {
	err = service.SysJob().DeleteJobByIds(ctx, req.Id)
	return
}
