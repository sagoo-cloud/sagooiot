package datahub

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sDataTemplateBusi struct{}

func init() {
	service.RegisterDataTemplateBusi(dataTemplateBusiNew())
}

func dataTemplateBusiNew() *sDataTemplateBusi {
	return &sDataTemplateBusi{}
}

func (s *sDataTemplateBusi) Add(ctx context.Context, in *model.DataTemplateBusiAddInput) (err error) {
	if len(in.BusiTypes) == 0 {
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	c := dao.DataTemplateBusi.Columns()
	for _, v := range in.BusiTypes {
		id, _ := dao.DataTemplateBusi.Ctx(ctx).
			Where(c.BusiTypes, v).
			Value("id")
		if id.Int() > 0 {
			continue
		}

		param := do.DataTemplateBusi{
			DataTemplateId: in.DataTemplateId,
			BusiTypes:      v,
			CreatedBy:      uint(loginUserId),
		}
		_, err = dao.DataTemplateBusi.Ctx(ctx).Data(param).Insert()
	}

	return
}

func (s *sDataTemplateBusi) GetInfos(ctx context.Context, busiTypes int) (data *entity.DataTemplateBusi, err error) {
	err = dao.DataTemplateBusi.Ctx(ctx).Where(g.Map{
		dao.DataTemplateBusi.Columns().BusiTypes: busiTypes,
		dao.DataTemplateBusi.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

func (s *sDataTemplateBusi) GetInfo(ctx context.Context, busiTypes int) (data *entity.DataTemplateBusi, err error) {
	err = dao.DataTemplateBusi.Ctx(ctx).Where(g.Map{
		dao.DataTemplateBusi.Columns().BusiTypes: busiTypes,
		dao.DataTemplateBusi.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

func (s *sDataTemplateBusi) GetTable(ctx context.Context, busiTypes int) (table string, err error) {
	busi, err := s.GetInfos(ctx, busiTypes)
	if err != nil {
		return
	}
	if busi == nil {
		err = gerror.New("未绑定数据模型")
		return
	}
	table = fmt.Sprintf("data_template_%d", busi.DataTemplateId)
	return
}
