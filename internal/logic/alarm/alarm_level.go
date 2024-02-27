package alarm

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
)

type sAlarmLevel struct{}

func init() {
	service.RegisterAlarmLevel(alarmLevelNew())
}

func alarmLevelNew() *sAlarmLevel {
	return &sAlarmLevel{}
}

func (s *sAlarmLevel) Detail(ctx context.Context, level uint) (out model.AlarmLevelOutput, err error) {
	err = dao.AlarmLevel.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     "AlarmLevelDetail",
		Force:    false,
	}).Where(dao.AlarmLevel.Columns().Level, level).Scan(&out)
	return
}

func (s *sAlarmLevel) All(ctx context.Context) (out *model.AlarmLevelListOutput, err error) {
	var p []*entity.AlarmLevel

	err = dao.AlarmLevel.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     "AlarmLevelAll",
		Force:    false,
	}).Scan(&p)
	if err != nil || p == nil {
		return
	}

	out = new(model.AlarmLevelListOutput)
	out.List = p

	return
}

func (s *sAlarmLevel) Edit(ctx context.Context, in []*model.AlarmLevelEditInput) (err error) {
	for _, v := range in {
		_, err = dao.AlarmLevel.Ctx(ctx).Data(g.Map{
			"name": v.Name,
		}).
			Where(dao.AlarmLevel.Columns().Level, v.Level).
			Update()
		if err != nil {
			return
		}
		_, err = cache.Instance().Remove(ctx, "AlarmLevelAll")
		if err != nil {
			continue
		}
	}
	return
}
