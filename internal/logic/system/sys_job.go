package system

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"
	"strings"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/util/gconv"
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
		err = m.Page(input.PageNum, input.PageSize).Order("job_id asc").Scan(&out)
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

func (s *sSysJob) AddJob(ctx context.Context, input *model.SysJobAddInput) (err error) {
	_, err = dao.SysJob.Ctx(ctx).Insert(input)
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
	_, err := dao.SysJob.Ctx(ctx).FieldsEx(dao.SysJob.Columns().JobId, dao.SysJob.Columns().CreateBy).Where(dao.SysJob.Columns().JobId, input.JobId).
		Update(input)

	// 同步定时任务到数据源和数据模型
	if input.JobGroup == "dataSourceJob" {
		if input.InvokeTarget == "dataSource" {
			err = service.DataSource().UpdateInterval(ctx, gconv.Uint64(input.JobParams), input.CronExpression)
		} else if input.InvokeTarget == "dataTemplate" {
			err = service.DataTemplate().UpdateInterval(ctx, gconv.Uint64(input.JobParams), input.CronExpression)
		}
	}
	return err
}

// JobStart 启动任务
func (s *sSysJob) JobStart(ctx context.Context, job *model.SysJobOut) error {
	//获取task目录下是否绑定对应的方法
	f := jobTask.TimeTaskList.GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("没有绑定对应的方法")
	}
	//传参
	paramArr := strings.Split(job.JobParams, "|")
	jobTask.TimeTaskList.EditParams(f.FuncName, paramArr)

	jname := fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId)

	rs := gcron.Search(jname)
	if rs == nil {
		newCtx := ctx
		if job.JobGroup == "dataSourceJob" {
			newCtx = s.WithValue(ctx, paramArr[0])
		}
		if job.MisfirePolicy == 1 {
			t, err := gcron.AddSingleton(newCtx, job.CronExpression, f.Run, jname)
			if err != nil {
				return err
			}
			if t == nil {
				return gerror.New("启动任务失败")
			}
		} else {
			t, err := gcron.AddOnce(newCtx, job.CronExpression, f.Run, jname)
			if err != nil {
				return err
			}
			if t == nil {
				return gerror.New("启动任务失败")
			}
		}
	}
	gcron.Start(jname)
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
func (s *sSysJob) JobStartMult(ctx context.Context, jobs []*model.SysJobOut) error {

	var jobIds = g.Slice{}
	for _, job := range jobs {
		//获取task目录下是否绑定对应的方法
		f := jobTask.TimeTaskList.GetByName(job.InvokeTarget)
		if f == nil {
			return gerror.New("没有绑定对应的方法")
		}
		//传参
		paramArr := strings.Split(job.JobParams, "|")
		jobTask.TimeTaskList.EditParams(f.FuncName, paramArr)

		jname := fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId)

		rs := gcron.Search(jname)
		if rs == nil {
			newCtx := ctx
			if job.JobGroup == "dataSourceJob" {
				newCtx = s.WithValue(ctx, paramArr[0])
			}
			if job.MisfirePolicy == 1 {
				t, err := gcron.AddSingleton(newCtx, job.CronExpression, f.Run, jname)
				if err != nil {
					return err
				}
				if t == nil {
					return gerror.New("启动任务失败")
				}
			} else {
				t, err := gcron.AddOnce(newCtx, job.CronExpression, f.Run, jname)
				if err != nil {
					return err
				}
				if t == nil {
					return gerror.New("启动任务失败")
				}
			}
		}
		gcron.Start(jname)
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
func (s *sSysJob) JobStop(ctx context.Context, job *model.SysJobOut) error {
	//获取task目录下是否绑定对应的方法
	f := jobTask.TimeTaskList.GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("没有绑定对应的方法")
	}

	jname := fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId)

	rs := gcron.Search(jname)
	if rs != nil {
		gcron.Remove(jname)
	}
	job.Status = 1
	_, err := dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, job.JobId).Unscoped().Update(g.Map{
		dao.SysJob.Columns().Status: job.Status,
	})
	return err
}

// JobRun 执行任务
func (s *sSysJob) JobRun(ctx context.Context, job *model.SysJobOut) error {
	//可以task目录下是否绑定对应的方法
	f := jobTask.TimeTaskList.GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}
	//传参
	paramArr := strings.Split(job.JobParams, "|")
	jobTask.TimeTaskList.EditParams(f.FuncName, paramArr)
	newCtx := ctx
	if job.JobGroup == "dataSourceJob" {
		newCtx = s.WithValue(ctx, paramArr[0])
	}
	task, err := gcron.AddOnce(newCtx, "@every 1s", f.Run)
	if err != nil || task == nil {
		return gerror.New("启动执行失败")
	}
	return nil
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
