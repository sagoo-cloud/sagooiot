package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sagooiot/api/v1/product"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/iotModel/sagooProtocol/north"
	"sagooiot/pkg/response"
	"sagooiot/pkg/tsd"
	"sagooiot/pkg/tsd/comm"
	"sagooiot/pkg/utility"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
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

// Get 获取设备详情
func (s *sDevDevice) Get(ctx context.Context, key string) (out *model.DeviceOutput, err error) {
	err = dao.DevDevice.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     consts.GetDetailDeviceOutput + key,
		Force:    false,
	}).WithAll().Where(dao.DevDevice.Columns().Key, key).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		err = errors.New("设备不存在")
		return
	}
	if out.Status != 0 {
		out.Status = dcache.GetDeviceStatus(ctx, out.Key) //查询设备状态
	}
	if out.Product != nil {
		out.ProductName = out.Product.Name
		if out.Product.Metadata != "" {
			err = json.Unmarshal([]byte(out.Product.Metadata), &out.TSL)
		}
	}
	return
}

// GetAll 获取所有设备
func (s *sDevDevice) GetAll(ctx context.Context) (out []*entity.DevDevice, err error) {
	m := dao.DevDevice.Ctx(ctx)
	err = m.Scan(&out)
	return
}
func (s *sDevDevice) Detail(ctx context.Context, key string) (out *model.DeviceOutput, err error) {
	err = dao.DevDevice.Ctx(ctx).WithAll().Where(dao.DevDevice.Columns().Key, key).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		err = errors.New("设备不存在")
		return
	}
	if out.Status != 0 {
		out.Status = dcache.GetDeviceStatus(ctx, out.Key) //查询设备状态
	}

	//如果未设置，获取系统设置的默认超时时间
	if out.OnlineTimeout == 0 {
		defaultTimeout, err := service.ConfigData().GetConfigByKey(ctx, consts.DeviceDefaultTimeoutTime)
		if err != nil || defaultTimeout == nil {
			defaultTimeout = &entity.SysConfig{
				ConfigValue: "30",
			}
		}
		out.OnlineTimeout = gconv.Int(defaultTimeout.ConfigValue)
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
		//获取当前在线的设备KEY列表
		onlineDevice, err := dcache.GetOnlineDeviceList()
		if in.Status == gconv.String(consts.DeviceStatueOnline) {
			if err == nil && len(onlineDevice) > 0 {
				m = m.Where(c.Key, onlineDevice)
			} else {
				return nil, nil
			}

		} else if in.Status == gconv.String(consts.DeviceStatueOffline) {
			if err == nil && len(onlineDevice) > 0 {
				m = m.WhereNotIn(c.Key, onlineDevice)
			}
		} else {
			m = m.Where(c.Status, gconv.Int(in.Status))
		}
	}
	if in.Key != "" {
		m = m.Where(c.Key, in.Key)
	}
	if in.ProductKey != "" {
		m = m.Where(c.ProductKey, in.ProductKey)
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

	m = m.WhereIn(
		dao.DevDevice.Columns().ProductKey,
		dao.DevProduct.Ctx(ctx).Fields(dao.DevProduct.Columns().Key),
	)

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
		if out.Device[i].Status != 0 {
			out.Device[i].Status = dcache.GetDeviceStatus(ctx, out.Device[i].Key)
		}
		out.Device[i].Product.Metadata = ""
	}

	return
}

// List 已发布产品的设备列表
func (s *sDevDevice) List(ctx context.Context, productKey string, keyWord string) (list []*model.DeviceOutput, err error) {
	m := dao.DevDevice.Ctx(ctx).
		Where(dao.DevDevice.Columns().Status+" > ?", model.DeviceStatusNoEnable)
	if productKey != "" {
		m = m.Where(dao.DevDevice.Columns().ProductKey, productKey)
	}
	if keyWord != "" {
		m = m.WhereLike(dao.DevDevice.Columns().Key, "%"+keyWord+"%").WhereOrLike(dao.DevDevice.Columns().Name, "%"+keyWord+"%")
	}

	err = m.WhereIn(dao.DevDevice.Columns().ProductKey,
		dao.DevProduct.Ctx(ctx).
			Fields(dao.DevProduct.Columns().Key).
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
		err = errors.New("设备标识重复")
		return
	}

	// 设备标识不能和产品标识重复
	pout, err := service.DevProduct().Detail(ctx, in.Key)
	if err != nil {
		return
	}
	if pout != nil {
		err = errors.New("设备标识不能和产品标识重复")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	rs, err := dao.DevDevice.Ctx(ctx).Data(do.DevDevice{
		DeptId:        service.Context().GetUserDeptId(ctx),
		Key:           in.Key,
		Name:          in.Name,
		ProductKey:    in.ProductKey,
		Desc:          in.Desc,
		Version:       in.Version,
		Lng:           in.Lng,
		Lat:           in.Lat,
		AuthType:      in.AuthType,
		AuthUser:      in.AuthUser,
		AuthPasswd:    in.AuthPasswd,
		AccessToken:   in.AccessToken,
		CertificateId: in.CertificateId,
		//ExtensionInfo: in.ExtensionInfo,
		Status:    0,
		CreatedBy: uint(loginUserId),
		CreatedAt: gtime.Now(),
		//Address:       in.Address,
	}).Insert()
	if err != nil {
		return
	}
	//北向设备添加消息
	north.WriteMessage(ctx, north.DeviceAddMessageTopic, nil, in.ProductKey, in.Key, iotModel.DeviceAddMessage{
		Timestamp: time.Now().UnixMilli(),
		Desc:      "",
	})
	newId, err := rs.LastInsertId()
	deviceId = uint(newId)

	// 设备标签
	if len(in.Tags) > 0 {
		for _, v := range in.Tags {
			intag := model.AddTagDeviceInput{
				DeviceId:  deviceId,
				DeviceKey: in.Key,
				Key:       v.Key,
				Name:      v.Name,
				Value:     v.Value,
			}
			if err = service.DevDeviceTag().Add(ctx, &intag); err != nil {
				return
			}
		}
	}

	return
}

func (s *sDevDevice) Edit(ctx context.Context, in *model.EditDeviceInput) (err error) {
	out, err := s.Detail(ctx, in.Key)
	if err != nil {
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevDevice
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdatedBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevDevice.Ctx(ctx).Data(param).Where(dao.DevDevice.Columns().Key, in.Key).Update()
	if err != nil {
		return
	}

	// 设备标签
	if len(in.Tags) > 0 {
		var l []model.AddTagDeviceInput
		for _, v := range in.Tags {
			intag := model.AddTagDeviceInput{
				DeviceId:  out.Id,
				DeviceKey: out.Key,
				Key:       v.Key,
				Name:      v.Name,
				Value:     v.Value,
			}
			l = append(l, intag)
		}
		if err = service.DevDeviceTag().Update(ctx, out.Id, l); err != nil {
			return
		}
	}
	//从缓存中删除
	_, err = cache.Instance().Remove(ctx, consts.CacheGfOrmPrefix+consts.GetDetailDeviceOutput+out.Key)

	return
}

// UpdateDeviceStatusInfo 更新设备状态信息，设备上线、离线、注册
func (s *sDevDevice) UpdateDeviceStatusInfo(ctx context.Context, deviceKey string, status int, timestamp time.Time) (err error) {
	device, err := s.Get(ctx, deviceKey)
	if err != nil {
		return
	}
	if device == nil {
		err = errors.New("设备不存在")
		return
	}
	var data = g.Map{}
	switch status {
	case consts.DeviceStatueOnline:
		if device.RegistryTime == nil {
			data[dao.DevDevice.Columns().RegistryTime] = timestamp
		} else {

			data[dao.DevDevice.Columns().LastOnlineTime] = timestamp
		}
		data[dao.DevDevice.Columns().Status] = consts.DeviceStatueOnline

	case consts.DeviceStatueOffline:
		data[dao.DevDevice.Columns().LastOnlineTime] = timestamp
		data[dao.DevDevice.Columns().Status] = consts.DeviceStatueOffline
	}
	_, err = dao.DevDevice.Ctx(ctx).
		Data(data).
		Where(dao.DevDevice.Columns().Key, deviceKey).
		Update()
	return
}

// BatchUpdateDeviceStatusInfo 批量更新设备状态信息，设备上线、离线、注册
func (s *sDevDevice) BatchUpdateDeviceStatusInfo(ctx context.Context, deviceStatusLogList []iotModel.DeviceStatusLog) (err error) {

	onLineData := g.Map{}
	onlineDeviceKeyList := make([]string, 0)

	offLineData := g.Map{}
	offLineDeviceKeyList := make([]string, 0)

	registryData := g.Map{}
	registryDeviceKeyList := make([]string, 0)

	for _, statusLog := range deviceStatusLogList {
		switch statusLog.Status {
		case consts.DeviceStatueOnline:
			device, err := dcache.GetDeviceDetailInfo(statusLog.DeviceKey)
			if err != nil {
				continue
			}

			if device.RegistryTime == nil {
				registryData[dao.DevDevice.Columns().RegistryTime] = statusLog.Timestamp
				registryDeviceKeyList = append(registryDeviceKeyList, statusLog.DeviceKey)
			}

			onLineData[dao.DevDevice.Columns().LastOnlineTime] = statusLog.Timestamp
			onLineData[dao.DevDevice.Columns().Status] = consts.DeviceStatueOnline
			onlineDeviceKeyList = append(onlineDeviceKeyList, statusLog.DeviceKey)
		case consts.DeviceStatueOffline:
			offLineData[dao.DevDevice.Columns().LastOnlineTime] = statusLog.Timestamp
			offLineData[dao.DevDevice.Columns().Status] = consts.DeviceStatueOffline
			offLineDeviceKeyList = append(offLineDeviceKeyList, statusLog.DeviceKey)
		}
	}
	if len(registryDeviceKeyList) > 0 {
		_, err = dao.DevDevice.Ctx(ctx).
			Data(registryData).
			WhereIn(dao.DevDevice.Columns().Key, registryDeviceKeyList).
			Update()
	}
	if len(onlineDeviceKeyList) > 0 {
		_, err = dao.DevDevice.Ctx(ctx).
			Data(onLineData).
			WhereIn(dao.DevDevice.Columns().Key, onlineDeviceKeyList).
			Update()
	}
	if len(offLineDeviceKeyList) > 0 {
		_, err = dao.DevDevice.Ctx(ctx).
			Data(offLineData).
			WhereIn(dao.DevDevice.Columns().Key, onlineDeviceKeyList).
			Update()
	}

	return
}

func (s *sDevDevice) UpdateExtend(ctx context.Context, in *model.DeviceExtendInput) (err error) {
	_, err = s.Detail(ctx, in.Key)
	if err != nil {
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevDevice
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdatedBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevDevice.Ctx(ctx).Data(param).Where(dao.DevDevice.Columns().Key, in.Key).Update()

	return
}

func (s *sDevDevice) Del(ctx context.Context, keys []string) (err error) {
	var p []*entity.DevDevice
	err = dao.DevDevice.Ctx(ctx).WhereIn(dao.DevDevice.Columns().Key, keys).Scan(&p)
	if err != nil {
		return
	}
	if len(p) == 0 {
		return errors.New("设备不存在")
	}

	// 状态校验
	for _, v := range p {
		if v.Status > model.DeviceStatusNoEnable {
			return errors.New(fmt.Sprintf("设备(%s)已启用，不能删除", v.Key))
		}
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	for _, key := range keys {
		res, _ := s.Detail(ctx, key)
		rs, err := dao.DevDevice.Ctx(ctx).
			Data(do.DevDevice{
				DeletedBy: uint(loginUserId),
				DeletedAt: gtime.Now(),
			}).
			Where(dao.DevDevice.Columns().Status, model.DeviceStatusNoEnable).
			Where(dao.DevDevice.Columns().Key, key).
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

			// 删除网关子设备TD子表
			if res.Product.DeviceType == model.DeviceTypeGateway {
				subList, err := s.bindList(ctx, res.Key)
				if err != nil {
					return err
				}
				for _, sub := range subList {
					if sub.MetadataTable == 1 {
						// 删除TD子表
						err = service.TSLTable().DropTable(ctx, sub.Key)
						if err != nil {
							return err
						}
					}
				}
			}
		}
		//北向设备删除消息
		north.WriteMessage(ctx, north.DeviceDeleteMessageTopic, nil, res.Product.Key, res.Key, iotModel.DeviceDeleteMessage{
			Timestamp: time.Now().UnixMilli(),
			Desc:      "",
		})
	}

	return
}

// Deploy 设备启用
func (s *sDevDevice) Deploy(ctx context.Context, key string) (err error) {
	//获取设备信息
	device, err := s.Detail(ctx, key)
	if err != nil {
		return
	}
	if device.Status > model.DeviceStatusNoEnable {
		return
	}

	pd, err := service.DevProduct().Detail(ctx, device.ProductKey)
	if err != nil {
		return err
	}
	if pd == nil {
		return errors.New("产品不存在")
	}
	if pd.Status != model.ProductStatusOn {
		return errors.New("产品未发布，请先发布产品")
	}

	err = dao.DevDevice.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 是否创建TD表结构
		var isCreate bool

		if device.MetadataTable == 0 {
			isCreate = true
		}
		// 检查TD表是否存在，不存在，则补建
		if device.MetadataTable == 1 {
			table := comm.DeviceTableName(device.Key)
			b, _ := service.TSLTable().CheckTable(ctx, table)
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
			err = service.TSLTable().CreateTable(ctx, pd.Key, device.Key)
			if err != nil {
				return err
			}
		}

		// 更新状态
		_, err = dao.DevDevice.Ctx(ctx).
			Data(g.Map{
				dao.DevDevice.Columns().Status:        model.DeviceStatusOff,
				dao.DevDevice.Columns().MetadataTable: 1,
			}).
			Where(dao.DevDevice.Columns().Key, key).
			Update()

		//设备启用后，更新缓存数据
		device.Status = model.DeviceStatusOff
		err = dcache.SetDeviceDetailInfo(device.Key, device)
		if err != nil {
			g.Log().Debug(ctx, "Deploy 设备数据存入缓存失败", err.Error())
		}

		// 网关启用子设备
		if pd.DeviceType == model.DeviceTypeGateway {
			subList, err := s.bindList(ctx, device.Key)
			if err != nil {
				return err
			}

			for _, sub := range subList {
				if err := s.Deploy(ctx, sub.Key); err != nil {
					return err
				}
				//设备启用后，更新缓存数据
				sub.Status = model.DeviceStatusOff
				err = dcache.SetDeviceDetailInfo(sub.Key, sub)
				if err != nil {
					g.Log().Debug(ctx, "Deploy 子设备数据存入缓存失败", err.Error())
				}
			}
		}

		return err
	})

	return
}

// Undeploy 设备禁用
func (s *sDevDevice) Undeploy(ctx context.Context, key string) (err error) {
	//获取设备信息
	device, err := s.Detail(ctx, key)
	if err != nil {
		return
	}
	if device.Status == model.DeviceStatusNoEnable {
		return
	}

	err = dao.DevDevice.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.DevDevice.Ctx(ctx).
			Data(g.Map{dao.DevDevice.Columns().Status: model.DeviceStatusNoEnable}).
			Where(dao.DevDevice.Columns().Key, key).
			Update()
		if err != nil {
			return err
		}

		//设备停用后，更新缓存数据
		device.Status = model.DeviceStatusNoEnable
		err = dcache.SetDeviceDetailInfo(device.Key, device)
		if err != nil {
			g.Log().Debug(ctx, "Deploy 设备数据存入缓存失败", err.Error())
		}

		// 网关禁用子设备
		if device.Product.DeviceType == model.DeviceTypeGateway {
			subList, err := s.bindList(ctx, device.Key)
			if err != nil {
				return err
			}

			for _, sub := range subList {
				if err := s.Undeploy(ctx, sub.Key); err != nil {
					return err
				}
				//设备停用后，更新缓存数据
				sub.Status = model.DeviceStatusNoEnable
				err = dcache.SetDeviceDetailInfo(sub.Key, sub)
				if err != nil {
					g.Log().Debug(ctx, "Deploy 子设备数据存入缓存失败", err.Error())
				}
			}
		}

		return err
	})

	return
}

// TotalByProductKey 统计产品下的设备数量
func (s *sDevDevice) TotalByProductKey(ctx context.Context, productKeys []string) (totals map[string]int, err error) {
	m := dao.DevDevice.Ctx(ctx)
	r, err := m.Fields(dao.DevDevice.Columns().ProductKey+", count(*) as total").
		WhereIn(dao.DevDevice.Columns().ProductKey, productKeys).
		Group(dao.DevDevice.Columns().ProductKey).
		All()
	if err != nil || r.Len() == 0 {
		return
	}

	totals = make(map[string]int, r.Len())
	for _, v := range r {
		t := gconv.Int(v["total"])
		productKey := gconv.String(v[dao.DevDevice.Columns().ProductKey])
		totals[productKey] = t
	}

	for _, key := range productKeys {
		if _, ok := totals[key]; !ok {
			totals[key] = 0
		}
	}

	return
}

// getTotalForMonthsData 获取统计设备月度消息数量
func (s *sDevDevice) getTotalForMonthsData(ctx context.Context) (data map[int]int, err error) {

	data = make(map[int]int, 12)
	for i := 0; i < 12; i++ {
		data[i+1] = 0
	}
	devices, err := s.GetAll(ctx)
	if err != nil {
		return
	}
	if devices != nil {
		var deviceKeys []string
		for _, device := range devices {
			deviceKeys = append(deviceKeys, "'"+device.Key+"'")
		}

		tsdDb := tsd.DB()
		defer tsdDb.Close()

		// 设备消息总量 TDengine
		var list gdb.Result
		sql := "select substr(to_iso8601(ts), 1, 7) as ym, count(*) as num from device_log where device in (?) and  ts >= '?' group by substr(to_iso8601(ts), 1, 7)"
		list, err = tsdDb.GetTableDataAll(ctx, sql, strings.Join(deviceKeys, ","), gtime.Now().Format("Y-01-01 00:00:00"))
		if err != nil {
			return
		}

		for _, v := range list {
			m := gstr.SubStr(v["ym"].String(), 5)
			data[gconv.Int(m)] = v["num"].Int()
		}
	}
	return
}

// getTotalForDayData 统计设备最近一月消息数量
func (s *sDevDevice) getTotalForDayData(ctx context.Context) (data map[string]int, err error) {

	start := time.Now().AddDate(0, -1, 1)
	data = make(map[string]int)
	for i := start; i.Before(time.Now()); {
		data[i.Format("2006-01-02")] = 0
		i = i.Add(24 * time.Hour)
	}

	devices, err := s.GetAll(ctx)
	if err != nil {
		return
	}
	if devices != nil {
		var deviceKeys []string
		for _, device := range devices {
			deviceKeys = append(deviceKeys, "'"+device.Key+"'")
		}

		tsdDb := tsd.DB()
		defer tsdDb.Close()

		// 设备消息总量 TDengine
		var list gdb.Result
		sql := "select substr(to_iso8601(ts), 1, 10) as ymd, count(*) as num from device_log where device in (?) partition by substr(to_iso8601(ts), 1, 10) interval(4w)"
		list, err = tsdDb.GetTableDataAll(ctx, sql, strings.Join(deviceKeys, ","))
		if err != nil {
			return
		}

		for _, v := range list {
			data[v["ymd"].String()] = v["num"].Int()
		}
	}

	return
}

// getAlarmTotalForMonthsData 统计设备月度告警数量
func (s *sDevDevice) getAlarmTotalForMonthsData(ctx context.Context) (data map[int]int, err error) {

	data = make(map[int]int, 12)
	for i := 0; i < 12; i++ {
		data[i+1] = 0
	}

	devices, err := s.GetAll(ctx)
	if err != nil {
		return
	}
	var deviceKeys []string
	for _, device := range devices {
		deviceKeys = append(deviceKeys, device.Key)
	}

	at := dao.AlarmLog.Columns().CreatedAt
	list, err := dao.AlarmLog.Ctx(ctx).
		Fields("date_format("+at+", '%Y%m') as ym, count(*) as num").
		Where(at+">=?", gtime.Now().Format("Y-01-01 00:00:00")).
		WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).
		Group("date_format(" + at + ", '%Y%m')").
		All()
	if err != nil {
		return
	}

	for _, v := range list {
		m := gstr.SubStr(v["ym"].String(), 4)
		data[gconv.Int(m)] = v["num"].Int()
	}

	return
}

// getAlarmTotalForDayData 获取统计设备最近一个月告警数量
func (s *sDevDevice) getAlarmTotalForDayData(ctx context.Context) (data map[string]int, err error) {

	start := time.Now().AddDate(0, -1, 1)
	data = make(map[string]int)
	for i := start; i.Before(time.Now()); {
		data[i.Format("2006-01-02")] = 0
		i = i.Add(24 * time.Hour)
	}

	devices, err := s.GetAll(ctx)
	if err != nil {
		return
	}
	var deviceKeys []string
	for _, device := range devices {
		deviceKeys = append(deviceKeys, "'"+device.Key+"'")
	}

	at := dao.AlarmLog.Columns().CreatedAt
	list, err := dao.AlarmLog.Ctx(ctx).
		Fields("date_format("+at+", '%Y-%m-%d') as ymd, count(*) as num").
		Where(at+">=?", start.Format("20060102")).
		WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).
		Group("date_format(" + at + ", '%Y-%m-%d')").
		All()
	if err != nil {
		return
	}

	for _, v := range list {
		data[v["ymd"].String()] = v["num"].Int()
	}

	return
}

// RunStatus 运行状态
func (s *sDevDevice) RunStatus(ctx context.Context, deviceKey string) (out *model.DeviceRunStatusOutput, err error) {

	device, err := dcache.GetDeviceDetailInfo(deviceKey)
	if err != nil {
		return nil, errors.New("设备不存在")
	}
	out = new(model.DeviceRunStatusOutput)
	out.Status = dcache.GetDeviceStatus(ctx, deviceKey)

	//获取数据
	deviceValueList := dcache.GetDeviceDetailData(context.Background(), deviceKey, consts.MsgTypePropertyReport, consts.MsgTypeGatewayBatch)
	if len(deviceValueList) == 0 {
		return
	}

	for _, d := range deviceValueList[0] {
		out.LastOnlineTime = gtime.New(d.CreateTime)
	}

	var properties []model.DevicePropertiy
	for _, v := range device.TSL.Properties {
		unit := ""
		if v.ValueType.Unit != nil {
			unit = *v.ValueType.Unit
		}

		var valueList = make([]*g.Var, 0)
		for _, d := range deviceValueList {
			valueList = append(valueList, g.NewVar(d[v.Key].Value))
		}

		pro := model.DevicePropertiy{
			Key:   v.Key,
			Name:  v.Name,
			Type:  v.ValueType.Type,
			Unit:  unit,
			Value: valueList[0],
			List:  valueList,
		}
		properties = append(properties, pro)
	}
	out.Properties = properties
	return
}

// GetLatestProperty 获取设备最新的属性值
func (s *sDevDevice) GetLatestProperty(ctx context.Context, key string) (list []model.DeviceLatestProperty, err error) {
	p, err := s.Get(ctx, key)
	if err != nil {
		return
	}
	if p.Status == model.DeviceStatusNoEnable {
		return
	}

	deviceTable := comm.DeviceTableName(p.Key)

	tsdDb := tsd.DB()
	defer tsdDb.Close()

	for _, v := range p.TSL.Properties {
		ckey := comm.TsdColumnName(v.Key)

		// 获取属性最近有效值
		sql := "select ? from ? where ? is not null order by ts desc limit 1"
		rs, err := tsdDb.GetTableDataOne(ctx, sql, ckey, deviceTable, ckey)
		if err != nil {
			return nil, err
		}
		value := rs[strings.ToLower(v.Key)]
		if value.IsEmpty() {
			continue
		}

		unit := ""
		if v.ValueType.Unit != nil {
			unit = *v.ValueType.Unit
		}

		pro := model.DeviceLatestProperty{
			Key:   v.Key,
			Name:  v.Name,
			Type:  v.ValueType.Type,
			Unit:  unit,
			Value: value,
		}
		list = append(list, pro)
	}
	return
}

// GetProperty 获取指定属性值
func (s *sDevDevice) GetProperty(ctx context.Context, in *model.DeviceGetPropertyInput) (out *model.DevicePropertiy, err error) {
	p, err := s.Detail(ctx, in.DeviceKey)
	if err != nil {
		return
	}
	if p.Status == model.DeviceStatusNoEnable {
		err = errors.New("设备未启用")
		return
	}

	tsdDb := tsd.DB()
	defer tsdDb.Close()

	sKey := in.PropertyKey
	in.PropertyKey = strings.ToLower(in.PropertyKey)
	col := comm.TsdColumnName(in.PropertyKey)

	deviceTable := comm.DeviceTableName(p.Key)

	// 属性上报时间
	ctime := in.PropertyKey + "_time"
	ctime = comm.TsdColumnName(ctime)

	// 属性值获取
	sql := "select ? from ? where ? is not null order by ? desc limit 1"
	rs, err := tsdDb.GetTableDataOne(ctx, sql, col, deviceTable, col, ctime)
	if err != nil {
		return
	}

	var name string
	var valueType string
	for _, v := range p.TSL.Properties {
		if strings.ToLower(v.Key) == in.PropertyKey {
			name = v.Name
			valueType = v.ValueType.Type
			break
		}
	}

	out = new(model.DevicePropertiy)
	out.Key = sKey
	out.Name = name
	out.Type = valueType
	out.Value = rs[in.PropertyKey]

	// 获取当天属性值列表
	sql = "select ? from ? where ? >= '?' order by ? desc"
	ls, _ := tsdDb.GetTableDataAll(ctx, sql, col, deviceTable, ctime, gtime.Now().Format("Y-m-d"), ctime)
	out.List = ls.Array(in.PropertyKey)

	return
}

// GetPropertyList 设备属性详情列表
func (s *sDevDevice) GetPropertyList(ctx context.Context, in *model.DeviceGetPropertyListInput) (out *model.DeviceGetPropertyListOutput, err error) {
	resultList, total, currentPage := dcache.GetDeviceDetailDataByPage(ctx, in.DeviceKey, in.PageNum, in.PageSize, consts.MsgTypePropertyReport, consts.MsgTypeGatewayBatch)
	if err != nil {
		return
	}
	out = new(model.DeviceGetPropertyListOutput)
	out.Total = total
	out.CurrentPage = currentPage

	var pro []*model.DevicePropertiyOut
	for _, d := range resultList {
		pro = append(pro, &model.DevicePropertiyOut{
			Ts:    gtime.New(d[in.PropertyKey].CreateTime),
			Value: gvar.New(d[in.PropertyKey].Value),
		})
	}
	out.List = pro

	return
}

// GetData 获取设备指定日期属性数据
func (s *sDevDevice) GetData(ctx context.Context, in *model.DeviceGetDataInput) (list []model.DevicePropertiyOut, err error) {
	p, err := s.Detail(ctx, in.DeviceKey)
	if err != nil {
		return
	}
	if p.Status == model.DeviceStatusNoEnable {
		err = errors.New("设备未启用")
		return
	}

	deviceTable := comm.DeviceTableName(p.Key)

	in.PropertyKey = strings.ToLower(in.PropertyKey)
	col := comm.TsdColumnName(in.PropertyKey)

	// 属性上报时间
	ctime := in.PropertyKey + "_time"
	ctime = comm.TsdColumnName(ctime)

	where := fmt.Sprintf("%s >= %q", ctime, gtime.Now().Format("Y-m-d"))
	if len(in.DateRange) > 1 {
		where = fmt.Sprintf("%s >= %q and %s <= %q", ctime, in.DateRange[0], ctime, in.DateRange[1])
	}

	desc := "asc"
	if in.IsDesc == 1 {
		desc = "desc"
	}

	tsdDb := tsd.DB()
	defer tsdDb.Close()

	// TDengine
	sql := "select ?, ? from ? where ? order by ? ?"
	ls, _ := tsdDb.GetTableDataAll(
		ctx,
		sql,
		fmt.Sprintf("distinct %s as ts", ctime),
		col,
		deviceTable,
		where,
		ctime,
		desc,
	)
	for _, v := range ls.List() {
		list = append(list, model.DevicePropertiyOut{
			Ts:    gvar.New(v["ts"]).GTime(),
			Value: gvar.New(v[in.PropertyKey]),
		})
	}
	return
}

// BindSubDevice 网关绑定子设备
func (s *sDevDevice) BindSubDevice(ctx context.Context, in *model.DeviceBindInput) error {
	gw, err := s.Get(ctx, in.GatewayKey)
	if err != nil {
		return err
	}
	if gw.Product.DeviceType != model.DeviceTypeGateway {
		return errors.New("非网关，不能绑定子设备")
	}
	if len(in.SubKeys) == 0 {
		return nil
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	for _, v := range in.SubKeys {
		sub, err := s.Get(ctx, v)
		if err != nil {
			return err
		}
		if sub.Product.DeviceType != model.DeviceTypeSub {
			return errors.New("非子设备类型，不能绑定")
		}

		rs, err := dao.DevDeviceGateway.Ctx(ctx).
			Where(dao.DevDeviceGateway.Columns().SubKey, v).
			One()
		if err != nil {
			return err
		}
		if !rs.IsEmpty() {
			return errors.New(fmt.Sprintf("%s，该子设备已被绑定", sub.Name))
		}

		_, err = dao.DevDeviceGateway.Ctx(ctx).Data(do.DevDeviceGateway{
			GatewayKey: in.GatewayKey,
			SubKey:     v,
			CreatedBy:  uint(loginUserId),
			CreatedAt:  gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

// UnBindSubDevice 网关解绑子设备
func (s *sDevDevice) UnBindSubDevice(ctx context.Context, in *model.DeviceBindInput) error {
	gw, err := s.Get(ctx, in.GatewayKey)
	if err != nil {
		return err
	}
	if gw.Product.DeviceType != model.DeviceTypeGateway {
		return errors.New("非网关设备")
	}
	if len(in.SubKeys) == 0 {
		return nil
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	_, err = dao.DevDeviceGateway.Ctx(ctx).
		Data(do.DevDeviceGateway{
			DeletedBy: uint(loginUserId),
			DeletedAt: gtime.Now(),
		}).
		Where(dao.DevDeviceGateway.Columns().GatewayKey, in.GatewayKey).
		WhereIn(dao.DevDeviceGateway.Columns().SubKey, in.SubKeys).
		Unscoped().
		Update()
	return err
}

// bindList 已绑定列表
func (s *sDevDevice) bindList(ctx context.Context, gatewayKey string) (list []*model.DeviceOutput, err error) {
	_, err = s.Get(ctx, gatewayKey)
	if err != nil {
		return
	}

	var dgw []*entity.DevDeviceGateway
	if err = dao.DevDeviceGateway.Ctx(ctx).Where(dao.DevDeviceGateway.Columns().GatewayKey, gatewayKey).Scan(&dgw); err != nil || len(dgw) == 0 {
		return
	}

	var subKeys []string
	for _, v := range dgw {
		subKeys = append(subKeys, v.SubKey)
	}

	err = dao.DevDevice.Ctx(ctx).
		WhereIn(dao.DevDevice.Columns().Key, subKeys).
		WithAll().
		OrderDesc(dao.DevDevice.Columns().Id).
		Scan(&list)
	return
}

// BindList 已绑定列表(分页)
func (s *sDevDevice) BindList(ctx context.Context, in *model.DeviceBindListInput) (out *model.DeviceBindListOutput, err error) {
	_, err = s.Get(ctx, in.GatewayKey)
	if err != nil {
		return
	}

	var dgw []*entity.DevDeviceGateway
	if err = dao.DevDeviceGateway.Ctx(ctx).Where(dao.DevDeviceGateway.Columns().GatewayKey, in.GatewayKey).Scan(&dgw); err != nil || len(dgw) == 0 {
		return
	}

	var subKeys []string
	for _, v := range dgw {
		subKeys = append(subKeys, v.SubKey)
	}

	m := dao.DevDevice.Ctx(ctx).
		WhereIn(dao.DevDevice.Columns().Key, subKeys).
		WithAll().
		OrderDesc(dao.DevDevice.Columns().Id)

	out = &model.DeviceBindListOutput{}
	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Scan(&out.List)
	if err != nil {
		return
	}

	for i, v := range out.List {
		out.List[i].Status = dcache.GetDeviceStatus(ctx, v.Key)
	}

	return
}

// ListForSub 子设备
func (s *sDevDevice) ListForSub(ctx context.Context, in *model.ListForSubInput) (out *model.ListDeviceForPageOutput, err error) {
	m := dao.DevDevice.Ctx(ctx)
	if in.ProductKey != "" {
		m = m.Where(dao.DevDevice.Columns().ProductKey, in.ProductKey)
	}
	if in.GatewayKey != "" {
		rs, err := dao.DevDeviceGateway.Ctx(ctx).
			Fields(dao.DevDeviceGateway.Columns().SubKey).
			Where(dao.DevDeviceGateway.Columns().GatewayKey, in.GatewayKey).
			Array()
		if err == nil && len(rs) > 0 {
			m = m.WhereNotIn(dao.DevDevice.Columns().Key, rs)
		}
	}

	m = m.WithAll().OrderDesc(dao.DevDevice.Columns().Id)
	out = &model.ListDeviceForPageOutput{}
	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Scan(&out.Device)
	return
}

// CheckBind 检查网关、子设备绑定关系
func (s *sDevDevice) CheckBind(ctx context.Context, in *model.CheckBindInput) (bool, error) {
	count, err := dao.DevDeviceGateway.Ctx(ctx).
		Where(dao.DevDeviceGateway.Columns().GatewayKey, in.GatewayKey).
		Where(dao.DevDeviceGateway.Columns().SubKey, in.SubKey).
		Count()
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}

// DelSub 子设备删除
func (s *sDevDevice) DelSub(ctx context.Context, key string) (err error) {
	subDev, err := s.Detail(ctx, key)
	if err != nil {
		return
	}
	if subDev.Product.DeviceType != model.DeviceTypeSub {
		return errors.New("该设备不是子设备类型")
	}
	if subDev.Status > model.DeviceStatusNoEnable {
		return errors.New("设备已启用，不能删除")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	now := gtime.Now()

	rs, err := dao.DevDevice.Ctx(ctx).
		Data(do.DevDevice{
			DeletedBy: uint(loginUserId),
			DeletedAt: now,
		}).
		Where(dao.DevDevice.Columns().Key, key).
		Unscoped().
		Update()
	if err != nil {
		return err
	}

	// 删除绑定关系
	_, err = dao.DevDeviceGateway.Ctx(ctx).
		Data(do.DevDeviceGateway{
			DeletedBy: uint(loginUserId),
			DeletedAt: now,
		}).
		Where(dao.DevDeviceGateway.Columns().SubKey, subDev.Key).
		Unscoped().
		Update()
	if err != nil {
		return
	}

	num, _ := rs.RowsAffected()
	if num > 0 && subDev.MetadataTable == 1 {
		// 删除TD子表
		err = service.TSLTable().DropTable(ctx, subDev.Key)
		if err != nil {
			return err
		}
	}

	return
}

// AuthInfo 获取认证信息
func (s *sDevDevice) AuthInfo(ctx context.Context, in *model.AuthInfoInput) (*model.AuthInfoOutput, error) {
	if in.DeviceKey == "" && in.ProductKey == "" {
		return nil, errors.New("缺少必要参数")
	}

	out := &model.AuthInfoOutput{}

	if in.DeviceKey != "" {
		device, err := s.Get(ctx, in.DeviceKey)
		if err != nil {
			return nil, err
		}
		out.AuthType = device.AuthType
		out.AuthUser = device.AuthUser
		out.AuthPasswd = device.AuthPasswd
		out.AccessToken = device.AccessToken
		out.CertificateId = device.CertificateId

		if out.AuthUser == "" && out.AuthPasswd == "" && out.AccessToken == "" && out.CertificateId == 0 {
			in.ProductKey = device.Product.Key
		} else {
			in.ProductKey = ""
		}
	}

	if in.ProductKey != "" {
		productData, err := service.DevProduct().Detail(ctx, in.ProductKey)
		if err != nil {
			return nil, err
		}
		out.AuthType = productData.AuthType
		out.AuthUser = productData.AuthUser
		out.AuthPasswd = productData.AuthPasswd
		out.AccessToken = productData.AccessToken
		out.CertificateId = productData.CertificateId
	}

	if out.CertificateId > 0 {
		if err := dao.SysCertificate.Ctx(ctx).Where(dao.SysCertificate.Columns().Id, out.CertificateId).Scan(&out.Certificate); err != nil {
			return nil, err
		}
	}

	return out, nil
}

// GetDeviceOnlineTimeOut 获取设备在线超时时长
func (s *sDevDevice) GetDeviceOnlineTimeOut(ctx context.Context, deviceKey string) (timeOut int) {
	// 获取设备在线超时时长
	timeOut = consts.DeviceOnlineTimeOut
	tv, err := dao.DevDevice.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 5,
		Name:     "DeviceOnlineTimeOut" + deviceKey,
		Force:    false,
	}).Where(dao.DevDevice.Columns().Key, deviceKey).Fields(dao.DevDevice.Columns().OnlineTimeout).Value()
	if err == nil && tv.Int() > 0 {
		timeOut = tv.Int()
	}
	return
}

// ExportDevices 导出设备
func (s *sDevDevice) ExportDevices(ctx context.Context, req *product.ExportDevicesReq) (res product.ExportDevicesRes, err error) {
	var data []model.DeviceOutput
	m := dao.DevDevice.Ctx(ctx).WithAll().Where(dao.DevDevice.Columns().ProductKey, req.ProductKey)
	if err = m.Scan(&data); err != nil {
		return
	}

	var outList []*model.DeviceExport
	for _, v := range data {
		var status string
		switch v.DevDevice.Status {
		case model.DeviceStatusNoEnable:
			status = "未启用"
		case model.DeviceStatusOff:
			status = "离线"
		case model.DeviceStatusOn:
			status = "在线"
		}
		var reqData = new(model.DeviceExport)
		reqData.Desc = v.DevDevice.Desc
		reqData.DeviceKey = v.DevDevice.Key
		reqData.DeviceName = v.DevDevice.Name
		reqData.ProductName = v.Product.Name
		reqData.Version = v.DevDevice.Version
		reqData.DeviceType = v.Product.DeviceType
		reqData.Status = status
		outList = append(outList, reqData)
	}
	if len(outList) == 0 {
		var reqData = new(model.DeviceExport)
		outList = append(outList, reqData)
	}

	//处理数据并导出
	var outData []interface{}
	for _, d := range outList {
		outData = append(outData, d)
	}
	dataRes := utility.ToExcel(outData)
	var request = g.RequestFromCtx(ctx)
	response.ToXls(request, dataRes, "设备列表")

	return
}

// ImportDevices 导入设备
func (s *sDevDevice) ImportDevices(ctx context.Context, req *product.ImportDevicesReq) (res product.ImportDevicesRes, err error) {
	file, err := req.File.Open()
	if err != nil {
		return
	}
	xlsx, err := excelize.OpenReader(file)

	if err != nil {
		return
	}

	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		return
	}
	device := new(entity.DevDevice)
	if len(rows) < 2 {
		err = errors.New("请添加设备")
		return
	}
	productData := new(entity.DevProduct)
	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, req.ProductKey).Scan(&productData)
	if err != nil {
		err = errors.New("产品不存在")
		return
	}
	for rIndex, row := range rows {
		if len(row) < 7 {
			continue
		}
		if rIndex == 0 {
			continue
		}
		bl := gregex.IsMatch("^[A-Za-z_]+[A-Za-z0-9_]*|[0-9]+$", []byte(row[1]))
		if !bl {
			res.Fail++
			res.DevicesKey = append(res.DevicesKey, row[1])
			continue
		}
		device.DeptId = service.Context().GetUserDeptId(ctx)
		device.ProductKey = productData.Key
		device.Name = row[0]
		device.Key = row[1]
		device.Desc = row[2]
		device.Version = row[3]
		device.Lng = row[4]
		device.Lat = row[5]
		device.OnlineTimeout = gconv.Int(row[6])

		//device.Version = row[5]

		_, err = dao.DevDevice.Ctx(ctx).Insert(device)
		if err != nil {
			res.Fail++
			res.DevicesKey = append(res.DevicesKey, device.Key)
		} else {
			res.Success++
		}
	}
	return
}

func (s *sDevDevice) SetDevicesStatus(ctx context.Context, req *product.SetDeviceStatusReq) (res product.SetDeviceStatusRes, err error) {
	if req.Status == 1 {
		for _, v := range req.Keys {
			err = s.Deploy(ctx, v)
			if err != nil {
				return
			}
		}
		return
	}
	if req.Status == 0 {
		for _, v := range req.Keys {
			err = s.Undeploy(ctx, v)
			if err != nil {
				return
			}
		}
	}
	return
}

// GetDeviceDataList 获取设备属性聚合数据列表
func (s *sDevDevice) GetDeviceDataList(ctx context.Context, in *model.DeviceDataListInput) (out *model.DeviceDataListOutput, err error) {
	device, err := s.Get(ctx, in.DeviceKey)
	if err != nil {
		return
	}
	propertys := device.TSL.Properties
	if len(propertys) == 0 {
		return
	}

	fields := []string{"_wstart", "_wend"}
	for _, v := range propertys {
		key := comm.TsdColumnName(v.Key)
		fields = append(fields, fmt.Sprintf("mode(%s) as %s", key, key))
	}

	timeUnit := "m"
	switch in.TimeUnit {
	case 1:
		timeUnit = "s"
	case 2:
		timeUnit = "m"
	case 3:
		timeUnit = "h"
	case 4:
		timeUnit = "d"
	}
	interval := fmt.Sprintf("interval(%d%s)", in.Interval, timeUnit)

	tsdDb := tsd.DB()
	defer tsdDb.Close()

	deviceTable := comm.DeviceTableName(device.Key)
	sql := "select count(*) as num from (select count(*) from ? ?)"
	rs, err := tsdDb.GetTableDataOne(ctx, sql, deviceTable, interval)
	if err != nil {
		return
	}
	out = new(model.DeviceDataListOutput)
	out.Total = rs["num"].Int()
	out.CurrentPage = in.PageNum

	sql = "select ? from ? ? order by _wstart desc limit ?, ?"
	out.List, err = tsdDb.GetTableDataAll(
		ctx,
		sql,
		strings.Join(fields, ","),
		deviceTable,
		interval,
		(in.PageNum-1)*in.PageSize,
		in.PageSize,
	)
	return
}

// GetAllForProduct 获取指定产品所有设备
func (s *sDevDevice) GetAllForProduct(ctx context.Context, productKey string) (list []*entity.DevDevice, err error) {
	err = dao.DevDevice.Ctx(ctx).Where(dao.DevDevice.Columns().ProductKey, productKey).Scan(&list)
	return
}

// CacheDeviceDetailList 缓存所有设备详情数据
func (s *sDevDevice) CacheDeviceDetailList(ctx context.Context) (err error) {
	productList, err := service.DevProduct().List(context.Background())
	if err != nil {
		return
	}
	for _, p := range productList {
		deviceList, err := service.DevDevice().List(context.Background(), p.Key, "")
		if err != nil {
			g.Log().Error(ctx, err.Error())
		}

		//缓存产品详细信息
		var detailProduct = new(model.DetailProductOutput)
		err = gconv.Scan(p, detailProduct)
		if err == nil {
			err = dcache.SetProductDetailInfo(p.Key, detailProduct)
		} else {
			g.Log().Error(ctx, err.Error())
		}

		for _, d := range deviceList {
			if d.Product.Metadata != "" {
				err = json.Unmarshal([]byte(d.Product.Metadata), &d.TSL)
				d.Product.Metadata = ""
				if err != nil {
					continue
				}
			}
			//缓存设备详细信息
			err := cache.Instance().Set(context.Background(), consts.DeviceDetailInfoPrefix+d.Key, d, 0)
			if err != nil {
				g.Log().Error(ctx, err.Error())
			}
		}
	}
	return

}
