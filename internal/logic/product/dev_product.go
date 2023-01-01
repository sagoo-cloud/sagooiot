package product

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDevProduct struct{}

func init() {
	service.RegisterDevProduct(productNew())
}

func productNew() *sDevProduct {
	return &sDevProduct{}
}

func (s *sDevProduct) Get(ctx context.Context, key string) (out *model.DetailProductOutput, err error) {
	err = dao.DevProduct.Ctx(ctx).WithAll().Where(dao.DevProduct.Columns().Key, key).Scan(&out)
	if err != nil || out == nil {
		return
	}

	if out.Metadata != "" {
		err = json.Unmarshal([]byte(out.Metadata), &out.TSL)
		if err != nil {
			return
		}
	}

	if out.Category != nil {
		out.CategoryName = out.Category.Name
	}

	// 获取产品的设备数量
	totals, err := service.DevDevice().TotalByProductId(ctx, []uint{out.Id})
	if err != nil {
		return
	}
	out.DeviceTotal = totals[out.Id]

	return
}

func (s *sDevProduct) Detail(ctx context.Context, id uint) (out *model.DetailProductOutput, err error) {
	err = dao.DevProduct.Ctx(ctx).WithAll().Where(dao.DevProduct.Columns().Id, id).Scan(&out)
	if err != nil || out == nil {
		return
	}

	if out.Metadata != "" {
		err = json.Unmarshal([]byte(out.Metadata), &out.TSL)
		if err != nil {
			return
		}
	}

	if out.Category != nil {
		out.CategoryName = out.Category.Name
	}

	// 获取产品的设备数量
	totals, err := service.DevDevice().TotalByProductId(ctx, []uint{out.Id})
	if err != nil {
		return
	}
	out.DeviceTotal = totals[out.Id]

	return
}

func (s *sDevProduct) GetNameByIds(ctx context.Context, productIds []uint) (names map[uint]string, err error) {
	var products []*entity.DevProduct
	c := dao.DevProduct.Columns()
	err = dao.DevProduct.Ctx(ctx).
		Fields(c.Id, c.Name).
		WhereIn(c.Id, productIds).
		Scan(&products)
	if err != nil || len(products) == 0 {
		return
	}

	names = make(map[uint]string, len(products))
	for _, v := range products {
		names[v.Id] = v.Name
	}

	for _, id := range productIds {
		if _, ok := names[id]; !ok {
			names[id] = ""
		}
	}

	return
}

func (s *sDevProduct) ListForPage(ctx context.Context, in *model.ListForPageInput) (out *model.ListForPageOutput, err error) {
	out = new(model.ListForPageOutput)
	c := dao.DevProduct.Columns()
	m := dao.DevProduct.Ctx(ctx).OrderDesc(c.Id)

	if in.Status != "" {
		m = m.Where(c.Status+" = ", gconv.Int(in.Status))
	}
	if in.CategoryId > 0 {
		m = m.Where(c.CategoryId, in.CategoryId)
	}
	if in.Name != "" {
		m = m.WhereLike(c.Name, "%"+in.Name+"%")
	}
	if len(in.MessageProtocols) > 0 {
		m = m.WhereIn(c.MessageProtocol, in.MessageProtocols)
	}
	if len(in.DeviceTypes) > 0 {
		m = m.WhereIn(c.DeviceType, in.DeviceTypes)
	}
	if len(in.DateRange) > 0 {
		m = m.WhereBetween(c.CreatedAt, in.DateRange[0], in.DateRange[1])
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.WithAll().Page(in.PageNum, in.PageSize).Scan(&out.Product)
	if err != nil || len(out.Product) == 0 {
		return
	}

	dLen := len(out.Product)
	var productIds = make([]uint, dLen)
	for i, v := range out.Product {
		productIds[i] = v.Id

		if v.Category != nil {
			out.Product[i].CategoryName = v.Category.Name
		}
	}

	// 获取产品的设备数量
	totals, err := service.DevDevice().TotalByProductId(ctx, productIds)
	if err != nil {
		return
	}
	for i, v := range out.Product {
		out.Product[i].DeviceTotal = totals[v.Id]
	}

	return
}

func (s *sDevProduct) List(ctx context.Context) (list []*model.ProductOutput, err error) {
	err = dao.DevProduct.Ctx(ctx).WithAll().
		Where(dao.DevProduct.Columns().Status, model.ProductStatusOn).
		OrderDesc(dao.DevProduct.Columns().Id).
		Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	dLen := len(list)
	var productIds = make([]uint, dLen)
	for i, v := range list {
		productIds[i] = v.Id

		if v.Category != nil {
			list[i].CategoryName = v.Category.Name
		}
	}

	// 获取产品的设备数量
	totals, err := service.DevDevice().TotalByProductId(ctx, productIds)
	if err != nil {
		return
	}
	for i, v := range list {
		list[i].DeviceTotal = totals[v.Id]
	}

	return
}

func (s *sDevProduct) Add(ctx context.Context, in *model.AddProductInput) (err error) {
	id, _ := dao.DevProduct.Ctx(ctx).
		Fields(dao.DevProduct.Columns().Id).
		Where(dao.DevProduct.Columns().Key, in.Key).
		Value()
	if id.Uint() > 0 {
		return gerror.New("产品标识重复")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevProduct
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)
	param.Status = 0

	tsl := &model.TSL{
		Key:  in.Key,
		Name: in.Name,
	}
	param.Metadata, err = json.Marshal(tsl)
	if err != nil {
		return
	}

	_, err = dao.DevProduct.Ctx(ctx).Data(param).Insert()

	return
}

func (s *sDevProduct) Edit(ctx context.Context, in *model.EditProductInput) (err error) {
	total, err := dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, in.Id).Count()
	if err != nil {
		return
	}
	if total == 0 {
		return gerror.New("产品不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevProduct
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevProduct.Ctx(ctx).Data(param).Where(dao.DevProduct.Columns().Id, in.Id).Update()

	return
}

func (s *sDevProduct) Del(ctx context.Context, ids []uint) (err error) {
	var p []*entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).WhereIn(dao.DevProduct.Columns().Id, ids).Scan(&p)
	if err != nil {
		return
	}
	if len(p) == 0 {
		return gerror.New("产品不存在")
	}
	if len(p) == 1 && p[0].Status == model.ProductStatusOn {
		return gerror.New("产品已发布，不能删除")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	for _, id := range ids {
		res, _ := s.Detail(ctx, id)

		rs, err := dao.DevProduct.Ctx(ctx).
			Data(do.DevProduct{
				DeletedBy: uint(loginUserId),
				DeletedAt: gtime.Now(),
			}).
			Where(dao.DevProduct.Columns().Status, model.ProductStatusOff).
			Where(dao.DevProduct.Columns().Id, id).
			Unscoped().
			Update()
		if err != nil {
			return err
		}

		num, _ := rs.RowsAffected()
		if num > 0 {
			// 删除TD表
			err = service.TSLTable().DropStable(ctx, res.Key)
			if err != nil {
				return err
			}
		}
	}

	return
}

func (s *sDevProduct) Deploy(ctx context.Context, id uint) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}
	if p.Status == model.ProductStatusOn {
		return gerror.New("产品已发布")
	}

	err = dao.DevProduct.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DevProduct.Ctx(ctx).
			Data(g.Map{dao.DevProduct.Columns().Status: model.ProductStatusOn}).
			Where(dao.DevProduct.Columns().Id, id).
			Update()
		if err != nil {
			return err
		}

		// 建立TD表
		if p.Metadata != "" && p.MetadataTable == 0 {
			var tsl *model.TSL
			err = json.Unmarshal([]byte(p.Metadata), &tsl)
			if err != nil {
				return err
			}

			err = service.TSLTable().CreateStable(ctx, tsl)
			if err != nil {
				return err
			}

			_, err = dao.DevProduct.Ctx(ctx).
				Data(g.Map{dao.DevProduct.Columns().MetadataTable: 1}).
				Where(dao.DevProduct.Columns().Id, id).
				Update()
			if err != nil {
				return err
			}
		}
		return nil
	})

	return
}

func (s *sDevProduct) Undeploy(ctx context.Context, id uint) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}
	if p.Status == model.ProductStatusOff {
		return gerror.New("产品已停用")
	}

	_, err = dao.DevProduct.Ctx(ctx).
		Data(g.Map{dao.DevProduct.Columns().Status: model.ProductStatusOff}).
		Where(dao.DevProduct.Columns().Id, id).
		Update()

	return
}
