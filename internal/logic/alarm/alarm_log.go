package alarm

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
)

type sAlarmLog struct{}

func init() {
	service.RegisterAlarmLog(alarmLogNew())
}

func alarmLogNew() *sAlarmLog {
	return &sAlarmLog{}
}

func (s *sAlarmLog) Detail(ctx context.Context, id uint64) (out *model.AlarmLogOutput, err error) {
	err = dao.AlarmLog.Ctx(ctx).WithAll().Where(dao.AlarmLog.Columns().Id, id).Scan(&out)
	return
}

func (s *sAlarmLog) Add(ctx context.Context, in *model.AlarmLogAddInput) (id uint64, err error) {
	rs, err := dao.AlarmLog.Ctx(ctx).Data(in).Insert()
	if err != nil {
		return
	}
	newId, err := rs.LastInsertId()
	id = uint64(newId)
	return
}

func (s *sAlarmLog) List(ctx context.Context, in *model.AlarmLogListInput) (out *model.AlarmLogListOutput, err error) {
	out = new(model.AlarmLogListOutput)
	c := dao.AlarmLog.Columns()
	m := dao.AlarmLog.Ctx(ctx).WithAll().OrderDesc(c.Id)

	if len(in.DateRange) > 0 {
		m = m.WhereBetween(c.CreatedAt, in.DateRange[0], in.DateRange[1])
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Scan(&out.List)
	return
}

func (s *sAlarmLog) Handle(ctx context.Context, in *model.AlarmLogHandleInput) (err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	c := dao.AlarmLog.Columns()
	_, err = dao.AlarmLog.Ctx(ctx).Data(g.Map{
		c.Status:   in.Status,
		c.UpdateBy: loginUserId,
		c.Content:  in.Content,
	}).Where(c.Id, in.Id).Update()

	return
}

func (s *sAlarmLog) TotalForLevel(ctx context.Context) (total []model.AlarmLogLevelTotal, err error) {
	rs, err := dao.AlarmLog.Ctx(ctx).Fields("level, count(*) as num").Group(dao.AlarmLevel.Columns().Level).All()
	if err != nil || rs.Len() == 0 {
		return
	}

	level, err := service.AlarmLevel().All(ctx)
	if err != nil {
		return
	}
	l := make(map[uint]string, len(level.List))
	for _, v := range level.List {
		l[v.Level] = v.Name
	}

	logTotal := 0
	for _, v := range rs {
		logTotal += v["num"].Int()

		total = append(total, model.AlarmLogLevelTotal{
			Level: v["level"].Uint(),
			Name:  l[v["level"].Uint()],
			Num:   v["num"].Int(),
		})
	}

	for k, v := range total {
		n, _ := strconv.ParseFloat(fmt.Sprintf("%.0f", float64(v.Num)/float64(logTotal)*100), 64)
		total[k].Ratio = n
	}

	return
}

//ClearLogByDays 按日期删除日志
func (s *sAlarmLog) ClearLogByDays(ctx context.Context, days int) (err error) {
	_, err = dao.AlarmLog.Ctx(ctx).Delete("to_days(now())-to_days(`created_at`) > ?", days+1)
	return
}
