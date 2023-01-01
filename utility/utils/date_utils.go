package utils

import (
	"github.com/gogf/gf/v2/os/gtime"
	"strconv"
	"time"
)

// GetWeekDay 获取本周的开始时间和结束时间
func GetWeekDay() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02 15:04:05"
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

// GetQuarterDay 获得当前季度的初始和结束日期
func GetQuarterDay() (string, string) {
	year := time.Now().Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
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
		begin = gtime.Now().AddDate(0, 0, -6).Format("Y-m-d 00:00:00")

		end = gtime.Now().Format("Y-m-d H:i:s")
		index = int(gtime.New(end).Sub(gtime.New(begin)).Hours()/24) + 1
		break
	case 3:
		//begin = gtime.Now().Format("Y-m-01 00:00:00")
		begin = gtime.Now().AddDate(0, 0, -23).Format("Y-m-d 00:00:00")
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
