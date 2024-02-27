package analysis

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/utility/utils"
	"time"
)

type sAnalysisDevice struct {
}

func init() {
	service.RegisterAnalysisDevice(analysisDeviceNew())
}

func analysisDeviceNew() *sAnalysisDevice {
	return &sAnalysisDevice{}
}

// GetDeviceDataTotalCount 获取设备消息总数统计,dataType :day,month,year
func (s *sAnalysisDevice) GetDeviceDataTotalCount(ctx context.Context, dataType string) (number int64, err error) {
	switch dataType {
	case "day":
		number = GetDeviceDataCountByToday()
	case "month":
		number = s.getDeviceDataCountTotal(ctx, "month")
	case "year":
		number = s.getDeviceDataCountTotal(ctx, "year")
	default:
		err = errors.New("dateType参数错误")
	}
	return
}

// GetDeviceOnlineOfflineCount 获取设备在线离线统计
func (s *sAnalysisDevice) GetDeviceOnlineOfflineCount(ctx context.Context) (res model.DeviceOnlineOfflineCount, err error) {
	//设备总量
	total, _ := cache.Instance().Get(ctx, consts.AnalysisDeviceCountPrefix+consts.DeviceTotal)
	if total.Val() != nil {
		res.Total = total.Int()
	} else {
		m := dao.DevDevice.Ctx(ctx)
		// 设备总量
		allNum, _ := m.Count()
		_ = cache.Instance().Set(ctx, consts.AnalysisDeviceCountPrefix+consts.DeviceTotal, allNum, 0)
		res.Total = allNum
	}

	//禁用设备数量
	disable, _ := cache.Instance().Get(ctx, consts.AnalysisDeviceCountPrefix+consts.DeviceDisable)
	if disable.Val() != nil {
		res.Disable = disable.Int()
	} else {
		m := dao.DevDevice.Ctx(ctx)
		disable, _ := m.Where(dao.DevDevice.Columns().Status, model.DeviceStatusNoEnable).Count()
		_ = cache.Instance().Set(ctx, consts.AnalysisDeviceCountPrefix+consts.DeviceDisable, disable, 0)
		res.Disable = disable
	}
	res.Online = dcache.CountDeviceOnlineNum()
	res.Offline = res.Total - res.Online
	return
}

// GetDeviceDataCountList 按年度每月设备消息统计，dataType 为统计数据类型 year:按年度,统计每个月的，month:按月份，统计每天的。当前的年与月
func (s *sAnalysisDevice) GetDeviceDataCountList(ctx context.Context, dateType string) (res []model.CountData, err error) {
	//res = make(map[string]int64)
	// 获取当前时间
	now := time.Now()
	// 获取当前年份
	year := now.Year()
	// 获取当前月份
	month := now.Month()

	switch dateType {
	case "month":
		// 获取当前月份的天数
		days := utils.CalcDaysFromYearMonth(year, int(month))
		for i := 1; i <= days; i++ {
			// 将日期格式化为字符串
			dayTag := fmt.Sprintf("%d-%02d-%02d", year, month, i)
			key := consts.AnalysisDeviceCountPrefix + consts.TodayMessageVolume + utils.GetTimeTagGroup() + dayTag
			countData, _ := cache.Instance().Get(context.Background(), key)
			var cd model.CountData
			cd.Title = gconv.String(i)
			cd.Value = countData.Int64()
			res = append(res, cd)
		}
	case "year":
		for i := 1; i <= 12; i++ {
			timeTag := fmt.Sprintf("%d:%02d", year, i)
			keys, err := dcache.SearchKey(consts.AnalysisDeviceCountPrefix + consts.TodayMessageVolume + timeTag)
			if err != nil {
				g.Log().Error(context.Background(), err.Error())
				continue
			}
			var number int64 = 0
			for _, v := range keys {
				countData, _ := cache.Instance().Get(context.Background(), v)
				number = number + countData.Int64()
			}
			var cd model.CountData
			cd.Title = gconv.String(i)
			cd.Value = number
			res = append(res, cd)
		}

	default:
		g.Log().Error(context.Background(), dateType, "dateType参数错误")
		return

	}
	return

}

// RemoveDeviceStatusCountCache 清除设备状态统计缓存
func RemoveDeviceStatusCountCache(ctx context.Context) {
	_, err := cache.Instance().Remove(ctx, consts.AnalysisDeviceCountPrefix+consts.DeviceTotal)
	if err != nil {
		g.Log().Debug(ctx, "清除设备状态统计缓存失败", err)
		return
	}
	_, err = cache.Instance().Remove(ctx, consts.AnalysisDeviceCountPrefix+consts.DeviceDisable)
	if err != nil {
		g.Log().Debug(ctx, "清除设备状态统计缓存失败", err)
		return
	}
}

// GetDeviceDataCountByToday 获取今日设备数据计数
func GetDeviceDataCountByToday() int64 {
	lastDate := utils.GetCurrentDateString()
	res, _ := cache.Instance().Get(context.Background(), consts.AnalysisDeviceCountPrefix+consts.TodayMessageVolume+utils.GetTimeTagGroup()+lastDate)
	return res.Int64()
}

func (s *sAnalysisDevice) getDeviceDataCountTotal(ctx context.Context, dateType string) (number int64) {
	switch dateType {
	case "month":
		dataList, err := s.GetDeviceDataCountList(ctx, "month")
		if err != nil {
			g.Log().Debug(ctx, "获取设备数据计数失败", err)
			return
		}
		for _, v := range dataList {
			var data = new(model.CountData)
			err := gconv.Scan(v, &data)
			if err != nil {
				continue
			}
			number = number + data.Value
		}

	case "year":
		dataList, err := s.GetDeviceDataCountList(ctx, "year")
		if err != nil {
			g.Log().Debug(ctx, "获取设备数据计数失败", err)
			return
		}
		for _, v := range dataList {
			var data = new(model.CountData)
			err := gconv.Scan(v, &data)
			if err != nil {
				continue
			}
			number = number + data.Value
		}

	default:
		keys, err := dcache.SearchKey(consts.AnalysisDeviceCountPrefix + consts.TodayMessageVolume)
		if err != nil {
			g.Log().Error(context.Background(), err.Error())
			return
		}
		for _, v := range keys {
			countData, _ := cache.Instance().Get(context.Background(), v)
			number = number + countData.Int64()
		}
	}
	return
}
