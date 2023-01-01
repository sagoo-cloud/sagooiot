package model

import "github.com/gogf/gf/v2/os/gtime"

// 设备日志 TDengine
type TdLog struct {
	Ts      *gtime.Time `json:"ts" dc:"时间"`
	Device  string      `json:"device" dc:"设备标识"`
	Type    string      `json:"type" dc:"日志类型"`
	Content string      `json:"content" dc:"日志内容"`
}

// 日志写入
type TdLogAddInput struct {
	Ts      *gtime.Time `json:"ts" dc:"时间"`
	Device  string      `json:"device" dc:"设备标识"`
	Type    string      `json:"type" dc:"日志类型"`
	Content string      `json:"content" dc:"日志内容"`
}
