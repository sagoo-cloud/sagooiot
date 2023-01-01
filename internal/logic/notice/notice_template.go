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

type sNoticeTemplate struct{}

func sNoticeTemplateNew() *sNoticeTemplate {
	return &sNoticeTemplate{}
}
func init() {
	service.RegisterNoticeTemplate(sNoticeTemplateNew())
}

//GetNoticeTemplateList 获取列表数据
func (s *sNoticeTemplate) GetNoticeTemplateList(ctx context.Context, in *model.GetNoticeTemplateListInput) (total, page int, list []*model.NoticeTemplateOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NoticeTemplate.Ctx(ctx)

		if in.KeyWord != "" {
			m = m.WhereLike(dao.NoticeTemplate.Columns().Title, "%"+in.KeyWord+"%")
			m = m.WhereOrLike(dao.NoticeTemplate.Columns().Content, "%"+in.KeyWord+"%")
		}
		if in.Code != "" {
			m = m.Where(dao.NoticeTemplate.Columns().Content, in.Code)
		}
		if in.ConfigId != "" {
			m = m.Where(dao.NoticeTemplate.Columns().ConfigId, in.ConfigId)
		}

		if in.SendGateway != "" {
			m = m.Where(dao.NoticeTemplate.Columns().SendGateway, in.SendGateway)
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

//GetNoticeTemplateById 获取指定ID数据
func (s *sNoticeTemplate) GetNoticeTemplateById(ctx context.Context, id string) (out *model.NoticeTemplateOutput, err error) {
	err = dao.NoticeTemplate.Ctx(ctx).Where(dao.NoticeTemplate.Columns().Id, id).Scan(&out)
	return
}

//GetNoticeTemplateByConfigId 获取指定ConfigID数据
func (s *sNoticeTemplate) GetNoticeTemplateByConfigId(ctx context.Context, configId string) (out *model.NoticeTemplateOutput, err error) {
	err = dao.NoticeTemplate.Ctx(ctx).Where(dao.NoticeTemplate.Columns().ConfigId, configId).Scan(&out)
	return
}

//AddNoticeTemplate 添加数据
func (s *sNoticeTemplate) AddNoticeTemplate(ctx context.Context, in model.NoticeTemplateAddInput) (err error) {
	_, err = dao.NoticeTemplate.Ctx(ctx).Insert(in)
	return
}

//EditNoticeTemplate 修改数据
func (s *sNoticeTemplate) EditNoticeTemplate(ctx context.Context, in model.NoticeTemplateEditInput) (err error) {
	_, err = dao.NoticeTemplate.Ctx(ctx).Where(dao.NoticeTemplate.Columns().Id, in.Id).Update(in)
	return
}

//SaveNoticeTemplate 直接更新数据
func (s *sNoticeTemplate) SaveNoticeTemplate(ctx context.Context, in model.NoticeTemplateAddInput) (err error) {
	_, err = dao.NoticeTemplate.Ctx(ctx).Where(dao.NoticeTemplate.Columns().ConfigId, in.ConfigId).Save(in)
	return
}

//DeleteNoticeTemplate 删除数据
func (s *sNoticeTemplate) DeleteNoticeTemplate(ctx context.Context, Ids []string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.NoticeTemplate.Ctx(ctx).Where(dao.NoticeTemplate.Columns().Id+" in(?)", Ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除模版数据失败")
	})
	return
}
