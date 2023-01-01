package main

import "reflect"

type DeviceData struct {
	HeadStr     string   //字头
	DeviceID    string   //设备ID
	Signal      string   //信号质量
	Battery     string   //电池电量
	Temperature string   //温度
	Humidity    string   //湿度
	Cycle       string   //周期
	Update      []string //待上传
}

func (d DeviceData) IsEmpty() bool {
	return reflect.DeepEqual(d, DeviceData{})
}
