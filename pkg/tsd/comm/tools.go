package comm

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/pkg/iotModel"
	"strings"
	"time"
)

// ProductTableName 获取TSD产品表名
func ProductTableName(key string) string {
	// td 表名加前缀，转义中划线
	return TdProductPrefix + strings.ToLower(strings.ReplaceAll(key, "-", "_"))
}

// DeviceTableName 获取TSD设备表名
func DeviceTableName(key string) string {
	// td 表名加前缀，转义中划线
	return TdDevicePrefix + strings.ToLower(strings.ReplaceAll(key, "-", "_"))
}

// DeviceLogTable 获取TSD设备日志表名
func DeviceLogTable(key string) string {
	// td 表名加前缀，转义中划线
	return TdLogPrefix + strings.ToLower(strings.ReplaceAll(key, "-", "_"))
}

// TsdColumnName 属性字段加前缀
func TsdColumnName(key string) string {
	key = strings.ToLower(key)
	if key == "ts" {
		return "ts"
	}
	return TdPropertyPrefix + key
}

// TsdTagName tag字段加前缀
func TsdTagName(key string) string {
	key = strings.ToLower(key)
	if key == "device" {
		return "device"
	}
	return TdTagPrefix + key
}

// GetDeviceField 获取设备数据key原始顺序列表与重构后的字段
func GetDeviceField(data iotModel.ReportPropertyData) (keys, field []string) {
	field = []string{"ts"}
	for k := range data {
		keys = append(keys, k)
		k = TsdColumnName(k)
		field = append(field, k)
		// 属性上报时间
		field = append(field, k+"_time")
	}
	return
}

// GetDeviceFieldAndValue 同时获取设备数据字段与值
func GetDeviceFieldAndValue(data iotModel.ReportPropertyData) (field, value []string) {
	field = []string{"ts"}
	for k, v := range data {
		value = append(value, "'"+gvar.New(v.Value).String()+"'")
		value = append(value, "'"+gtime.New(v.CreateTime).Format("Y-m-d H:i:s")+"'")
		k = TsdColumnName(k)
		field = append(field, k)
		// 属性上报时间
		field = append(field, k+"_time")
	}
	return
}

// GetDeviceValue 获取设备数据值
func GetDeviceValue(field []string, data iotModel.ReportPropertyData) []string {
	var value []string
	//跟据统一的key列表顺序，对数据值排序输出
	for _, key := range field {
		for k, v := range data {
			_, ok := data[key]
			if ok {
				if k == key {

					value = append(value, "'"+gvar.New(v.Value).String()+"'")
					value = append(value, "'"+gtime.New(v.CreateTime).Format("Y-m-d H:i:s")+"'")
				}
			}
		}
	}
	return value
}

// ChangeTime 连接时区处理
func ChangeTime(v *g.Var) (rs *g.Var) {
	driver := g.Cfg().MustGet(context.TODO(), "tdengine.type")
	if driver.String() == "taosRestful" {
		if t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", v.String()); err == nil {
			rs = gvar.New(t.Local().Format("2006-01-02 15:04:05"))
			return
		}
	}

	rs = v
	return
}
