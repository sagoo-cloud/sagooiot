package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
	"sagooiot/internal/tasks"
	"sagooiot/pkg/worker"
)

type sSysJob struct {
}

func sysJobNew() *sSysJob {
	return &sSysJob{}
}

func init() {
	service.RegisterSysJob(sysJobNew())
}

// JobList 获取任务列表
func (s *sSysJob) JobList(ctx context.Context, input *model.GetJobListInput) (total int, out []*model.SysJobOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysJob.Ctx(ctx)
		if input != nil {
			if input.Status != "" {
				m = m.Where("status", gconv.Int(input.Status))
			}
			if input.JobGroup != "" {
				m = m.Where("job_group", input.JobGroup)
			}
			if input.JobName != "" {
				m = m.Where("job_name like ?", "%"+input.JobName+"%")
			}
		}
		total, err = m.Count()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		if input.PageNum == 0 {
			input.PageNum = 1
		}
		if input.PageSize == 0 {
			input.PageSize = consts.PageSize
		}
		err = m.Page(input.PageNum, input.PageSize).Order("job_id desc").Scan(&out)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})
	return
}

// GetJobs 获取已开启执行的任务
func (s *sSysJob) GetJobs(ctx context.Context) (jobs []*model.SysJobOut, err error) {
	err = dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().Status, 0).Scan(&jobs)
	return
}

// GetJobFuns 获取任务可用方法列表
func (s *sSysJob) GetJobFuns(ctx context.Context) (jobsList []*model.SysJobFunListOut, err error) {
	funList := worker.TasksInstance().GetTaskJobNameList()
	for k, v := range funList {
		var fun = new(model.SysJobFunListOut)
		fun.FunName = k
		fun.Explain = v
		jobsList = append(jobsList, fun)
	}
	return
}

func (s *sSysJob) AddJob(ctx context.Context, input *model.SysJobAddInput) (err error) {
	//获取task目录下是否绑定对应的方法
	checkName := worker.TasksInstance().CheckFuncName(input.InvokeTarget)
	if !checkName {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", input.InvokeTarget)
		return gerror.New(errInfo)
	}

	_, err = dao.SysJob.Ctx(ctx).Data(do.SysJob{
		JobName:        input.JobName,
		JobParams:      input.JobParams,
		JobGroup:       input.JobGroup,
		InvokeTarget:   input.InvokeTarget,
		CronExpression: input.CronExpression,
		MisfirePolicy:  input.MisfirePolicy,
		Concurrent:     input.Concurrent,
		Status:         input.Status,
		CreatedBy:      input.CreateBy,
		Remark:         input.Remark,
		CreatedAt:      gtime.Now(),
	}).Insert()
	return
}

func (s *sSysJob) GetJobInfoById(ctx context.Context, id int) (job *model.SysJobOut, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.SysJob.Ctx(ctx).Where("job_id", id).Scan(&job)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	if job == nil || err != nil {
		err = gerror.New("获取任务信息失败")
	}
	return
}

func (s *sSysJob) EditJob(ctx context.Context, input *model.SysJobEditInput) error {
	_, err := dao.SysJob.Ctx(ctx).FieldsEx(dao.SysJob.Columns().JobId, dao.SysJob.Columns().CreatedBy).Where(dao.SysJob.Columns().JobId, input.JobId).
		Update(input)
	return err
}

// JobStart 启动任务
func (s *sSysJob) JobStart(ctx context.Context, job *model.SysJobOut) error {
	//获取task目录下是否绑定对应的方法
	checkName := worker.TasksInstance().CheckFuncName(job.InvokeTarget)
	if !checkName {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", job.InvokeTarget)
		return gerror.New(errInfo)
	}

	//传参解析
	paramArr, err := worker.TasksInstance().ParseParameters(job.JobParams)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	taskData := tasks.TaskJob{
		ID:         fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId),
		TaskType:   "Type-" + gconv.String(job.MisfirePolicy),
		MethodName: job.InvokeTarget,
		Params:     paramArr,
		Explain:    job.JobName,
	}
	runPayload, _ := json.Marshal(taskData)
	if job.MisfirePolicy == 1 {
		err := worker.TasksInstance().Cron(
			worker.WithRunCtx(context.Background()),
			worker.WithRunUuid(taskData.ID),          // 任务ID
			worker.WithRunGroup(taskData.MethodName), // 任务组
			worker.WithRunExpr(job.CronExpression),
			worker.WithRunTimeout(10),
			worker.WithRunReplace(true),
			worker.WithRunPayload(runPayload),
		)
		if err != nil {
			g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务失败")
			return err
		}
	} else {
		err := worker.TasksInstance().Once(
			worker.WithRunCtx(context.Background()),
			worker.WithRunUuid(taskData.ID),          // 任务ID
			worker.WithRunGroup(taskData.MethodName), // 任务组
			worker.WithRunTimeout(10),
			worker.WithRunNow(true),
			worker.WithRunReplace(true),
			worker.WithRunPayload(runPayload),
		)
		if err != nil {
			g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务失败")
			return err
		}
	}

	if job.MisfirePolicy == 1 {
		job.Status = 0
		_, err := dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, job.JobId).Unscoped().Update(g.Map{
			dao.SysJob.Columns().Status: job.Status,
		})
		return err
	}
	return nil
}

// JobStartMult 批量启动任务
func (s *sSysJob) JobStartMult(ctx context.Context, jobsList []*model.SysJobOut) error {
	var jobIds = g.Slice{}
	for _, job := range jobsList {
		//获取task目录下是否绑定对应的方法
		checkName := worker.TasksInstance().CheckFuncName(job.InvokeTarget)
		if !checkName {
			g.Log().Debugf(ctx, "没有绑定对应的方法:%s", job.InvokeTarget)
			continue
		}

		//传参解析
		paramArr, err := worker.TasksInstance().ParseParameters(job.JobParams)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}

		taskData := tasks.TaskJob{
			ID:         fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId),
			TaskType:   "Type-" + gconv.String(job.MisfirePolicy),
			MethodName: job.InvokeTarget,
			Params:     paramArr,
			Explain:    job.JobName,
		}
		runPayload, _ := json.Marshal(taskData)

		if job.MisfirePolicy == 1 {
			err := worker.TasksInstance().Cron(
				worker.WithRunCtx(ctx),
				worker.WithRunUuid(taskData.ID),          // 任务ID
				worker.WithRunGroup(taskData.MethodName), // 任务组
				worker.WithRunExpr(job.CronExpression),
				worker.WithRunTimeout(10),
				worker.WithRunReplace(true),
				worker.WithRunPayload(runPayload),
			)
			if err != nil {
				g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务失败")
				continue
			}
		} else {
			err := worker.TasksInstance().Once(
				worker.WithRunCtx(ctx),
				worker.WithRunUuid(taskData.ID),          // 任务ID
				worker.WithRunGroup(taskData.MethodName), // 任务组
				worker.WithRunTimeout(10),
				worker.WithRunNow(true),
				worker.WithRunReplace(true),
				worker.WithRunPayload(runPayload),
			)
			if err != nil {
				g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务失败")
				continue
			}
		}

		if job.MisfirePolicy == 1 {
			jobIds = append(jobIds, job.JobId)
		}
	}

	_, err := dao.SysJob.Ctx(ctx).WhereIn(dao.SysJob.Columns().JobId, jobIds).Unscoped().Update(g.Map{
		dao.SysJob.Columns().Status: 0,
	})
	return err
}

// JobStop 停止任务
func (s *sSysJob) JobStop(ctx context.Context, job *model.SysJobOut) (err error) {
	//获取task目录下是否绑定对应的方法
	checkName := worker.TasksInstance().CheckFuncName(job.InvokeTarget)
	if !checkName {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", job.InvokeTarget)
		return errors.New(errInfo)
	}

	taskJobId := fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId)
	_ = worker.TasksInstance().Remove(ctx, taskJobId)
	job.Status = 1
	_, err = dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, job.JobId).Unscoped().Update(g.Map{
		dao.SysJob.Columns().Status: job.Status,
	})
	return
}

// JobRun 执行任务
func (s *sSysJob) JobRun(ctx context.Context, job *model.SysJobOut) (err error) {
	//可以task目录下是否绑定对应的方法
	checkName := worker.TasksInstance().CheckFuncName(job.InvokeTarget)
	if !checkName {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", job.InvokeTarget)
		return errors.New(errInfo)
	}

	//传参解析
	paramArr, err := worker.TasksInstance().ParseParameters(job.JobParams)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	taskData := tasks.TaskJob{
		ID:         fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId),
		TaskType:   "Type-" + gconv.String(job.MisfirePolicy),
		MethodName: job.InvokeTarget,
		Params:     paramArr,
		Explain:    job.JobName,
	}
	runPayload, _ := json.Marshal(taskData)

	err = worker.TasksInstance().Once(
		worker.WithRunCtx(context.Background()),
		worker.WithRunUuid(taskData.ID),          // 任务ID
		worker.WithRunGroup(taskData.MethodName), // 任务组
		worker.WithRunTimeout(10),
		worker.WithRunNow(true),
		worker.WithRunReplace(true),
		worker.WithRunPayload(runPayload),
	)
	if err != nil {
		errInfo := fmt.Sprintf(taskData.MethodName, taskData.Explain, "启动任务失败")
		return errors.New(errInfo)
	}
	return
}

// DeleteJobByIds 删除任务
func (s *sSysJob) DeleteJobByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	gst := gset.NewFrom(ids)
	var jobs []*model.SysJobOut
	jobs, err = s.GetJobs(ctx)
	if err != nil {
		return
	}
	for _, job := range jobs {
		if gst.Contains(int(job.JobId)) {
			err = gerror.New("运行中的任务不能删除")
			return
		}
	}
	_, err = dao.SysJob.Ctx(ctx).Delete(dao.SysJob.Columns().JobId+" in (?)", ids)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("删除失败")
	}
	return
}

// 数据源、数据模型使用
type contextKey string

func (s *sSysJob) WithValue(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, contextKey("id"), value)
}

func (s *sSysJob) Value(ctx context.Context) uint64 {
	value := ctx.Value(contextKey("id"))
	return gconv.Uint64(value)
}
