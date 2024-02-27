package influxdb

import (
	"sagooiot/pkg/iotModel"
)

// InsertDeviceData 插入设备数据
func (m *Influxdb) InsertDeviceData(deviceKey string, data iotModel.ReportPropertyData, subKey ...string) (err error) {

	return
}

// BatchInsertDeviceData 批量插入单设备的数据
func (m *Influxdb) BatchInsertDeviceData(deviceKey string, deviceDataList []iotModel.ReportPropertyData) (resultNum int, err error) {
	return
}

// BatchInsertMultiDeviceData 批量插入多设备的数据
func (m *Influxdb) BatchInsertMultiDeviceData(multiDeviceDataList map[string][]iotModel.ReportPropertyData) (resultNum int, err error) {
	return
}

// 监听设备数据日志
func (m *Influxdb) WatchDeviceData(deviceKey string, callback func(data iotModel.ReportPropertyData)) (err error) {

	return
}
