package analysis

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/tsd"
	"sagooiot/pkg/tsd/comm"
	"sort"
	"strings"
	"time"
)

type sAnalysisTsdData struct {
}

func init() {
	service.RegisterAnalysisTsdData(analysisTsdDataNew())
}

func analysisTsdDataNew() *sAnalysisTsdData {
	return &sAnalysisTsdData{}
}

type RawData struct {
	Ts    time.Time // 上报时间
	value float64   // 参数值

}

// GetDeviceIndicatorTrend 获取指标趋势
func (s *sAnalysisTsdData) GetDeviceIndicatorTrend(ctx context.Context, req model.DeviceIndicatorTrendReq) (rs []*model.DeviceIndicatorTrendRes, err error) {
	// 创建数据库连接
	table := comm.DeviceTableName(req.DeviceKey)
	db := tsd.GetDB()
	defer db.Close()
	if "" == req.ProductKey {
		err = fmt.Errorf("deviceKey is blank")
		return
	}
	//参数校验 是否存在产品；设备是否启用；起止时间是否有值
	flag, err := checkParam(ctx, req, 1)
	if err != nil && flag {
		return
	}
	//判断起止时间
	if req.StartDate == "" || req.EndDate == "" {
		return nil, errors.New("起止时间不能为空!")
	}

	sqlStr := fmt.Sprintf("select %s,%s from %s where ts >= '%s' and ts<= '%s'", "p_"+req.Properties, "p_"+req.Properties+"_time", table, req.StartDate, req.EndDate)

	rows, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	defer rows.Close()

	list, err := g.DB().GetCore().RowsToResult(ctx, rows)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}
	if !list.IsEmpty() {
		for _, v := range list {
			out := &model.DeviceIndicatorTrendRes{
				DataValue: v["p_"+req.Properties].Float64(),
				Date:      v["p_"+req.Properties+"_time"].String(),
			}
			rs = append(rs, out)
		}
	}
	return
}

// GetDeviceIndicatorPolymerize 获取指标聚合
func (s *sAnalysisTsdData) GetDeviceIndicatorPolymerize(ctx context.Context, req model.DeviceIndicatorPolymerizeReq) (rs []*model.DeviceIndicatorPolymerizeRes, err error) {
	startDate := req.StartDate
	endDate := req.EndDate
	startTime, err := time.Parse("2006-01-02 00:00:00", startDate)
	if err != nil {
		return rs, errors.New("开始时间格式错误")
	}
	endTime, err := time.Parse("2006-01-02 00:00:00", endDate)
	if err != nil {
		return rs, errors.New("结束时间格式错误")
	}
	fmt.Println(startTime, endTime)
	//参数校验 是否存在产品；设备是否启用；
	flag, err := checkParam(ctx, req, 2)
	if err != nil && flag {
		return
	}
	//判断起止时间
	if req.StartDate == "" || req.EndDate == "" {
		return nil, errors.New("起止时间不能为空!")
	}

	//1.读取设备参数的原始数据。
	// 创建数据库连接
	table := comm.DeviceTableName(req.DeviceKey)
	db := tsd.GetDB()
	defer db.Close()
	if "" == req.ProductKey {
		err = fmt.Errorf("deviceKey is blank")
		return
	}

	sqlStr := fmt.Sprintf("select %s,%s from %s where ts >= '%s' and ts<= '%s'", "p_"+req.Properties, "p_"+req.Properties+"_time", table, req.StartDate, req.EndDate)

	rows, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	defer rows.Close()

	list, err := g.DB().GetCore().RowsToResult(ctx, rows)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}
	var resList []RawData
	if !list.IsEmpty() {
		for _, v := range list {
			out := RawData{
				value: v["p_"+req.Properties].Float64(),
				Ts:    v["p_"+req.Properties+"_time"].Time(),
			}
			resList = append(resList, out)
		}
	}

	//2.根据聚合力度（5分钟、1小时、1天）和时间范围对数据进行分组。
	aggregationLevel := req.DateType
	timeRangeStart, _ := time.Parse("2006-01-02 15:04:05", req.StartDate)
	timeRangeEnd, _ := time.Parse("2006-01-02 15:04:05", req.EndDate)

	results, err := aggregateData(resList, aggregationLevel, timeRangeStart, timeRangeEnd)

	return results, err
}

// 聚合计算
func aggregateData(rawData []RawData, aggregationLevel string, timeRangeStart, timeRangeEnd time.Time) ([]*model.DeviceIndicatorPolymerizeRes, error) {
	interval, err := calculateInterval(aggregationLevel)
	if err != nil {
		return nil, err
	}

	// 2.按时间分组
	groups := make(map[string][]RawData)
	for _, data := range rawData {
		ts := data.Ts
		if err != nil {
			return nil, err
		}
		if ts.Before(timeRangeStart) || ts.After(timeRangeEnd) {
			continue // 过滤掉不在时间范围内的数据
		}
		roundedTime := ts.Round(interval)
		key := roundedTime.Format("2006-01-02 15:04")
		groups[key] = append(groups[key], data)
	}

	// 3.计算每个分组的最大值、最小值和平均值
	var results []*model.DeviceIndicatorPolymerizeRes
	for timeKey, group := range groups {
		var maxVal, minVal, sumVal float64
		for _, data := range group {
			if data.value > maxVal {
				maxVal = data.value
			}
			if data.value < minVal || minVal == 0 { // 初始化minVal为第一个值
				minVal = data.value
			}
			sumVal += data.value
		}
		avgVal := sumVal / float64(len(group))

		//4.将计算结果转换为analysis.DeviceIndicatorPolymerizeRes，并放入结果切片中。
		results = append(results, &model.DeviceIndicatorPolymerizeRes{
			Date:             timeKey,
			DataAverageValue: avgVal,
			DataMaxValue:     maxVal,
			DataMinValue:     minVal,
		})
	}

	// 按时间顺序排序
	sort.Slice(results, func(i, j int) bool {
		return results[i].Date < results[j].Date
	})

	return results, nil
}

// 根据聚合力度计算时间间隔
func calculateInterval(aggregationLevel string) (time.Duration, error) {
	switch aggregationLevel {
	case "1":
		return 5 * time.Minute, nil
	case "2":
		return time.Hour, nil
	case "3":
		return 24 * time.Hour, nil
	default:
		return 0, fmt.Errorf("unsupported aggregation level: %s", aggregationLevel)
	}
}

// 从rows中读取数据并转换为RawData切片
func readDataFromRows(rows *sql.Rows, column string) (res []RawData, err error) {
	var rawData []RawData
	columns, _ := rows.Columns()
	for rows.Next() {
		values := make([]any, len(columns))
		for i := range values {
			values[i] = new(any)
		}
		// 扫描行数据到 values 切片
		err = rows.Scan(values...)
		if err != nil {
			panic(err)
		}
		var row RawData
		for i, colName := range columns {
			if strings.HasPrefix(colName, column) {
				row.value = values[i].(float64)
			}
			if colName == "ts" {
				row.Ts = values[i].(time.Time)
			}
		}
		rawData = append(rawData, RawData{Ts: row.Ts, value: row.value})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rawData, err
}

func checkParam(ctx context.Context, req interface{}, ty int) (flag bool, err error) {
	var productKey = ""
	var deviceCode = ""
	if ty == 1 {
		trendReq := req.(model.DeviceIndicatorTrendReq)
		productKey = trendReq.ProductKey
		deviceCode = trendReq.DeviceKey
	} else if ty == 2 {
		polymerizeReq := req.(model.DeviceIndicatorPolymerizeReq)
		productKey = polymerizeReq.ProductKey
		deviceCode = polymerizeReq.DeviceKey
	} else {
		return false, errors.New("类型错误")
	}
	//1.判断是否存在产品
	out, err := service.DevProduct().Detail(ctx, productKey)
	if out == nil || err != nil {
		return false, errors.New("获取产品信息错误!")
	}
	//2.判断设备是否启用
	deviceOut, err := service.DevDevice().Get(ctx, deviceCode)
	if deviceOut == nil || err != nil {
		if deviceOut.Status == 0 {
			return false, errors.New("设备不在线!")
		}
		return false, errors.New("获取设备信息错误!")
	}
	return true, nil
}
