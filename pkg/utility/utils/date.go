package utils

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"strconv"
	"time"
)

// GetWeekDay 获取本周的开始时间和结束时间
func GetWeekDay() (string, string) {
	now := time.Now()
	start := now.Truncate(24*time.Hour).AddDate(0, 0, int(time.Monday-now.Weekday())).Format("2006-01-02") + " 00:00:00"
	end := now.Truncate(24*time.Hour).AddDate(0, 0, int(time.Monday-now.Weekday())+6).Format("2006-01-02") + " 23:59:59"
	return start, end
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func GetHourBetweenDates(sdate, edate string) []string {
	var d []string
	date := gtime.New(sdate)
	date2 := gtime.New(edate)
	date2Str := date2.Format("Y-m-d H:00:00")

	d = append(d, date.Format("Y-m-d H:00:00"))
	for {
		date = gtime.New(date).Add(time.Hour)
		dateStr := date.Format("Y-m-d H:00:00")
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// GetQuarterDay 获得当前季度的初始和结束日期
func GetQuarterDay() (string, string) {
	now := time.Now()
	year, quarter := now.Year(), (now.Month()-1)/3+1
	start := time.Date(year, (quarter-1)*3+1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	end := time.Date(year, (quarter-1)*3+1+2, daysIn((quarter-1)*3+1+2, year), 23, 59, 59, 0, time.Local).Format("2006-01-02 15:04:05")
	return start, end
}
func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// GetTimeByType 根据类型获取开始时间、结束时间及差值 1 天 2 周 3 月 4 年
func GetTimeByType(types int) (index int, begin string, end string) {
	switch types {
	case 1:
		begin = gtime.Now().Format("Y-m-d 00:00:00")
		end = gtime.Now().Format("Y-m-d H:i:s")
		index = gtime.Now().Hour() + 1
		break
	case 2:
		//begin, _ = GetWeekDay()
		begin = gtime.Now().AddDate(0, 0, -7).Format("Y-m-d 00:00:00")

		end = gtime.Now().Format("Y-m-d H:i:s")
		index = int(gtime.New(end).Sub(gtime.New(begin)).Hours()/24) + 1
		break
	case 3:
		//begin = gtime.Now().Format("Y-m-01 00:00:00")
		begin = gtime.Now().AddDate(0, 0, -30).Format("Y-m-d 00:00:00")
		//end = gtime.Now().AddDate(0, 1, 0).Format("Y-m-01 00:00:00")
		end = gtime.Now().Format("Y-m-d H:i:s")
		//index = gtime.Now().Day()
		index = int(gtime.New(end).Sub(gtime.New(begin)).Hours()/24) + 1
		break
	case 4:
		begin = gtime.Now().Format("Y-01-01 00:00:00")
		end = gtime.Now().AddDate(1, 0, 0).Format("Y-01-01 00:00:00")
		index = gtime.Now().Month()
		break
	default:
		begin = gtime.Now().Format("Y-m-d 00:00:00")
		end = gtime.Now().Format("Y-m-d H:i:s")
		index = gtime.Now().Hour() + 1
		break
	}
	return
}

// GetTime 根据类型和开始时间获取时间段及长度
func GetTime(i int, types int, begin string) (startTime string, endTime string, duration int, unit string) {
	switch types {
	case 1:
		h, _ := time.ParseDuration(strconv.Itoa(i) + "h")
		startTime = gtime.New(begin).Add(h).Format("Y-m-d H:i:s")
		endTime = gtime.New(startTime).Add(time.Hour).Format("Y-m-d H:i:s")
		duration = gtime.New(startTime).Hour()
		unit = "时"
		break
	case 2, 3:
		startTime = gtime.New(begin).AddDate(0, 0, i).Format("Y-m-d H:i:s")
		endTime = gtime.New(startTime).AddDate(0, 0, 1).Format("Y-m-d H:i:s")
		duration = gtime.New(startTime).Day()
		unit = "日"
	case 4:
		startTime = gtime.New(begin).AddDate(0, i, 0).Format("Y-m-d H:i:s")
		endTime = gtime.New(startTime).AddDate(0, 1, 0).Format("Y-m-d H:i:s")
		duration = gtime.New(startTime).Month()
		unit = "月"
		break
	default:
		startTime = gtime.New(begin).Add(time.Duration(i)).Format("Y-m-d H:i:s")
		endTime = gtime.New(startTime).Add(time.Hour).Format("Y-m-d H:i:s")
		duration = gtime.New(startTime).Hour()
		unit = "时"
		break
	}
	return
}

// CalcDaysFromYearMonth 返回给定年份和月份的天数
func CalcDaysFromYearMonth(year int, month int) int {
	// 获取指定月份的第一天
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	// 获取下一个月的第一天
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)
	// 获取当前月的最后一天
	lastDayOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	// 返回当前月份的天数
	return lastDayOfMonth.Day()
}

// GetTimeUnix 转为时间戳->秒数
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

// GetTimeMills 转为时间戳->毫秒数
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// GetTimeByInt 时间戳转时间
func GetTimeByInt(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

// GetHourDiffer  计算俩个时间差多少小时
func GetHourDiffer(startTime, endTime string) float64 {
	var hour float64
	t1, err := time.ParseInLocation(consts.TimeFormatCompact, startTime, time.Local)
	t2, err := time.ParseInLocation(consts.TimeFormatCompact, endTime, time.Local)
	if err == nil && CompareTime(t1, t2) {
		diff := GetTimeUnix(t2) - GetTimeUnix(t1)
		hour = float64(diff) / 3600
		return hour
	}
	return hour
}

// GetMinutesDiffer  计算俩个时间差多少分钟
func GetMinutesDiffer(startTime, endTime string) int {
	// 两个时间点
	t1 := gtime.NewFromStr(startTime)
	t2 := gtime.NewFromStr(endTime)

	// 计算时间差并转换为分钟数
	diff := t2.Sub(t1)
	return int(diff.Minutes())
}

// CompareTime 比较两个时间大小
func CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// IsSameDay 是否为同一天
func IsSameDay(t1, t2 int64) bool {
	y1, m1, d1 := time.Unix(t1, 0).Date()
	y2, m2, d2 := time.Unix(t2, 0).Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameMinute 是否为同一分钟
func IsSameMinute(t1, t2 int64) bool {
	d1 := time.Unix(t1, 0).Format("2006-01-02 15:04")
	d2 := time.Unix(t2, 0).Format("2006-01-02 15:04")
	return d1 == d2
}

// GetCurrentDateString 获取当前日期字符串
func GetCurrentDateString() string {
	return time.Now().Format("2006-01-02")
}

// GetTimeTagGroup 获取当前日期字符串,结果：2024:05
func GetTimeTagGroup() string {
	// 获取当前时间
	now := time.Now()
	// 获取当前年份
	year := now.Year()
	// 获取当前月份
	month := now.Month()
	return fmt.Sprintf("%d:%02d:", year, month)
}
