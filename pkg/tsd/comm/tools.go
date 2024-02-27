package comm

import (
	"strings"
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
