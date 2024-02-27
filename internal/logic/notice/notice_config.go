package notice

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
)

type sNoticeConfig struct{}

func sNoticeConfigNew() *sNoticeConfig {
	return &sNoticeConfig{}
}
func init() {
	service.RegisterNoticeConfig(sNoticeConfigNew())
}

// GetNoticeConfigList 获取列表数据
func (s *sNoticeConfig) GetNoticeConfigList(ctx context.Context, in *model.GetNoticeConfigListInput) (total, page int, list []*model.NoticeConfigOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NoticeConfig.Ctx(ctx)

		if in.KeyWord != "" {
			m = m.WhereLike(dao.NoticeConfig.Columns().Title, "%"+in.KeyWord+"%")
		}
		if in.Types != "" {
			m = m.Where(dao.NoticeConfig.Columns().Types, in.Types)
		}
		if in.SendGateway != "" {
			m = m.Where(dao.NoticeConfig.Columns().SendGateway, in.SendGateway)
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
	})
	return
}

// GetNoticeConfigById 获取指定ID数据
func (s *sNoticeConfig) GetNoticeConfigById(ctx context.Context, id string) (out *model.NoticeConfigOutput, err error) {
	err = dao.NoticeConfig.Ctx(ctx).Where(dao.NoticeConfig.Columns().Id, id).Scan(&out)
	return
}

// AddNoticeConfig 添加数据
func (s *sNoticeConfig) AddNoticeConfig(ctx context.Context, in model.NoticeConfigAddInput) (err error) {
	_, err = dao.NoticeConfig.Ctx(ctx).Data(do.NoticeConfig{
		Id:          guid.S(),
		DeptId:      service.Context().GetUserDeptId(ctx),
		Title:       in.Title,
		SendGateway: in.SendGateway,
		Types:       in.Types,
		CreatedAt:   gtime.Now(),
	}).Insert()
	return
}

// EditNoticeConfig 修改数据
func (s *sNoticeConfig) EditNoticeConfig(ctx context.Context, in model.NoticeConfigEditInput) (err error) {
	noticeConfig, err := s.GetNoticeConfigById(ctx, in.Id)
	if err != nil {
		return
	}
	if noticeConfig == nil {
		return gerror.New("通知配置不存在")
	}

	_, err = dao.NoticeConfig.Ctx(ctx).FieldsEx(dao.NoticeConfig.Columns().Id).Where(dao.NoticeConfig.Columns().Id, in.Id).Update(in)
	return
}

// DeleteNoticeConfig 删除数据
func (s *sNoticeConfig) DeleteNoticeConfig(ctx context.Context, Ids []string) (err error) {
	_, err = dao.NoticeConfig.Ctx(ctx).Where(dao.NoticeConfig.Columns().Id+" in(?)", Ids).Delete()
	if err != nil {
		return errors.New("删除失败")
	}
	return

}
