package product

import (
	"context"
	"encoding/json"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/tsd/comm"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDevInit struct{}

func init() {
	service.RegisterDevInit(devInitNew())
}

func devInitNew() *sDevInit {
	return &sDevInit{}
}

// InitProductForTd 产品表结构初始化
func (s *sDevInit) InitProductForTd(ctx context.Context) (err error) {
	// 资源锁
	lockKey := "tdLock:initProductTable"
	lockVal, err := g.Redis().Do(ctx, "SET", lockKey, gtime.Now().Unix(), "NX", "EX", "3600")
	if err != nil {
		return
	}
	if lockVal.String() != "OK" {
		return
	}
	defer func() {
		_, err = g.Redis().Do(ctx, "DEL", lockKey)
	}()

	var list []*entity.DevProduct
	c := dao.DevProduct.Columns()
	err = dao.DevProduct.Ctx(ctx).Where(c.Status, model.ProductStatusOn).Where(c.MetadataTable, 1).Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	// 检测td表结构是否存在，不存在则创建
	for _, p := range list {
		stable := comm.ProductTableName(p.Key)
		b, _ := service.TSLTable().CheckStable(ctx, stable)

		if b {
			continue
		}

		var tsl *model.TSL
		err = json.Unmarshal([]byte(p.Metadata), &tsl)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		if len(tsl.Properties) == 0 {
			g.Log().Errorf(ctx, "产品 %s 物模型数据异常", p.Key)
			continue
		}

		err = service.TSLTable().CreateStable(ctx, tsl)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
	}

	return nil
}

// InitDeviceForTd 设备表结构初始化
func (s *sDevInit) InitDeviceForTd(ctx context.Context) (err error) {
	// 资源锁
	lockKey := "tdLock:initDeviceTable"
	lockVal, err := g.Redis().Do(ctx, "SET", lockKey, gtime.Now().Unix(), "NX", "EX", "3600")
	if err != nil {
		return
	}
	if lockVal.String() != "OK" {
		return
	}
	defer func() {
		_, err = g.Redis().Do(ctx, "DEL", lockKey)
	}()

	var list []*entity.DevDevice
	c := dao.DevDevice.Columns()
	err = dao.DevDevice.Ctx(ctx).WhereGT(c.Status, model.DeviceStatusNoEnable).Where(c.MetadataTable, 1).Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	// 检测td表结构是否存在，不存在则创建
	for _, d := range list {

		// 检测设备表是否创建TD表的标识
		/*if d.MetadataTable == 1 {
			continue
		}*/

		pd, err := service.DevProduct().Detail(ctx, d.ProductKey)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		if pd == nil {
			g.Log().Errorf(ctx, "设备 %s 所属产品不存在", d.Key)
			continue
		}

		table := comm.DeviceTableName(d.Key)
		b, _ := service.TSLTable().CheckTable(ctx, table)

		if b {
			continue
		}

		d := d
		go func() {
			err = service.TSLTable().CreateTable(ctx, pd.Key, d.Key)
			if err != nil {
				g.Log().Errorf(ctx, "设备 %s(%s) 建表失败: %s", d.Key, pd.Key, err.Error())
			}
		}()
	}

	return nil
}
