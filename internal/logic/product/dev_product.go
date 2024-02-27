package product

import (
	"context"
	"encoding/json"
	"fmt"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/tsd/comm"
	"time"

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

func (s *sDevProduct) Detail(ctx context.Context, key string) (out *model.DetailProductOutput, err error) {
	err = dao.DevProduct.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 30,
		Name:     consts.GetDetailProductOutput + key,
		Force:    false,
	}).WithAll().Where(dao.DevProduct.Columns().Key, key).Scan(&out)
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

	return
}

func (s *sDevProduct) GetInfoById(ctx context.Context, id uint) (out *entity.DevProduct, err error) {
	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Id, id).Scan(&out)
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
		m = m.Where(c.Status, gconv.Int(in.Status))
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
	var productIds = make([]string, dLen)
	for i, v := range out.Product {
		productIds[i] = v.Key

		if v.Category != nil {
			out.Product[i].CategoryName = v.Category.Name
		}
	}

	// 获取产品的设备数量
	totals, err := service.DevDevice().TotalByProductKey(ctx, productIds)
	if err != nil {
		return
	}
	for i, v := range out.Product {
		out.Product[i].DeviceTotal = totals[v.Key]
		out.Product[i].Metadata = ""

	}

	return
}

func (s *sDevProduct) List(ctx context.Context) (list []*model.ProductOutput, err error) {
	m := dao.DevProduct.Ctx(ctx)
	err = m.WithAll().
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
		list[i].Metadata = ""
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

	tsl := &model.TSL{
		Key:  in.Key,
		Name: in.Name,
	}
	metadata, err := json.Marshal(tsl)
	if err != nil {
		return
	}

	_, err = dao.DevProduct.Ctx(ctx).Data(do.DevProduct{
		DeptId:            service.Context().GetUserDeptId(ctx),
		Key:               in.Key,
		Name:              in.Name,
		CategoryId:        in.CategoryId,
		MessageProtocol:   in.MessageProtocol,
		TransportProtocol: in.TransportProtocol,
		DeviceType:        in.DeviceType,
		Desc:              in.Desc,
		Icon:              in.Icon,
		Metadata:          metadata,
		Status:            0,
		AuthType:          in.AuthType,
		AuthUser:          in.AuthUser,
		AuthPasswd:        in.AuthPasswd,
		AccessToken:       in.AccessToken,
		CertificateId:     in.CertificateId,
		ScriptInfo:        in.ScriptInfo,
		CreatedBy:         uint(loginUserId),
		CreatedAt:         gtime.Now(),
	}).Insert()

	return
}

func (s *sDevProduct) Edit(ctx context.Context, in *model.EditProductInput) (err error) {
	devProduct, err := s.Detail(ctx, in.Key)
	if devProduct == nil {
		return gerror.New("产品不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevProduct
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdatedBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevProduct.Ctx(ctx).Data(param).Where(dao.DevProduct.Columns().Key, in.Key).Update()

	//从缓存中删除
	_, err = cache.Instance().Remove(ctx, consts.CacheGfOrmPrefix+consts.GetDetailProductOutput+devProduct.Key)

	return
}

func (s *sDevProduct) UpdateExtend(ctx context.Context, in *model.ExtendInput) (err error) {
	devProduct, err := s.Detail(ctx, in.Key)
	if devProduct == nil {
		return gerror.New("产品不存在")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevProduct
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdatedBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevProduct.Ctx(ctx).Data(param).Where(dao.DevProduct.Columns().Key, in.Key).Update()
	//从缓存中删除
	_, err = cache.Instance().Remove(ctx, consts.CacheGfOrmPrefix+consts.GetDetailProductOutput+devProduct.Key)

	return
}

func (s *sDevProduct) Del(ctx context.Context, keys []string) (err error) {
	var p []*entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).WhereIn(dao.DevProduct.Columns().Key, keys).Scan(&p)
	if err != nil {
		return
	}
	if len(p) == 0 {
		return gerror.New("产品不存在")
	}

	// 状态校验
	for _, v := range p {
		if v.Status == model.ProductStatusOn {
			return gerror.Newf("产品(%s)已发布，不能删除", v.Key)
		}
	}

	//判断产品下是否有未删除的设备
	var devices []*entity.DevDevice
	err = dao.DevDevice.Ctx(ctx).WhereIn(dao.DevDevice.Columns().Key, keys).Scan(&devices)
	if err != nil {
		return
	}
	for _, device := range devices {
		var productName string
		for _, v := range p {
			if device.ProductKey == v.Key {
				productName = v.Name
				break
			}
		}
		return gerror.Newf("产品(%s)下有未删除的设备，不能删除", productName)
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	for _, key := range keys {
		res, _ := s.Detail(ctx, key)

		rs, err := dao.DevProduct.Ctx(ctx).
			Data(do.DevProduct{
				DeletedBy: uint(loginUserId),
				DeletedAt: gtime.Now(),
			}).
			Where(dao.DevProduct.Columns().Status, model.ProductStatusOff).
			Where(dao.DevProduct.Columns().Key, key).
			Unscoped().
			Update()
		if err != nil {
			return err
		}

		num, _ := rs.RowsAffected()
		if num > 0 && res.MetadataTable == 1 {
			// 删除TD表
			err = service.TSLTable().DropStable(ctx, res.Key)
			if err != nil {
				return err
			}
		}
		//从缓存中删除
		_, err = cache.Instance().Remove(ctx, consts.CacheGfOrmPrefix+consts.GetDetailProductOutput+res.Key)
	}

	return
}

// Deploy 产品发布
func (s *sDevProduct) Deploy(ctx context.Context, productKey string) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, productKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}

	if p.Status == model.ProductStatusOn {
		return
	}
	if p.Metadata == "" {
		return gerror.New("请创建物模型属性")
	}

	var tsl *model.TSL
	err = json.Unmarshal([]byte(p.Metadata), &tsl)
	if err != nil {
		return err
	}
	if len(tsl.Properties) == 0 {
		return gerror.New("请创建物模型属性")
	}

	err = dao.DevProduct.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 是否创建TD表结构
		var isCreate bool

		if p.MetadataTable == 0 {
			isCreate = true
		}
		// 检查TD表是否存在
		if p.MetadataTable == 1 {
			stable := comm.ProductTableName(p.Key)
			b, err := service.TSLTable().CheckStable(ctx, stable)
			if err != nil {
				if err.Error() == "sql: no rows in result set" {
					isCreate = true
				} else {
					return err
				}
			}
			if !b {
				isCreate = true
			}
		}

		// 创建TD表结构
		if isCreate {
			err = service.TSLTable().CreateStable(ctx, tsl)
			if err != nil {
				return err
			}
		}

		// 更新状态
		_, err = dao.DevProduct.Ctx(ctx).
			Data(g.Map{
				dao.DevProduct.Columns().Status:        model.ProductStatusOn,
				dao.DevProduct.Columns().MetadataTable: 1,
			}).
			Where(dao.DevProduct.Columns().Key, p.Key).
			Update()

		return err
	})
	if err != nil {
		return err
	}

	//从缓存中删除
	_, err = cache.Instance().Remove(ctx, consts.CacheGfOrmPrefix+consts.GetDetailProductOutput+p.Key)
	return
}

// Undeploy 产品停用
func (s *sDevProduct) Undeploy(ctx context.Context, productKey string) (err error) {
	var p *entity.DevProduct

	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, productKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}

	if p.Status == model.ProductStatusOff {
		return
	}

	// 检查是否有启用设备
	in := model.ListDeviceInput{
		ProductKey: p.Key,
	}
	devList, err := service.DevDevice().List(ctx, gconv.String(in), "")
	if err != nil {
		return
	}
	if len(devList) > 0 {
		return gerror.New("该产品有启用设备，请先停用设备")
	}

	_, err = dao.DevProduct.Ctx(ctx).
		Data(g.Map{dao.DevProduct.Columns().Status: model.ProductStatusOff}).
		Where(dao.DevProduct.Columns().Key, p.Key).
		Update()
	if err != nil {
		return err
	}

	//从缓存中删除
	_, err = cache.Instance().Remove(ctx, consts.CacheGfOrmPrefix+consts.GetDetailProductOutput+p.Key)
	return
}

// ListForSub 子设备类型产品
func (s *sDevProduct) ListForSub(ctx context.Context) (list []*model.ProductOutput, err error) {
	m := dao.DevProduct.Ctx(ctx)
	err = m.WithAll().
		Where(dao.DevProduct.Columns().DeviceType, model.DeviceTypeSub).
		OrderDesc(dao.DevProduct.Columns().Id).
		Scan(&list)
	return
}

// UpdateScriptInfo 脚本更新
func (s *sDevProduct) UpdateScriptInfo(ctx context.Context, in *model.ScriptInfoInput) (err error) {
	var devProduct *entity.DevProduct
	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, in.Key).Scan(&devProduct)
	if err != nil {
		return
	}
	if devProduct == nil {
		return gerror.New("产品不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	devProduct.ScriptInfo = in.ScriptInfo
	devProduct.UpdatedBy = uint(loginUserId)
	devProduct.UpdatedAt = gtime.Now()
	_, err = dao.DevProduct.Ctx(ctx).Data(devProduct).Where(dao.DevProduct.Columns().Key, in.Key).Update()
	//从缓存中删除
	_, err = cache.Instance().Remove(ctx, consts.CacheGfOrmPrefix+consts.GetDetailProductOutput+devProduct.Key)
	return
}

// ConnectIntro 获取设备接入信息
func (s *sDevProduct) ConnectIntro(ctx context.Context, productKey string) (out *model.DeviceConnectIntroOutput, err error) {
	var product *entity.DevProduct
	if err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, productKey).Scan(&product); err != nil || product == nil {
		return
	}
	if product.MessageProtocol == "" {
		return
	}

	// 默认服务
	if product.MessageProtocol == "SagooMqtt" {
		mqtt, err := g.Cfg().Get(ctx, "mqtt")
		if err != nil {
			return nil, err
		}
		mm := mqtt.MapStrVar()
		au := mm["auth"].MapStrStr()

		out = &model.DeviceConnectIntroOutput{
			Name:        "默认服务",
			Protocol:    "SagooMqtt",
			Description: "",
			Link:        fmt.Sprintf("mqtt://%s", mm["addr"].String()),
			AuthType:    1,
			AuthUser:    au["userName"],
			AuthPasswd:  au["userPassWorld"],
		}
		return out, nil
	}

	// 网络服务
	var networkS *entity.NetworkServer
	// TODO status=1 启用
	if err = dao.NetworkServer.Ctx(ctx).WhereLike(dao.NetworkServer.Columns().Protocol, "%"+product.MessageProtocol+"%").Scan(&networkS); err != nil || networkS == nil {
		return
	}

	out = &model.DeviceConnectIntroOutput{
		Name:        networkS.Name,
		Protocol:    product.MessageProtocol,
		Description: networkS.Remark,
		Link:        fmt.Sprintf("%s://0.0.0.0:%s", networkS.Types, networkS.Addr),
	}

	out.AuthType = networkS.AuthType
	out.AuthUser = networkS.AuthUser
	out.AuthPasswd = networkS.AuthPasswd
	out.AccessToken = networkS.AccessToken
	out.CertificateId = networkS.CertificateId

	if product.AuthType > 0 {
		out.AuthType = product.AuthType
		out.AuthUser = product.AuthUser
		out.AuthPasswd = product.AuthPasswd
		out.AccessToken = product.AccessToken
		out.CertificateId = product.CertificateId
	}

	if out.CertificateId > 0 {
		if cert, err := service.SysCertificate().GetInfoById(ctx, out.CertificateId); err == nil && cert != nil {
			out.CertificateName = cert.Name
		}
	}

	return
}
