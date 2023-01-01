package alarm

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sAlarmLevel struct{}

func init() {
	service.RegisterAlarmLevel(alarmLevelNew())
}

func alarmLevelNew() *sAlarmLevel {
	return &sAlarmLevel{}
}

func (s *sAlarmLevel) Detail(ctx context.Context, level uint) (out model.AlarmLevelOutput, err error) {
	err = dao.AlarmLevel.Ctx(ctx).Where(dao.AlarmLevel.Columns().Level, level).Scan(&out)
	return
}

func (s *sAlarmLevel) All(ctx context.Context) (out *model.AlarmLevelListOutput, err error) {
	var p []*entity.AlarmLevel

	err = dao.AlarmLevel.Ctx(ctx).Scan(&p)
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
	}

	return
}
