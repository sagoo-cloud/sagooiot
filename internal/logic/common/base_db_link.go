package common

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBaseDbLink struct {
}

func BaseDbLink() *sBaseDbLink {
	return &sBaseDbLink{}
}

func init() {
	service.RegisterBaseDbLink(BaseDbLink())
}

// GetList 获取数据源数据列表
func (s *sBaseDbLink) GetList(ctx context.Context, input *model.BaseDbLinkDoInput) (total int, out []*model.BaseDbLinkOut, err error) {
	m := dao.BaseDbLink.Ctx(ctx)
	if input.Host != "" {
		m = m.WhereLike(dao.BaseDbLink.Columns().Host, "%"+input.Host+"%")
	}
	if input.Name != "" {
		m = m.WhereLike(dao.BaseDbLink.Columns().Name, "%"+input.Name+"%")
	}
	if input.Types != "" {
		m = m.WhereLike(dao.BaseDbLink.Columns().Types, "%"+input.Types+"%")
	}
	if input.Port != "" {
		m = m.WhereLike(dao.BaseDbLink.Columns().Port, "%"+input.Port+"%")
	}
	if input.UserName != "" {
		m = m.WhereLike(dao.BaseDbLink.Columns().UserName, "%"+input.UserName+"%")
	}
	if input.Status != -1 {
		m = m.Where(dao.BaseDbLink.Columns().Status, input.Status)
	}
	m = m.Where(dao.BaseDbLink.Columns().IsDeleted, 0)
	//获取总数
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取数据源列表数据失败")
		return
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.DefaultPageSize
	}
	//获取数据源列表信息
	err = m.Page(input.PageNum, input.PageSize).OrderDesc(dao.BaseDbLink.Columns().CreatedAt).Scan(&out)
	if err != nil {
		err = gerror.New("获取数据源列表失败")
		return
	}
	return
}

// Add 添加数据源
func (s *sBaseDbLink) Add(ctx context.Context, input *model.AddBaseDbLinkInput) (err error) {
	var baseDbLink *entity.BaseDbLink
	//根据名称查看角色是否存在
	baseDbLink = checkBaseDbLinkName(ctx, input.Name, baseDbLink, 0)
	if baseDbLink != nil {
		return gerror.New("数据源已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	baseDbLink = new(entity.BaseDbLink)
	if err := gconv.Scan(input, &baseDbLink); err != nil {
		return err
	}
	baseDbLink.IsDeleted = 0
	baseDbLink.CreatedBy = uint(loginUserId)
	_, err = dao.BaseDbLink.Ctx(ctx).Data(baseDbLink).Insert()
	if err != nil {
		return err
	}
	return
}

// Detail 数据源详情
func (s *sBaseDbLink) Detail(ctx context.Context, baseDbLinkId int) (entity *entity.BaseDbLink, err error) {
	_ = dao.BaseDbLink.Ctx(ctx).Where(g.Map{
		dao.BaseDbLink.Columns().Id: baseDbLinkId,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("ID错误")
	}
	return
}

// Edit 修改数据源
func (s *sBaseDbLink) Edit(ctx context.Context, input *model.EditBaseDbLinkInput) (err error) {
	var baseDbLink, BaseDbLink2 *entity.BaseDbLink
	//根据ID查看数据源是否存在
	baseDbLink = checkBaseDbLinkId(ctx, input.Id, baseDbLink)
	if baseDbLink == nil {
		return gerror.New("数据源不存在")
	}
	BaseDbLink2 = checkBaseDbLinkName(ctx, input.Name, BaseDbLink2, input.Id)
	if BaseDbLink2 != nil {
		return gerror.New("相同数据源已存在,无法修改")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if err := gconv.Scan(input, &baseDbLink); err != nil {
		return err
	}
	baseDbLink.UpdatedBy = int(uint(loginUserId))
	_, err = dao.BaseDbLink.Ctx(ctx).Data(baseDbLink).
		Where(dao.BaseDbLink.Columns().Id, input.Id).Update()
	if err != nil {
		return gerror.New("修改失败")
	}
	return
}

// 检查相同数据源名称的数据是否存在
func checkBaseDbLinkName(ctx context.Context, BaseDbLinkName string, BaseDbLink *entity.BaseDbLink, tag int) *entity.BaseDbLink {
	m := dao.BaseDbLink.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.BaseDbLink.Columns().Id, tag)
	}
	_ = m.Where(g.Map{
		dao.BaseDbLink.Columns().Name:      BaseDbLinkName,
		dao.BaseDbLink.Columns().IsDeleted: 0,
	}).Scan(&BaseDbLink)
	return BaseDbLink
}

// 检查指定ID的数据是否存在
func checkBaseDbLinkId(ctx context.Context, BaseDbLinkId int, BaseDbLink *entity.BaseDbLink) *entity.BaseDbLink {
	_ = dao.BaseDbLink.Ctx(ctx).Where(g.Map{
		dao.BaseDbLink.Columns().Id:        BaseDbLinkId,
		dao.BaseDbLink.Columns().IsDeleted: 0,
	}).Scan(&BaseDbLink)
	return BaseDbLink
}

// Del 根据ID删除数据源信息
func (s *sBaseDbLink) Del(ctx context.Context, BaseDbLinkId int) (err error) {
	var BaseDbLink *entity.BaseDbLink
	_ = dao.BaseDbLink.Ctx(ctx).Where(g.Map{
		dao.BaseDbLink.Columns().Id: BaseDbLinkId,
	}).Scan(&BaseDbLink)
	if BaseDbLink == nil {
		return gerror.New("ID错误")
	}
	loginUserId := service.Context().GetUserId(ctx)
	//删除数据源信息
	_, err = dao.BaseDbLink.Ctx(ctx).
		Data(g.Map{
			dao.BaseDbLink.Columns().DeletedBy: uint(loginUserId),
			dao.BaseDbLink.Columns().IsDeleted: 1,
		}).Where(dao.BaseDbLink.Columns().Id, BaseDbLinkId).
		Delete()
	return
}
