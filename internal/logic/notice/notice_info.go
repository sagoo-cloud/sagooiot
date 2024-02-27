package notice

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
)

type sNoticeInfo struct{}

func sNoticeInfoNew() *sNoticeInfo {
	return &sNoticeInfo{}
}
func init() {
	service.RegisterNoticeInfo(sNoticeInfoNew())
}

// GetNoticeInfoList 获取列表数据
func (s *sNoticeInfo) GetNoticeInfoList(ctx context.Context, in *model.GetNoticeInfoListInput) (total, page int, list []*model.NoticeInfoOutput, err error) {
	m := dao.NoticeInfo.Ctx(ctx)

	if in.KeyWord != "" {
		m = m.WhereLike(dao.NoticeInfo.Columns().MsgTitle, "%"+in.KeyWord+"%")
		m = m.WhereOrLike(dao.NoticeInfo.Columns().MsgBody, "%"+in.KeyWord+"%")
	}
	if in.Method != "" {
		m = m.Where(dao.NoticeInfo.Columns().Method, in.Method)
	}
	if in.ConfigId != "" {
		m = m.Where(dao.NoticeInfo.Columns().ConfigId, in.ConfigId)
	}

	if in.ComeFrom != "" {
		m = m.Where(dao.NoticeInfo.Columns().ComeFrom, in.ComeFrom)
	}

	if in.Status != -1 {
		m = m.Where(dao.NoticeInfo.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取总行数失败")
		return
	}
	page = in.PageNum
	if in.PageSize == 0 {
		in.PageSize = consts.PageSize
	}
	err = m.Page(page, in.PageSize).Order("created_at desc").Scan(&list)
	if err != nil {
		err = gerror.New("获取数据失败")
	}
	return
}

// GetNoticeInfoById 获取指定ID数据
func (s *sNoticeInfo) GetNoticeInfoById(ctx context.Context, id int) (out *model.NoticeInfoOutput, err error) {
	err = dao.NoticeInfo.Ctx(ctx).Where(dao.NoticeInfo.Columns().Id, id).Scan(&out)
	return
}

// AddNoticeInfo 添加数据
func (s *sNoticeInfo) AddNoticeInfo(ctx context.Context, in model.NoticeInfoAddInput) (err error) {
	_, err = dao.NoticeInfo.Ctx(ctx).Data(do.NoticeInfo{
		ConfigId:   in.ConfigId,
		ComeFrom:   in.ComeFrom,
		Method:     in.Method,
		MsgTitle:   in.MsgTitle,
		MsgBody:    in.MsgBody,
		MsgUrl:     in.MsgUrl,
		UserIds:    in.UserIds,
		OrgIds:     in.OrgIds,
		Totag:      in.Totag,
		Status:     in.Status,
		MethodCron: in.MethodCron,
		MethodNum:  in.MethodNum,
		CreatedAt:  gtime.Now(),
	}).Insert()
	return
}

// EditNoticeInfo 修改数据
func (s *sNoticeInfo) EditNoticeInfo(ctx context.Context, in model.NoticeInfoEditInput) (err error) {
	_, err = dao.NoticeInfo.Ctx(ctx).FieldsEx(dao.NoticeInfo.Columns().Id).Where(dao.NoticeInfo.Columns().Id, in.Id).Update(in)
	return
}

// DeleteNoticeInfo 删除数据
func (s *sNoticeInfo) DeleteNoticeInfo(ctx context.Context, Ids []int) (err error) {
	_, err = dao.NoticeInfo.Ctx(ctx).Where(dao.NoticeInfo.Columns().Id+" in(?)", Ids).Delete()
	if err != nil {
		return errors.New("删除通知数据失败")
	}
	return
}
