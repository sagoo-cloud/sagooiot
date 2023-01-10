package product

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDevDevice struct{}

func init() {
	service.RegisterDevDevice(deviceNew())
}

func deviceNew() *sDevDevice {
	return &sDevDevice{}
}

func (s *sDevDevice) Get(ctx context.Context, key string) (out *model.DeviceOutput, err error) {
	err = dao.DevDevice.Ctx(ctx).WithAll().Where(dao.DevDevice.Columns().Key, key).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		err = gerror.New("设备不存在")
		return
	}
	if out.Product != nil {
		out.ProductName = out.Product.Name
		if out.Product.Metadata != "" {
			err = json.Unmarshal([]byte(out.Product.Metadata), &out.TSL)
		}
	}
	return
}

func (s *sDevDevice) Detail(ctx context.Context, id uint) (out *model.DeviceOutput, err error) {
	err = dao.DevDevice.Ctx(ctx).WithAll().Where(dao.DevDevice.Columns().Id, id).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		err = gerror.New("设备不存在")
		return
	}
	if out.Product != nil {
		out.ProductName = out.Product.Name
		if out.Product.Metadata != "" {
			err = json.Unmarshal([]byte(out.Product.Metadata), &out.TSL)
		}
	}
	return
}

func (s *sDevDevice) ListForPage(ctx context.Context, in *model.ListDeviceForPageInput) (out *model.ListDeviceForPageOutput, err error) {
	out = new(model.ListDeviceForPageOutput)
	c := dao.DevDevice.Columns()
	m := dao.DevDevice.Ctx(ctx).OrderDesc(c.Id)

	if in.Status != "" {
		m = m.Where(c.Status+" = ", gconv.Int(in.Status))
	}
	if in.Key != "" {
		m = m.Where(c.Key, in.Key)
	}
	if in.ProductId > 0 {
		m = m.Where(c.ProductId, in.ProductId)
	}
	if in.TunnelId > 0 {
		m = m.Where(c.TunnelId, in.TunnelId)
	}
	if in.Name != "" {
		m = m.WhereLike(c.Name, "%"+in.Name+"%")
	}
	if len(in.DateRange) > 0 {
		m = m.WhereBetween(c.CreatedAt, in.DateRange[0], in.DateRange[1])
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.WithAll().Page(in.PageNum, in.PageSize).Scan(&out.Device)
	if err != nil {
		return
	}

	for i, v := range out.Device {
		if v.Product != nil {
			out.Device[i].ProductName = v.Product.Name
		}
	}

	return
}

// List 已发布产品的设备列表
func (s *sDevDevice) List(ctx context.Context, in *model.ListDeviceInput) (list []*model.DeviceOutput, err error) {
	m := dao.DevDevice.Ctx(ctx).
		Where(dao.DevDevice.Columns().Status+" > ?", model.DeviceStatusNoEnable)
	if in.ProductId > 0 {
		m = m.Where(dao.DevDevice.Columns().ProductId, in.ProductId)
	}
	err = m.WhereIn(dao.DevDevice.Columns().ProductId,
		dao.DevProduct.Ctx(ctx).
			Fields(dao.DevProduct.Columns().Id).
			Where(dao.DevProduct.Columns().Status, model.ProductStatusOn)).
		WithAll().
		OrderDesc(dao.DevDevice.Columns().Id).
		Scan(&list)
	if err != nil {
		return
	}

	for i, v := range list {
		if v.Product != nil {
			list[i].ProductName = v.Product.Name
		}
	}

	return
}

func (s *sDevDevice) Add(ctx context.Context, in *model.AddDeviceInput) (deviceId uint, err error) {
	id, _ := dao.DevDevice.Ctx(ctx).
		Fields(dao.DevDevice.Columns().Id).
		Where(dao.DevDevice.Columns().Key, in.Key).
		Value()
	if id.Uint() > 0 {
		err = gerror.New("设备标识重复")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevDevice
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)
	param.Status = 0

	rs, err := dao.DevDevice.Ctx(ctx).Data(param).Insert()
	if err != nil {
		return
	}
	newId, err := rs.LastInsertId()
	deviceId = uint(newId)

	return
}

func (s *sDevDevice) Edit(ctx context.Context, in *model.EditDeviceInput) (err error) {
	total, err := dao.DevDevice.Ctx(ctx).Where(dao.DevDevice.Columns().Id, in.Id).Count()
	if err != nil {
		return
	}
	if total == 0 {
		return gerror.New("设备不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevDevice
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevDevice.Ctx(ctx).Data(param).Where(dao.DevDevice.Columns().Id, in.Id).Update()
	return
}

func (s *sDevDevice) Del(ctx context.Context, ids []uint) (err error) {
	var p []*entity.DevDevice
	err = dao.DevDevice.Ctx(ctx).WhereIn(dao.DevDevice.Columns().Id, ids).Scan(&p)
	if err != nil {
		return
	}
	if len(p) == 0 {
		return gerror.New("设备不存在")
	}
	if len(p) == 1 && p[0].Status > model.DeviceStatusNoEnable {
		return gerror.New("设备已启用，不能删除")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	for _, id := range ids {
		res, _ := s.Detail(ctx, id)

		rs, err := dao.DevDevice.Ctx(ctx).
			Data(do.DevDevice{
				DeletedBy: uint(loginUserId),
				DeletedAt: gtime.Now(),
			}).
			Where(dao.DevDevice.Columns().Status, model.DeviceStatusNoEnable).
			Where(dao.DevDevice.Columns().Id, id).
			Unscoped().
			Update()
		if err != nil {
			return err
		}

		num, _ := rs.RowsAffected()
		if num > 0 && res.MetadataTable == 1 {
			// 删除TD子表
			err = service.TSLTable().DropTable(ctx, res.Key)
			if err != nil {
				return err
			}
		}
	}

	return
}

// Deploy 设备启用
func (s *sDevDevice) Deploy(ctx context.Context, id uint) (err error) {
	out, err := s.Detail(ctx, id)
	if err != nil {
		return
	}
	if out.Status > model.DeviceStatusNoEnable {
		return gerror.New("设备已启用")
	}

	pd, err := service.DevProduct().Detail(ctx, out.ProductId)
	if err != nil {
		return err
	}
	if pd == nil {
		return gerror.New("产品不存在")
	}

	err = dao.DevDevice.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DevDevice.Ctx(ctx).
			Data(g.Map{dao.DevDevice.Columns().Status: model.DeviceStatusOff}).
			Where(dao.DevDevice.Columns().Id, id).
			Update()
		if err != nil {
			return err
		}

		// 建立TD子表
		if out.MetadataTable == 0 {
			err = service.TSLTable().CreateTable(ctx, out.Product.Key, out.Key)
			if err != nil {
				return err
			}

			_, err = dao.DevDevice.Ctx(ctx).
				Data(g.Map{dao.DevDevice.Columns().MetadataTable: 1}).
				Where(dao.DevDevice.Columns().Id, id).
				Update()
		}
		return err
	})

	return
}

// Undeploy 设备禁用
func (s *sDevDevice) Undeploy(ctx context.Context, id uint) (err error) {
	var p *entity.DevDevice

	err = dao.DevDevice.Ctx(ctx).Where(dao.DevDevice.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("设备不存在")
	}
	if p.Status == model.DeviceStatusNoEnable {
		return gerror.New("设备已禁用")
	}

	_, err = dao.DevDevice.Ctx(ctx).
		Data(g.Map{dao.DevDevice.Columns().Status: model.DeviceStatusNoEnable}).
		Where(dao.DevDevice.Columns().Id, id).
		Update()

	return
}

// Online 设备上线
func (s *sDevDevice) Online(ctx context.Context, key string) (err error) {
	out, err := s.Get(ctx, key)
	if err != nil {
		return
	}
	if out.Status == model.DeviceStatusOn {
		return gerror.New("设备已上线")
	}

	err = dao.DevDevice.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DevDevice.Ctx(ctx).
			Data(g.Map{dao.DevDevice.Columns().Status: model.DeviceStatusOn}).
			Where(dao.DevDevice.Columns().Id, out.Id).
			Update()
		if err != nil {
			return err
		}

		// 建立TD子表
		if out.MetadataTable == 0 {
			err = service.TSLTable().CreateTable(ctx, out.Product.Key, out.Key)
			if err != nil {
				return err
			}

			_, err = dao.DevDevice.Ctx(ctx).
				Data(g.Map{dao.DevDevice.Columns().MetadataTable: 1}).
				Where(dao.DevDevice.Columns().Id, out.Id).
				Update()
		}
		return err
	})

	return
}

// Offline 设备下线
func (s *sDevDevice) Offline(ctx context.Context, key string) (err error) {
	out, err := s.Get(ctx, key)
	if err != nil {
		return
	}
	if out.Status == model.DeviceStatusOff {
		return gerror.New("设备已离线")
	}

	_, err = dao.DevDevice.Ctx(ctx).
		Data(g.Map{dao.DevDevice.Columns().Status: model.DeviceStatusOff}).
		Where(dao.DevDevice.Columns().Id, out.Id).
		Update()

	return
}

// 统计产品下的设备数量
func (s *sDevDevice) TotalByProductId(ctx context.Context, productIds []uint) (totals map[uint]int, err error) {
	r, err := dao.DevDevice.Ctx(ctx).Fields(dao.DevDevice.Columns().ProductId+", count(*) as total").
		WhereIn(dao.DevDevice.Columns().ProductId, productIds).
		Group(dao.DevDevice.Columns().ProductId).
		All()
	if err != nil || r.Len() == 0 {
		return
	}

	totals = make(map[uint]int, r.Len())
	for _, v := range r {
		t := gconv.Int(v["total"])
		id := gconv.Uint(v[dao.DevDevice.Columns().ProductId])
		totals[id] = t
	}

	for _, id := range productIds {
		if _, ok := totals[id]; !ok {
			totals[id] = 0
		}
	}

	return
}

// 统计设备数量
func (s *sDevDevice) Total(ctx context.Context) (data model.DeviceTotalOutput, err error) {
	key := "device:total"
	tag := "device"
	value := common.Cache().GetOrSetFunc(ctx, key, func(ctx context.Context) (value interface{}, err error) {
		var rs model.DeviceTotalOutput
		// 设备总量
		rs.DeviceTotal, err = dao.DevDevice.Ctx(ctx).Count()
		if err != nil {
			return
		}

		// 离线设备数量
		rs.DeviceOffline, err = dao.DevDevice.Ctx(ctx).Where(dao.DevDevice.Columns().Status, model.DeviceStatusOff).Count()
		if err != nil {
			return
		}

		// 产品总量
		rs.ProductTotal, err = dao.DevProduct.Ctx(ctx).Count()
		if err != nil {
			return
		}

		// 产品新增数量
		rs.ProductAdded, err = dao.DevProduct.Ctx(ctx).
			Where(dao.DevProduct.Columns().CreatedAt+">=?", gtime.Now().Format("Y-m-d")).
			Count()
		if err != nil {
			return
		}

		// 设备消息总量 TDengine
		sql := "select count(*) as num from device_log"
		data, err := service.TdEngine().GetOne(ctx, sql)
		if err != nil {
			return
		}
		rs.MsgTotal = data["num"].Int()

		// 设备消息新增数量 TDengine
		sql = "select count(*) as num from device_log where ts >= '?'"
		data, err = service.TdEngine().GetOne(ctx, sql, gtime.Now().Format("Y-m-d"))
		if err != nil {
			return
		}
		rs.MsgAdded = data["num"].Int()

		// 设备报警总量
		rs.AlarmTotal, err = dao.AlarmLog.Ctx(ctx).Count()
		if err != nil {
			return
		}

		// 设备报警增量
		rs.AlarmAdded, err = dao.AlarmLog.Ctx(ctx).
			Where(dao.AlarmLog.Columns().CreatedAt+">=?", gtime.Now().Format("Y-m-d")).
			Count()
		if err != nil {
			return
		}

		value = rs
		return
	}, 7200*time.Second, tag)

	err = gconv.Struct(value, &data)
	return
}

// 统计设备月度数量
func (s *sDevDevice) TotalForMonths(ctx context.Context) (data map[int]int, err error) {
	key := "device:totalForMonths"
	tag := "device"
	value := common.Cache().GetOrSetFunc(ctx, key, func(ctx context.Context) (value interface{}, err error) {
		rs := make(map[int]int, 12)
		for i := 0; i < 12; i++ {
			rs[i+1] = 0
		}

		// 设备消息总量 TDengine
		sql := "select substr(to_iso8601(ts), 1, 7) as ym, count(*) as num from device_log where ts >= '?' group by substr(to_iso8601(ts), 1, 7)"
		list, err := service.TdEngine().GetAll(ctx, sql, gtime.Now().Format("Y-01-01 00:00:00"))
		if err != nil {
			return
		}

		for _, v := range list {
			m := gstr.SubStr(v["ym"].String(), 5)
			rs[gconv.Int(m)] = v["num"].Int()
		}

		value = rs
		return
	}, 7200*time.Second, tag)

	err = gconv.Scan(value, &data)
	return
}

// 统计设备月度告警数量
func (s *sDevDevice) AlarmTotalForMonths(ctx context.Context) (data map[int]int, err error) {
	key := "device:alarmTotalForMonths"
	tag := "device"
	value := common.Cache().GetOrSetFunc(ctx, key, func(ctx context.Context) (value interface{}, err error) {
		rs := make(map[int]int, 12)
		for i := 0; i < 12; i++ {
			rs[i+1] = 0
		}

		at := dao.AlarmLog.Columns().CreatedAt
		list, err := dao.AlarmLog.Ctx(ctx).
			Fields("date_format("+at+", '%Y%m') as ym, count(*) as num").
			Where(at+">=?", gtime.Now().Format("Y-01-01 00:00:00")).
			Group("date_format(" + at + ", '%Y%m')").
			All()
		if err != nil {
			return
		}

		for _, v := range list {
			m := gstr.SubStr(v["ym"].String(), 4)
			rs[gconv.Int(m)] = v["num"].Int()
		}

		value = rs
		return
	}, 7200*time.Second, tag)

	err = gconv.Scan(value, &data)
	return
}

// RunStatus 运行状态
func (s *sDevDevice) RunStatus(ctx context.Context, id uint) (out *model.DeviceRunStatusOutput, err error) {
	p, err := s.Detail(ctx, id)
	if err != nil {
		return
	}

	out = new(model.DeviceRunStatusOutput)
	out.Status = p.Status
	// out.LastOnlineTime = p.LastOnlineTime

	if p.Status == model.DeviceStatusNoEnable {
		return
	}

	// 属性值获取
	sql := "select * from ? order by ts desc limit 1"
	rs, err := service.TdEngine().GetOne(ctx, sql, p.Key)
	if err != nil {
		return
	}
	out.LastOnlineTime = rs["ts"].GTime()

	var properties []model.DevicePropertiy
	for _, v := range p.TSL.Properties {
		// 获取当天属性值列表
		var ls gdb.Result
		if _, ok := rs[v.Key]; ok {
			sql := "select ? from ? where ts >= '?' order by ts desc"
			ls, _ = service.TdEngine().GetAll(ctx, sql, v.Key, p.Key, gtime.Now().Format("Y-m-d"))
		}

		unit := ""
		if v.ValueType.Unit != nil {
			unit = *v.ValueType.Unit
		}
		pro := model.DevicePropertiy{
			Key:   v.Key,
			Name:  v.Name,
			Type:  v.ValueType.Type,
			Unit:  unit,
			Value: rs[v.Key],
			List:  ls.Array(v.Key),
		}
		properties = append(properties, pro)
	}
	out.Properties = properties

	return
}

// GetProperty 获取指定属性值
func (s *sDevDevice) GetProperty(ctx context.Context, in *model.DeviceGetPropertyInput) (out *model.DevicePropertiy, err error) {
	p, err := s.Detail(ctx, in.Id)
	if err != nil {
		return
	}
	if p.Status == model.DeviceStatusNoEnable {
		err = gerror.New("设备未启用")
		return
	}

	// 属性值获取
	sql := "select ? from ? order by ts desc limit 1"
	rs, err := service.TdEngine().GetOne(ctx, sql, in.PropertyKey, p.Key)
	if err != nil {
		return
	}

	var name string
	var valueType string
	for _, v := range p.TSL.Properties {
		if v.Key == in.PropertyKey {
			name = v.Name
			valueType = v.ValueType.Type
			break
		}
	}

	out = new(model.DevicePropertiy)
	out.Key = in.PropertyKey
	out.Name = name
	out.Type = valueType
	out.Value = rs[in.PropertyKey]

	// 获取当天属性值列表
	sql = "select ? from ? where ts >= '?' order by ts desc"
	ls, _ := service.TdEngine().GetAll(ctx, sql, in.PropertyKey, p.Key, gtime.Now().Format("Y-m-d"))
	out.List = ls.Array(in.PropertyKey)

	return
}

// GetPropertyList 设备属性详情列表
func (s *sDevDevice) GetPropertyList(ctx context.Context, in *model.DeviceGetPropertyListInput) (out *model.DeviceGetPropertyListOutput, err error) {
	p, err := s.Detail(ctx, in.Id)
	if err != nil {
		return
	}
	if p.Status == model.DeviceStatusNoEnable {
		err = gerror.New("设备未启用")
		return
	}

	out = new(model.DeviceGetPropertyListOutput)

	// TDengine
	sql := "select count(*) as num from ? where ts >= '?'"
	rs, err := service.TdEngine().GetOne(ctx, sql, p.Key, gtime.Now().Format("Y-m-d"))
	if err != nil {
		return
	}
	out.Total = rs["num"].Int()
	out.CurrentPage = in.PageNum

	sql = "select ts, ? from ? where ts >= '?' order by ts desc limit ?, ?"
	ls, _ := service.TdEngine().GetAll(
		ctx, sql,
		in.PropertyKey,
		p.Key,
		gtime.Now().Format("Y-m-d"),
		(in.PageNum-1)*in.PageSize,
		in.PageSize,
	)

	var pro []*model.DevicePropertiyOut
	for _, v := range ls.List() {

		pro = append(pro, &model.DevicePropertiyOut{
			Ts:    gvar.New(v["ts"]).GTime(),
			Value: gvar.New(v[in.PropertyKey]),
		})
	}
	out.List = pro

	return
}
