package system

import (
	"context"
	"fmt"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
	"sagooiot/pkg/response"
	"sagooiot/pkg/utility"
	"sagooiot/pkg/worker"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysJobLog struct {
}

func sysJobLogNew() *sSysJobLog {
	return &sSysJobLog{}
}
func init() {
	service.RegisterSysJobLog(sysJobLogNew())
}

func (s *sSysJobLog) GetJobLog(ctx context.Context, id int) (out *model.SysJobLogOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		out = new(model.SysJobLogOut)
		err = dao.SysJobLog.Ctx(ctx).Where(dao.SysJobLog.Columns().IsDeleted, 0).Where(dao.SysJobLog.Columns().Id, id).Scan(out)
		if err != nil {
			err = gerror.New("获取任务日志详情失败")
		}
	})
	return
}

// JobLogList 任务日志列表
func (s *sSysJobLog) JobLogList(ctx context.Context, input *model.GetJobLogListInput) (total int, out []*model.SysJobLogOut, err error) {
	m := dao.SysJobLog.Ctx(ctx)
	m = m.Where(dao.SysJobLog.Columns().IsDeleted, 0) // 只查询未删除的记录
	if input != nil {
		if input.Status != "" {
			m = m.Where(dao.SysJobLog.Columns().Status, gconv.Int(input.Status))
		}
		if input.JobName != "" {
			m = m.WhereLike(dao.SysJob.Columns().JobName, "%"+input.JobName+"%")
		}
		if len(input.DateRange) > 1 {
			m = m.WhereGTE(dao.SysJobLog.Columns().CreatedAt, input.DateRange[0]+" 00:00:00")
			m = m.WhereLTE(dao.SysJobLog.Columns().CreatedAt, input.DateRange[1]+" 23:59:59")
		}

		if input.PageNum == 0 {
			input.PageNum = 1
		}
		if input.PageSize == 0 {
			input.PageSize = consts.PageSize
		}

		m = m.Page(input.PageNum, input.PageSize)

	}
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取总行数失败")
		return
	}

	err = m.OrderDesc(dao.SysJobLog.Columns().CreatedAt).Scan(&out)
	if err != nil {
		err = gerror.New("获取任务日志列表失败")
	}
	return
}

// AddJobLog 添加任务日志
func (s *sSysJobLog) AddJobLog(ctx context.Context, input *model.SysJobLogAddInput) error {
	//获取task目录下是否绑定对应的方法
	checkName := worker.TasksInstance().CheckFuncName(input.InvokeTarget)
	if !checkName {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", input.InvokeTarget)
		return gerror.New(errInfo)
	}

	_, err := dao.SysJobLog.Ctx(ctx).Data(do.SysJobLog{
		JobName:        input.JobName,
		InvokeTarget:   input.InvokeTarget,
		CronExpression: input.CronExpression,
		StartTime:      input.StartTime,
		EndTime:        input.EndTime,
		JobMessage:     input.JobMessage,
		Status:         input.Status,
		ExceptionInfo:  input.ExceptionInfo,
		CreatedAt:      gtime.Now(),
	}).Insert()

	return err
}

func (s *sSysJobLog) DelJobLogByIds(ctx context.Context, ids []int) (err error) {
	userId := service.Context().GetUserId(ctx)

	_, err = dao.SysJobLog.Ctx(ctx).Data(g.Map{
		dao.SysJobLog.Columns().DeletedBy: userId,
		dao.SysJobLog.Columns().DeletedAt: gtime.Now(),
		dao.SysJobLog.Columns().IsDeleted: 1,
	}).Where(dao.SysJobLog.Columns().Id, ids).Update()
	if err != nil {
		err = gerror.New("删除任务日志失败")
	}
	return
}

func (s *sSysJobLog) Export(ctx context.Context, input *model.GetJobLogListInput) (err error) {
	m := dao.SysJobLog.Ctx(ctx)
	m = m.Where("is_deleted", 0)
	if input != nil {
		if input.Status != "" {
			m = m.Where(dao.SysJobLog.Columns().Status, gconv.Int(input.Status))
		}
		if input.JobName != "" {
			m = m.WhereLike(dao.SysJob.Columns().JobName, "%"+input.JobName+"%")
		}
		if len(input.DateRange) > 1 {
			m = m.WhereGTE(dao.SysJobLog.Columns().CreatedAt, input.DateRange[0]+" 00:00:00")
			m = m.WhereLTE(dao.SysJobLog.Columns().CreatedAt, input.DateRange[1]+" 23:59:59")
		}
	}
	//获取任务日志列表信息
	var outList []*model.SysJobLogOut
	err = m.Scan(&outList)
	if err != nil {
		err = gerror.New("获取任务日志列表失败")
		return
	}

	var resData []interface{}
	for _, out := range outList {
		var exportOut = new(model.SysJobLogExportOut)
		if err = gconv.Scan(out, exportOut); err != nil {
			return
		}
		if out.Status == 0 {
			exportOut.Status = "成功"
		} else if out.Status == 1 {
			exportOut.Status = "失败"
		}
		resData = append(resData, exportOut)
	}
	data := utility.ToExcel(resData)
	var request = g.RequestFromCtx(ctx)
	response.ToXls(request, data, "SysJobLog")
	return
}

func (s *sSysJobLog) ClearJobLogByDays(ctx context.Context, days int) (err error) {
	_, err = dao.SysJobLog.Ctx(ctx).Delete("to_days(now())-to_days(`created_at`) > ?", days+1)
	return
}
