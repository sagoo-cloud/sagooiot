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
	"time"
)

type sAnalysisAlarm struct{}

func init() {
	service.RegisterAnalysisAlarm(analysisAlarmNew())
}
func analysisAlarmNew() *sAnalysisAlarm {
	return &sAnalysisAlarm{}
}

// GetDeviceAlertCountByYearMonth 按年度每月设备告警数统计
func (s *sAnalysisAlarm) GetDeviceAlertCountByYearMonth(ctx context.Context, year string) (res []model.CountData, err error) {
	timeTag := fmt.Sprintf("%s:%s", "year", year)
	resData, err := cache.Instance().GetOrSetFunc(ctx, consts.AnalysisAlarmCountPrefix+consts.AlarmMonthsMessageVolume+timeTag, func(ctx context.Context) (value interface{}, err error) {
		value, err = s.getAlarmDataCount(ctx, "year", year)
		return
	}, time.Second*15)
	err = gconv.Scan(resData.Val(), &res)
	return
}

// GetDeviceAlertCountByMonthDay 按月度每日设备告警数统计
func (s *sAnalysisAlarm) GetDeviceAlertCountByMonthDay(ctx context.Context, month string) (res []model.CountData, err error) {
	year := time.Now().Year()
	timeTag := fmt.Sprintf("%s:%s", gconv.String(year), month)
	resData, err := cache.Instance().GetOrSetFunc(ctx, consts.AnalysisAlarmCountPrefix+consts.AlarmMonthsMessageVolume+timeTag, func(ctx context.Context) (value interface{}, err error) {
		year := time.Now().Year()
		value, err = s.getAlarmDataCount(ctx, "month", gconv.String(year), month)
		return
	}, time.Hour)
	err = gconv.Scan(resData.Val(), &res)
	return
}

// GetDeviceAlertCountByDayHour 按日每小时设备告警数统计
func (s *sAnalysisAlarm) GetDeviceAlertCountByDayHour(ctx context.Context, day string) (res []model.CountData, err error) {
	year := time.Now().Year()
	month := time.Now().Month()
	timeTag := fmt.Sprintf("%s:%s:%s", gconv.String(year), month, day)
	resData, err := cache.Instance().GetOrSetFunc(ctx, consts.AnalysisAlarmCountPrefix+consts.AlarmMonthsMessageVolume+timeTag, func(ctx context.Context) (value interface{}, err error) {
		value, err = s.getAlarmDataCount(ctx, "month", gconv.String(year), gconv.String(month), day)
		return
	}, time.Second*15)

	err = gconv.Scan(resData.Val(), &res)
	return
}

// getAlarmDataCount 统计设备告警数据，dataType：year，month，day ，value：year，month，day
// dataType：year，value：year ，获取到指定年的每个月的统计
// dataType：month，value：year，month 获取指定的年、月的每天的统计
// dataType：day，value：year，month，day 获取指定年、月、日的每个小时的统计
func (s *sAnalysisAlarm) getAlarmDataCount(ctx context.Context, dataType string, value ...string) (res []model.CountData, err error) {

	m := dao.AlarmLog.Ctx(ctx)
	deviceKeys := getDeviceKeys(ctx, "")
	m = m.WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys)

	switch dataType {
	case "year":
		m = m.Fields("month(created_at) as title, count(*) as count").
			Where("year(created_at)=?", value[0]).
			Group("month(created_at)")

	case "month":
		m = m.Fields("day(created_at) AS title, count(*) as count").
			Where("year(created_at)=?", value[0]).
			Where("month(created_at)=?", value[1]).
			Group("day(created_at)")

	case "day":
		m = m.Fields("hour(created_at) as title, count(*) as count").
			Where("year(created_at)=?", value[0]).
			Where("month(created_at)=?", value[1]).
			Where("day(created_at)=?", value[2]).
			Group("hour(created_at)")
	}

	list, err := m.All()
	if err != nil {
		return
	}

	for _, v := range list {
		var countData model.CountData
		countData.Title = v["title"].String()
		countData.Value = v["count"].Int64()
		res = append(res, countData)
	}
	return
}

// GetAlarmTotalCount 告警总数统计（当年、当月、当日）,dataType :day,month,year ,date:2021 or 01 or21
func (s *sAnalysisAlarm) GetAlarmTotalCount(ctx context.Context, dataType, date string) (number int64, err error) {
	resData, err := cache.Instance().GetOrSetFunc(ctx, consts.AnalysisAlarmCountPrefix+consts.AlarmMonthsMessageVolume+dataType, func(ctx context.Context) (value interface{}, err error) {
		year := time.Now().Year()
		month := time.Now().Month()
		switch dataType {
		case "year":
			value = s.getAlarmTotalCount(ctx, "year", date)
		case "month":
			value = s.getAlarmTotalCount(ctx, "month", gconv.String(year), date)
		case "day":
			value = s.getAlarmTotalCount(ctx, "day", gconv.String(year), gconv.String(month), date)
		default:
			err = errors.New("参数错误")
		}
		return
	}, time.Minute*1)
	number = gconv.Int64(resData)
	return
}

func (s *sAnalysisAlarm) getAlarmTotalCount(ctx context.Context, dataType string, value ...string) (number int64) {
	m := dao.AlarmLog.Ctx(ctx)
	deviceKeys := getDeviceKeys(ctx, "")
	m = m.WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys)

	switch dataType {
	case "year":
		m = m.Fields("count(*) as count").
			Where("year(created_at)=?", value[0])
	case "month":
		m = m.Fields("count(*) as count").
			//WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).
			Where("year(created_at)=?", value[0]).
			Where("month(created_at)=?", value[1])
	case "day":
		m = m.Fields("count(*) as count").
			//WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).
			Where("year(created_at)=?", value[0]).
			Where("month(created_at)=?", value[1]).
			Where("day(created_at)=?", value[2])
	}
	res, err := m.One()
	if err != nil {
		g.Log().Debug(ctx, err.Error())
		return
	}
	number = res["count"].Int64()
	return

}

// GetAlarmLevelCount 告警级别统计
func (s *sAnalysisAlarm) GetAlarmLevelCount(ctx context.Context, dataType, date string) (res []model.CountData, err error) {
	timeTag := fmt.Sprintf("%s:%s", dataType, date)
	resData, err := cache.Instance().GetOrSetFunc(ctx, consts.AnalysisAlarmCountPrefix+consts.AlarmLevelMessageVolume+timeTag, func(ctx context.Context) (value interface{}, err error) {
		year := time.Now().Year()
		month := int(time.Now().Month())
		switch dataType {
		case "year":
			value, err = s.getAlarmLevelCount(ctx, "year", date)
		case "month":
			value, err = s.getAlarmLevelCount(ctx, "month", gconv.String(year), date)
		case "day":
			value, err = s.getAlarmLevelCount(ctx, "day", gconv.String(year), gconv.String(month), date)
		default:
			err = errors.New("参数错误")
		}
		return
	}, time.Second*15)
	err = gconv.Scan(resData.Val(), &res)
	return
}

func (s *sAnalysisAlarm) getAlarmLevelCount(ctx context.Context, dataType string, value ...string) (res []model.CountData, err error) {
	m := dao.AlarmLog.Ctx(ctx).Fields("level as title,count(*) as count")
	deviceKeys := getDeviceKeys(ctx, "")
	m = m.WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys)
	switch dataType {
	case "year":
		//m.WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).
		m = m.Where("year(created_at)=?", value[0])
	case "month":
		//m.WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).
		m = m.Where("year(created_at)=?", value[0]).
			Where("month(created_at)=?", value[1])
	case "day":
		//m.WhereIn(dao.AlarmLog.Columns().DeviceKey, deviceKeys).
		m = m.Where("year(created_at)=?", value[0]).
			Where("month(created_at)=?", value[1]).
			Where("day(created_at)=?", value[2])
	}
	resData, err := m.Group("level").All()
	if err != nil {
		g.Log().Debug(ctx, err.Error())
		return
	}

	for _, v := range resData {
		var countData model.CountData
		countData.Title = v["title"].String()
		countData.Value = v["count"].Int64()
		res = append(res, countData)
	}

	return
}
