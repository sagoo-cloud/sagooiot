package notice

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sNoticeLog struct{}

func init() {
	service.RegisterNoticeLog(noticeLogNew())
}

func noticeLogNew() *sNoticeLog {
	return &sNoticeLog{}
}

// 通知日志记录
func (s *sNoticeLog) Add(ctx context.Context, in *model.NoticeLogAddInput) (err error) {
	_, err = dao.NoticeLog.Ctx(ctx).Data(in).Insert()
	return
}

// 删除日志
func (s *sNoticeLog) Del(ctx context.Context, ids []uint64) (err error) {
	_, err = dao.NoticeLog.Ctx(ctx).WhereIn(dao.NoticeLog.Columns().Id, ids).Delete()
	return
}

// 搜索
func (s *sNoticeLog) Search(ctx context.Context, in *model.NoticeLogSearchInput) (out *model.NoticeLogSearchOutput, err error) {
	out = new(model.NoticeLogSearchOutput)
	m := dao.NoticeLog.Ctx(ctx)
	if in.Status != -1 {
		m = m.Where(dao.NoticeLog.Columns().Status, in.Status)
	}
	if len(in.DateRange) > 0 {
		m = m.WhereBetween(dao.NoticeLog.Columns().SendTime, in.DateRange[0], in.DateRange[1])
	}

	out.CurrentPage = in.PageNum
	if out.Total, err = m.Count(); err != nil || out.Total == 0 {
		return
	}
	var list []model.NoticeLogList
	if err = m.Page(in.PageNum, in.PageSize).OrderDesc(dao.NoticeLog.Columns().Id).Scan(&list); err != nil {
		return
	}

	// 获取发送通道中文配置
	var dtList []model.SysDictDataOut
	err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictType.Columns().DictType, "notice_send_gateway").Scan(&dtList)

	for i, v := range list {
		for _, t := range dtList {
			if v.SendGateway == t.DictValue {
				list[i].Gateway = t.DictLabel
			}
		}
	}
	out.List = list

	return
}

//ClearLogByDays 按日期删除日志
func (s *sNoticeLog) ClearLogByDays(ctx context.Context, days int) (err error) {
	_, err = dao.NoticeLog.Ctx(ctx).Delete("to_days(now())-to_days(`send_time`) > ?", days+1)
	return
}
