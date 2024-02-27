package product

import (
	"context"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDevCategory struct{}

func init() {
	service.RegisterDevCategory(categoryNew())
}

func categoryNew() *sDevCategory {
	return &sDevCategory{}
}

func (s *sDevCategory) Detail(ctx context.Context, id uint) (out *model.ProductCategoryOutput, err error) {
	var p *entity.DevProductCategory

	err = dao.DevProductCategory.Ctx(ctx).Where(dao.DevProductCategory.Columns().Id, id).Scan(&p)
	if err != nil || p == nil {
		return
	}

	out = &model.ProductCategoryOutput{
		DevProductCategory: p,
	}

	return
}

func (s *sDevCategory) GetNameByIds(ctx context.Context, categoryIds []uint) (names map[uint]string, err error) {
	var categorys []*entity.DevProductCategory
	c := dao.DevProductCategory.Columns()
	err = dao.DevProductCategory.Ctx(ctx).
		Fields(c.Id, c.Name).
		WhereIn(c.Id, categoryIds).
		OrderAsc(c.Sort).
		OrderDesc(c.Id).
		Scan(&categorys)
	if err != nil || len(categorys) == 0 {
		return
	}

	names = make(map[uint]string, len(categorys))
	for _, v := range categorys {
		names[v.Id] = v.Name
	}

	for _, id := range categoryIds {
		if _, ok := names[id]; !ok {
			names[id] = ""
		}
	}

	return
}

// ListForPage 产品分类列表
func (s *sDevCategory) ListForPage(ctx context.Context, page, limit int, name string) (out []*model.ProductCategoryTreeOutput, total int, err error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = consts.PageSize
	}

	m := dao.DevProductCategory.Ctx(ctx).
		OrderAsc(dao.DevProductCategory.Columns().Sort).
		OrderDesc(dao.DevProductCategory.Columns().Id)

	if name != "" {
		m = m.WhereLike(dao.DevProductCategory.Columns().Name, "%"+name+"%")
	}

	var categorys []*entity.DevProductCategory
	err = m.Scan(&categorys)
	if err != nil {
		return
	}
	out = Tree(categorys, 0)
	total = len(out)
	start := (page - 1) * limit
	end := page * limit
	out = out[start:end]

	return
}

func (s *sDevCategory) List(ctx context.Context, name string) (out []*model.ProductCategoryTreeOutput, err error) {
	var categorys []*entity.DevProductCategory
	m := dao.DevProductCategory.Ctx(ctx).OrderAsc(dao.DevProductCategory.Columns().Sort).OrderDesc(dao.DevProductCategory.Columns().Id)
	if name != "" {
		m = m.WhereLike(dao.DevProductCategory.Columns().Name, "%"+name+"%")
	}

	err = m.Scan(&categorys)
	if err != nil || len(categorys) == 0 {
		return
	}

	if name != "" {
		if err = gconv.Scan(categorys, &out); err != nil {
			return
		}
	} else {
		out = Tree(categorys, 0)
	}

	return
}

func Tree(all []*entity.DevProductCategory, pid uint) (rs []*model.ProductCategoryTreeOutput) {
	for _, v := range all {
		if v.ParentId == pid {
			var treeRes *model.ProductCategoryTreeOutput
			if err := gconv.Scan(v, &treeRes); err != nil {
				return
			}
			treeRes.Children = Tree(all, v.Id)

			rs = append(rs, treeRes)
		}
	}
	return
}

func (s *sDevCategory) Add(ctx context.Context, in *model.AddProductCategoryInput) (err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	_, err = dao.DevProductCategory.Ctx(ctx).Data(do.DevProductCategory{
		DeptId:    service.Context().GetUserDeptId(ctx),
		ParentId:  in.ParentId,
		Key:       in.Key,
		Name:      in.Name,
		Sort:      in.Sort,
		Desc:      in.Desc,
		CreatedBy: uint(loginUserId),
	}).Insert()
	if err != nil {
		return
	}

	return
}

func (s *sDevCategory) Edit(ctx context.Context, in *model.EditProductCategoryInput) (err error) {
	var category *entity.DevProductCategory

	err = dao.DevProductCategory.Ctx(ctx).Where(dao.DevProductCategory.Columns().Id, in.Id).Scan(&category)
	if err != nil {
		return
	}
	if category == nil {
		return gerror.New("分类不存在")
	}

	if category.Id == in.ParentId {
		return gerror.New("上级分类节点选择错误,请重新选择!")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	_, err = dao.DevProductCategory.Ctx(ctx).Data(do.DevProductCategory{
		ParentId:  in.ParentId,
		Key:       in.Key,
		Name:      in.Name,
		Sort:      in.Sort,
		Desc:      in.Desc,
		UpdatedBy: uint(loginUserId),
	}).Where(dao.DevProductCategory.Columns().Id, in.Id).Update()

	return
}

func (s *sDevCategory) Del(ctx context.Context, id uint) (err error) {
	var categorys []*entity.DevProductCategory

	err = dao.DevProductCategory.Ctx(ctx).
		Where(dao.DevProductCategory.Columns().Id, id).
		WhereOr(dao.DevProductCategory.Columns().ParentId, id).
		Scan(&categorys)
	if err != nil {
		return
	}
	if len(categorys) == 0 {
		return gerror.New("分类不存在")
	}
	if len(categorys) > 1 {
		return gerror.New("请先删除子分类")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	_, err = dao.DevProductCategory.Ctx(ctx).
		Data(do.DevProductCategory{
			DeletedBy: uint(loginUserId),
			DeletedAt: gtime.Now(),
		}).
		Where(dao.DevProductCategory.Columns().Id, id).
		Unscoped().
		Update()

	return
}
