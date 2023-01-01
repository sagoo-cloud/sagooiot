package notice

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/liberr"
)

type sNoticeConfig struct{}

func sNoticeConfigNew() *sNoticeConfig {
	return &sNoticeConfig{}
}
func init() {
	service.RegisterNoticeConfig(sNoticeConfigNew())
}

//GetNoticeConfigList 获取列表数据
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

//GetNoticeConfigById 获取指定ID数据
func (s *sNoticeConfig) GetNoticeConfigById(ctx context.Context, id int) (out *model.NoticeConfigOutput, err error) {
	err = dao.NoticeConfig.Ctx(ctx).Where(dao.NoticeConfig.Columns().Id, id).Scan(&out)
	return
}

//AddNoticeConfig 添加数据
func (s *sNoticeConfig) AddNoticeConfig(ctx context.Context, in model.NoticeConfigAddInput) (err error) {
	_, err = dao.NoticeConfig.Ctx(ctx).Insert(in)
	return
}

//EditNoticeConfig 修改数据
func (s *sNoticeConfig) EditNoticeConfig(ctx context.Context, in model.NoticeConfigEditInput) (err error) {
	_, err = dao.NoticeConfig.Ctx(ctx).FieldsEx(dao.NoticeConfig.Columns().Id).Where(dao.NoticeConfig.Columns().Id, in.Id).Update(in)
	return
}

//DeleteNoticeConfig 删除数据
func (s *sNoticeConfig) DeleteNoticeConfig(ctx context.Context, Ids []string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.NoticeConfig.Ctx(ctx).Where(dao.NoticeConfig.Columns().Id+" in(?)", Ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除配置数据失败")
	})
	return

}
