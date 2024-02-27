package analysis

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"time"
)

type sAnalysisProduct struct {
}

func init() {
	service.RegisterAnalysisProduct(analysisProductNew())
}

func analysisProductNew() *sAnalysisProduct {
	return &sAnalysisProduct{}
}

// GetDeviceCountForProduct 获取产品下的设备数量
func (s *sAnalysisProduct) GetDeviceCountForProduct(ctx context.Context, productKey string) (number int, err error) {
	key := consts.AnalysisProductCountPrefix + consts.ProductDeviceCount + gconv.String(productKey)
	resData, err := cache.Instance().GetOrSetFunc(ctx, key, func(ctx context.Context) (value interface{}, err error) {
		m := dao.DevDevice.Ctx(ctx)
		value, err = m.Where(dao.DevDevice.Columns().ProductKey, productKey).Count()
		if err != nil {
			return
		}
		return
	}, time.Hour*1)
	if err != nil {
		return
	}
	number = resData.Int()
	return
}

// GetProductCount 获取产品数量统计
func (s *sAnalysisProduct) GetProductCount(ctx context.Context) (res model.ProductCountRes, err error) {
	key := consts.AnalysisProductCountPrefix + "total"
	resData, err := cache.Instance().GetOrSetFunc(ctx, key, func(ctx context.Context) (value interface{}, err error) {
		value, err = s.getTotalData(ctx)
		return
	}, time.Minute*1)
	err = gconv.Struct(resData.Val(), &res)
	return
}

// getTotalData 从数据库中获取统计数据
func (s *sAnalysisProduct) getTotalData(ctx context.Context) (data model.ProductCountRes, err error) {
	m := dao.DevProduct.Ctx(ctx)
	// 产品总量
	data.Total, err = m.Count()
	if err != nil {
		return
	}
	// 产品新增数量
	data.Added, err = m.
		Where(dao.DevProduct.Columns().CreatedAt+">=?", gtime.Now().Format("Y-m-d")).
		Count()
	if err != nil {
		return
	}
	//产品启用
	data.Enable, err = m.Where(dao.DevProduct.Columns().Status, 1).Count()
	if err != nil {
		return
	}
	data.Disable = data.Total - data.Enable
	return
}
