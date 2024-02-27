package alarm

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"strconv"
	"time"

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
	var deptId int
	if in.Type != 2 {
		var alarmRule *model.AlarmRuleOutput
		//获取规则信息
		alarmRule, err = service.AlarmRule().Detail(ctx, in.RuleId)
		if err != nil {
			return
		}
		if alarmRule != nil {
			deptId = alarmRule.DeptId
		}
	}
	rs, err := dao.AlarmLog.Ctx(ctx).Data(do.AlarmLog{
		DeptId:     deptId,
		Type:       in.Type,
		RuleId:     in.RuleId,
		RuleName:   in.RuleName,
		Level:      in.Level,
		Data:       in.Data,
		ProductKey: in.ProductKey,
		DeviceKey:  in.DeviceKey,
		Status:     0,
		CreatedAt:  gtime.Now(),
		Content:    "",
	}).Insert()
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
	m := dao.AlarmLog.Ctx(ctx).WithAll().
		OrderDesc(c.Id)

	if len(in.DateRange) > 0 {
		m = m.WhereBetween(c.CreatedAt, in.DateRange[0], in.DateRange[1])
	}

	if len(in.Status) > 0 {
		m = m.Where(dao.AlarmLog.Columns().Status, in.Status)
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Scan(&out.List)
	return
}

func (s *sAlarmLog) Handle(ctx context.Context, in *model.AlarmLogHandleInput) (err error) {
	alarmLog, err := s.Detail(ctx, in.Id)
	if err != nil {
		return
	}
	if alarmLog == nil {
		return gerror.New("告警日志不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	c := dao.AlarmLog.Columns()
	_, err = dao.AlarmLog.Ctx(ctx).Data(g.Map{
		c.Status:    in.Status,
		c.UpdatedBy: loginUserId,
		c.Content:   in.Content,
	}).Where(c.Id, in.Id).Update()

	//更新缓存
	alarmLog.Status = in.Status
	key := consts.DeviceAlarmLogPrefix + alarmLog.ProductKey + alarmLog.DeviceKey + alarmLog.Expression
	err = cache.Instance().Set(ctx, key, alarmLog, time.Minute*10)

	return
}

func (s *sAlarmLog) TotalForLevel(ctx context.Context) (total []model.AlarmLogLevelTotal, err error) {
	//TODO 缓存时间需要优化 ====================
	rs, err := dao.AlarmLog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 1000,
		Name:     "AlarmLogTotalForLevel",
		Force:    false,
	}).Fields("level, count(*) as num").Group(dao.AlarmLevel.Columns().Level).All()
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

// ClearLogByDays 按日期删除日志
func (s *sAlarmLog) ClearLogByDays(ctx context.Context, days int) (err error) {
	_, err = dao.AlarmLog.Ctx(ctx).Delete("to_days(now())-to_days(`created_at`) > ?", days+1)
	return
}
